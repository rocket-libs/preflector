package main

import (
	"fmt"
	"log"
	"rocketlibs/preflector/workers"
	"rocketlibs/preflector/workers/text_files"
)

func main() {
	preflectorFiles := make([]string, 0)
	workers.InjectPreflectorFiles(".", &preflectorFiles)
	fmt.Println(fmt.Sprintf("Found %d preflector files", len(preflectorFiles)))
	for _, file := range preflectorFiles {
		classDefinition, err := text_files.GetClassDefinition(file)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Println(classDefinition.Location)
	}
}
