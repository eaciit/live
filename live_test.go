package live

import (
	"fmt"
	//"github.com/eaciit/toolkit"
	"testing"
	"time"
)

func TestMongo(t *testing.T) {
	fmt.Println("Test Check for MongoDb Service")

	svc := NewService()
	svc.Ping = func() *Ping {
		p := new(Ping)
		p.Type = PingType_HttpStatus
		p.Host = "http://localhost:27123"
		return p
	}()

	svc.RestartAfterNCritical = 3
	svc.Interval = 1 * time.Second
	svc.CommandStart = &Command{
		Type:         CommandType_Local,
		CommandText:  "sudo",
		CommandParms: []string{"mongod", "--config", "/data/mdb/3.0/service.conf", "--fork"},
	}

	svc.KeepAlive()

	if svc.MonitorStatus != "Running" {
		t.Errorf("Error, service status monitor check is %s \n", svc.MonitorStatus)
	}

	for svc.MonitorStatus == "Running" {
		time.Sleep(100 * time.Millisecond)
	}
}
