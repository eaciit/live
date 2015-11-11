package live

import (
	"fmt"
	"github.com/eaciit/toolkit"
	//"io/ioutil"
	//"net/http"
	"net"
	"regexp"
	"strings"
)

type PingTypeEnum int

const (
	PingType_Network PingTypeEnum = iota
	PingType_HttpStatus
	PingType_HttpBody
	PingType_Command
	PingType_Custom
)

type HttpBodyEnum int

const (
	HttpBody_Contains HttpBodyEnum = iota
	HttpBody_Equals
)

type ResponseEnum int

const (
	Response_Contains ResponseEnum = iota
	Response_Equals
	Response_RegEx
)

type Ping struct {
	Type       PingTypeEnum
	User       string
	Password   string
	Host       string
	LastStatus string

	//--- these attributes are used for command check
	Command       string
	CommandParms  []string
	ResponseType  ResponseEnum
	ResponseValue string

	HttpBodyType   HttpBodyEnum
	HttpBodySearch string

	FnPing func(*Ping) error
}

func (p *Ping) Check() error {
	var e error
	pingType := p.Type
	if pingType == PingType_Network {
		e = p.checkNetwork()
	} else if pingType == PingType_HttpStatus {
		e = p.checkHttpStatus()
	} else if pingType == PingType_HttpBody {
		e = p.checkHttpBody()
	} else if pingType == PingType_Command {
		e = p.checkCommand()
	} else if pingType == PingType_Custom {
		e = p.checkCustom()
	}
	return e
}

func (p *Ping) checkNetwork() error {
	_, e := net.Dial("tcp", p.Host)
	if e != nil {
		return fmt.Errorf("Unable to access %s, %s", p.Host, e.Error())
	}

	return nil
}

func (p *Ping) checkHttpStatus() error {
	r, e := toolkit.HttpCall(p.Host, "GET", nil, false, "", "")
	if e != nil {
		return fmt.Errorf("Unable to access %s, %s", p.Host, e.Error())
	}
	if r.StatusCode != 200 {
		return fmt.Errorf("Unable to access %s, code: %d status: %s", p.Host, r.StatusCode, r.Status)
	}
	return nil
}

func (p *Ping) checkHttpBody() error {
	r, e := toolkit.HttpCall(p.Host, "GET", nil, false, "", "")
	if e != nil {
		return e
	}
	body := toolkit.HttpContentString(r)
	if p.HttpBodyType == HttpBody_Contains {
		if !strings.Contains(body, p.HttpBodySearch) {
			return fmt.Errorf("Phrase %s could not be found on response body", p.HttpBodySearch)
		}
	} else if p.HttpBodyType == HttpBody_Equals {
		if body != p.HttpBodySearch {
			return fmt.Errorf("Response is not valid. Expecting for %s", p.HttpBodySearch)
		}
	} else {
		return fmt.Errorf("Invalid parameter")
	}
	return nil
}

func (p *Ping) checkCommand() error {

	ps := []string{}

	if p.CommandParms != nil {
		ps = p.CommandParms
	}
	res, e := toolkit.RunCommand(p.Command, ps...)

	if e != nil {
		return e
	}

	if p.ResponseType == Response_Equals {
		if res != p.ResponseValue {
			return fmt.Errorf("Response is not valid. Expecting for %s", p.ResponseValue)
		}
	} else if p.ResponseType == Response_Contains {
		if !strings.Contains(res, p.ResponseValue) {
			return fmt.Errorf("Phrase %s could not be found on response", p.ResponseValue)
		}
	} else if p.ResponseType == Response_RegEx {
		match, _ := regexp.MatchString(p.ResponseValue, res)
		if !match {
			return fmt.Errorf("Response is not valid. Not match with pattern %s", p.ResponseValue)
		}
	}
	return nil
}

func (p *Ping) checkCustom() error {
	if p.FnPing == nil {
		return nil
	}

	return p.FnPing(p)
}
