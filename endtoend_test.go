package main

import (
	"fmt"
	"go/build"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestEndToEnd(t *testing.T) {
	dir, err := ioutil.TempDir("", "exemplar")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)
	// Create exemplar in temporary directory.
	exemplar := filepath.Join(dir, "exemplar.exe")
	err = run("go", "build", "-o", exemplar, "main.go")
	if err != nil {
		t.Fatalf("building exemplar: %s", err)
	}
	testPropertizer(t, exemplar, dir)
}

func testPropertizer(t *testing.T, exemplar, dir string) {
	// Read the testdata directory.
	fd, err := os.Open("testdata/propertizer")
	if err != nil {
		t.Fatal(err)
	}
	defer fd.Close()
	names, err := fd.Readdirnames(-1)
	if err != nil {
		t.Fatalf("Readdirnames: %s", err)
	}
	// Generate, compile, and run the test programs.
	for _, name := range names {
		if !strings.HasSuffix(name, ".go") {
			//t.Errorf("%s is not a Go file", name)
			continue
		}
		if name == "cgo.go" && !build.Default.CgoEnabled {
			t.Logf("cgo is no enabled for %s", name)
			continue
		}
		// Names are known to be ASCII and long enough. Might make this
		//typeName := fmt.Sprintf("%c%s", name[0]+'A'-'a', name[1:len(name)-len(".go")])
		typeName := fmt.Sprintf("%c%s", name[0]+'A'-'a', name[1:len(name)-len(".go")])
		exemplarCompileAndRun(t, dir, "testdata/propertizer", exemplar, typeName, name, "_properties.go")
	}
}

func exemplarCompileAndRun(t *testing.T, dir, srcDir, exemplar, typeName, fileName, suffix string) {
	t.Logf("run: %s %s\n", fileName, typeName)
	source := filepath.Join(dir, fileName)
	err := copy(source, filepath.Join(srcDir, fileName))
	if err != nil {
		t.Fatalf("copying file to temporary directory: %s", err)
	}
	stringSource := filepath.Join(dir, typeName+suffix)
	// Run exemplar in temporary directory.
	err = run(exemplar, "propertizer", "--type", typeName, "--output", stringSource, source)
	if err != nil {
		t.Fatal(err)
	}
	// Run the binary in the temporary directory.
	err = run("go", "run", stringSource, source)
	if err != nil {
		t.Fatal(err)
	}
}

// copy copies the from file to the to file.
func copy(to, from string) error {
	toFd, err := os.Create(to)
	if err != nil {
		return err
	}
	defer toFd.Close()
	fromFd, err := os.Open(from)
	if err != nil {
		return err
	}
	defer fromFd.Close()
	_, err = io.Copy(toFd, fromFd)
	return err
}

// run runs a single command and returns an error if it does not succeed.
// os/exec should have this function, to be honest.
func run(name string, arg ...string) error {
	fmt.Printf("cmd: %s, args: %+v \n", name, arg)
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
