package querying

import (
	"errors"
	"fmt"
	"log"
	"path/filepath"
	"rocketlibs/preflector/definitions"
	"rocketlibs/preflector/querying/querying_workers"
	"strings"
)

func GetClassDefinition(file string, jsonProcessor querying_workers.IJsonProcessor) (classDefinition definitions.ClassDefinition, err error) {

	log.Println("Parsing: " + file)

	classDefinition, err = jsonProcessor.GetClassDefinitionFromFile(file)
	if err != nil {
		return classDefinition, err
	}
	classDefinition.ClassName, err = getClassName(file)
	return classDefinition, err
}

func getClassName(filename string) (className string, err error) {
	qualifiedFilename := filepath.Base(filename)
	bits := strings.Split(qualifiedFilename, ".")
	filenameBitsCount := len(bits)
	errInvalidNameFormat := errors.New(fmt.Sprintf("File '%s' does not follow naming conventions of preflector files", filename))
	if filenameBitsCount == 0 {
		return "", errInvalidNameFormat
	} else {
		secondLastBitIsPreflector := bits[filenameBitsCount-2] == "preflector"
		invalidSecondLastBit := secondLastBitIsPreflector == false
		if invalidSecondLastBit {
			return "", errInvalidNameFormat
		} else {
			return bits[0], nil
		}
	}

}
