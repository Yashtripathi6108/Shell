package builtins

import (
	"fmt"

	"github.com/Yashtripathi6108/Shell/internal/shell/commands"
)

type TypeCommand struct {
	commands.BaseCommand
	Args []string
}

func (c *TypeCommand) Execute() error {
	if len(c.Args) == 0 {
		return fmt.Errorf("type: missing argument")
	}
	fmt.Fprintf(c.Stdout, "%s is a builtin command\n", c.Args[0])
	return nil
}
