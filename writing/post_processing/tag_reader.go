package post_processing

import (
	"errors"
	"fmt"
	"rocketlibs/preflector/utils/string_utils"
	"strings"
)

func GetDirective(line string, candidatePostProcessor string) (tag string, err error) {
	const tagTerminator = "/>"
	postProcessorTag := getPostProcessorTag(candidatePostProcessor)
	indexOfTagStart := strings.Index(line, postProcessorTag)
	if indexOfTagStart < 0 {
		return "", errors.New(fmt.Sprintf("Could not find tag beginning with %s in line %s", postProcessorTag, line))
	}
	textAfterTagStart, err := string_utils.SubString(line, indexOfTagStart, len(line)-indexOfTagStart)
	if err != nil {
		return "", err
	}

	indexOfTagEnd := strings.Index(textAfterTagStart, tagTerminator)
	if indexOfTagEnd < 0 {
		return "", errors.New(fmt.Sprintf("Could not find tag terminator '%s'", tagTerminator))
	}
	tag, err = string_utils.SubString(textAfterTagStart, 0, indexOfTagEnd+len(tagTerminator))
	if err != nil {
		return "", err
	}
	return tag, nil
}
