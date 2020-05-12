package line_managers

import (
	"bufio"
	"os"
)

type Receiver struct {
	LineProvider ILineProvider
}

type ILineProvider interface {
	Get() (line string, err error)
	Initialize(pathToFile string) (err error)
}

var scanner *bufio.Scanner

func (receiver Receiver) Initialize(pathToFile string) (err error) {
	targetFile, err := os.Open(pathToFile)
	if err != nil {
		return err
	}
	defer targetFile.Close()
	scanner = bufio.NewScanner(targetFile)
	return nil
}

func Get() (line string, err error) {
	return scanner.Text(), nil
}
