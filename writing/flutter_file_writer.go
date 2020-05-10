package writing

import (
	"fmt"
	"rocketlibs/preflector/definitions"
)

func WriteFlutterClass(manyClassDefinitions []definitions.ClassDefinition) {
	for index, singleClassDefinition := range manyClassDefinitions {
		println(fmt.Sprintf("Definition %d of %d", index, len(manyClassDefinitions)))
		GenerateClass(singleClassDefinition)
	}
}
