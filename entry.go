package main

import (
	"fmt"
	"rocketlibs/preflector/workers"
	"rocketlibs/preflector/workers/text_files"
)

func main() {
	preflectorFiles := make([]string, 0)

	workers.InjectPreflectorFiles(".", &preflectorFiles)
	for _, file := range preflectorFiles {
		fileContents := text_files.GetFileContents(file)
		fmt.Print(fileContents)
	}
}
