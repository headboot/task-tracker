package filecontroller

import (
	"errors"
	"fmt"
	"os"
)

const fileName string = "tasks.json"

func isFileExist(filePath string) bool {
	_, error := os.Stat(filePath)
	return errors.Is(error, os.ErrExist)
}

func CreateFile() (*os.File, error) {

	if isFileExist(fileName) {
		return OpenFile(), os.ErrExist
	}

	file, fileCreateError := os.Create(fileName)

	if fileCreateError != nil {
		return nil, fmt.Errorf("error: %s", fileCreateError)
	}

	return file, nil
}


func OpenFile() (*os.File) {
	file, error := os.Open(fileName)
	
	if error != nil {
		fmt.Println(error)
		os.Exit(1)
	}

	return file
}

func CloseFile(file *os.File) {
	file.Close()
}

func WriteToFile() {

}

func DeleteFromFile() {

}

func ReadFile(file *os.File) {
	
}
