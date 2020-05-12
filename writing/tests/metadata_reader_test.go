package tests

import (
	"encoding/json"
	"rocketlibs/preflector/writing"
	"testing"
)

func TestMetaDataReadCorrectly(t *testing.T) {
	jsonString := `{"Blah":"Value","Metadata":{"Template":"basic.htm","OutputFile":"hello.htm"}}`
	var jsonMap map[string]interface{}
	json.Unmarshal([]byte(jsonString), &jsonMap)
	metadata, err := writing.GetMetadata(jsonMap)
	if err != nil {
		t.Errorf("Error extracting metadata from json")
		t.Error(err)
	} else {
		if metadata.OutputFile != "hello.htm" {
			t.Errorf("Expected output file to be 'hello.htm' but instead was '%s'", metadata.OutputFile)
		}
		if metadata.Template != "basic.htm" {
			t.Errorf("Expected template to be 'basic.htm' but instead was '%s'", metadata.Template)
		}
	}
}
