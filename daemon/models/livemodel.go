package models

import (
	"time"
)

type Service struct {
	ID            int       `json:"ID"`
	Title         string    `json:"Title"`
	Description   string    `json:"Description"`
	Host          string    `json:"Host"`
	Port          string    `json:"Port"`
	Path          string    `json:"Path"`
	CreateService time.Time `json:"CreateService"`
	Status        string    `json:"Status"`
	Username      string    `json:"Username"`
	Password      string    `json:"Password"`
	LastUpdate    time.Time `json:"LastUpdate"`
}

type ServiceNew struct {
	Service          ServiceModel     `json:"Service"`
	Ping             Ping             `json:"Ping"`
	ExedCommandStart ExedCommandStart `json:"ExedCommandStart"`
	ExedCommandStop  ExedCommandStop  `json:"ExedCommandStop"`
}

type ExedCommandStart struct {
	Type            string   `json:"Type"`
	CommandText     string   `json:"CommandText"`
	CommandTextSsh  string   `json:"CommandTextSsh"`
	CommandParm     []string `json:"CommandParm"`
	RestUrl         string   `json:"RestUrl"`
	RestMenthod     string   `json:"RestMenthod"`
	RestUser        string   `json:"RestUser"`
	RestPassword    string   `json:"RestPassword"`
	RestAuthType    string   `json:"RestAuthType"`
	SshHost         string   `json:"SshHost"`
	SshPort         string   `json:"SshPort"`
	SshUser         string   `json:"SshUser"`
	SshPassword     string   `json:"SshPassword"`
	SshKeyLocation  string   `json:"SshKeyLocation"`
	SshAuthType     string   `json:"SshAuthType"`
	ValidationType  string   `json:"ValidationType"`
	ValidationValue string   `json:"ValidationValue"`
}

type ExedCommandStop struct {
	Type            string   `json:"Type"`
	CommandText     string   `json:"CommandText"`
	CommandTextSsh  string   `json:"CommandTextSsh"`
	CommandParm     []string `json:"CommandParm"`
	RestUrl         string   `json:"RestUrl"`
	RestMenthod     string   `json:"RestMenthod"`
	RestUser        string   `json:"RestUser"`
	RestPassword    string   `json:"RestPassword"`
	RestAuthType    string   `json:"RestAuthType"`
	SshHost         string   `json:"SshHost"`
	SshPort         string   `json:"SshPort"`
	SshUser         string   `json:"SshUser"`
	SshPassword     string   `json:"SshPassword"`
	SshKeyLocation  string   `json:"SshKeyLocation"`
	SshAuthType     string   `json:"SshAuthType"`
	ValidationType  string   `json:"ValidationType"`
	ValidationValue string   `json:"ValidationValue"`
}

type ServiceModel struct {
	ID                    int       `json:"ID"`
	Title                 string    `json:"Title"`
	Description           string    `json:"Description"`
	RestartAfterNCritical int       `json:"RestartAfterNCritical"`
	Interval              int       `json:"Interval"`
	PathLog               string    `json:"PathLog"`
	TypeLog               string    `json:"TypeLog"`
	StatusService         string    `json:"StatusService"`
	EmailWarning          []string  `json:"EmailWarning"`
	EmailError            []string  `json:"EmailError"`
	LogStatus             string    `json:"LogStatus"`
	DateStatus            time.Time `json:"DateStatus"`
	LastUpdate            time.Time `json:"LastUpdate"`
}

type Ping struct {
	Type           string   `json:"Type"`
	User           string   `json:"User"`
	Password       string   `json:"Password"`
	Host           string   `json:"Host"`
	Port           string   `json:"Port"`
	LastStatus     string   `json:"LastStatus"`
	Command        string   `json:"Command"`
	CommandParm    []string `json:"CommandParm"`
	ResponseType   string   `json:"ResponseType"`
	ResponseValue  string   `json:"ResponseValue"`
	HttpBodyType   string   `json:"HttpBodyType"`
	HttpBodySearch string   `json:"HttpBodySearch"`
}

type ExedCommand struct {
	Type                 string   `json:"Type"`
	CommandText          string   `json:"CommandText"`
	CommandTextStart     string   `json:"CommandTextStart"`
	CommandTextStop      string   `json:"CommandTextStop"`
	CommandParmStart     []string `json:"CommandParmStart"`
	CommandParmStop      []string `json:"CommandParmStop"`
	RestUrl              string   `json:"RestUrl"`
	RestUrlStop          string   `json:"RestUrlStop"`
	RestMenthod          string   `json:"RestMenthod"`
	RestUser             string   `json:"RestUser"`
	RestPassword         string   `json:"RestPassword"`
	RestAuthType         string   `json:"RestAuthType"`
	SshHost              string   `json:"SshHost"`
	SshPort              string   `json:"SshPort"`
	SshUser              string   `json:"SshUser"`
	SshPassword          string   `json:"SshPassword"`
	SshKeyLocation       string   `json:"SshKeyLocation"`
	SshAuthType          string   `json:"SshAuthType"`
	ValidationTypeStart  string   `json:"ValidationTypeStart"`
	ValidationValueStart string   `json:"ValidationValueStart"`
	ValidationTypeStop   string   `json:"ValidationTypeStop"`
	ValidationValueStop  string   `json:"ValidationValueStop"`
}

type SequenceService struct {
	Title string `json:"Title"`
	Seq   int    `json:"Seq"`
}
