package post_processing

import (
	"errors"
	"fmt"
	"strings"
)

type PostProcessorDirective struct {
	Command string
	Args    string
}

func GetPostProcessors(line string) []PostProcessorDirective {
	if line == "" {
		return nil
	}
	appliedPostProcessors := make([]PostProcessorDirective, 0)
	supportedPostProcessors := make([]string, 0)
	supportedPostProcessors = append(supportedPostProcessors, "trim")
	for _, candidatePostProcessor := range supportedPostProcessors {
		if containsPostProcessor(line, candidatePostProcessor) {
			appliedPostProcessors = append(appliedPostProcessors)
		}
	}
}

func getDirective(line string, candidatePostProcessor string) (postProcessorDirective PostProcessorDirective, err error) {
	postProcessorTag := getPostProcessorTag(candidatePostProcessor)
	indexOfTagStart := strings.Index(line, postProcessorTag)
	if indexOfTagStart < 0 {
		return postProcessorDirective, errors.New(fmt.Sprintf("Could not find tag beginning with %s in line %s", postProcessorTag, line))
	}
	indexOfTagEnd := line[indexOfTagStart]

}

func containsPostProcessor(line string, candidatePostProcessor string) bool {
	postProcessorTag := getPostProcessorTag(candidatePostProcessor)
	return strings.Contains(line, postProcessorTag)
}

func getPostProcessorTag(candidatePostProcessor string) string {
	return "<" + candidatePostProcessor
}
