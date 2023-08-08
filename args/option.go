package args

type SimpleOption struct {
	Name      string
	Arguments []*SimpleArgument
	parsed    bool
}
