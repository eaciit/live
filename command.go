package live

type Command struct {
	Type    string
	Command string
}

func (c *Command) Exec() error {
	return nil
}
