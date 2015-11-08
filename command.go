package live

import (
	"fmt"
	"github.com/eaciit/toolkit"
)

type CommandTypeEnum int
type ValidationTypeEnum int

const (
	CommandType_Local CommandTypeEnum = 1
	CommandType_SSH   CommandTypeEnum = 2
	CommandType_REST  CommandTypeEnum = 3

	ValidationType_Contain ValidationTypeEnum = 1
	ValidationType_Equal   ValidationTypeEnum = 2
	ValidationType_Regex   ValidationTypeEnum = 10
)

type Command struct {
	Type         CommandTypeEnum
	CommandText  string
	CommandParms []string

	ValidationType  ValidationTypeEnum
	ValidationValue string
}

func (c *Command) Exec() error {
	var (
		_ string
		e error
	)
	e = fmt.Errorf("Command %s %s can't be executed. No valid implementation can be found")
	if c.Type == CommandType_Local {
		ps := []string{}
		if c.CommandParms != nil {
			ps = c.CommandParms
		}
		_, e = toolkit.RunCommand(c.CommandText, ps...)

		//--- validation should be here
	}
	return e
}
