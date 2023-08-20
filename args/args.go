package args

import (
	"fmt"
	collections "github.com/inflamously/goextensions/collections"
	"log"
)

type SimpleCommandFunc func(args []SimpleArgument, options []SimpleOption)

type SimpleCommand struct {
	Name       string
	Function   SimpleCommandFunc
	subcommand *SimpleCommand
	arguments  []*SimpleArgument
	options    []*SimpleOption
	parent     *SimpleCommand
}

func (c *SimpleCommand) Root() *SimpleCommand {
	var parent = c.parent
	for {
		if parent.parent != nil {
			parent = parent.parent
		} else {
			return parent
		}
	}
}

func (c *SimpleCommand) Subcommand(command *SimpleCommand) *SimpleCommand {
	if c.arguments != nil {
		log.Panicf("Cannot create subcommand '%s' of '%s' in combination with arguments.", command.Name, c.Name)
	}
	c.subcommand = command
	c.subcommand.parent = c
	return command
}

func (c *SimpleCommand) Argument(argument *SimpleArgument) *SimpleCommand {
	if c.subcommand != nil {
		log.Panicf("Cannot add arguments in command '%s in combination with subcommand of '%s'", c.Name, c.subcommand.Name)
	}
	c.arguments = append(c.arguments, argument)
	return c
}

func (c *SimpleCommand) Option(name string, arguments []*SimpleArgument) *SimpleCommand {
	if c.options == nil {
		c.options = []*SimpleOption{}
	}
	if IsOption(name) {
		panic("Options must start with '--'")
	}
	option := SimpleOption{
		Name:      name,
		Arguments: arguments,
	}
	c.options = append(c.options, &option)
	return c
}

/*
Executes function if exist on SimpleCommand
*/
func (c *SimpleCommand) Execute() {
	if c.Function != nil {
		var clonedArguments = make([]SimpleArgument, 0)
		var clonedOptions = make([]SimpleOption, 0)

		for _, arg := range c.arguments {
			if arg.value == "" {
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

/*
Parse string array of arguments into SimpleCommand structure
*/
func (c *SimpleCommand) Parse(args []string) bool {
	mutArgs := args

	// Check if command matches else ignore
	if c.Name != mutArgs[0] {
		return false
	}
	mutArgs = mutArgs[1:]

	if len(mutArgs) <= 0 {
		c.Execute()
		return true
	}

	// Check if subcommand match exists, switch to subcommand else ignore
	if len(mutArgs) > 0 && c.subcommand != nil && c.subcommand.Name == mutArgs[0] {
		c.subcommand.Parse(mutArgs)
		c.Execute()
		return true
	}

	mutArgs = c.parseArguments(mutArgs)
	mutArgs = c.parseOptions(mutArgs)

	if len(mutArgs) > 0 {
		log.Printf("\nToo many parameters passed into command '%s' -> %s\n", c.Name, mutArgs)
	}

	c.Execute()

	return true
}

func (c *SimpleCommand) parseArguments(mutArgs []string) []string {
	if c.arguments != nil && len(c.arguments) > 0 {
		argsCount := len(c.arguments)
		if argsCount > len(mutArgs) {
			panic("Missing command arguments")
		}
		parseArguments := mutArgs[:argsCount]
		for argIndex, arg := range c.arguments {
			arg.value = parseArguments[argIndex]
		}
	}
	mutArgs = mutArgs[len(c.arguments):]
	return mutArgs
}

func (c *SimpleCommand) parseOptions(mutArgs []string) []string {
	var removeIndexes []int
	if c.options != nil {
		for _, option := range c.options {
			for argIndex, arg := range mutArgs {
				if IsOption(arg) {
					continue
				}
				if arg == option.Name {
					option.parsed = true
					removeIndexes = append(removeIndexes, argIndex)
					for optionIndex, optionArg := range option.Arguments {
						fmt.Printf("ArgIndex %d, OptionIndex %d, OptionArg %s, \n", argIndex, optionIndex, optionArg)
						optionArgIndex := 1 + argIndex + optionIndex
						if len(mutArgs) > optionArgIndex {
							optionArg.value = mutArgs[optionArgIndex]
							removeIndexes = append(removeIndexes, optionArgIndex)
						} else {
							log.Panicf("Not enough arguments for option '%s' in command '%s'", option.Name, c.Name)
						}
					}
				}
			}
		}

		min, max := collections.MinMaxInt(removeIndexes)
		mutArgs = collections.SliceOutwards[string](mutArgs, min, max)
	}

	return mutArgs
}
