package live

import (
	"fmt"
	"testing"
	"time"
)

func TestMongo(t *testing.T) {
	fmt.Println("Test Check for MongoDb Service")

	/*	svc := NewService()
		svc.Ping = func() *Ping {
			p := new(Ping)
			p.Type = PingType_HttpStatus
			p.Host = "http://localhost:27017"
			return p
		}()

		svc.RestartAfterNCritical = 3
		svc.Interval = 1 * time.Second
		svc.CommandStart = &Command{
			Type:         CommandType_Local,
			CommandText:  "sudo",
			CommandParms: []string{"mongod", "--config", "/data/mdb/3.0/service.conf", "--fork"},
		}*/
	/*
		svc := NewService()
		svc.Ping = func() *Ping {
			p := new(Ping)
			p.Type = PingType_HttpBody
			p.Host = "http://192.168.56.101:8080/status"
			p.HttpBodySearch = "RUNNING"
			p.HttpBodyType = HttpBody_Contains
			return p
		}()

		svc.RestartAfterNCritical = 3
		svc.Interval = 1 * time.Second
		svc.CommandStart = &Command{
			Type:            CommandType_REST,
			RESTUrl:         "http://192.168.56.101:8080/start",
			RESTMethod:      "GET", //POST,GET
			RESTUser:        "ALIP",
			RESTPassword:    "qwerty",
			RESTAuthType:    RESTAuthType_None,
			ValidationType:  ValidationType_Contain,
			ValidationValue: "SUCCESS",
		}*/

	svc := NewService()
	svc.Ping = func() *Ping {
		p := new(Ping)
		p.Type = PingType_HttpBody
		p.Host = "http://192.168.56.101:8080/status"
		p.HttpBodySearch = "RUNNING"
		p.HttpBodyType = HttpBody_Contains
		return p
	}()

	svc.RestartAfterNCritical = 3
	svc.Interval = 1 * time.Second
	svc.CommandStart = &Command{
		Type: CommandType_SSH,
		SshClient: &SshParm{
			SSHHost:     "192.168.56.101:22",
			SSHUser:     "alip",
			SSHPassword: "Bismillah",
			//SSHKeyLocation string
			SSHAuthType: SSHAuthType_Password,
		},
		CommandText:     "sudo service mysql start",
		ValidationType:  ValidationType_Contain,
		ValidationValue: "running",
	}
	/*	svc := NewService()
		svc.Ping = func() *Ping {
			p := new(Ping)
			p.Type = PingType_Command
			p.Host = "http://localhost:27017"
			p.Command = "cmd"
			p.CommandParms = []string{"/C", "sc", "query", "mongodb"}
			p.ResponseType = Response_Contains
			p.ResponseValue = "RUNNING"
			return p
		}()

		svc.RestartAfterNCritical = 3
		svc.Interval = 1 * time.Second
		svc.CommandStart = &Command{
			Type:            CommandType_Local,
			CommandText:     "cmd",
			CommandParms:    []string{"/C", "sc", "start", "mongodb"},
			ValidationType:  ValidationType_Contain,
			ValidationValue: "RUNNING",
		}*/

	svc.KeepAlive()

	if svc.MonitorStatus != "Running" {
		t.Errorf("Error, service status monitor check is %s \n", svc.MonitorStatus)
	}

	for svc.MonitorStatus == "Running" {
		time.Sleep(100 * time.Millisecond)
	}
}
