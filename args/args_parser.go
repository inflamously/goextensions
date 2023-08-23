package args

import (
	"github.com/inflamously/goextensions/collections"
	"log"
)

/*
Checks subcommands and if matches, parses args further.
*/
func (c *SimpleCommand) parseSubcommand(mutArgs []string) bool {
	for _, command := range c.subcommands {
		if len(mutArgs) > 0 && command.Name == mutArgs[0] {
			command.Parse(mutArgs)
			return true
		}
	}

	return false
}

func (c *SimpleCommand) parseArguments(mutArgs []string) ([]string, bool) {
	if c.arguments != nil && len(c.arguments) > 0 {
		argsCount := len(c.arguments)
		if argsCount > len(mutArgs) {
			log.Printf("Command '%s' is missing '%d' arguments", c.Name, argsCount)
			return mutArgs, false
		}
		parseArguments := mutArgs[:argsCount]
		for argIndex, arg := range c.arguments {
			arg.Value = parseArguments[argIndex]
		}
	}
	mutArgs = mutArgs[len(c.arguments):]
	return mutArgs, true
}

func (c *SimpleCommand) parseOptions(mutArgs []string) ([]string, bool) {
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
							log.Printf("Not enough arguments for option '%s' in command '%s'", option.Name, c.Name)
							return mutArgs, false
						}
					}
				}
			}
		}

		if removeIndexes != nil {
			min, max := collections.MinMaxInt(removeIndexes)
			mutArgs = collections.SliceOutwards[string](mutArgs, min, max)
		}
	}

	return mutArgs, true
}
