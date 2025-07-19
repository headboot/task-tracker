package controller

import (
	"fmt"
	"strconv"

	"github.com/headboot/track-task/common"
	fileController "github.com/headboot/track-task/controller/file_controller"
	jsonconverter "github.com/headboot/track-task/controller/json_converter"
	"github.com/headboot/track-task/model"
	"github.com/headboot/track-task/model/commands"
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

func StartCommand(command commands.Command, commandArg []string) {

	switch command {

	case commands.Add:
		{
			lastArgument := commandArg[len(commandArg)-1]
			add(lastArgument)
		}

	case commands.Delete:
		{
			lastArgument := commandArg[len(commandArg)-1]
			delete(lastArgument)
		}

	case commands.List:
		{
			list()
		}

	case commands.Update:
		{
			id, err := strconv.ParseUint(commandArg[len(commandArg)-2], 10, 64)
			if err != nil {
				common.EndWithErr(err)
			}
			lastArg := commandArg[len(commandArg)-1]

			update(id, lastArg)
		}

	case commands.MarkDone:
		{
			id, err := strconv.ParseUint(commandArg[len(commandArg)-1], 10, 64)
			if err != nil {
				common.EndWithErr(err)
			}
			mark(id, model.Done)
		}

	case commands.MarkProgress:
		{
			id, err := strconv.ParseUint(commandArg[len(commandArg)-1], 10, 64)
			if err != nil {
				common.EndWithErr(err)
			}
			mark(id, model.Progress)
		}
	case commands.Help:
		{
			help()
		}
	default:
		fmt.Println("Unknown command for task-controller")
	}

}

func add(lastArgument string) {

	task := model.New(lastArgument)

	readedData := fileController.ReadFile()

	data := jsonconverter.ParseFromJson(readedData)

	data = append(data, task)

	result := jsonconverter.ParseToJson(data)

	fileController.WriteToFile(result)

}

func list() {
	data := fileController.ReadFile()
	
	result := jsonconverter.ParseFromJson(data)

	var tableData pterm.TableData = pterm.TableData{
		{"id", "task", "created", "updated", "status"},
	}


	for _, task := range result {
		taskString := []string{strconv.FormatUint(uint64(task.Id), 10), task.Description,task.CreatedAt.String(), task.UpdatedAt.String(), string(task.Status)}

		tableData = append(tableData, taskString)
	}

	pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(tableData).Render()


}


func delete(lastArg string) {
		id, err := strconv.ParseUint(lastArg,10, 64)

		if err != nil {
			common.EndWithErr(err)
		}

		fileData := fileController.ReadFile()

		data := jsonconverter.ParseFromJson(fileData)

		for index, task  := range data {
			if task.Id != id {
				continue
			}

			data = append(data[:index], data[index+1:]... )
			break
		}

		result := jsonconverter.ParseToJson(data)
		fileController.WriteToFile(result)
}


func update(id uint64, newText string) {
	data := fileController.ReadFile()

	tasks := jsonconverter.ParseFromJson(data)

	for index, task  := range tasks {
		if task.Id != id {
			continue
		}
		tasks[index] = task.Update(model.UpdateOptions{
			Text: newText,
		})
		break
	}

	result := jsonconverter.ParseToJson(tasks)

	fileController.WriteToFile(result)
}

func mark(id uint64, status model.Status) {
	data := fileController.ReadFile()

	tasks := jsonconverter.ParseFromJson(data)

	for index, task := range tasks {
		if id != task.Id {
			continue
		}
		tasks[index] = task.Update(
			model.UpdateOptions{
				Status: status,
			})
			break
	}

	result := jsonconverter.ParseToJson(tasks)

	fileController.WriteToFile(result)
}

func help() {
	s, _ := pterm.DefaultBigText.WithLetters(putils.LettersFromString("TaskT")).Srender()

	pterm.DefaultCenter.Println(s)

	pterm.DefaultCenter.Println(
		"use add <taskname> - to add new task \nlist - for list tasks\nupdate <id> <new task name>\nmark-done <id> - to mark status of task done\nmark-progress <id> - to mark status progress\ndelete <id> - delete task")
}