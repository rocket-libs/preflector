package querying_workers

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"rocketlibs/preflector/definitions"
)

type FileReceiver struct {
}

type IJsonProcessor interface {
	GetClassDefinitionFromFile(filename string) (classDefinition definitions.ClassDefinition, err error)
}

func (fileReceiver FileReceiver) GetClassDefinitionFromFile(filename string) (classDefinition definitions.ClassDefinition, err error) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		return classDefinition, err
	}
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return classDefinition, err
	}
	err = json.Unmarshal(byteValue, &classDefinition)
	return classDefinition, err
}
