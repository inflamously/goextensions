package args

import "strings"

/*
Checks if given string (arg) contains -- as it marks it as an option
*/
func IsOption(arg string) bool {
	return !strings.HasPrefix(arg, "--")
}

/*
Create a new command
*/
func NewCommand(command *SimpleCommand) *SimpleCommand {
	return command
}
