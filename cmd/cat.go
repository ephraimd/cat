package main

import (
	"fmt"
	"os"

	"github.com/ephraimd/cat/pkg"
)

func main() {
	// check the args for file name
	// if its there, then read the file and display it in stdout

	args := os.Args
	if len(args) < 2 {
		fmt.Fprintln(os.Stdout, "Expected an argument")
		return
	}

	filenames := args[1:]

	for _, filename := range filenames {

		if filename == "-" {
			pkg.ReadAndDisplayStandardInput()
			continue
		}

		pkg.ReadAndDisplayTextFile(filename, os.Stdout, false)
	}

}
