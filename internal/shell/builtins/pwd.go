package builtins

import (
	"fmt"
	"os"

	"github.com/Yashtripathi6108/Shell/internal/shell/commands"
)

type PwdCommand struct {
	commands.BaseCommand
}

func (c *PwdCommand) Execute() error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	fmt.Fprintln(c.Stdout, dir)
	return nil
}
