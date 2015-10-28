package live

import (
	"strings"
	"time"
)

type Service struct {
	Ping             *Ping
	Status           string
	MonitorStatus    string
	Interval         time.Duration
	CriticalInterval time.Duration

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
	return s
}

func (s *Service) KeepAlive() {
	go func(s *Service) {
		s.MonitorStatus = "Run"
		for s.MonitorStatus == "Run" {
			select {
			case <-time.After(s.Interval):
				e := s.Ping.Check()
				if e != nil {
					s.criticalFound++
					s.Status = s.Ping.LastStatus
				} else {
					s.criticalFound = 0
				}
			}
		}
	}(s)
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
}
