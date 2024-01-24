package main

import (
	"fmt"
	"os"

	"github.com/hoenn/ministub/pkg/ministub"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: ministub <file.go> <interface name>")
		os.Exit(1)
	}

	filepath := os.Args[1]
	interfaceName := os.Args[2]
	bs, err := ministub.ParseAndStub(interfaceName, filepath, "", true)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bs))
}
