package tests

import (
	"rocketlibs/preflector/utils/string_utils"
	"testing"
)

func TestSubStringWorks(t *testing.T) {
	const inputString = "donkey"
	outputString, _ := string_utils.SubString(inputString, 2, 4)
	if outputString != "nkey" {
		t.Errorf("Expected substring to be 'nkey' but instead was '%s'", outputString)
	}
}
