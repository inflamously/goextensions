package args

import (
	collections "github.com/inflamously/goextensions/collections"
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

func (c *SimpleCommand) Parent() *SimpleCommand {
	return c.parent
}

func (c *SimpleCommand) RootCommand() *SimpleCommand {
	if c.parent == nil {
		return c
	}
	var parent = c.parent
	for {
		if parent.parent != nil {
			parent = parent.parent
		} else {
			return parent
		}
	}
}

/*
Adds new subcommand which can be called
*/
//TODO: Support multiple subcommands?
func (c *SimpleCommand) Subcommand(command *SimpleCommand) *SimpleCommand {
	if c.arguments != nil {
		log.Panicf("Cannot create subcommand '%s' of '%s' in combination with arguments.", command.Name, c.Name)
	}
	command.parent = c
	c.subcommands = append(c.subcommands, command)
	return command
}

/*
Adds new argument to command which is required to be entered
*/
func (c *SimpleCommand) Argument(argument *SimpleArgument) *SimpleCommand {
	if len(c.subcommands) > 0 {
		log.Panicf("Cannot add arguments in command '%s if subcommands exists", c.Name)
	}
	c.arguments = append(c.arguments, argument)
	return c
}

/*
Adds new option to command e.g. "--plain"
*/
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

/*
Parse string array of arguments into SimpleCommand structure
*/
func (c *SimpleCommand) Parse(args []string) bool {
	mutArgs := args

	if mutArgs == nil || len(mutArgs) <= 0 {
		log.Panicf("No arguments passed into command '%s'", c.Name)
	}

	// Check if command matches else ignore
	if c.Name != mutArgs[0] {
		c.Help()
		return false
	}
	mutArgs = mutArgs[1:]

	// Check if subcommand match exists, switch to subcommand else ignore
	c.parseSubcommand(mutArgs)

	// Parse further args
	mutArgs = c.parseArguments(mutArgs)
	mutArgs = c.parseOptions(mutArgs)

	if len(mutArgs) > 0 {
		log.Panicf("Too many parameters passed into command '%s' -> %s\n", c.Name, mutArgs)
	}

	c.Execute()

	return true
}

func (c *SimpleCommand) Help() {

}

func (c *SimpleCommand) parseSubcommand(mutArgs []string) {
	for _, command := range c.subcommands {
		if len(mutArgs) > 0 && command.Name == mutArgs[0] {
			command.Parse(mutArgs)

			//TODO: Should we execute command after subcommand parse?
			c.Execute()
		}
	}
}

func (c *SimpleCommand) parseArguments(mutArgs []string) []string {
	if c.arguments != nil && len(c.arguments) > 0 {
		argsCount := len(c.arguments)
		if argsCount > len(mutArgs) {
			log.Panicf("Command '%s' is missing '%d' arguments", c.Name, argsCount)
		}
		parseArguments := mutArgs[:argsCount]
		for argIndex, arg := range c.arguments {
			arg.Value = parseArguments[argIndex]
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
						optionArgIndex := 1 + argIndex + optionIndex
						if len(mutArgs) > optionArgIndex {
							optionArg.Value = mutArgs[optionArgIndex]
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
