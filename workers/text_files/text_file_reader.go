package text_files

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"rocketlibs/preflector/definitions"
)

func GetClassDefinition(file string) (classDefinition definitions.ClassDefinition, err error) {
	log.Println("Parsing: " + file)
	jsonFile, err := os.Open(file)
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
