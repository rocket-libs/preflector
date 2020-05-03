package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	preflectorFiles := make([]string, 0)
	injectPreflectorFiles(".", &preflectorFiles)
	for _, file := range preflectorFiles {
		fileContents := getFileContents(file)
		fmt.Print(fileContents)
	}
}

func getFileContents(file string) string {
	jsonFile, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()
	fileContents, err := bufio.NewScanner(jsonFile)
	return fileContents
}

func injectPreflectorFiles(directoryName string, preflectorFiles *[]string) {
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
					injectPreflectorFiles(subDirectory, preflectorFiles)
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
	isTargetFile, _ := filepath.Match("*.preflect.json", fileName)
	return isTargetFile
}
