package args

import (
	"fmt"
	"github.com/inflamously/goextensions/tests"
	"testing"
)

func TestCommand_ParseTooManyArguments(t *testing.T) {
	args := []string{
		"tags",
		"random",
		"10",
		"x",
		"--test",
		"123",
	}

	command := CreateCommand(
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
		RootCommand()

	defer tests.StopPanic("Command did not panic due to too many arguments")

	command.Parse(args)
}

func TestCommand_ParseSubcommand(t *testing.T) {
	args := []string{
		"tags",
		"random",
		"10",
	}

	command := CreateCommand(
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
		RootCommand()

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

	command := CreateCommand(
		&SimpleCommand{
			Name: "tags",
			Function: func(args []SimpleArgument, options []SimpleOption) {
				fmt.Printf("Running 'tags' command with args: %s \n", args)
			},
		},
	).Option(
		"--test",
		[]*SimpleArgument{
			{}, {},
		},
	).Subcommand(
		&SimpleCommand{
			Name: "random",
			Function: func(args []SimpleArgument, options []SimpleOption) {
				fmt.Printf("Running 'random' command with args: %s \n", args)
			},
		},
	).RootCommand()

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

	command := CreateCommand(
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
		RootCommand()

	command.Parse(args)
}

func TestCommand_Argument(t *testing.T) {
	args := []string{
		"tags",
		"random",
	}

	command := CreateCommand(
		&SimpleCommand{
			Name: "tags",
		},
	).Option(
		"--test",
		[]*SimpleArgument{
			{}, {},
		},
	).Option(
		"--x",
		nil,
	).Subcommand(
		&SimpleCommand{
			Name: "random",
		},
	).Argument(
		&SimpleArgument{},
	).RootCommand()

	defer tests.StopPanic("Command did not panic due to missing arguments")

	command.Parse(args)
}
