package controller

import (
	"fmt"

	"github.com/headboot/track-task/commands"
)

func StartCommand(command commands.Command, commandArg []string) {

	switch command {

	case commands.Add:
		{
			add(commandArg)
		}

	case commands.Delete:
		{
			fmt.Println("deleted", commandArg[len(commandArg)-1])
		}

	case commands.List:
		{
			fmt.Println("listed")
		}

	case commands.Update:
		{
			fmt.Println("updated", commandArg[len(commandArg)-1])
		}

	case commands.MarkDone:
		{
			fmt.Println("marked-done")
		}

	case commands.MarkProgress:
		{
			fmt.Println("marked-progress")
		}
	case commands.Help: {
		fmt.Println("help:\nto add task use command: add 'task'\nto get list of tasks use command: list")
	}
	default:
		fmt.Println("Unknown command for controller")
	}

}

func add(commandArg []string) {
	lastArgument := commandArg[len(commandArg)-1]

	fmt.Printf("added %s \n", lastArgument)
}


