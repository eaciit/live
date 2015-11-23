package live

import (
	"fmt"
	"github.com/eaciit/toolkit"
	"net/http"
	"regexp"
	"strings"
)

type CommandTypeEnum int
type ValidationTypeEnum int
type RESTAuthTypeEnum int

const (
	// inisialisation constanta for type execution
	CommandType_Local CommandTypeEnum = 1
	CommandType_SSH   CommandTypeEnum = 2
	CommandType_REST  CommandTypeEnum = 3

	ValidationType_Contain ValidationTypeEnum = 1
	ValidationType_Equal   ValidationTypeEnum = 2
	ValidationType_Regex   ValidationTypeEnum = 10

	RESTAuthType_None  RESTAuthTypeEnum = 1
	RESTAuthType_Basic RESTAuthTypeEnum = 2
)

type Command struct {
	Type CommandTypeEnum

	// these attributes are used for local command
	CommandText  string
	CommandParms []string

	// these attributes are used for REST
	RESTUrl      string
	RESTMethod   string
	RESTUser     string
	RESTPassword string
	RESTAuthType RESTAuthTypeEnum

	// these attributes used for SSH
	SshClient *SshSetting

	// these attributes used for validation after command
	ValidationType  ValidationTypeEnum
	ValidationValue string
}

/*
Validation message after command execute
*/
func (c *Command) Validate(res string) error {
	if c.ValidationType == ValidationType_Equal {
		if res != c.ValidationValue {
			return fmt.Errorf("Response is not valid. Expecting for %s", c.ValidationValue)
		}
	} else if c.ValidationType == ValidationType_Contain {
		if !strings.Contains(res, c.ValidationValue) {
			return fmt.Errorf("Phrase %s could not be found on response", c.ValidationValue)
		}
	} else if c.ValidationType == ValidationType_Regex {
		match, _ := regexp.MatchString(c.ValidationValue, res)
		if !match {
			return fmt.Errorf("Response is not valid. Not match with pattern %s", c.ValidationValue)
		}
	}

	return nil
}

/*
Execute command depend on type that has been declared before
*/
func (c *Command) Exec() (string, error) {
	var (
		res     string
		e       error
		httpRes *http.Response
	)
	res = "initial"

	e = fmt.Errorf("Command %s %s can't be executed. No valid implementation can be found")

	if c.Type == CommandType_Local {
		ps := []string{}
		if c.CommandParms != nil {
			ps = c.CommandParms
		}

		res, e = toolkit.RunCommand(c.CommandText, ps...)

	} else if c.Type == CommandType_SSH {

		ps := []string{c.CommandText}
		res, e = c.SshClient.RunCommandSsh(ps...)

	} else if c.Type == CommandType_REST {
		if c.RESTAuthType == RESTAuthType_None {
			httpRes, e = toolkit.HttpCall(c.RESTUrl, c.RESTMethod, nil, nil)
		} else if c.RESTAuthType == RESTAuthType_Basic {
			var config = map[string]interface{}{"auth": "basic", "user": c.RESTUser, "password": c.RESTPassword}
			httpRes, e = toolkit.HttpCall(c.RESTUrl, c.RESTMethod, nil, config)
			// httpRes, e = toolkit.HttpCall(c.RESTUrl, c.RESTMethod, nil, true, c.RESTUser, c.RESTPassword)
		}
		res = toolkit.HttpContentString(httpRes)
	}
	return res, e
}
