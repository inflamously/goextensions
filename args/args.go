package args

import (
	"errors"
	"log"
)

type SimpleCommandFunc func(args []SimpleArgument, options []SimpleOption)

type SimpleCommand struct {
	Name        string
	Function    SimpleCommandFunc
	subcommands []*SimpleCommand
	arguments   []*SimpleArgument
	options     []*SimpleOption
	parent      *SimpleCommand
}

func (c *SimpleCommand) verify() error {
	if c.Function != nil && len(c.subcommands) > 0 {
		return errors.New("Cannot mix Function with Subcommand")
	}

	return nil
}

/*
Executes function if exist on command
*/
func (c *SimpleCommand) Execute() {
	if c.Function != nil {
		var clonedArguments = make([]SimpleArgument, 0)
		var clonedOptions = make([]SimpleOption, 0)

		for _, arg := range c.arguments {
			if arg.Value == "" {
				continue
			}
			clonedArguments = append(clonedArguments, *arg)
		}

		for _, option := range c.options {
			newOption := *option
			newOption.Arguments = []*SimpleArgument{}

			for _, arg := range option.Arguments {
				newArg := *arg // Deref to copy (create new) and ref pointer to new arg item in new array
				newOption.Arguments = append(newOption.Arguments, &newArg)
			}

			clonedOptions = append(clonedOptions, newOption)
		}

		(c.Function)(clonedArguments, clonedOptions)
	}
}

func (c *SimpleCommand) commandTreeString() string {
	commandParent := c
	commands := make([]string, 0)
	for {
		if commandParent == nil {
			break
		}
		commands = append(commands, commandParent.Name)
		commandParent = commandParent.parent
	}

	return ""
}

func (c *SimpleCommand) Help(args []string) {
	if len(args) > 0 {
		log.Printf("Command with args \"%s\" not found.\n", args)
	} else {
		log.Println("No command provided.")
	}

	log.Printf("Available subcommands for \"%s\":\n", c.commandTreeString())
	for _, subcommand := range c.subcommands {
		log.Printf("*\t\"%s\"", subcommand.Name)
	}
}

/*
Parse given arguments and execute command
*/
func (c *SimpleCommand) Parse(args []string) bool {
	mutArgs := args

	if mutArgs == nil || len(mutArgs) <= 0 {
		c.Help(args)
		return false
	}

	// Check if command matches else ignore
	if c.Name != mutArgs[0] {
		c.Help(args)
		return false
	}
	mutArgs = mutArgs[1:]

	// Check if subcommand match exists, switch to subcommand else ignore
	if len(c.subcommands) > 0 && c.parseSubcommand(mutArgs) {
		return true
	}

	// Parse further args
	mutArgs = c.parseArguments(mutArgs)
	mutArgs = c.parseOptions(mutArgs)

	if len(mutArgs) > 0 {
		log.Panicf("Too many parameters passed into command '%s' -> %s\n", c.Name, mutArgs)
	}

	if c.Function != nil {
		log.Printf("Executing command '%s' with args '%s'", c.Name, args[1:])
		c.Execute()
	} else {
		c.Help(mutArgs)
	}

	return true
}
