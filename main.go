package main

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/hoenn/mimic/pkg/mimic"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: mimic {optional <file.go>} <interfaceA,interfaceB>")
		os.Exit(1)
	}

	// Use GOFILE if filename is not specified. Specifying filenames is useful
	// for a single generate.go file in a package with multiple generate directives.
	filepath := os.Getenv("GOFILE")
	interfaceNamesStr := ""
	if len(os.Args) == 3 {
		filepath = os.Args[1]
		interfaceNamesStr = os.Args[2]
	} else {
		interfaceNamesStr = os.Args[1]
	}
	interfaceNames := strings.Split(interfaceNamesStr, ",")
	filename := path.Base(filepath)
	fmt.Printf("Generating fakes for %v from %s\n", interfaceNames, filename)

	for i, n := range interfaceNames {
		interfaceNames[i] = strings.TrimSpace(n)
	}
	bs, err := mimic.ParseAndStubFromFile(interfaceNames, filepath, true)
	if err != nil {
		fmt.Println("Unable to parse and create fake:", err)
		os.Exit(1)
	}

	dir := path.Dir(filepath)
	fakeFilepath := path.Join(dir, fmt.Sprintf("%s_fake.go", strings.TrimSuffix(filename, ".go")))
	fmt.Printf("Writing %s\n", fakeFilepath)
	f, err := os.Create(fakeFilepath)
	if err != nil {
		fmt.Printf("Unable to open %s for writing: %s\n", fakeFilepath, err.Error())
		os.Exit(1)
	}
	defer func() {
		err := f.Close()
		if err != nil {
			fmt.Printf("Unable to close %s: %v", fakeFilepath, err)
		}
	}()

	_, err = f.Write(bs)
	if err != nil {
		fmt.Println("Unable to write fake:", err)
		os.Exit(1)
	}
}
