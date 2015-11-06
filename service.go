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
	LogEngine    *toolkit.LogEngine

	criticalFound int

	cstate    chan string
	lastCheck time.Time
}

func NewService() *Service {
	s := new(Service)
	s.cstate = make(chan string)
	s.Interval = 15 * time.Millisecond
	s.CriticalInterval = 1 * time.Millisecond
	s.RestartAfterNCritical = 3
	return s
}

func (s *Service) addLog(logtxt string, logtype string) {
	if s.LogEngine == nil {
		le, _ := toolkit.NewLog(true, false, "", "", "")
		s.LogEngine = le
	}
	s.LogEngine.AddLog(logtxt, logtype)
}

func (s *Service) KeepAlive() {
	s.MonitorStatus = "Running"
	s.Status = "OK"
	s.addLog(fmt.Sprintf("Service %s live monitor started", s.Name), "INFO")
	go func(s *Service) {
		for s.MonitorStatus == "Running" {
			select {
			case <-time.After(s.Interval):
				if s.criticalFound < s.RestartAfterNCritical {
					e := s.Ping.Check()
					if e != nil {
						s.Status = s.Ping.LastStatus
						s.criticalFound++
						s.addLog(fmt.Sprintf("Service %s check fails - %d. Error: %s \n",
							s.Name, s.criticalFound, e.Error()), "WARNING")
						if s.criticalFound == s.RestartAfterNCritical {
							e = s.bringItUp()
							if e != nil {
								s.addLog(
									fmt.Sprintf("Service %s restart fails - %d. Error: %s \n", s.Name, 1, e.Error()),
									"ERROR")
							} else {
								s.addLog(
									fmt.Sprintf("Service %s restarted successfully \n", s.Name),
									"INFO")
								s.criticalFound = 0
								s.Status = "OK"
							}
						}
					} else {
						s.criticalFound = 0
					}
				} else if s.criticalFound == s.RestartAfterNCritical {
					s.addLog(
						fmt.Sprintf("Max critical event (%d) has been exceeded. Service monitor will be stopped\n",
							s.RestartAfterNCritical), "ERROR")
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
	var e error

	if s.Status == "OK" {
		if s.CommandStop != nil {
			s.CommandStop.Exec()
		}
	}

	e = s.CommandStart.Exec()
	if e != nil {
		return e
	}

	return nil
}
