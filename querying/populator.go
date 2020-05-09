package querying

import (
	"io/ioutil"
	"path/filepath"
)

func InjectPreflectorFiles(directoryName string, preflectorFiles *[]string) {
	items, _ := ioutil.ReadDir(directoryName)
	for _, item := range items {
		if item.IsDir() {
			subitems, _ := ioutil.ReadDir(directoryName + "/" + item.Name())
			for _, subitem := range subitems {
				if !subitem.IsDir() {
					currentDirectory := directoryName + "/" + item.Name()
					handleFile(currentDirectory, subitem.Name(), preflectorFiles)
				} else {
					var subDirectory = directoryName + "/" + item.Name() + "/" + subitem.Name()
					InjectPreflectorFiles(subDirectory, preflectorFiles)
				}
			}
		} else {
			handleFile(directoryName, item.Name(), preflectorFiles)

		}
	}
}

func handleFile(currentDirectory, fileName string, preflectorFiles *[]string) {
	fullFileName := currentDirectory + "/" + fileName
	if isPreflectorFile(fileName) {
		*preflectorFiles = append(*preflectorFiles, fullFileName)
	}
}

func isPreflectorFile(fileName string) bool {
	isTargetFile, _ := filepath.Match("*.preflector.json", fileName)
	return isTargetFile
}
