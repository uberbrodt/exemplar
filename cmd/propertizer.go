// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"bytes"
	"log"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
	"github.com/uberbrodt/exemplar/parse"
)

var typeFlag string
var outputFlag string
var getterPrefixFlag bool
var interfaceNameFlag string

func init() {
	RootCmd.AddCommand(propertizerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// propertizerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	propertizerCmd.Flags().StringVarP(&typeFlag, "type", "t", "", "comma-separated list of type names; must be set")
	propertizerCmd.Flags().StringVarP(&outputFlag, "output", "o", "", "output file name; default srcdir/<stype>_properties.go")
	propertizerCmd.Flags().BoolVar(&getterPrefixFlag, "getterPrefix", false, "Prefix the accessor methods with Get")
	propertizerCmd.Flags().StringVar(&interfaceNameFlag, "interfaceName", "", "If set, creates an interface with the name provided for the struct methods")

}

// propertizerCmd represents the propertizer command
var propertizerCmd = &cobra.Command{
	Use:   "propertizer",
	Short: "Create getters and setters for a struct",
	Long: `Generate getters, setters and interfaces for a struct
	`,
	Run: func(cmd *cobra.Command, args []string) {
		g := new(parse.Generator)

		path := cmd.Flags().Args()

		propertizer := func(typeName string, fields []parse.Field, imports []parse.Import) {

			g.Printf("import \"fmt\"\n") // Used by all methods.
			if len(fields) == 0 {
				log.Fatalf("no fields defined for type %s", typeName)
			}

			for _, imprt := range imports {
				g.Printf("import \"%s\" \n", imprt.ImportedName)
			}

			for _, field := range fields {
				propertizerTag := field.Tags["propertizer"]
				if propertizerTag.Value != "private" {
					buildAccessor(g, field, typeName)
					buildMutator(g, field, typeName)
				}
			}

			if interfaceNameFlag != "" {
				g.Printf("\n")
				g.Printf("type %s interface {\n", interfaceNameFlag)
				for _, field := range fields {
					propertizerTag := field.Tags["propertizer"]
					if propertizerTag.Value != "private" {
						buildInterfaceAccessor(g, field, typeName)
						buildInterfaceMutator(g, field, typeName)
					}

				}

				g.Printf("}")
			}
		}

		g.Run(path, typeFlag, outputFlag, propertizer)

	},
}

func buildInterfaceAccessor(g *parse.Generator, f parse.Field, typeName string) {
	tempStr := strings.Replace(strings.Title(CamelCase(f.Name)), "Id", "ID", -1)
	if getterPrefixFlag {
		g.Printf("Get%s() %s\n", tempStr, f.TypeName)
	} else {
		g.Printf("%s() %s\n", tempStr, f.TypeName)
	}
}

func buildInterfaceMutator(g *parse.Generator, f parse.Field, typeName string) {
	tempStr := strings.Replace(strings.Title(CamelCase(f.Name)), "Id", "ID", -1)
	g.Printf("Set%s(x %s) \n", tempStr, f.TypeName)
}

func buildAccessor(g *parse.Generator, f parse.Field, typeName string) {
	//Make Id ID per Go std, upper case private fields
	tempStr := strings.Replace(strings.Title(CamelCase(f.Name)), "Id", "ID", -1)
	if getterPrefixFlag {
		g.Printf("func (this *%s) Get%s() %s {\n", typeName, tempStr,
			f.TypeName)
	} else {
		g.Printf("func (this *%s) %s() %s {\n", typeName, tempStr,
			f.TypeName)
	}
	g.Printf("  return this.%s \n", f.Name)
	g.Printf("}\n")
}

func buildMutator(g *parse.Generator, f parse.Field, typeName string) {
	tempStr := strings.Replace(strings.Title(CamelCase(f.Name)), "Id", "ID", -1)
	g.Printf("func (this *%s) Set%s(x %s) {\n", typeName, tempStr, f.TypeName)
	g.Printf("  this.%s = x \n", f.Name)
	g.Printf("}\n")
}

var camelingRegex = regexp.MustCompile("[0-9A-Za-z]+")

func CamelCase(src string) string {
	byteSrc := []byte(src)
	chunks := camelingRegex.FindAll(byteSrc, -1)
	for idx, val := range chunks {
		if idx > 0 {
			chunks[idx] = bytes.Title(val)
		}
	}
	return string(bytes.Join(chunks, nil))
}
