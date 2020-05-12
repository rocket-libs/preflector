package tests

import (
	"rocketlibs/preflector/writing/post_processing"
	"testing"
)

func TestSelfClosingTagReturnedCorrectly(t *testing.T) {
	const line = "blah blah<trim args=',' />some more stuff"
	tag, err := post_processing.GetSelfClosingTag(line, "trim")
	if err != nil {
		t.Error(err)
	} else if tag != "<trim args=',' />" {
		t.Errorf("Tag not retrieved correctly. Result '%s'", tag)
	}
}

func TestVerboseTagReturnedCorrectly(t *testing.T) {
	const line = "blah<loop args='Properties'>final this.{{Name}},</loop>blah"
	tag, err := post_processing.GetVerboseTag(line, "loop")
	if err != nil {
		t.Error(err)
	} else if tag != "<loop args='Properties'>final this.{{Name}},</loop>" {
		t.Errorf("Tag not retrieved correctly. Result '%s'", tag)
	}
}
