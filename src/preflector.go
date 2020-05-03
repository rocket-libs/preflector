package main

import (
	"fmt"
)

func main() {
	preflectorFiles := make([]string, 0)

	preflector_files_populator.InjectPreflectorFiles(".", &preflectorFiles)
	for _, file := range preflectorFiles {
		fileContents := getFileContents(file)
		fmt.Print(fileContents)
	}
}
