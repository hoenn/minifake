package main

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/hoenn/ministub/pkg/ministub"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: ministub <file.go> <interfaceA,interfaceB>")
		os.Exit(1)
	}

	filepath := os.Args[1]
	interfaceNamesStr := os.Args[2]
	interfaceNames := strings.Split(interfaceNamesStr, ",")

	for i, n := range interfaceNames {
		interfaceNames[i] = strings.TrimSpace(n)
	}
	bs, err := ministub.ParseAndStub(interfaceNames, filepath, "", true)
	if err != nil {
		panic(err)
	}

	filename := path.Base(filepath)
	parts := strings.Split(filepath, filename)
	fakeFilepath := path.Join(parts[0], fmt.Sprintf("%s_fake.go", strings.TrimSuffix(filename, ".go")))
	f, err := os.Create(fakeFilepath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = f.Write(bs)
	if err != nil {
		panic(err)
	}
}
