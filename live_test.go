package live

import (
	"fmt"
	"github.com/eaciit/toolkit"
	"testing"
	"time"
)

/*
	CommandType_Local CommandTypeEnum = 1
	CommandType_SSH   CommandTypeEnum = 2
	CommandType_REST  CommandTypeEnum = 3

	PingType_Network PingTypeEnum = iota
	PingType_HttpStatus
	PingType_HttpBody
	PingType_Command

	COMMAND
	ValidationType_Contain ValidationTypeEnum = 1
	ValidationType_Equal   ValidationTypeEnum = 2
	ValidationType_Regex   ValidationTypeEnum = 10



*/

func TestLocalCommandPingCommand(t *testing.T) {

	fmt.Println("[START](Ping Command|Exec LocalCommand)")
	//	fmt.Println("Test Check for MongoDb Service Local With Execution Command And Ping Command")
	var (
		err error
		i   int = 30
	)

	svc := NewService()
	svc.Ping = func() *Ping {
		p := new(Ping)
		p.Type = PingType_Command
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
	}

	svc.CommandStop = &Command{
		Type:            CommandType_Local,
		CommandText:     "cmd",
		CommandParms:    []string{"/C", "sc", "stop", "mongodb"},
		ValidationType:  ValidationType_Contain,
		ValidationValue: "STOP_PENDING",
	}

	svc.Log, err = toolkit.NewLog(false, true, "E:\\goproject\\LOG", "TestLocalCommandPingCommand", "20060102")

	if err != nil {
		t.Errorf("Error Start Log : %s", err.Error())
	}

	svc.KeepAlive()

	if svc.MonitorStatus != "Running" {
		t.Errorf("Error, service status monitor check is %s \n", svc.MonitorStatus)
	}

	for svc.MonitorStatus == "Running" {
		time.Sleep(2000 * time.Millisecond)
		i = i - 1
		if i < 0 {
			svc.StopMonitor()
		}

		if i == 10 {
			svc.CommandStop.Exec()
		}

		if svc.criticalFound == 0 {
			fmt.Print(".")
		} else {
			fmt.Print("S")
		}

	}
	fmt.Println("[END] \n\n")
}

func TestLocalCommandPingHttpStatus(t *testing.T) {

	fmt.Println("[START](Ping HttpStatus|Exec LocalCommand)")

	//	fmt.Println("Test Check for MongoDb Service Local With Execution Command And Ping HttpStatus")
	var (
		err error
		i   int = 30
	)

	svc := NewService()
	svc.Ping = func() *Ping {
		p := new(Ping)
		p.Type = PingType_HttpStatus
		p.Host = "http://localhost:27017"
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
	}

	svc.CommandStop = &Command{
		Type:            CommandType_Local,
		CommandText:     "cmd",
		CommandParms:    []string{"/C", "sc", "stop", "mongodb"},
		ValidationType:  ValidationType_Contain,
		ValidationValue: "STOP_PENDING",
	}

	svc.Log, err = toolkit.NewLog(false, true, "E:\\goproject\\LOG", "TestLocalCommandPingHttpStatus", "20060102")

	if err != nil {
		t.Errorf("Error Start Log : %s", err.Error())
	}

	svc.KeepAlive()

	if svc.MonitorStatus != "Running" {
		t.Errorf("Error, service status monitor check is %s \n", svc.MonitorStatus)
	}

	for svc.MonitorStatus == "Running" {
		time.Sleep(2000 * time.Millisecond)
		i = i - 1
		if i < 0 {
			svc.StopMonitor()
		}

		if i == 10 {

			svc.CommandStop.Exec()
		}

		if svc.criticalFound == 0 {
			fmt.Print(".")
		} else {
			fmt.Print("S")
		}

	}
	fmt.Println("[END] \n\n")
}

func TestLocalCommandPingHttpBody(t *testing.T) {

	fmt.Println("[START](Ping HttpBody|Exec LocalCommand)")

	//	fmt.Println("Test Check for MongoDb Service Local With Execution Command And Ping HttpBody")
	var (
		err error
		i   int = 30
	)

	svc := NewService()
	svc.Ping = func() *Ping {
		p := new(Ping)
		p.Type = PingType_HttpBody
		p.Host = "http://localhost:27017"
		p.HttpBodySearch = "access MongoDB over HTTP"
		p.HttpBodyType = HttpBody_Contains
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
	}

	svc.CommandStop = &Command{
		Type:            CommandType_Local,
		CommandText:     "cmd",
		CommandParms:    []string{"/C", "sc", "stop", "mongodb"},
		ValidationType:  ValidationType_Contain,
		ValidationValue: "STOP_PENDING",
	}

	svc.Log, err = toolkit.NewLog(false, true, "E:\\goproject\\LOG", "TestLocalCommandPingHttpBody", "20060102")

	if err != nil {
		t.Errorf("Error Start Log : %s", err.Error())
	}

	svc.KeepAlive()

	if svc.MonitorStatus != "Running" {
		t.Errorf("Error, service status monitor check is %s \n", svc.MonitorStatus)
	}

	for svc.MonitorStatus == "Running" {
		time.Sleep(2000 * time.Millisecond)
		i = i - 1
		if i < 0 {
			svc.StopMonitor()
		}

		if i == 10 {

			svc.CommandStop.Exec()
		}

		if svc.criticalFound == 0 {
			fmt.Print(".")
		} else {
			fmt.Print("S")
		}

	}

	fmt.Println("[END] \n\n")

}

func TestSshExecPingHttpBody(t *testing.T) {

	fmt.Println("[START](Ping HttpBody|Exec Ssh)")

	//	fmt.Println("Test Check for MySql Service Local With Execution Ssh And Ping HttpBody[REST]")
	var (
		err error
		i   int = 30
	)

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
			SSHAuthType: SSHAuthType_Password,
		},
		CommandText:     "sudo service mysql start",
		ValidationType:  ValidationType_Contain,
		ValidationValue: "running",
	}

	svc.CommandStop = &Command{
		Type: CommandType_SSH,
		SshClient: &SshParm{
			SSHHost:     "192.168.56.101:22",
			SSHUser:     "alip",
			SSHPassword: "Bismillah",
			SSHAuthType: SSHAuthType_Password,
		},
		CommandText:     "sudo service mysql stop",
		ValidationType:  ValidationType_Contain,
		ValidationValue: "running",
	}

	svc.Log, err = toolkit.NewLog(false, true, "E:\\goproject\\LOG", "TestSshExecPingHttpBody", "20060102")

	if err != nil {
		t.Errorf("Error Start Log : %s", err.Error())
	}

	svc.KeepAlive()

	if svc.MonitorStatus != "Running" {
		t.Errorf("Error, service status monitor check is %s \n", svc.MonitorStatus)
	}

	for svc.MonitorStatus == "Running" {
		time.Sleep(2000 * time.Millisecond)
		i = i - 1
		if i < 0 {
			svc.StopMonitor()
		}

		if i == 10 {
			svc.CommandStop.Exec()
		}

		if svc.criticalFound == 0 {
			fmt.Print(".")
		} else {
			fmt.Print("S")
		}

	}
	fmt.Println("[END] \n\n")
}

func TestRestExecPingHttpBody(t *testing.T) {

	fmt.Println("[START](Ping HttpBody|Exec REST)")

	//	fmt.Println("Test Check for MySql Service With Execution Rest And Ping HttpBody[REST]")
	var (
		err error
		i   int = 30
	)

	svc := NewService()
	svc.Ping = func() *Ping {
		p := new(Ping)
		p.Type = PingType_HttpBody
		p.Host = "http://192.168.56.101:8080/status"
		p.HttpBodySearch = "RUNNING"
		p.HttpBodyType = HttpBody_Contains
		return p
	}()

	svc.Name = "MongoDb 3.0 WT Port 27123"
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
	}

	svc.CommandStop = &Command{
		Type:            CommandType_REST,
		RESTUrl:         "http://192.168.56.101:8080/stop",
		RESTMethod:      "GET", //POST,GET
		RESTUser:        "ALIP",
		RESTPassword:    "qwerty",
		RESTAuthType:    RESTAuthType_None,
		ValidationType:  ValidationType_Contain,
		ValidationValue: "SUCCESS",
	}

	svc.Log, err = toolkit.NewLog(false, true, "E:\\goproject\\LOG", "TestRestExecPingHttpBody", "20060102")

	if err != nil {
		t.Errorf("Error Start Log : %s", err.Error())
	}

	svc.KeepAlive()

	if svc.MonitorStatus != "Running" {
		t.Errorf("Error, service status monitor check is %s \n", svc.MonitorStatus)
	}

	for svc.MonitorStatus == "Running" {
		time.Sleep(2000 * time.Millisecond)
		i = i - 1
		if i < 0 {
			svc.StopMonitor()
		}

		if i == 10 {
			svc.CommandStop.Exec()
		}

		if svc.criticalFound == 0 {
			fmt.Print(".")
		} else {
			fmt.Print("S")
		}
	}
	fmt.Println("[END] \n\n")
}

func TestSsh(t *testing.T) {
	var SshClient SshParm

	SshClient.SSHAuthType = SSHAuthType_Certificate
	SshClient.SSHHost = "192.168.56.101:22"
	SshClient.SSHUser = "alip"
	SshClient.SSHKeyLocation = "C:\\Users\\User\\.ssh\\id_rsa"

	ps := []string{"sudo service mysql status"}
	res, e := SshClient.RunCommandSsh(ps...)

	if e != nil {
		t.Errorf("Error, %s \n", e)
	} else {
		t.Logf("RUN, %s \n", res)
	}
}

// func TestLocalAll(t *testing.T) {
// 	fmt.Println("Test Check for MongoDb Service Local")
// 	var err error

// 	i := 10

// 		svc := NewService()
// 		svc.Ping = func() *Ping {
// 			p := new(Ping)
// 			p.Type = PingType_HttpBody
// 			p.Host = "http://192.168.56.101:8080/status"
// 			p.HttpBodySearch = "RUNNING"
// 			p.HttpBodyType = HttpBody_Contains
// 			return p
// 		}()

// 		svc.RestartAfterNCritical = 3
// 		svc.Interval = 1 * time.Second
// 		svc.CommandStart = &Command{
// 			Type:            CommandType_REST,
// 			RESTUrl:         "http://192.168.56.101:8080/start",
// 			RESTMethod:      "GET", //POST,GET
// 			RESTUser:        "ALIP",
// 			RESTPassword:    "qwerty",
// 			RESTAuthType:    RESTAuthType_None,
// 			ValidationType:  ValidationType_Contain,
// 			ValidationValue: "SUCCESS",
// 		}

// 	svc.Log, err = toolkit.NewLog(false, true, "E:\\goproject\\LOG", "TestMongoDB", "20060102")

// 	if err != nil {
// 		t.Errorf("Error Start Log : %s", err.Error())
// 	}

// 	svc.KeepAlive()

// 	if svc.MonitorStatus != "Running" {
// 		t.Errorf("Error, service status monitor check is %s \n", svc.MonitorStatus)
// 	}

// 	for svc.MonitorStatus == "Running" {
// 		time.Sleep(2000 * time.Millisecond)
// 		fmt.Println("RUN", i)

// 		i = i - 1

// 		if i < 0 {
// 			svc.StopMonitor()
// 			//			t.SkipNow()
// 		}
// 	}

// }
