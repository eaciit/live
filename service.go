package live

import (
	"fmt"
	"github.com/eaciit/toolkit"
	"strings"
	"time"
)

type Service struct {
	Ping                  *Ping
	Name                  string
	Status                string
	MonitorStatus         string
	Interval              time.Duration
	CriticalInterval      time.Duration
	RestartAfterNCritical int

	CommandStart *Command
	CommandStop  *Command

	criticalFound int

	cstate    chan string
	lastCheck time.Time

	Log *toolkit.LogEngine
}

func NewService() *Service {
	s := new(Service)
	s.cstate = make(chan string)
	s.Interval = 15 * time.Millisecond
	s.CriticalInterval = 1 * time.Millisecond
	s.RestartAfterNCritical = 3
	return s
}

func (s *Service) AddLog(logtext, logtype string) {
	if s.Log == nil {
		s.Log, _ = toolkit.NewLog(true, false, "", "", "")
	}
	s.Log.AddLog(logtext, logtype)
}

func (s *Service) KeepAlive() {
	s.MonitorStatus = "Running"
	s.Status = "OK"
	go func(s *Service) {
		for s.MonitorStatus == "Running" {
			select {
			case <-time.After(s.Interval):
				if s.criticalFound < s.RestartAfterNCritical {
					e := s.Ping.Check()
					if e != nil {
						s.Status = s.Ping.LastStatus
						s.criticalFound++
						s.AddLog(fmt.Sprintf("[Service %s check fails - %d. Error: %s]", s.Name, s.criticalFound, e.Error()), "ERROR")
						if s.criticalFound == s.RestartAfterNCritical {
							e = s.bringItUp()
							if e != nil {
								s.AddLog(fmt.Sprintf("[Service %s restart fails - %d. Error: %s]", s.Name, 1, e.Error()), "ERROR")
							} else {
								s.AddLog(fmt.Sprintf("[Service %s restarted successfully]", s.Name), "INFO")
								s.criticalFound = 0
								s.Status = "OK"
							}
						}
					} else {
						s.criticalFound = 0
						s.AddLog(fmt.Sprintf("[Service %s ping successfully]", s.Name), "INFO")
					}
				} else if s.criticalFound == s.RestartAfterNCritical {
					s.AddLog(fmt.Sprintf("[Max critical event (%d) has been exceeded. Service monitor will be stopped]", s.RestartAfterNCritical), "WARNING")
					s.criticalFound++
					s.StopMonitor()
				}
			}
		}
	}(s)
}

func (s *Service) StopMonitor() {
	s.MonitorStatus = "Stop"
}

func (s *Service) receiveState() {
	go func(s *Service) {
		run := true
		for run {
			select {
			case newState := <-s.cstate:
				s.Status = newState
				newState = strings.ToLower(newState)
				if newState == "stop" {
					run = false
				}
			}
		}
	}(s)
}

func (s *Service) bringItUp() error {
	var (
		e   error
		res string
	)

	if s.Status == "OK" {
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
