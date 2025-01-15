package parser

import (
	"strings"

	"github.com/Yashtripathi6108/Shell/internal/shell/commands"
	"github.com/Yashtripathi6108/Shell/internal/shell/registry"
)

type CommandParser struct {
	factory *registry.CommandFactory
}

func NewCommandParser(factory *registry.CommandFactory) *CommandParser {
	return &CommandParser{factory: factory}
}

func (p *CommandParser) Parse(input string) (commands.Command, error) {
	parts := strings.Fields(input)
	if len(parts) == 0 {
		return nil, nil
	}

	cmdName := parts[0]
	cmdArgs := parts[1:]

	cmd, err := p.factory.Create(cmdName, cmdArgs)
	if err != nil {
		return nil, err
	}

	return cmd, nil
}
