package post_processing

import (
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
	return nil
}

func containsPostProcessor(line string, candidatePostProcessor string) bool {
	postProcessorTag := getPostProcessorTag(candidatePostProcessor)
	return strings.Contains(line, postProcessorTag)
}

func getPostProcessorTag(candidatePostProcessor string) string {
	return "<" + candidatePostProcessor
}
