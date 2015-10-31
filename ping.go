package live

import (
	"fmt"
	"ioutil"
	"net/http"
)

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

func httpCall(r *http.Request) (*http.Response, error) {
	client := new(http.Client)
	resp, err := client.Do(&req)
	return resp, err
}

func (p *Ping) httpStatus() error {
	rq := http.NewRequest("GET", p.Host, nil)
	rs, err := httpCall(r)
	if err != nil {
		return fmt.Errorf("PINGCONNECTFAIL: unable to access %s, %s", p.Host, err.Error())
	}

	if rs.StatusCode != 200 {
		return fmt.Errorf("PINGCONNECTFAIL: unable to access %s, returned code is %d", p.Host, rs.StatusCode)
	}
}

func (p *Ping) httpBody() {
	rq := http.NewRequest("GET", p.Host, nil)
	rs, err := httpCall(r)
	if err != nil {
		return fmt.Errorf("PINGCONNECTFAIL: unable to access %s, %s", p.Host, err.Error())
	}
	defer rs.Close()

}
