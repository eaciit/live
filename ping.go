package live

import (
	"fmt"
	"github.com/eaciit/toolkit"
	"ioutil"
	"net/http"
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

type Ping struct {
	Type       PingTypeEnum
	User       string
	Password   string
	Host       string
	Command    string
	LastStatus string

	HttpBodyType   HttpBodyEnum
	HttpBodySearch string

	FnPing func(*ping) error
}

func (p *Ping) Check() error {
	pingType := p.Type
	if pingType == PingType_Network {
		return p.checkNetwork()
	} else if pingType == PingType_HttpStatus {
		return p.checkHttpStatus()
	} else if pingType == PingType_HttpBody {
		return p.checkHttpStatus()
	} else if pingType == PingType_Command {
		return p.checkCommand()
	} else if pingType == PingType_Custom {
		return p.checkCustom(p.FnPing)
	}
	return nil
}

func (p *Ping) checkNetwork() error {
	return nil
}

func (p *Ping) checkHttpStatus() error {
	r, e := toolkit.HttpCall(p.Host, "GET", nil, false, "", "")
	if e != nil {
		return e
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
	} else if p.HttpBodySearch == HttpBody_Equals {
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

func (p *Ping) checkCustom(fn func(*ping) error) {
	if p.FnPing == nil {
		return nil
	}

	return p.FnPing(p)
}
