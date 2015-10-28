package live

type Ping struct {
	Type     string
	User     string
	Password string
	Host     string
	Command  string

	LastStatus string
}

func (p *Ping) Check() error {
	return nil
}
