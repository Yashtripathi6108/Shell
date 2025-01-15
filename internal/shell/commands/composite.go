package commands

import (
	"io"
)

type CompositeCommand struct {
	Commands []Command // Changed 'commands' to 'Commands' to make it exported
}

func (c *CompositeCommand) Add(cmd Command) {
	c.Commands = append(c.Commands, cmd) // Updated to use 'Commands'
}

func (c *CompositeCommand) Execute() error {
	for _, cmd := range c.Commands { // Updated to use 'Commands'
		if err := cmd.Execute(); err != nil {
			return err
		}
	}
	return nil
}

func (c *CompositeCommand) SetStdin(r io.Reader) {
	for _, cmd := range c.Commands { // Updated to use 'Commands'
		cmd.SetStdin(r)
	}
}

func (c *CompositeCommand) SetStdout(w io.Writer) {
	for _, cmd := range c.Commands { // Updated to use 'Commands'
		cmd.SetStdout(w)
	}
}

func (c *CompositeCommand) StdinPipe() (io.WriteCloser, error) {
	return nil, nil
}

func (c *CompositeCommand) StdoutPipe() (io.ReadCloser, error) {
	return nil, nil
}
