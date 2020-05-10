package querying_workers

import (
	"encoding/json"
	"rocketlibs/preflector/definitions"
	"rocketlibs/preflector/utils/text_files"
)

type FileReceiver struct {
	JsonProcessor IJsonProcessor
}

type IJsonProcessor interface {
	GetClassDefinitionFromFile(filename string) (classDefinition definitions.ClassDefinition, err error)
}

func (fileReceiver FileReceiver) GetClassDefinitionFromFile(filename string) (classDefinition definitions.ClassDefinition, err error) {
	byteValue, err := text_files.ReadBytes(filename)
	if err != nil {
		return classDefinition, err
	}
	err = json.Unmarshal(byteValue, &classDefinition)
	return classDefinition, err
}
