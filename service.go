package live

import (
	"fmt"
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
}

func NewService() *Service {
	s := new(Service)
	s.cstate = make(chan string)
	s.Interval = 15 * time.Millisecond
	s.CriticalInterval = 1 * time.Millisecond
	s.RestartAfterNCritical = 3
	return s
}

func (s *Service) KeepAlive() {
	s.MonitorStatus = "Running"
	go func(s *Service) {
		for s.MonitorStatus == "Running" {
			select {
			case <-time.After(s.Interval):
				if s.criticalFound < s.RestartAfterNCritical {
					e := s.Ping.Check()
					if e != nil {
						s.Status = s.Ping.LastStatus
						s.criticalFound++
						fmt.Printf("[%v] Service %s check fails - %d. Error: %s \n", time.Now(), s.Name, s.criticalFound, e.Error())
						if s.criticalFound == s.RestartAfterNCritical {
							s.bringItUp()
						}
					} else {
						s.criticalFound = 0
					}
				} else if s.criticalFound == s.RestartAfterNCritical {
					fmt.Printf("[%v] Max critical event (%d) has been exceeded. Service monitor will be stopped\n",
						time.Now(),
						s.RestartAfterNCritical)
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

	if s.Status != "Stop" {
		e = s.CommandStart.Exec()
		if e != nil {
			return e
		}
	}

	e = s.CommandStart.Exec()
	if e != nil {
		return e
	}
	return nil
}
