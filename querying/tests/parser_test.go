package tests

import (
	"rocketlibs/preflector/querying"
	"testing"
)

func TestClassNameReturnedCorrectly(t *testing.T) {
	const expectedClassName = "sample"
	className, _ := querying.GetClassName("sample.preflector.json")
	if className != expectedClassName {
		t.Errorf("Expected class name to be '%s' but instead was '%s'", expectedClassName, className)
	}
}
