package controllers

import (
	// "github.com/astaxie/beego"
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
var arrsvc []modelsvc

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
	filename := "static/servicedb/sequenceService.json"
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

		f, err = os.Create("static/servicedb/service" + strconv.Itoa(idService) + ".json")
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

		f, err = os.Create("static/servicedb/service1.json")
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

	f, err := os.Create("static/servicedb/service" + strconv.Itoa(dataservice.Service.ID) + ".json")
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
	var service []m.ServiceNew
	// helper.PopulateAsObject(&service, "Service", bson.M{}, nil, 0, 0)
	// this.Json(helper.BuildResponse(true, service, ""))
	data := m.ServiceNew{}
	files, _ := ioutil.ReadDir("static/servicedb/")
	for _, f := range files {
		var extension = filepath.Ext(f.Name())
		if f.Name() != "sequenceService.json" && extension == ".json" {
			serviceFile, err := os.Open("static/servicedb/" + f.Name())
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
				ConfigService(data)
			}
		}
	}
	this.Json(helper.BuildResponse(true, service, ""))
	// this.Json(helper.BuildResponse(true, nil, ""))
}

func (this *HomeController) GetDetailService() {
	r := this.Ctx.Request
	data := m.ServiceNew{}
	serviceFile, err := os.Open("static/servicedb/service" + r.FormValue("ID") + ".json")
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
	err := os.Remove("static/servicedb/service" + r.FormValue("ID") + ".json")
	if err != nil {
		fmt.Println(err)
	}
	this.Json(helper.BuildResponse(true, r.FormValue("ID"), ""))
}

func (this *HomeController) StartService() {
	r := this.Ctx.Request
	data := m.ServiceNew{}
	serviceFile, err := os.Open("static/servicedb/service" + r.FormValue("ID") + ".json")
	if err != nil {
		fmt.Println(err)
	}
	jsonParser := json.NewDecoder(serviceFile)
	if err = jsonParser.Decode(&data); err != nil {
		fmt.Println(err)
	}
	defer serviceFile.Close()

	f, err := os.Create("static/servicedb/service" + strconv.Itoa(data.Service.ID) + ".json")
	if err != nil {
		fmt.Println(err)
	}

	data.Service.DateStatus = time.Now()
	data.Service.StatusService = "Start"
	b, err := json.Marshal(data)
	n, err := io.WriteString(f, string(b))
	if err != nil {
		fmt.Println(n, err)
	}

	defer f.Close()

	ConfigService(data)

	this.Json(helper.BuildResponse(true, nil, ""))
}

func ConfigService(data m.ServiceNew) {
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

	if data.ExedCommand.ValidationTypeStart == "ValidationType_Contain" {
		valtypestart = lv.ValidationType_Contain
	} else if data.ExedCommand.ValidationTypeStart == "ValidationType_Equal" {
		valtypestart = lv.ValidationType_Equal
	} else {
		valtypestart = lv.ValidationType_Regex
	}
	if data.ExedCommand.ValidationTypeStop == "ValidationType_Contain" {
		valtypestop = lv.ValidationType_Contain
	} else if data.ExedCommand.ValidationTypeStop == "ValidationType_Equal" {
		valtypestop = lv.ValidationType_Equal
	} else {
		valtypestop = lv.ValidationType_Regex
	}

	if data.ExedCommand.Type == "CommandType_Local" {
		svc.CommandStart = &lv.Command{
			Type:            lv.CommandType_Local,
			CommandText:     data.ExedCommand.CommandText,
			CommandParms:    data.ExedCommand.CommandParmStart,
			ValidationType:  valtypestart,
			ValidationValue: "RUNNING",
		}

		svc.CommandStop = &lv.Command{
			Type:            lv.CommandType_Local,
			CommandText:     data.ExedCommand.CommandText,
			CommandParms:    data.ExedCommand.CommandParmStop,
			ValidationType:  valtypestop,
			ValidationValue: "STOP_PENDING",
		}
	} else if data.ExedCommand.Type == "CommandType_SSH" {
		if data.ExedCommand.SshAuthType == "SSHAuthType_Password" {
			svc.CommandStart = &lv.Command{
				Type: lv.CommandType_SSH,
				SshClient: &lv.SshParm{
					SSHHost:     data.ExedCommand.SshHost + ":" + data.ExedCommand.SshPort,
					SSHUser:     data.ExedCommand.SshUser,
					SSHPassword: data.ExedCommand.SshPassword,
					SSHAuthType: lv.SSHAuthType_Password,
				},
				CommandText:     data.ExedCommand.CommandTextStart,
				ValidationType:  valtypestart,
				ValidationValue: "running",
			}

			svc.CommandStop = &lv.Command{
				Type: lv.CommandType_SSH,
				SshClient: &lv.SshParm{
					SSHHost:     data.ExedCommand.SshHost + ":" + data.ExedCommand.SshPort,
					SSHUser:     data.ExedCommand.SshUser,
					SSHPassword: data.ExedCommand.SshPassword,
					SSHAuthType: lv.SSHAuthType_Password,
				},
				CommandText:     data.ExedCommand.CommandTextStop,
				ValidationType:  valtypestop,
				ValidationValue: "running",
			}
		} else {
			svc.CommandStart = &lv.Command{
				Type: lv.CommandType_SSH,
				SshClient: &lv.SshParm{
					SSHHost:        data.ExedCommand.SshHost + ":" + data.ExedCommand.SshPort,
					SSHUser:        data.ExedCommand.SshUser,
					SSHKeyLocation: data.ExedCommand.SshKeyLocation,
					SSHAuthType:    lv.SSHAuthType_Certificate,
				},
				CommandText:     data.ExedCommand.CommandTextStart,
				ValidationType:  valtypestart,
				ValidationValue: "running",
			}

			svc.CommandStop = &lv.Command{
				Type: lv.CommandType_SSH,
				SshClient: &lv.SshParm{
					SSHHost:        data.ExedCommand.SshHost + ":" + data.ExedCommand.SshPort,
					SSHUser:        data.ExedCommand.SshUser,
					SSHKeyLocation: data.ExedCommand.SshKeyLocation,
					SSHAuthType:    lv.SSHAuthType_Certificate,
				},
				CommandText:     data.ExedCommand.CommandTextStop,
				ValidationType:  valtypestop,
				ValidationValue: "running",
			}
		}
	} else {
		var valrestaunth lv.RESTAuthTypeEnum
		if data.ExedCommand.RestAuthType == "RESTAuthType_None" {
			valrestaunth = lv.RESTAuthType_None
		} else {
			valrestaunth = lv.RESTAuthType_Basic
		}
		svc.CommandStart = &lv.Command{
			Type:            lv.CommandType_REST,
			RESTUrl:         data.ExedCommand.RestUrl,
			RESTMethod:      data.ExedCommand.RestMenthod, //POST,GET
			RESTUser:        data.ExedCommand.RestUser,
			RESTPassword:    data.ExedCommand.RestPassword,
			RESTAuthType:    valrestaunth,
			ValidationType:  valtypestart,
			ValidationValue: "SUCCESS",
		}

		svc.CommandStop = &lv.Command{
			Type:            lv.CommandType_REST,
			RESTUrl:         data.ExedCommand.RestUrlStop,
			RESTMethod:      data.ExedCommand.RestMenthod, //POST,GET
			RESTUser:        data.ExedCommand.RestUser,
			RESTPassword:    data.ExedCommand.RestPassword,
			RESTAuthType:    valrestaunth,
			ValidationType:  valtypestop,
			ValidationValue: "SUCCESS",
		}
	}

	svc.Log, err = toolkit.NewLog(false, true, "static/logservice/", "LogService"+strconv.Itoa(data.Service.ID), "20060102")
	if err != nil {
		fmt.Println("Error Start Log : %s", err.Error())
	}
	datasvc := modelsvc{}
	datasvc.ID = data.Service.ID
	datasvc.svc = svc
	if len(arrsvc) == 0 && data.Service.StatusService == "Start" {
		svc.KeepAlive()
		arrsvc = append(arrsvc, datasvc)
	}
	for j := 0; j < len(arrsvc); j++ {
		if arrsvc[j].ID != data.Service.ID && data.Service.StatusService == "Start" {
			svc.KeepAlive()
			arrsvc = append(arrsvc, datasvc)
		} else if data.Service.StatusService == "Start" {
			svc.KeepAlive()
			arrsvc[j] = datasvc
		}
	}
}

func (this *HomeController) StopService() {
	r := this.Ctx.Request
	data := m.ServiceNew{}
	serviceFile, err := os.Open("static/servicedb/service" + r.FormValue("ID") + ".json")
	if err != nil {
		fmt.Println(err)
	}
	jsonParser := json.NewDecoder(serviceFile)
	if err = jsonParser.Decode(&data); err != nil {
		fmt.Println(err)
	}
	defer serviceFile.Close()

	f, err := os.Create("static/servicedb/service" + strconv.Itoa(data.Service.ID) + ".json")
	if err != nil {
		fmt.Println(err)
	}

	data.Service.DateStatus = time.Now()
	data.Service.StatusService = "Stop"
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

func (this *HomeController) GetLogService() {
	r := this.Ctx.Request
	var valLog string
	b, err := ioutil.ReadFile("static/logservice/LogService" + r.FormValue("ID") + "%!(EXTRA string=" + r.FormValue("DateFilter") + ")")
	if err != nil {
		valLog = ""
		fmt.Println(err)
	} else {
		valLog = string(b)
	}
	this.Json(helper.BuildResponse(true, valLog, ""))
}
