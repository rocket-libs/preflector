package tests

import (
	"rocketlibs/preflector/writing/post_processing"
	"testing"
)

func TestTagReturnedCorrectly(t *testing.T){
	const line = "blah blah<trim args=',' />some more stuff"
	tag, err := post_processing.GetDirective(line,"trim");
	if(err != nil){
		t.Error(err)
	}else if (tag != "<trim args=',' />"){
		t.Errorf("Tag not retrieved correctly. Result '%s'",tag)
	}
}