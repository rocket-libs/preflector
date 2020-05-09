package writing

import (
	"fmt"
	"rocketlibs/preflector/definitions"
)

func Write(manyClassDefinitions []definitions.ClassDefinition) {
	for index, singleClassDefinition := range manyClassDefinitions {
		println(fmt.Sprintf("Definition %d of %d", index, len(manyClassDefinitions)))
		writeSingleClass(singleClassDefinition)
	}
}

func writeSingleClass(singleClassDefinition definitions.ClassDefinition) {
	println(singleClassDefinition.ClassName)
}
