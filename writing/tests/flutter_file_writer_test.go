package tests

import (
	"rocketlibs/preflector/definitions"
	"rocketlibs/preflector/writing"
	"strings"
	"testing"
)

func TestSimplePlaceHoldersInsertedCorrectly(t *testing.T) {
	const className = "User"
	singleClassDefinition := definitions.ClassDefinition{
		Template:  "model",
		ClassName: className,
	}

	generatedClass, err := writing.GenerateClass(singleClassDefinition)
	if err != nil {
		t.Error(err)
	} else {
		doesNotContainClassName := strings.Contains(generatedClass, className) == false
		if doesNotContainClassName {
			t.Errorf("Class name was not found in string %s", generatedClass)
		}
	}
}
