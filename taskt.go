package main

import (
	"fmt"
	"os"
	"slices"

	"github.com/headboot/track-task/common"
	"github.com/headboot/track-task/controller"
	"github.com/headboot/track-task/model/commands"
)

func main() {

	args := os.Args

	if len(args) < 2 {
		fmt.Println("nothing to do, use --help command")
		os.Exit(1)
	}

	command, err := getCommandFromArg(args[commands.CommandIndex])

	if err != nil || command == commands.UnknownCommand {
		common.EndWithErr(err)
	}

	if error := checkCommandArg(command, &args); error != nil {
		common.EndWithErr(error)
	}

	controller.StartCommand(command, args[commands.CommandIndex+1:])

}

func getCommandFromArg(commandArg string) (commands.Command, error) {
	command, ok := commands.ComandsMap[commandArg]

	if !ok {
		return commands.UnknownCommand, fmt.Errorf("unknown command: %s", commandArg)
	}

	return command, nil
}

func checkCommandArg(cmd commands.Command, args *[]string) error {
	if slices.Contains(commands.CommandsNeed3Args, cmd) && len(*args) == 3 {
		return nil
	} else if (len(*args) == 2) && (!slices.Contains(commands.CommandsNeed3Args, cmd)) {
		return nil
	} else if len(*args) == 4 && (slices.Contains(commands.CommandsNeed4Args, cmd)) {
		  return nil
	}
	return fmt.Errorf("too many or not enough arguments!!!!! use --help")
}
