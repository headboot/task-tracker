package filecontroller

import (
	"errors"
	"os"

	"github.com/headboot/track-task/common"
	jsonconverter "github.com/headboot/track-task/controller/json_converter"
	"github.com/headboot/track-task/model"
)

const (
	fileName string = "tasks.json"
)

/* returns true if file is exist or false if file is not exist*/
func isFileExist() bool {

	_, err := os.Stat(fileName)

	if errors.Is(err, os.ErrNotExist) {
		return false
	} else {
		return true
	}
}

func OpenFile() *os.File {
	var (
		file *os.File
		err  error
	)

	if isFileExist() {
		file, err = os.OpenFile(fileName, os.O_RDWR|os.O_APPEND|os.O_CREATE|os.O_TRUNC, os.ModeAppend)

		if err != nil {
			common.EndWithErr(err)
		}

		return file
	}

	file, err = os.Create(fileName)

	if err != nil {
		common.EndWithErr(err)
	}

	data := jsonconverter.ParseToJson([]model.Task{})

	WriteToFile(data)

	return file
}

func ReadFile() ([]byte) {

	if !(isFileExist()) {
		OpenFile()
	}

	content, err := os.ReadFile(fileName)

	if err != nil && !errors.Is(err, os.ErrNotExist) {
		common.EndWithErr(err)
	}

	if errors.Is(err, os.ErrNotExist) {
		return []byte{}
	}

	return content
}

func WriteToFile(data []byte) {
	file := OpenFile()

	result := append(data, []byte("\n")...)

	if _, err := file.Write(result); err != nil {
		common.EndWithErr(err)
	}


	defer file.Close()
}
