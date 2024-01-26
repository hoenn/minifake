package main

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/hoenn/minifake/pkg/minifake"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: minifake <file.go> <interfaceA,interfaceB>")
		os.Exit(1)
	}

	filepath := os.Args[1]
	interfaceNamesStr := os.Args[2]
	interfaceNames := strings.Split(interfaceNamesStr, ",")
	filename := path.Base(filepath)
	fmt.Printf("Generating fakes for %v from %s\n", interfaceNames, filename)

	for i, n := range interfaceNames {
		interfaceNames[i] = strings.TrimSpace(n)
	}
	bs, err := minifake.ParseAndStubFromFile(interfaceNames, filepath, true)
	if err != nil {
		fmt.Println("Unable to parse and create fake:", err)
		os.Exit(1)
	}

	parts := strings.Split(filepath, filename)
	fakeFilepath := path.Join(parts[0], fmt.Sprintf("%s_fake.go", strings.TrimSuffix(filename, ".go")))
	fmt.Printf("Writing %s\n", fakeFilepath)
	f, err := os.Create(fakeFilepath)
	if err != nil {
		fmt.Printf("Unable to open %s for writing: %s\n", fakeFilepath, err.Error())
		os.Exit(1)
	}
	defer f.Close()

	_, err = f.Write(bs)
	if err != nil {
		fmt.Println("Unable to write fake:", err)
		os.Exit(1)
	}
}
