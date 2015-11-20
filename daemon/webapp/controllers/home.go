package controllers

import (
	"github.com/astaxie/beego"
	"github.com/eaciit/ecbg"
	lv "github.com/eaciit/live"
	"github.com/eaciit/live/daemon/webapp/helper"
	m "github.com/eaciit/live/daemon/webapp/models"
	"github.com/eaciit/toolkit"
	// "gopkg.in/mgo.v2/bson"
	// "bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	// "strings"
	"time"
	// tk "github.com/eaciit/toolkit"
)

// var arrsvc []*lv.Service
var (
	arrsvc         []modelsvc
	logservicepath = beego.AppConfig.String("logservice_path")
	servicedbpath  = beego.AppConfig.String("servicedb_path")
)

type modelsvc struct {
	ID  int
	svc *lv.Service
}

type HomeController struct {
	// PrivateController
	ecbg.Controller
}

func (this *HomeController) Get() {
	// viewLocation := beego.AppConfig.String("view_dir")

	this.Data["PageId"] = "Dashboard"
	this.Data["PageTitle"] = "Dashboard"
	this.Layout = "shared/layout.tpl"
	this.TplNames = "home/index.tpl"
}

func (this *HomeController) AddService() {
	// r := this.Ctx.Request
	filename := servicedbpath + "/sequenceService.json"
	if _, err := os.Stat(filename); err == nil {
		fmt.Printf("file exists; processing...")
		f, err := os.Open(filename)
		if err != nil {
			fmt.Println(err)
		}
		var seqData []m.SequenceService

		jsonParser := json.NewDecoder(f)
		if err = jsonParser.Decode(&seqData); err != nil {
			fmt.Println(err)
		}
		idService := seqData[0].Seq + 1

		f, err = os.Create(servicedbpath + "/service" + strconv.Itoa(idService) + ".json")
		if err != nil {
			fmt.Println(err)
		}

		dataservice := m.ServiceNew{}
		this.GetPayload(&dataservice)
		dataservice.Service.ID = idService
		dataservice.Service.LastUpdate = time.Now()
		dataservice.Service.DateStatus = time.Now()
		dataservice.Service.StatusService = "Stop"
		b, err := json.Marshal(dataservice)
		n, err := io.WriteString(f, string(b))

		if err != nil {
			fmt.Println(n, err)
		}

		f, err = os.Create(filename)
		if err != nil {
			fmt.Println(err)
		}
		n, err = io.WriteString(f, `[{"Title":"Service", "Seq": `+strconv.Itoa(idService)+`}]`)
		if err != nil {
			fmt.Println(n, err)
		}
		defer f.Close()
	} else {
		f, err := os.Create(filename)
		if err != nil {
			fmt.Println(err)
		}
		n, err := io.WriteString(f, `[{"Title":"Service", "Seq": 1}]`)
		if err != nil {
			fmt.Println(n, err)
		}

		f, err = os.Create(servicedbpath + "/service1.json")
		if err != nil {
			fmt.Println(err)
		}

		dataservice := m.ServiceNew{}
		this.GetPayload(&dataservice)
		fmt.Println(dataservice.Ping.CommandParm)
		dataservice.Service.ID = 1
		dataservice.Service.LastUpdate = time.Now()
		dataservice.Service.DateStatus = time.Now()
		dataservice.Service.StatusService = "Stop"
		b, err := json.Marshal(dataservice)
		n, err = io.WriteString(f, string(b))
		if err != nil {
			fmt.Println(n, err)
		}
		defer f.Close()
	}

	this.Json(helper.BuildResponse(true, nil, ""))
}

func (this *HomeController) UpdateService() {
	dataservice := m.ServiceNew{}
	this.GetPayload(&dataservice)
	fmt.Println(dataservice.Ping.CommandParm)
	dataservice.Service.LastUpdate = time.Now()
	dataservice.Service.StatusService = "Stop"

	f, err := os.Create(servicedbpath + "/service" + strconv.Itoa(dataservice.Service.ID) + ".json")
	if err != nil {
		fmt.Println(err)
	}
	b, err := json.Marshal(dataservice)
	n, err := io.WriteString(f, string(b))
	if err != nil {
		fmt.Println(n, err)
	}
	this.Json(helper.BuildResponse(true, nil, ""))
}

func (this *HomeController) GetService() {
	r := this.Ctx.Request
	var service []m.ServiceNew
	// helper.PopulateAsObject(&service, "Service", bson.M{}, nil, 0, 0)
	// this.Json(helper.BuildResponse(true, service, ""))
	data := m.ServiceNew{}
	files, _ := ioutil.ReadDir(servicedbpath + "/")
	for _, f := range files {
		var extension = filepath.Ext(f.Name())
		if f.Name() != "sequenceService.json" && extension == ".json" {
			serviceFile, err := os.Open(servicedbpath + "/" + f.Name())
			if err != nil {
				fmt.Println(err)
			}
			jsonParser := json.NewDecoder(serviceFile)
			if err = jsonParser.Decode(&data); err != nil {
				fmt.Println(err)
			}
			service = append(service, data)
			defer serviceFile.Close()
			if len(arrsvc) == 0 {
				LogStatus := ConfigService(data, r.FormValue("Statuslive"))
				data.Service.LogStatus = LogStatus
			}
		}
	}
	this.Json(helper.BuildResponse(true, service, ""))
	// this.Json(helper.BuildResponse(true, nil, ""))
}

func (this *HomeController) GetDetailService() {
	r := this.Ctx.Request
	data := m.ServiceNew{}
	serviceFile, err := os.Open(servicedbpath + "/service" + r.FormValue("ID") + ".json")
	if err != nil {
		fmt.Println(err)
	}
	jsonParser := json.NewDecoder(serviceFile)
	if err = jsonParser.Decode(&data); err != nil {
		fmt.Println(err)
	}
	defer serviceFile.Close()
	this.Json(helper.BuildResponse(true, data, ""))
}

func (this *HomeController) RemoveService() {
	r := this.Ctx.Request
	err := os.Remove(servicedbpath + "/service" + r.FormValue("ID") + ".json")
	if err != nil {
		fmt.Println(err)
	}
	this.Json(helper.BuildResponse(true, r.FormValue("ID"), ""))
}

func (this *HomeController) StartService() {
	r := this.Ctx.Request
	data := m.ServiceNew{}
	serviceFile, err := os.Open(servicedbpath + "/service" + r.FormValue("ID") + ".json")
	if err != nil {
		fmt.Println(err)
	}
	jsonParser := json.NewDecoder(serviceFile)
	if err = jsonParser.Decode(&data); err != nil {
		fmt.Println(err)
	}
	defer serviceFile.Close()

	data.Service.DateStatus = time.Now()
	data.Service.StatusService = "Start"

	LogStatus := ConfigService(data, r.FormValue("Statuslive"))
	if LogStatus == "OK" {
		f, err := os.Create(servicedbpath + "/service" + strconv.Itoa(data.Service.ID) + ".json")
		if err != nil {
			fmt.Println(err)
		}
		data.Service.LogStatus = "Success"
		b, err := json.Marshal(data)
		n, err := io.WriteString(f, string(b))
		if err != nil {
			fmt.Println(n, err)
		}
		defer f.Close()
	} else if LogStatus == "Fail" {
		f, err := os.Create(servicedbpath + "/service" + strconv.Itoa(data.Service.ID) + ".json")
		if err != nil {
			fmt.Println(err)
		}
		data.Service.LogStatus = "Fail"
		data.Service.StatusService = "Stop"
		b, err := json.Marshal(data)
		n, err := io.WriteString(f, string(b))
		if err != nil {
			fmt.Println(n, err)
		}
		defer f.Close()
	}

	this.Json(helper.BuildResponse(true, LogStatus, ""))
}

func ConfigService(data m.ServiceNew, statuslive string) string {
	var (
		err          error
		pingtype     lv.PingTypeEnum
		valtypestart lv.ValidationTypeEnum
		valtypestop  lv.ValidationTypeEnum
		// exedtype lv.CommandTypeEnum
	)
	svc := lv.NewService()

	if data.Ping.Type == "PingType_Command" {
		pingtype = lv.PingType_Command
	} else if data.Ping.Type == "PingType_Network" {
		pingtype = lv.PingType_Network
	} else if data.Ping.Type == "PingType_HttpStatus" {
		pingtype = lv.PingType_HttpStatus
	} else if data.Ping.Type == "PingType_HttpBody" {
		pingtype = lv.PingType_HttpBody
	} else {
		pingtype = lv.PingType_Custom
	}

	if data.Ping.Type == "PingType_Command" {
		var resptype lv.ResponseEnum
		if data.Ping.ResponseType == "Response_Contains" {
			resptype = lv.Response_Contains
		} else if data.Ping.ResponseType == "Response_Equals" {
			resptype = lv.Response_Equals
		} else {
			resptype = lv.Response_RegEx
		}
		svc.Ping = func() *lv.Ping {
			p := new(lv.Ping)
			p.Type = pingtype
			p.Command = data.Ping.Command
			p.CommandParms = data.Ping.CommandParm
			p.ResponseType = resptype
			p.ResponseValue = "RUNNING"
			return p
		}()
	} else if data.Ping.Type == "PingType_Network" || data.Ping.Type == "PingType_HttpStatus" {
		svc.Ping = func() *lv.Ping {
			p := new(lv.Ping)
			p.Type = pingtype
			p.Host = "http://" + data.Ping.Host + ":" + data.Ping.Port
			return p
		}()
	} else if data.Ping.Type == "PingType_HttpBody" {
		var bodytype lv.HttpBodyEnum
		if data.Ping.HttpBodyType == "HttpBody_Contains" {
			bodytype = lv.HttpBody_Contains
		} else {
			bodytype = lv.HttpBody_Equals
		}
		svc.Ping = func() *lv.Ping {
			p := new(lv.Ping)
			p.Type = pingtype
			p.Host = "http://" + data.Ping.Host + ":" + data.Ping.Port
			p.HttpBodySearch = "RUNNING"
			p.HttpBodyType = bodytype
			return p
		}()
	}

	svc.RestartAfterNCritical = data.Service.RestartAfterNCritical
	svc.Interval = time.Duration(data.Service.Interval) * time.Second

	if data.ExedCommandStart.ValidationType == "ValidationType_Contain" {
		valtypestart = lv.ValidationType_Contain
	} else if data.ExedCommandStart.ValidationType == "ValidationType_Equal" {
		valtypestart = lv.ValidationType_Equal
	} else {
		valtypestart = lv.ValidationType_Regex
	}
	if data.ExedCommandStop.ValidationType == "ValidationType_Contain" {
		valtypestop = lv.ValidationType_Contain
	} else if data.ExedCommandStop.ValidationType == "ValidationType_Equal" {
		valtypestop = lv.ValidationType_Equal
	} else {
		valtypestop = lv.ValidationType_Regex
	}

	// Exec Start
	if data.ExedCommandStart.Type == "CommandType_Local" {
		svc.CommandStart = &lv.Command{
			Type:            lv.CommandType_Local,
			CommandText:     data.ExedCommandStart.CommandText,
			CommandParms:    data.ExedCommandStart.CommandParm,
			ValidationType:  valtypestart,
			ValidationValue: "RUNNING",
		}
	} else if data.ExedCommandStart.Type == "CommandType_SSH" {
		if data.ExedCommandStart.SshAuthType == "SSHAuthType_Password" {
			svc.CommandStart = &lv.Command{
				Type: lv.CommandType_SSH,
				SshClient: &lv.SshSetting{
					SSHHost:     data.ExedCommandStart.SshHost + ":" + data.ExedCommandStart.SshPort,
					SSHUser:     data.ExedCommandStart.SshUser,
					SSHPassword: data.ExedCommandStart.SshPassword,
					SSHAuthType: lv.SSHAuthType_Password,
				},
				CommandText:     data.ExedCommandStart.CommandText,
				ValidationType:  valtypestart,
				ValidationValue: "running",
			}
		} else {
			svc.CommandStart = &lv.Command{
				Type: lv.CommandType_SSH,
				SshClient: &lv.SshSetting{
					SSHHost:        data.ExedCommandStart.SshHost + ":" + data.ExedCommandStart.SshPort,
					SSHUser:        data.ExedCommandStart.SshUser,
					SSHKeyLocation: data.ExedCommandStart.SshKeyLocation,
					SSHAuthType:    lv.SSHAuthType_Certificate,
				},
				CommandText:     data.ExedCommandStart.CommandText,
				ValidationType:  valtypestart,
				ValidationValue: "running",
			}
		}
	} else {
		var valrestaunth lv.RESTAuthTypeEnum
		if data.ExedCommandStart.RestAuthType == "RESTAuthType_None" {
			valrestaunth = lv.RESTAuthType_None
		} else {
			valrestaunth = lv.RESTAuthType_Basic
		}
		svc.CommandStart = &lv.Command{
			Type:            lv.CommandType_REST,
			RESTUrl:         data.ExedCommandStart.RestUrl,
			RESTMethod:      data.ExedCommandStart.RestMenthod, //POST,GET
			RESTUser:        data.ExedCommandStart.RestUser,
			RESTPassword:    data.ExedCommandStart.RestPassword,
			RESTAuthType:    valrestaunth,
			ValidationType:  valtypestart,
			ValidationValue: "SUCCESS",
		}
	}

	// Exec Stop
	if data.ExedCommandStop.Type == "CommandType_Local" {
		svc.CommandStop = &lv.Command{
			Type:            lv.CommandType_Local,
			CommandText:     data.ExedCommandStop.CommandText,
			CommandParms:    data.ExedCommandStop.CommandParm,
			ValidationType:  valtypestop,
			ValidationValue: "STOP_PENDING",
		}
	} else if data.ExedCommandStop.Type == "CommandType_SSH" {
		if data.ExedCommandStop.SshAuthType == "SSHAuthType_Password" {
			svc.CommandStop = &lv.Command{
				Type: lv.CommandType_SSH,
				SshClient: &lv.SshSetting{
					SSHHost:     data.ExedCommandStop.SshHost + ":" + data.ExedCommandStop.SshPort,
					SSHUser:     data.ExedCommandStop.SshUser,
					SSHPassword: data.ExedCommandStop.SshPassword,
					SSHAuthType: lv.SSHAuthType_Password,
				},
				CommandText:     data.ExedCommandStop.CommandText,
				ValidationType:  valtypestop,
				ValidationValue: "running",
			}
		} else {
			svc.CommandStop = &lv.Command{
				Type: lv.CommandType_SSH,
				SshClient: &lv.SshSetting{
					SSHHost:        data.ExedCommandStop.SshHost + ":" + data.ExedCommandStop.SshPort,
					SSHUser:        data.ExedCommandStop.SshUser,
					SSHKeyLocation: data.ExedCommandStop.SshKeyLocation,
					SSHAuthType:    lv.SSHAuthType_Certificate,
				},
				CommandText:     data.ExedCommandStop.CommandText,
				ValidationType:  valtypestop,
				ValidationValue: "running",
			}
		}
	} else {
		var valrestaunth lv.RESTAuthTypeEnum
		if data.ExedCommandStop.RestAuthType == "RESTAuthType_None" {
			valrestaunth = lv.RESTAuthType_None
		} else {
			valrestaunth = lv.RESTAuthType_Basic
		}
		svc.CommandStop = &lv.Command{
			Type:            lv.CommandType_REST,
			RESTUrl:         data.ExedCommandStop.RestUrl,
			RESTMethod:      data.ExedCommandStop.RestMenthod, //POST,GET
			RESTUser:        data.ExedCommandStop.RestUser,
			RESTPassword:    data.ExedCommandStop.RestPassword,
			RESTAuthType:    valrestaunth,
			ValidationType:  valtypestop,
			ValidationValue: "SUCCESS",
		}
	}

	svc.EmailError = data.Service.EmailWarning
	svc.EmailWarning = data.Service.EmailError

	svc.Mail = &lv.EmailSetting{
		SenderEmail:   "admin.support@eaciit.com",
		HostEmail:     "smtp.office365.com",
		PortEmail:     587,
		UserEmail:     "admin.support@eaciit.com",
		PasswordEmail: "B920Support",
	}

	svc.Log, err = toolkit.NewLog(false, true, logservicepath+"/", "LogService"+strconv.Itoa(data.Service.ID), "20060102")
	if err != nil {
		fmt.Println("Error Start Log : %s", err.Error())
	}
	datasvc := modelsvc{}
	datasvc.ID = data.Service.ID
	datasvc.svc = svc

	if statuslive == "Start" && data.Service.StatusService == "Start" {
		if len(arrsvc) == 0 && data.Service.StatusService == "Start" {
			svc.KeepAlive()
			arrsvc = append(arrsvc, datasvc)
			return "Preparing"
		}
		for j := 0; j < len(arrsvc); j++ {
			if arrsvc[j].ID != data.Service.ID && data.Service.StatusService == "Start" {
				if statuslive == "Start" {
					svc.KeepAlive()
					arrsvc = append(arrsvc, datasvc)
					return "Preparing"
				}
			} else if data.Service.StatusService == "Start" {
				if statuslive == "Start" {
					svc.KeepAlive()
					arrsvc[j] = datasvc
					return "Preparing"
				}
			}
		}
	} else if statuslive == "Live" && data.Service.StatusService == "Start" {
		for j := 0; j < len(arrsvc); j++ {
			if arrsvc[j].ID == data.Service.ID {
				if arrsvc[j].svc.MonitorStatus == "Running" {
					if arrsvc[j].svc.PingStatus == "OK" {
						return arrsvc[j].svc.PingStatus
					} else if arrsvc[j].svc.PingStatus == "Fail" {
						return arrsvc[j].svc.PingStatus
					} else {
						return arrsvc[j].svc.PingStatus
					}
				}
			}
		}
	}
	return "Fail"
}

func (this *HomeController) StopService() {
	r := this.Ctx.Request
	data := m.ServiceNew{}
	serviceFile, err := os.Open(servicedbpath + "/service" + r.FormValue("ID") + ".json")
	if err != nil {
		fmt.Println(err)
	}
	jsonParser := json.NewDecoder(serviceFile)
	if err = jsonParser.Decode(&data); err != nil {
		fmt.Println(err)
	}
	defer serviceFile.Close()

	f, err := os.Create(servicedbpath + "/service" + strconv.Itoa(data.Service.ID) + ".json")
	if err != nil {
		fmt.Println(err)
	}

	data.Service.DateStatus = time.Now()
	data.Service.StatusService = "Stop"
	data.Service.LogStatus = "Success"
	b, err := json.Marshal(data)
	n, err := io.WriteString(f, string(b))
	if err != nil {
		fmt.Println(n, err)
	}

	fmt.Println(arrsvc)
	for j := 0; j < len(arrsvc); j++ {
		if arrsvc[j].ID == data.Service.ID && data.Service.StatusService == "Stop" {
			arrsvc[j].svc.StopMonitor()
		}
	}
	this.Json(helper.BuildResponse(true, nil, ""))
}

func (this *HomeController) StopServer() {
	r := this.Ctx.Request
	data := m.ServiceNew{}
	serviceFile, err := os.Open(servicedbpath + "/service" + r.FormValue("ID") + ".json")
	if err != nil {
		fmt.Println(err)
	}
	jsonParser := json.NewDecoder(serviceFile)
	if err = jsonParser.Decode(&data); err != nil {
		fmt.Println(err)
	}
	defer serviceFile.Close()

	f, err := os.Create(servicedbpath + "/service" + strconv.Itoa(data.Service.ID) + ".json")
	if err != nil {
		fmt.Println(err)
	}

	data.Service.DateStatus = time.Now()
	b, err := json.Marshal(data)
	n, err := io.WriteString(f, string(b))
	if err != nil {
		fmt.Println(n, err)
	}

	fmt.Println(arrsvc)
	for j := 0; j < len(arrsvc); j++ {
		if arrsvc[j].ID == data.Service.ID {
			arrsvc[j].svc.CommandStop.Exec()
		}
	}
	this.Json(helper.BuildResponse(true, nil, ""))
}

func (this *HomeController) GetLogService() {
	r := this.Ctx.Request
	var valLog string
	b, err := ioutil.ReadFile(servicedbpath + "/LogService" + r.FormValue("ID") + "%!(EXTRA string=" + r.FormValue("DateFilter") + ")")
	if err != nil {
		valLog = ""
		fmt.Println(err)
	} else {
		valLog = string(b)
	}
	this.Json(helper.BuildResponse(true, valLog, ""))
}
