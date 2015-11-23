package live

import (
	"fmt"
	"github.com/eaciit/toolkit"
	"net"
	"regexp"
	"strings"
)

type PingTypeEnum int

const (
	// these variable used for determine ping to execute
	PingType_Network PingTypeEnum = iota
	PingType_HttpStatus
	PingType_HttpBody
	PingType_Command
	PingType_Custom
)

type HttpBodyEnum int

const (
	// these variable used for validation after ping http body
	HttpBody_Contains HttpBodyEnum = iota
	HttpBody_Equals
)

type ResponseEnum int

const (
	// these variable used for validation response after ping http rest
	Response_Contains ResponseEnum = iota
	Response_Equals
	Response_RegEx
)

type Ping struct {
	// these attributes are used for initialization ping such us server location, and ping type
	Type       PingTypeEnum
	User       string
	Password   string
	Host       string
	LastStatus string

	// these attributes are used for command check
	Command      string
	CommandParms []string

	// these attributes are used for validate response after ping
	ResponseType  ResponseEnum
	ResponseValue string

	// these attributes are used for validate httpbody after ping
	HttpBodyType   HttpBodyEnum
	HttpBodySearch string

	FnPing func(*Ping) error
}

/*
Select ping type before execute
*/
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

/*
Ping service over network tcp status
*/
func (p *Ping) checkNetwork() error {
	_, e := net.Dial("tcp", p.Host)
	p.LastStatus = "Fail"
	if e != nil {
		return fmt.Errorf("Unable to access %s, %s", p.Host, e.Error())
	}
	p.LastStatus = "OK"
	return nil
}

/*
Ping service over http status
*/
func (p *Ping) checkHttpStatus() error {
	r, e := toolkit.HttpCall(p.Host, "GET", nil, false, "", "")
	p.LastStatus = "Fail"
	if e != nil {
		return fmt.Errorf("Unable to access %s, %s", p.Host, e.Error())
	}
	if r.StatusCode != 200 {
		return fmt.Errorf("Unable to access %s, code: %d status: %s", p.Host, r.StatusCode, r.Status)
	}
	p.LastStatus = "OK"
	return nil
}

/*
Ping service over http body and validate response
*/
func (p *Ping) checkHttpBody() error {
	p.LastStatus = "Fail"
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
	p.LastStatus = "OK"
	return nil
}

/*
Ping service over local command execute
*/
func (p *Ping) checkCommand() error {

	ps := []string{}

	if p.CommandParms != nil {
		ps = p.CommandParms
	}
	res, e := toolkit.RunCommand(p.Command, ps...)
	p.LastStatus = "Fail"
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
	p.LastStatus = "OK"
	return nil
}

/*
NEXT Feature, ping with custom method
*/
func (p *Ping) checkCustom() error {
	if p.FnPing == nil {
		return nil
	}

	return p.FnPing(p)
}
