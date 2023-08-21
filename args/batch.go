package args

import "log"

/*
This batch contains multiple Commands which accept args to be parsed
*/
type SimpleCommandBatch struct {
	Commands []*SimpleCommand
}

func (scb *SimpleCommandBatch) Parse(args []string) bool {
	if len(scb.Commands) == 0 {
		log.Panicf("SimpleCommandBatch cannot be empty")
	}

	if len(args) == 0 {
		log.Printf("No command provided.")
		log.Printf("Available commands:")
		for _, c := range scb.Commands {
			log.Printf("*\t\"%s\"", c.Name)
		}
		return false
	}

	for _, c := range scb.Commands {
		if c.Parse(args) {
			return true
		}
	}

	return false
}
