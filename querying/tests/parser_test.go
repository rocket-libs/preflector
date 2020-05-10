package tests

import (
	"rocketlibs/preflector/definitions"
	"rocketlibs/preflector/querying"
	"testing"

	"github.com/stretchr/testify/mock"
)

type MockFileReceiver struct {
	mock.Mock
}

func (fileReceiver *MockFileReceiver) GetClassDefinitionFromFile(filename string) (classDefinition definitions.ClassDefinition, err error) {
	returnVals := fileReceiver.Called(filename)
	return returnVals.Get(0).(definitions.ClassDefinition), returnVals.Error(1)
}

func TestClassNameReturnedCorrectly(t *testing.T) {
	const expectedClassName = "sample"
	const inputFilename = "sample.preflector.json"
	fileReceiverMock := new(MockFileReceiver)
	fileReceiverMock.On("GetClassDefinitionFromFile", inputFilename).Return(definitions.ClassDefinition{}, nil)
	fileReceiverInstance := MockFileReceiver{fileReceiverMock.Mock}
	classDefinition, _ := querying.GetClassDefinition(inputFilename, &fileReceiverInstance)
	if classDefinition.ClassName != expectedClassName {
		t.Errorf("Expected class name to be '%s' but instead was '%s'", expectedClassName, classDefinition.ClassName)
	}
}
