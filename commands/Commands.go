package commands

type Command uint8

const CommandIndex = 1

const (
	Add          Command = 1
	List         Command = 2
	Update       Command = 3
	Delete       Command = 4
	MarkProgress Command = 5
	MarkDone     Command = 6
	Help         Command = 7

	UnknownCommand Command = 0
)

var CommandsNeedsThreeArgs = []Command{
	Add, Update, Delete, MarkProgress, MarkDone,
}

var ComandsMap = map[string]Command{
	"add":           Add,
	"list":          List,
	"update":        Update,
	"delete":        Delete,
	"mark-progress": MarkProgress,
	"mark-done":     MarkDone,
	"--help":        Help,
}
