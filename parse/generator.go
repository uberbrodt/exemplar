package parse

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/build"
	"go/format"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	_ "golang.org/x/tools/go/gcimporter"
	"golang.org/x/tools/go/types"
)

//The caller will send a function of this type to do all the actual
//modification of the target package
type GeneratorFunc func(typeName string, fields []Field, imports []Import)

// File holds a single parsed file and associated data.
type File struct {
	pkg  *Package  // Package to which this file belongs.
	file *ast.File // Parsed AST.
	// These fields are reset for each type being generated.
	typeName string  // Name of the struct type.
	fields   []Field // Accumulator for the structs fields
	imports  []Import
}

// Value represents a struct field
type Field struct {
	Name     string // The name of the field.
	TypeName string //string representation of the Go Type of the field
	Tags     map[string]StructTag
}

type StructTag struct {
	Name  string
	Value string
}

type Import struct {
	ImportedName string
}

type Package struct {
	dir      string
	name     string
	defs     map[*ast.Ident]types.Object
	files    []*File
	typesPkg *types.Package
}

// Generator holds the state of the analysis. Primarily used to buffer
// the output for format.Source.
type Generator struct {
	buf bytes.Buffer // Accumulated output.
	pkg *Package     // Package we are scanning.
	dir string
}

//Run parses the target package and generates the code, verifying the package before and after generation.
//pathArgs is a list of file paths, to either individual files or whole directories.
//typeName is the name of a struct we're working on. outputName is where the generated code should go.
//genFn is the most important part, and recieves all the meta info about the targeted Type
func (g *Generator) Run(pathArgs []string, typeName string, outputName string, genFn GeneratorFunc) error {
	//Parse the package
	g.Prepare(pathArgs)

	// Print the header and package clause.
	g.Printf("// Code generated by \"exemplar%s\"; DO NOT EDIT\n", strings.Join(os.Args[1:], " "))
	g.Printf("\n")
	g.Printf("package %s", g.pkg.name)
	g.Printf("\n")

	g.collectAndGenerate(typeName, genFn)

	//format output
	src := g.format()

	// Write to file.
	//TODO: Fix this to not be tied to propertizer
	if outputName == "" {
		baseName := fmt.Sprintf("%s_properties.go", typeName)
		outputName = filepath.Join(g.dir, strings.ToLower(baseName))
	}
	err := ioutil.WriteFile(outputName, src, 0644)
	if err != nil {
		log.Fatalf("writing output: %s", err)
	}
	return nil
}

func (g *Generator) Prepare(args []string) {
	if len(args) == 1 && isDirectory(args[0]) {
		g.dir = args[0]
		g.parsePackageDir(args[0])
	} else {
		g.dir = filepath.Dir(args[0])
		g.parsePackageFiles(args)
	}
}

//collect gathers all the info from all the package's files about the type
//necessary to do fun things
func (g *Generator) collectAndGenerate(typeName string, genFn GeneratorFunc) {
	fields := make([]Field, 0, 100)
	imports := make([]Import, 0, 100)
	for _, file := range g.pkg.files {
		// Set the state for this run of the walker.
		file.typeName = typeName
		file.fields = nil
		if file.file != nil {
			ast.Inspect(file.file, file.genDecl)
			fields = append(fields, file.fields...)
			imports = append(imports, file.imports...)
		}
	}

	genFn(typeName, fields, imports)

}

// genDecl processes one declaration clause.
func (f *File) genDecl(node ast.Node) bool {
	decl, ok := node.(*ast.GenDecl)
	if !ok || decl.Tok != token.TYPE { // We only care about Type declarations.
		return true
	}
	// The name of the type of the constants we are declaring.
	// Can change if this is a multi-element declaration.
	typ := ""
	// Loop over the elements of the declaration. Each element is a ValueSpec:
	// a list of names possibly followed by a type, possibly followed by values.
	// If the type and value are both missing, we carry down the type (and value,
	// but the "go/types" package takes care of that).
	for _, spec := range decl.Specs {
		tspec := spec.(*ast.TypeSpec) // Guaranteed to succeed as this is TYPE.
		if tspec.Type != nil {
			// "X T". We have a type. Remember it.
			typ = tspec.Name.Name
		}
		if typ != f.typeName {
			// This is not the type we're looking for.
			continue
		}
		// We now have a list of names (from one line of source code) all being
		// declared with the desired type.

		structType, ok := tspec.Type.(*ast.StructType)
		if !ok {
			//not a struct type
			continue
		}

		typesObj, typeObjOk := f.pkg.defs[tspec.Name]
		if !typeObjOk {
			log.Fatalf("no type info found for struct %s", typ)
		}

		for _, fieldLine := range structType.Fields.List {
			for _, field := range fieldLine.Names {
				//skip struct padding
				if field.Name == "_" {
					continue
				}
				fieldObj, _, _ := types.LookupFieldOrMethod(typesObj.Type(), false, f.pkg.typesPkg, field.Name)

				typeStr := fieldObj.Type().String()
				tags := parseFieldTags(fieldLine.Tag)

				processedTypeStr, importPath := processTypeStr(typeStr)
				//log.Printf("processedTypeStr: %s, importPath: %s", processedTypeStr, importPath)

				if importPath != "" && !importExists(importPath, f.imports) {

					f.imports = append(f.imports, Import{importPath})

				}

				v := Field{
					Name:     field.Name,
					Tags:     tags,
					TypeName: processedTypeStr,
				}
				f.fields = append(f.fields, v)
			}
		}
	}
	return false
}

//We need to make sure that we get the type used right, with a package prefix
func processTypeStr(typeStr string) (typeName, importPath string) {
	if strings.Contains(typeStr, "/") {
		slashSplit := strings.Split(typeStr, "/")
		pkgNameAndType := slashSplit[len(slashSplit)-1]
		pkgName := strings.Split(pkgNameAndType, ".")[0]
		importPath := fmt.Sprintf("%s/%s", strings.Join(slashSplit[0:len(slashSplit)-1], "/"), pkgName)
		return pkgNameAndType, importPath
	} else if strings.Contains(typeStr, ".") {
		dotSplit := strings.Split(typeStr, ".")
		importPath := dotSplit[0]
		pkgNameAndType := typeStr
		return pkgNameAndType, importPath
	} else {
		return typeStr, ""
	}
}

//Check to see if a path already exists in the []Import
func importExists(pathName string, imports []Import) bool {
	for _, val := range imports {
		if pathName == val.ImportedName {
			return true
		}
	}
	return false
}

// parsePackageDir parses the package residing in the directory.
func (g *Generator) parsePackageDir(directory string) {
	pkg, err := build.Default.ImportDir(directory, 0)
	if err != nil {
		log.Fatalf("cannot process directory %s: %s", directory, err)
	}
	var names []string
	names = append(names, pkg.GoFiles...)
	names = append(names, pkg.CgoFiles...)
	// TODO: Need to think about constants in test files. Maybe write type_string_test.go
	// in a separate pass? For later.
	// names = append(names, pkg.TestGoFiles...) // These are also in the "foo" package.
	names = append(names, pkg.SFiles...)
	names = prefixDirectory(directory, names)
	g.parsePackage(directory, names, nil)
}

// parsePackage analyzes the single package constructed from the named files.
// If text is non-nil, it is a string to be used instead of the content of the file,
// to be used for testing. parsePackage exits if there is an error.
func (g *Generator) parsePackage(directory string, names []string, text interface{}) {
	var files []*File
	var astFiles []*ast.File
	g.pkg = new(Package)
	fs := token.NewFileSet()
	for _, name := range names {
		if !strings.HasSuffix(name, ".go") {
			continue
		}
		parsedFile, err := parser.ParseFile(fs, name, text, 0)
		if err != nil {
			log.Fatalf("parsing package: %s: %s", name, err)
		}
		astFiles = append(astFiles, parsedFile)
		files = append(files, &File{
			file: parsedFile,
			pkg:  g.pkg,
		})
	}
	if len(astFiles) == 0 {
		log.Fatalf("%s: no buildable Go files", directory)
	}
	g.pkg.name = astFiles[0].Name.Name
	g.pkg.files = files
	g.pkg.dir = directory
	// Type check the package.
	g.pkg.check(fs, astFiles)
}

// parsePackageFiles parses the package occupying the named files.
func (g *Generator) parsePackageFiles(names []string) {
	g.parsePackage(".", names, nil)
}

// prefixDirectory places the directory name on the beginning of each name in the list.
func prefixDirectory(directory string, names []string) []string {
	if directory == "." {
		return names
	}
	ret := make([]string, len(names))
	for i, name := range names {
		ret[i] = filepath.Join(directory, name)
	}
	return ret
}

func (g *Generator) Printf(format string, args ...interface{}) {
	fmt.Fprintf(&g.buf, format, args...)
}

//format returns the gofmt-ed contents of the Generator's buffer.
func (g *Generator) format() []byte {
	src, err := format.Source(g.buf.Bytes())
	if err != nil {
		// Should never happen, but can arise when developing this code.
		// The user can compile the output to see the error.
		log.Printf("warning: internal error: invalid Go generated: %s", err)
		log.Printf("warning: compile the package to analyze the error")
		return g.buf.Bytes()
	}
	return src
}

// check type-checks the package. The package must be OK to proceed.
func (pkg *Package) check(fs *token.FileSet, astFiles []*ast.File) {
	pkg.defs = make(map[*ast.Ident]types.Object)
	config := types.Config{FakeImportC: true, Error: func(err error) {
		cErr := err.(types.Error)
		if cErr.Soft {
			return
		}
		if strings.Contains(cErr.Msg, "has no field or method") ||
			strings.Contains(cErr.Msg, "invalid operation: cannot call non-function") {
			log.Printf("IGNORED: during package check: %s", cErr.Msg)
			return
		}
		log.Fatalf("checking package: %s", cErr.Msg)
	}}
	info := &types.Info{
		Defs: pkg.defs,
	}
	typesPkg, _ := config.Check(pkg.dir, fs, astFiles, info)
	pkg.typesPkg = typesPkg
}

// isDirectory reports whether the named file is a directory.
func isDirectory(name string) bool {
	info, err := os.Stat(name)
	if err != nil {
		log.Fatal(err)
	}
	return info.IsDir()
}

func parseFieldTags(tagString *ast.BasicLit) map[string]StructTag {
	tagMap := make(map[string]StructTag)
	if tagString != nil {
		sanitized := strings.Replace(tagString.Value, "`", "", -1)

		var buffer []byte = make([]byte, 0, 10)
		var key string
		var inTag bool
		for i := 0; i < len(sanitized); i++ {
			if sanitized[i] == ':' {
				key = bytes.NewBuffer(buffer).String()
				buffer = make([]byte, 0, 10)
				continue
			}
			if sanitized[i] == '"' {
				if inTag {
					tagMap[key] = StructTag{Name: key, Value: bytes.NewBuffer(buffer).String()}
					buffer, key = make([]byte, 0, 10), ""
					//key = ""
					inTag = false
					continue
				} else {
					inTag = true
					continue
				}
			}
			buffer = append(buffer, sanitized[i])
		}
	}
	return tagMap
}
