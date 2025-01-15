package pipeline

import "github.com/Yashtripathi6108/Shell/internal/shell/commands"

type Pipeline struct {
	Commands []commands.Command
}

func (p *Pipeline) AddCommand(cmd commands.Command) {
	p.Commands = append(p.Commands, cmd)
}

func (p *Pipeline) Execute() error {
	composite := &commands.CompositeCommand{
		Commands: p.Commands,
	}
	return composite.Execute()
}
