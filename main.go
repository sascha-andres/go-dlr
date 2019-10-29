package main

import (
	"flag"
	"fmt"
	"livingit.de/code/go-dlr/internal"
	"os"
)

var fileToParse = "go.mod"

func init() {
	flag.StringVar(&fileToParse, "file", "go.mod", "Set file to read")

	flag.Parse()
}

func main() {
	if _, err := os.Stat(fileToParse); err == nil {
		result, err := internal.HasLocals(fileToParse)
		if err != nil {
			fmt.Printf("error reading go.mod: %s", err)
			os.Exit(1)
		}
		if result {
			os.Exit(1)
		}
	}
}