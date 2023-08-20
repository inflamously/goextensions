package args

/*
This batch contains multiple Commands which accept args to be parsed
*/
type SimpleCommandBatch struct {
	Commands []*SimpleCommand
}

func (scb *SimpleCommandBatch) Parse(args []string) bool {
	for _, command := range scb.Commands {
		if command.Parse(args) {
			return true
		}
	}

	return false
}
