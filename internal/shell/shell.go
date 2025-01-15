package shell

import (
	"bufio"
	"os"
	"strings"

	"github.com/Yashtripathi6108/Shell/internal/shell/registry"

	"github.com/Yashtripathi6108/Shell/internal/shell/parser"

	"github.com/fatih/color"
)

type Shell struct {
	parser *parser.CommandParser
}

func NewShell() *Shell {
	factory := registry.NewCommandFactory(os.Stdout)
	registry.RegisterBuiltins(factory)
	parser := parser.NewCommandParser(factory)
	return &Shell{
		parser: parser,
	}
}

// Run starts the shell and applies custom styling to the prompt and messages
func (sh *Shell) Run() {
	reader := bufio.NewReader(os.Stdin)

	// Define colors for styling
	promptColor := color.New(color.FgCyan).Add(color.Underline) // Cyan with underline for prompt
	successColor := color.New(color.FgGreen)                    // Green for success messages
	errorColor := color.New(color.FgRed)                        // Red for error messages
	infoColor := color.New(color.FgYellow)                      // Yellow for info messages

	// Display a welcome message
	infoColor.Println("Welcome to the Go Shell!")
	infoColor.Println("Type 'exit' to quit the shell.")

	for {
		// Print prompt with styling
		promptColor.Print("$ ")

		// Read input from user
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// Exit condition
		if input == "exit" {
			successColor.Println("Exiting the shell. Goodbye!")
			break
		}

		// Parse the command
		cmd, err := sh.parser.Parse(input)
		if err != nil {
			errorColor.Println("Error parsing command:", err)
			continue
		}

		// Skip empty input
		if cmd == nil {
			continue
		}

		// Execute the command and handle errors
		if err := cmd.Execute(); err != nil {
			errorColor.Println("Error executing command:", err)
		} else {
			successColor.Println("Command executed successfully!")
		}
	}
}
