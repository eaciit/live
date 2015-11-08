package live

import (
	"fmt"
	"github.com/eaciit/toolkit"
	//"io/ioutil"
	//"net/http"
	"strings"
)

type PingTypeEnum int
type HttpBodyEnum int

const (
	PingType_Network PingTypeEnum = iota
	PingType_HttpStatus
	PingType_HttpBody
	PingType_Command
	PingType_Custom

	HttpBody_Contains HttpBodyEnum = iota
	HttpBody_Equals
)

type Ping struct {
	Type       PingTypeEnum
	User       string
	Password   string
	Host       string
	Command    string
	LastStatus string

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
		e = p.checkHttpStatus()
	} else if pingType == PingType_Command {
		e = p.checkCommand()
	} else if pingType == PingType_Custom {
		e = p.checkCustom()
	}
	return e
}

func (p *Ping) checkNetwork() error {
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
	return nil
}

func (p *Ping) checkCustom() error {
	if p.FnPing == nil {
		return nil
	}

	return p.FnPing(p)
}
