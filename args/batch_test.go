package args

import (
	"github.com/inflamously/goextensions/tests"
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

func TestSimpleCommandBatch_ParseNoArgs(t *testing.T) {
	scb := SimpleCommandBatch{
		Commands: []*SimpleCommand{
			CreateCommand(&SimpleCommand{
				Name: "version",
				Function: func(args []SimpleArgument, options []SimpleOption) {

				},
			}),
		},
	}
	defer tests.StopPanic("No arguments passed but command parse was still called!")
	scb.Parse([]string{})
}

func TestSimpleCommandBatch_ParseNoCommand(t *testing.T) {
	scb := SimpleCommandBatch{}
	defer tests.StopPanic("Empty command list but did not panic!")
	scb.Parse([]string{"test"})
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
