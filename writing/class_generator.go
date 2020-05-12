package writing

import (
	"fmt"
	"rocketlibs/preflector/definitions"
	"rocketlibs/preflector/utils/text_files"
	"strings"
)

func GenerateClass(classDefinition definitions.ClassDefinition) (generatedFileContent string, err error) {
	println(fmt.Sprintf("Writing class: %s", classDefinition.ClassName))
	byteValue, err := text_files.ReadBytes("/home/nyingi/work/code/rocket/go/preflector/writing/templates/" + classDefinition.Template + ".htm")
	if err != nil {
		return "", err
	}
	templateString := string(byteValue)
	templateLines := strings.Split(templateString, "\n")
	generatedFileContent = ""
	//currentPreprocessors := make([]string, 0)
	for index, currentLine := range templateLines {
		println(fmt.Sprintf("\tProcessing line %d of %d", index, len(templateLines)))
		generatedFileContent += replaceSimplePlaceholders(currentLine, classDefinition)
		generatedFileContent += "\n"
	}
	return generatedFileContent, nil
}

func replaceSimplePlaceholders(line string, classDefinition definitions.ClassDefinition) string {
	line = strings.ReplaceAll(line, "{{ClassName}}", classDefinition.ClassName)
	return line
}
