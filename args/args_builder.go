package args

import "log"

/*
Returns parent command or nil
*/
func (c *SimpleCommand) Parent() *SimpleCommand {
	return c.parent
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
Returns root command of tree from a given subcommand
*/
func (c *SimpleCommand) RootCommand() *SimpleCommand {
	if err := c.verify(); err != nil {
		log.Panicf("Command failed due to '%s'", err)
	}

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
