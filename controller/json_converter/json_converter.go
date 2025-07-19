package jsonconverter

import (
	"encoding/json"
	"io"

	"github.com/headboot/track-task/common"
	"github.com/headboot/track-task/model"
)

func ParseToJson(task []model.Task) ([]byte) {
	data, err := json.Marshal(task)

	if err != nil {
		common.EndWithErr(err)
	}

	return data
}


 func ParseFromJson(data []byte) ([]model.Task) {
	 tasks := []model.Task{
		
	 }
	 err := json.Unmarshal(data, &tasks)

	 if err != nil && err != io.EOF {
			common.EndWithErr(err)
	 }

	 return tasks
 }