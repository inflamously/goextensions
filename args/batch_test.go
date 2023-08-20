package args

import (
	"log"
	"testing"
)

func TestSimpleCommandBatch_Parse(t *testing.T) {
	scb := SimpleCommandBatch{
		Commands: []*SimpleCommand{
			CreateCommand(
				&SimpleCommand{
					Name: "version",
				},
			).RootCommand(),
		},
	}

	testArgs := []string{
		"version",
	}

	scb.Parse(testArgs)
}

func TestSimpleCommandBatch_ParseMultiCommand(t *testing.T) {
	scb := SimpleCommandBatch{
		Commands: []*SimpleCommand{
			CreateCommand(
				&SimpleCommand{
					Name: "version",
				},
			).RootCommand(),
			CreateCommand(
				&SimpleCommand{
					Name: "expectus",
					Function: func(args []SimpleArgument, options []SimpleOption) {
						log.Panicf("Expectus called!")
					},
				}),
		},
	}

	testArgs := []string{
		"version",
	}

	scb.Parse(testArgs)
}
