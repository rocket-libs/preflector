package tests

import (
	"os"
	"path/filepath"
	"rocketlibs/preflector/delegate"
	"testing"
)

func TestEndToEnd(t *testing.T) {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	delegate.Work(dir)
}
