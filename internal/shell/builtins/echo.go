package builtins

import (
	"fmt"
	"strings"

	"github.com/Yashtripathi6108/go-shell/internal/shell/commands"
)

type EchoCommand struct {
	commands.BaseCommand
	Args []string
}

func (c *EchoCommand) Execute() error {
	fmt.Fprintln(c.Stdout, strings.Join(c.Args, " "))
	return nil
}
