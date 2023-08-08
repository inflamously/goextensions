package args

import (
	"fmt"
	"testing"
)

func TestCommand_Parse(t *testing.T) {
	args := []string{
		"tags",
		"random",
		"10",
		"x",
		"--test",
		"123",
	}

	command := NewCommand(
		&SimpleCommand{
			Name: "tags",
			Function: func(args []SimpleArgument, options []SimpleOption) {
				fmt.Printf("Running 'tags' command with args: %s \n", args)
			},
		}).
		Option(
			"--test",
			[]*SimpleArgument{
				{}, {},
			},
		).
		Subcommand(&SimpleCommand{
			Name: "random",
			Function: func(args []SimpleArgument, options []SimpleOption) {
				fmt.Printf("Running 'random' command with args: %s \n", args)
			},
		}).
		Argument(&SimpleArgument{}).
		Root()
	command.Parse(args)
}

func TestCommand_ParseTags(t *testing.T) {
	args := []string{
		"tags",
		"random",
		"10",
		"x",
		"--test",
		"123",
	}

	command := NewCommand(
		&SimpleCommand{
			Name: "tags",
			Function: func(args []SimpleArgument, options []SimpleOption) {
				fmt.Printf("Running 'tags' command with args: %s \n", args)
			},
		}).
		Option(
			"--test",
			[]*SimpleArgument{
				{}, {},
			}).
		Subcommand(&SimpleCommand{
			Name: "random",
			Function: func(args []SimpleArgument, options []SimpleOption) {
				fmt.Printf("Running 'random' command with args: %s \n", args)
			},
		}).
		Root()

	for i := len(args); i > 0; i-- {
		command.Parse(args[:i])
	}
}

func TestCommand_ParseOptions(t *testing.T) {
	args := []string{
		"tags",
		"random",
		"--abc",
		"123",
	}

	command := NewCommand(
		&SimpleCommand{
			Name: "tags",
		}).
		Option(
			"--test",
			[]*SimpleArgument{
				{}, {},
			},
		).
		Option(
			"--x",
			nil,
		).
		Subcommand(&SimpleCommand{
			Name: "random",
		}).
		Option(
			"--abc",
			[]*SimpleArgument{
				{},
			},
		).
		Root()

	command.Parse(args)
}
