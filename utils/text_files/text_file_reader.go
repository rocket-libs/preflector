package text_files

import (
	"fmt"
	"io/ioutil"
	"os"
)

func ReadBytes(filename string) (byteValue []byte, err error) {
	println(fmt.Sprintf("Loading text file: %s", filename))
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return ioutil.ReadAll(file)
}
