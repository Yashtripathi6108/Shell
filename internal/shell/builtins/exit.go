package builtins

import (
	"os"

	"github.com/Yashtripathi6108/Shell/internal/shell/commands"
)

type ExitCommand struct {
	commands.BaseCommand
	Args []string
}

func (c *ExitCommand) Execute() error {
	os.Exit(0)
	return nil
}
