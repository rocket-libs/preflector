package delegate

import (
	"fmt"
	"rocketlibs/preflector/definitions"
	"rocketlibs/preflector/querying"
)

func Work(directory string) error {
	if directory == "" {
		directory = "."
	}
	preflectorFiles := make([]string, 0)
	querying.InjectPreflectorFiles(directory, &preflectorFiles)
	fmt.Println(fmt.Sprintf("Found %d preflector files", len(preflectorFiles)))
	manyClassDefinitions := make([]definitions.ClassDefinition, 0)
	for _, file := range preflectorFiles {
		singleClassDefinition, err := querying.GetClassDefinition(file)
		if err != nil {
			return err
		}
		manyClassDefinitions = append(manyClassDefinitions, singleClassDefinition)
	}
	return nil
}
