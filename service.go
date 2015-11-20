package live

import (
	"fmt"
	"github.com/eaciit/toolkit"
	"strings"
	"time"
)

type Service struct {
	// these attributes are used for initialization service
	Name                  string
	PingStatus            string //Preparing,Fail,[STOP,OK]
	MonitorStatus         string
	Interval              time.Duration
	CriticalInterval      time.Duration
	RestartAfterNCritical int

	// these attributes are used for limitation warning ping found
	criticalFound int

	// these attributes are used for ping service
	Ping *Ping

	// these attributes are used for start and stop service over live library
	CommandStart *Command
	CommandStop  *Command

	cstate    chan string
	lastCheck time.Time

	// these attributes are used for log setting
	Log *toolkit.LogEngine

	// these attributes are used for mail setting when warning or error found
	Mail         *EmailSetting
	EmailWarning []string
	EmailError   []string
}

/*
Initialization new service with some atribute value
*/
func NewService() *Service {
	s := new(Service)
	s.cstate = make(chan string)
	s.Interval = 15 * time.Millisecond
	s.CriticalInterval = 1 * time.Millisecond
	s.RestartAfterNCritical = 3
	return s
}

/*
Generate log as condition inputed, if warning and error found will send email too
*/
func (s *Service) AddLog(logtext, logtype string) {
	if s.Log == nil {
		s.Log, _ = toolkit.NewLog(true, false, "", "", "")
	}
	s.Log.AddLog(logtext, logtype)

	subj := fmt.Sprintf("[%s] From Service %s at %s", logtype, s.Name, time.Now().Format("20060102_15:04:05"))
	msg := fmt.Sprintf("Message From Live Service : %s", logtext)

	if logtype == "ERROR" {
		s.Mail.sendEmail(subj, msg, s.EmailError)
	} else if logtype == "WARNING" {
		s.Mail.sendEmail(subj, msg, s.EmailWarning)
	}
}

/*
Set monitor status to Running and ping service as atribute has been determined
if critical condition found, keep service up.
if Error found when start the service, set monitor status to stop
*/
func (s *Service) KeepAlive() {
	s.MonitorStatus = "Running"
	s.PingStatus = "Preparing"
	go func(s *Service) {
		for s.MonitorStatus == "Running" {
			select {
			case <-time.After(s.Interval):
				if s.criticalFound < s.RestartAfterNCritical {
					e := s.Ping.Check()
					s.PingStatus = s.Ping.LastStatus
					if e != nil {
						s.criticalFound++
						s.AddLog(fmt.Sprintf("[Service %s check fails - %d. Error: %s]", s.Name, s.criticalFound, e.Error()), "WARNING")
						if s.criticalFound == s.RestartAfterNCritical {
							e = s.bringItUp()
							if e != nil {
								s.AddLog(fmt.Sprintf("[Service %s restart fails - %d. Error: %s]", s.Name, 1, e.Error()), "ERROR")
							} else {
								s.AddLog(fmt.Sprintf("[Service %s restarted successfully]", s.Name), "INFO")
								s.criticalFound = 0
								s.PingStatus = "OK"
							}
						}
					} else {
						s.criticalFound = 0
						s.AddLog(fmt.Sprintf("[Service %s ping successfully]", s.Name), "INFO")
					}
				} else if s.criticalFound == s.RestartAfterNCritical {
					s.AddLog(fmt.Sprintf("[Max critical event (%d) has been exceeded. Service monitor will be stopped]", s.RestartAfterNCritical), "ERROR")
					s.criticalFound++
					s.StopMonitor()
				}
			}
		}
	}(s)
}

/*
Stop Monitor status
*/
func (s *Service) StopMonitor() {
	s.MonitorStatus = "Stop"
}

/*
Set service status
*/
func (s *Service) receiveState() {
	go func(s *Service) {
		run := true
		for run {
			select {
			case newState := <-s.cstate:
				s.PingStatus = newState
				newState = strings.ToLower(newState)
				if newState == "stop" {
					run = false
				}
			}
		}
	}(s)
}

/*
Keep the service live with restart or turn on the service
*/
func (s *Service) bringItUp() error {
	var (
		e   error
		res string
	)

	if s.PingStatus == "OK" {
		if s.CommandStop != nil {
			_, e = s.CommandStop.Exec()
		}
	}

	res, e = s.CommandStart.Exec()

	if e != nil {
		return e
	}

	if s.CommandStart.ValidationValue != "" {
		e = s.CommandStart.Validate(res)
	}

	if e != nil {
		return e
	}

	return nil
}
