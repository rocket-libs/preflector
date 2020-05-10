package delegate

import (
	"fmt"
	"rocketlibs/preflector/definitions"
	"rocketlibs/preflector/querying"
	"rocketlibs/preflector/querying/querying_workers"
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
		singleClassDefinition, err := querying.GetClassDefinition(file, &querying_workers.FileReceiver{})
		if err != nil {
			return err
		}
		manyClassDefinitions = append(manyClassDefinitions, singleClassDefinition)
	}
	return nil
}
