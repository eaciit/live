
<link rel="shortcut icon" href="../favicon.ico">
<link rel="stylesheet" type="text/css" href="/static/gridlib/css/normalize.css" />
<link rel="stylesheet" type="text/css" href="/static/gridlib/fonts/font-awesome-4.3.0/css/font-awesome.min.css" />
<link rel="stylesheet" type="text/css" href="/static/gridlib/css/style1.css" />
<script src="/static/gridlib/js/modernizr.custom.js"></script>
<script type="text/javascript" src="/static/tokeninput/jquery.tokeninput.js"></script>
<link rel="stylesheet" href="/static/tokeninput/token-input.css" type="text/css" />
<link rel="stylesheet" href="/static/tokeninput/token-input-bootstrap.css" type="text/css" />
<link href="/static/tokeninput/token-input-facebook.css" type="text/css" rel="stylesheet" />

<script>
	model.BreadCrumbs.push(new BreadCrumb("Home", "Home", "#", "active", ""));
  var modelServiceNew = {
    Service:{
      ID: 0,
      Title: "",
      Description: "",
      RestartAfterNCritical : 3,
      Interval: 1,
      PathLog: "",
      TypeLog: "Daily",
      StatusService: "",
      EmailWarning: [],
      EmailError: [],
      LogStatus: "Success",
    },
    Ping : {
      Type : "PingType_Network",
      User : "",
      Password : "",
      Host : "",
      Port : "",
      LastStatus : "",
      Command : "",
      CommandParm : [],
      ResponseType : "Response_Contains",
      ResponseValue : "",
      HttpBodyType : "HttpBody_Contains",
      HttpBodySearch : ""
    },
    ExedCommandStart : {
      Type : "CommandType_Local",
      CommandText : "",
      CommandTextSsh: "",
      CommandParm : [],
      RestUrl : "",
      RestMenthod : "",
      RestUser : "",
      RestPassword : "",
      RestAuthType : "RESTAuthType_None",
      SshHost : "",
      SshPort : "",
      SshUser : "",
      SshPassword : "",
      SshKeyLocation : "",
      SshAuthType : "SSHAuthType_Password",
      ValidationType: "ValidationType_Contain",
      ValidationValue: "",
    },
    ExedCommandStop : {
      Type : "CommandType_Local",
      CommandText : "",
      CommandTextSsh: "",
      CommandParm : [],
      RestUrl : "",
      RestMenthod : "",
      RestUser : "",
      RestPassword : "",
      RestAuthType : "RESTAuthType_None",
      SshHost : "",
      SshPort : "",
      SshUser : "",
      SshPassword : "",
      SshKeyLocation : "",
      SshAuthType : "SSHAuthType_Password",
      ValidationType: "ValidationType_Contain",
      ValidationValue: "",
    }
  }

	var Home = {
    Processing:ko.observable(true),
    Mode: ko.observable(true),
    ModeAdd: ko.observable("Service"),
    ModeSave: ko.observable("Save"),
    IndexService: ko.observable(),
    RecordServices: ko.observableArray([]),
    RecordServiceNew: ko.mapping.fromJS(modelServiceNew),
    ArrService: ko.observableArray([]),
    filterKeyword: ko.observable(''),
    PingType: ko.observableArray(['PingType_Network','PingType_HttpStatus','PingType_HttpBody','PingType_Command','PingType_Custom']),
    ResponseType: ko.observableArray(['Response_Contains','Response_Equals','Response_RegEx']),
    ExedCommandType: ko.observableArray(['CommandType_Local','CommandType_SSH','CommandType_REST']),
    ValidationType: ko.observableArray(['ValidationType_Contain','ValidationType_Equal','ValidationType_Regex']),
    SshAuthType: ko.observableArray(['SSHAuthType_Password','SSHAuthType_Certificate']),
    TypeServiceLog: ko.observableArray(['Daily','Monthly','Yearly']),
    HttpBodyType: ko.observableArray(['HttpBody_Contains','HttpBody_Equals']),
    RESTAuthType : ko.observableArray(['RESTAuthType_None','RESTAuthType_Basic']),
    DateFilter: ko.observable(''),
    DateStatus: ko.observable(),
    DateStatusString: ko.observable(),
	}

  Home.gridColumns = ko.observableArray([
    {
      field: "Date", title: "",
      headerAttributes: { style: 'text-align: center' },
      headerTemplate: '<label class="gridHeaderLabel">Date</label>',
    },
    {
      field: "Type", title: "",
      headerAttributes: { style: 'text-align: center' },
      headerTemplate: '<label class="gridHeaderLabel">Type</label>',
    },
    {
      field: "Log", title: "",
      headerAttributes: { style: 'text-align: center' },
      headerTemplate: '<label class="gridHeaderLabel">Log</label>',
    },
  ]);

  function diffDateTime(earlierdate){
    // var difference = laterdate.getTime() - earlierdate.getTime();
    var datenow = new Date(), dateval = new Date(earlierdate),difference = datenow.getTime() - dateval.getTime();
 
    var daysDifference = Math.floor(difference/1000/60/60/24);
    difference -= daysDifference*1000*60*60*24;
 
    var hoursDifference = Math.floor(difference/1000/60/60);
    difference -= hoursDifference*1000*60*60;
 
    var minutesDifference = Math.floor(difference/1000/60);
    difference -= minutesDifference*1000*60;
 
    var secondsDifference = Math.floor(difference/1000);

    return daysDifference + 'd ' + hoursDifference + 'h ' + minutesDifference + 'm ' + secondsDifference + 's '
  }

</script>

<script id="gridService" type="text/html">
    <div class="grid__item" href="#" data-bind="attr:{indexGrid:$index}">
      <div class="content-itemgrid">
        <div class="col-md-6 item-removegrid item-headerleft" style="text-align:left;">
          <span class="glyphicon glyphicon-play btneditgrid" data-bind="click:function(){Home.DetailService(Service.ID(),'Start', $index())}"></span>
          <span class="glyphicon glyphicon-stop btneditgrid" data-bind="click:function(){Home.ServiceStop(Service.ID())}"></span>
          <span class="glyphicon glyphicon-list-alt" data-bind="click:function(){Home.DetailService(Service.ID(),'Log', $index())}"></span>
        </div>
        <div class="col-md-6 item-removegrid item-headerright">
          <span class="glyphicon glyphicon-pencil btneditgrid" data-bind="click:function(){Home.DetailService(Service.ID(),'Grid', $index())}"></span>
          <span class="glyphicon glyphicon-remove" data-bind="click:function(){Home.RemoveService(Service.ID())}"></span>
        </div>
        <div class="item-bodygrid" data-bind="click:function(){Home.DetailService(Service.ID(), 'Detail', $index())}">
          <h2 class="title title--preview" data-bind="text:Service.Title"></h2>
          <div class="loader"></div>
          <span class="category" data-bind="text:Ping.Host() + ' : ' + Ping.Port()"></span>
          <div class="meta meta--preview">
            <img class="meta__avatar" src="/static/img/symbol_check.png" data-bind="attr:{src:Service.StatusService() == 'Start' ? '/static/img/symbol_check.png' : '/static/img/stop1normalred.png'}" alt="author01" width="50" height="50" />
            <span class="meta__avatar glyphicon glyphicon-info-sign" style="border-radius: 20px; width:20px; height:20px; font-size:20px; margin:1em auto;" data-bind="visible:Service.LogStatus() == 'Fail' || Service.LogStatus() == 'Preparing', style : {color: Service.LogStatus() == 'Fail' ? '#F73434':'#F0ED46'}"></span>
            <span class="meta__date"><i class="fa fa-calendar-o"></i> <span data-bind="text:moment(Service.DateStatus()).format('DD MMM YYYY')">9 Apr</span></span>
            <span class="meta__reading-time"><i class="fa fa-clock-o"></i> <span data-bind="text:diffDateTime(Service.DateStatus())"> </span></span>
          </div>
        </div>
      </div>
    </div>
</script>
<!-- F73434 && F0ED46 -->

<div class="panel panel-warning" data-bind="with:Home">
	<div class="panel-body">
    <div class="row" data-bind="visible:Processing()">
        <div class="col-md-12 align-center">
        {{template "shared/processing.tpl"}}
        </div>
    </div>
  </div>
  <div class="row" data-bind="visible:!Processing() && Mode()" id="gridservice-list">
    <!-- Content Grid Service -->
  	<div class="col-md-12 subCenter">
			<input type="text" class="form-control form-filter" placeholder="Search Service Here !!" data-bind="value:filterKeyword, valueUpdate: 'keyup'" />
			<button class="btn btn-primary btn-sm" data-bind="click:AddService"><span class="glyphicon glyphicon-plus"></span>&nbsp;Add Service</button>
  	</div>
  	<div class="col-md-12" style="margin-top:10px; margin-bottom:10px;">
  		<!-- Content Grid Service -->
      <div class="container">
        <div id="theGrid" class="main">
          <section class="grid" data-bind="template:{name:'gridService', foreach:ArrService}">
          </section>
        </div>
      </div>
  	</div>
    <!-- End Content Grid Service -->
  </div>
  <div class="row" data-bind="visible:!Processing() && !Mode(), with:RecordServiceNew" id="logservice-list">
    <!-- Content Log -->
    <div class="col-md-12 header-log">
      <div class="contentbtnback" data-bind="click:Home.BackGrid">
        <span class="glyphicon glyphicon-chevron-left btnbackservice"></span> <div>Back</div>
      </div>
      <div class="header-titlelog">
        <span data-bind="text:Service.Title"></span>
      </div>
    </div>
    <div class="col-md-11 col-sm-offset-1 log-detail" data-bind="text:'Log for '+Service.Title()+' '+Ping.Host()+' : '+Ping.Port()+' ,'+Service.Description()"></div>
    <div class="col-md-12" class="content-log">
      <div class="col-md-2 log-status">
        <div class="col-md-12 btnfilterlog">
          <label class="col-md-4">Date</label>
          <div class="col-md-8 periodfilter">
              <input name="txtfilterdate" type="text" data-bind="kendoDatePicker:{format: 'dd MMM yyyy', parseFormat: 'dd MMM yyyy',value:Home.DateFilter, change:function(){Home.GetLogService(Service.ID())}}" />
          </div>
        </div>
        <!-- <div class="col-md-12 btnfilterlog"><button class="btn btn-sm btn-primary" data-bind="click:function(){Home.GetLogService(Service.ID())}"><span class="glyphicon glyphicon-search"></span> Filter</button></div> -->
        <span class="col-md-12" style="padding-bottom:10px;padding-top:10px;">Status</span>
        <span>
          <img class="meta__avatar" src="/static/img/symbol_check.png" data-bind="attr:{src:Service.StatusService() == 'Start' ? '/static/img/symbol_check.png' : '/static/img/stop1normalred.png'}" alt="author01" width="60" height="60" />
        </span>
        <span class="col-md-12">Live</span>
        <span class="col-md-12" data-bind="text:Home.DateStatus"></span>
      </div>
      <div class="col-md-9" id="gridlog" data-bind="kendoGrid:{data:[], filterable:true,pageable: true, groupable:false, columns:Home.gridColumns}"></div>
    </div>
    <!-- End Content Log -->
  </div>
</div>

<div class="modal fade" id="modalAddService" tabindex="-1" role="dialog" aria-labelledby="myModalLabel">
  <div class="modal-dialog" role="document">
    <div class="modal-content" data-bind="with:Home">
      <div class="modal-header">
        <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
        <h4 class="modal-title" id="myModalLabel">Add Service</h4>
      </div>
      <div class="modal-body" data-bind="with:RecordServiceNew">
        <div class="col-md-12 subCenter">
          <span class="service-titleadd" data-bind="text:Home.ModeAdd"></span>
        </div>
        <div class="col-md-12" data-bind="visible:Home.ModeAdd() == 'Service'">
          <div class="row">
            <label class="col-md-3 filter-label">Title</label>
            <div class="col-md-8">
                <input name="txtTitle" class="form-input form-control" type="text" data-bind="value:Service.Title" />
            </div>
          </div>
          <div class="row">
            <label class="col-md-3 filter-label">Description</label>
            <div class="col-md-8">
                <input name="txtDescription" class="form-input form-control" type="text" data-bind="value:Service.Description" />
            </div>
          </div>

          <div class="row">
            <label class="col-md-3 filter-label">Restart</label>
            <div class="col-md-8">
                <input name="txtRestart" class="form-input form-control" type="text" data-bind="value:Service.RestartAfterNCritical" />
            </div>
          </div>

          <div class="row">
            <label class="col-md-3 filter-label">Interval</label>
            <div class="col-md-8 dd-ping">
                <input name="txtInterval" class="form-input form-control" type="text" data-bind="value:Service.Interval" />
            </div>
          </div>

          <div class="row">
            <label class="col-md-3 filter-label">Email Warning</label>
            <div class="col-md-8 dd-ping">
              <input name="txtemailwarning" id="txtemailwarning" class="form-input form-control"/>
            </div>
          </div>

          <div class="row">
            <label class="col-md-3 filter-label">Email Error</label>
            <div class="col-md-8">
              <input name="txtemailerror" id="txtemailerror" class="form-input form-control"/>
            </div>
          </div>
        </div>
        <div class="col-md-12" data-bind="visible:Home.ModeAdd() == 'Ping'">
          <div class="row">
            <label class="col-md-2 filter-label">Type</label>
            <div class="col-md-9 dd-ping">
              <input name="ddtype" style="width:100%" data-bind="kendoDropDownList:{data:Home.PingType, value:Ping.Type}"/>
            </div>
          </div>

          <div class="row" data-bind="visible:Ping.Type() == 'PingType_Network' || Ping.Type() == 'PingType_HttpStatus' || Ping.Type() == 'PingType_HttpBody'">
            <label class="col-md-2 filter-label">Host</label>
            <div class="col-md-4">
                <input name="txtHost" class="form-input form-control" type="text" data-bind="value:Ping.Host" />
            </div>
            <label class="col-md-2 filter-label">Port</label>
            <div class="col-md-3">
                <input name="txtPort" class="form-input form-control" type="text" data-bind="value:Ping.Port" />
            </div>
          </div>

          <div class="row" data-bind="visible:Ping.Type() == 'PingType_Command'">
            <label class="col-md-2 filter-label">Command</label>
            <div class="col-md-9">
              <input name="txtcommand" class="form-input form-control" data-bind="value:Ping.Command"/>
            </div>
          </div>

          <div class="row" data-bind="visible:Ping.Type() == 'PingType_Command'">
            <label class="col-md-2 filter-label">Cmd.Parm</label>
            <div class="col-md-9">
              <input name="txtcommandparm" id="txtcommandparmping" class="form-input form-control"/>
            </div>
          </div>

          <div class="row" data-bind="visible:Ping.Type() == 'PingType_Command'">
            <label class="col-md-2 filter-label">Resp.Type</label>
            <div class="col-md-9 dd-ping">
              <input name="ddresponsetype" style="width:100%" data-bind="kendoDropDownList:{data:Home.ResponseType, value:Ping.ResponseType}"/>
            </div>
          </div>

          <div class="row" data-bind="visible:Ping.Type() == 'PingType_Command'">
            <label class="col-md-2 filter-label">Resp.Value</label>
            <div class="col-md-9">
              <input name="txtresponsevalue" class="form-input form-control" data-bind="value:Ping.ResponseValue"/>
            </div>
          </div>

          <div class="row" data-bind="visible:Ping.Type() == 'PingType_HttpBody'">
            <label class="col-md-2 filter-label">Http Type</label>
            <div class="col-md-9 dd-ping">
              <input name="ddresponsetype" style="width:100%" data-bind="kendoDropDownList:{data:Home.HttpBodyType, value:Ping.HttpBodyType}"/>
            </div>
          </div>

          <div class="row" data-bind="visible:Ping.Type() == 'PingType_HttpBody'">
            <label class="col-md-2 filter-label">Http Value</label>
            <div class="col-md-9">
              <input name="txtresponsevalue" class="form-input form-control" data-bind="value:Ping.HttpBodySearch"/>
            </div>
          </div>

        </div>

        <!-- Exec Start -->
        <div class="col-md-12" data-bind="visible:Home.ModeAdd() == 'EXEC COMMAND START'">
          <div class="row">
            <label class="col-md-3 filter-label">Type</label>
            <div class="col-md-8 dd-ping">
              <input name="ddexedtype" style="width:100%" data-bind="kendoDropDownList:{data:Home.ExedCommandType, value: ExedCommandStart.Type}"/>
            </div>
          </div>

          <div class="row" data-bind="visible:ExedCommandStart.Type() == 'CommandType_Local'">
            <label class="col-md-3 filter-label">Command Text</label>
            <div class="col-md-8">
              <input name="txtcommand" class="form-input form-control" data-bind="value:ExedCommandStart.CommandText"/>
            </div>
          </div>

          <div class="row" data-bind="visible:ExedCommandStart.Type() == 'CommandType_SSH'">
            <label class="col-md-3 filter-label">Command Start</label>
            <div class="col-md-8">
              <input name="txtcommandstart" class="form-input form-control" data-bind="value:ExedCommandStart.CommandText"/>
            </div>
          </div>

          <div class="row" data-bind="visible:ExedCommandStart.Type() == 'CommandType_Local'">
            <label class="col-md-3 filter-label">Command Start</label>
            <div class="col-md-8">
              <input name="txtcommandparm" id="txtcommandparmexedstart" class="form-input form-control"/>
            </div>
          </div>

          <div class="row" data-bind="visible:ExedCommandStart.Type() == 'CommandType_REST'">
            <label class="col-md-3 filter-label">Url Start</label>
            <div class="col-md-8">
              <input name="txtresturl" class="form-input form-control" data-bind="value:ExedCommandStart.RestUrl"/>
            </div>
          </div>

          <div class="row" data-bind="visible:ExedCommandStart.Type() == 'CommandType_REST'">
            <label class="col-md-3 filter-label">Rest Method</label>
            <div class="col-md-8">
              <input name="txtrestmethod" class="form-input form-control" data-bind="value:ExedCommandStart.RestMenthod"/>
            </div>
          </div>

          <div class="row" data-bind="visible:ExedCommandStart.Type() == 'CommandType_REST'">
            <label class="col-md-3 filter-label">Auth Type</label>
            <div class="col-md-8 dd-ping">
              <input name="ddrestaunthtype" style="width:100%" data-bind="kendoDropDownList:{data:Home.RESTAuthType, value: ExedCommandStart.RestAuthType}"/>
            </div>
          </div>

          <div class="row" data-bind="visible:ExedCommandStart.Type() == 'CommandType_REST' && ExedCommandStart.RestAuthType() == 'RESTAuthType_Basic'">
            <label class="col-md-3 filter-label">Rest User</label>
            <div class="col-md-3">
                <input name="txtrestuser" class="form-input form-control" type="text" data-bind="value:ExedCommandStart.RestUser" />
            </div>
            <label class="col-md-2 filter-label">Rest Pass</label>
            <div class="col-md-3">
                <input name="txtrestpass" class="form-input form-control" type="password" data-bind="value:ExedCommandStart.RestPassword" />
            </div>
          </div>

          <div class="row" data-bind="visible:ExedCommandStart.Type() == 'CommandType_SSH'">
            <label class="col-md-3 filter-label">SSH Host</label>
            <div class="col-md-3">
                <input name="txtSshHost" class="form-input form-control" type="text" data-bind="value:ExedCommandStart.SshHost" />
            </div>
            <label class="col-md-2 filter-label">SSH Port</label>
            <div class="col-md-3">
                <input name="txtSshPort" class="form-input form-control" type="text" data-bind="value:ExedCommandStart.SshPort" />
            </div>
          </div>

          <div class="row" data-bind="visible:ExedCommandStart.Type() == 'CommandType_SSH'">
            <label class="col-md-3 filter-label">Auth Type</label>
            <div class="col-md-8 dd-ping">
              <input name="ddsshauthtype" style="width:100%" data-bind="kendoDropDownList:{data:Home.SshAuthType, value: ExedCommandStart.SshAuthType}"/>
            </div>
          </div>

          <div class="row" data-bind="visible:ExedCommandStart.Type() == 'CommandType_SSH'">
            <label class="col-md-3 filter-label">SSH User</label>
            <div class="col-md-3">
                <input name="txtsshuser" class="form-input form-control" type="text" data-bind="value:ExedCommandStart.SshUser" />
            </div>
            <label class="col-md-2 filter-label" data-bind="visible:ExedCommandStart.SshAuthType() == 'SSHAuthType_Password'">SSH Pass</label>
            <div class="col-md-3" data-bind="visible:ExedCommandStart.SshAuthType() == 'SSHAuthType_Password'">
                <input name="txtsshpassword" class="form-input form-control" type="password" data-bind="value:ExedCommandStart.SshPassword" />
            </div>
            <label class="col-md-2 filter-label" data-bind="visible:ExedCommandStart.SshAuthType() != 'SSHAuthType_Password'">Key Loc</label>
            <div class="col-md-3" data-bind="visible:ExedCommandStart.SshAuthType() != 'SSHAuthType_Password'">
                <input name="txtsshkeyloc" class="form-input form-control" type="text" data-bind="value:ExedCommandStart.SshKeyLocation" />
            </div>
          </div>

          <div class="row">
            <label class="col-md-3 filter-label">Validation Start</label>
            <div class="col-md-8 dd-ping">
              <input name="ddvaltype" style="width:100%" data-bind="kendoDropDownList:{data:Home.ValidationType, value: ExedCommandStart.ValidationType}"/>
            </div>
          </div>

          <div class="row">
            <label class="col-md-3 filter-label">Value Start</label>
            <div class="col-md-8">
              <input name="txtvalidationvalue" class="form-input form-control" data-bind="value:ExedCommandStart.ValidationValue"/>
            </div>
          </div>
        </div>

        <!-- Exed STOP -->
        <div class="col-md-12" data-bind="visible:Home.ModeAdd() == 'EXEC COMMAND STOP'">
          <div class="row">
            <label class="col-md-3 filter-label">Type</label>
            <div class="col-md-8 dd-ping">
              <input name="ddexedtype" style="width:100%" data-bind="kendoDropDownList:{data:Home.ExedCommandType, value: ExedCommandStop.Type}"/>
            </div>
          </div>

          <div class="row" data-bind="visible:ExedCommandStop.Type() == 'CommandType_Local'">
            <label class="col-md-3 filter-label">Command Text</label>
            <div class="col-md-8">
              <input name="txtcommand" class="form-input form-control" data-bind="value:ExedCommandStop.CommandText"/>
            </div>
          </div>

          <div class="row" data-bind="visible:ExedCommandStop.Type() == 'CommandType_SSH'">
            <label class="col-md-3 filter-label">Command Stop</label>
            <div class="col-md-8">
              <input name="txtcommandstop" class="form-input form-control" data-bind="value:ExedCommandStop.CommandText"/>
            </div>
          </div>

          <div class="row" data-bind="visible:ExedCommandStop.Type() == 'CommandType_Local'">
            <label class="col-md-3 filter-label">Command Stop</label>
            <div class="col-md-8">
              <input name="txtcommandparm" id="txtcommandparmexedstop" class="form-input form-control"/>
            </div>
          </div>

          <div class="row" data-bind="visible:ExedCommandStop.Type() == 'CommandType_REST'">
            <label class="col-md-3 filter-label">Url Stop</label>
            <div class="col-md-8">
              <input name="txtresturl" class="form-input form-control" data-bind="value:ExedCommandStop.RestUrl"/>
            </div>
          </div>

          <div class="row" data-bind="visible:ExedCommandStop.Type() == 'CommandType_REST'">
            <label class="col-md-3 filter-label">Rest Method</label>
            <div class="col-md-8">
              <input name="txtrestmethod" class="form-input form-control" data-bind="value:ExedCommandStop.RestMenthod"/>
            </div>
          </div>

          <div class="row" data-bind="visible:ExedCommandStop.Type() == 'CommandType_REST'">
            <label class="col-md-3 filter-label">Auth Type</label>
            <div class="col-md-8 dd-ping">
              <input name="ddrestaunthtype" style="width:100%" data-bind="kendoDropDownList:{data:Home.RESTAuthType, value: ExedCommandStop.RestAuthType}"/>
            </div>
          </div>

          <div class="row" data-bind="visible:ExedCommandStop.Type() == 'CommandType_REST' && ExedCommandStop.RestAuthType() == 'RESTAuthType_Basic'">
            <label class="col-md-3 filter-label">Rest User</label>
            <div class="col-md-3">
                <input name="txtrestuser" class="form-input form-control" type="text" data-bind="value:ExedCommandStop.RestUser" />
            </div>
            <label class="col-md-2 filter-label">Rest Pass</label>
            <div class="col-md-3">
                <input name="txtrestpass" class="form-input form-control" type="password" data-bind="value:ExedCommandStop.RestPassword" />
            </div>
          </div>

          <div class="row" data-bind="visible:ExedCommandStop.Type() == 'CommandType_SSH'">
            <label class="col-md-3 filter-label">SSH Host</label>
            <div class="col-md-3">
                <input name="txtSshHost" class="form-input form-control" type="text" data-bind="value:ExedCommandStop.SshHost" />
            </div>
            <label class="col-md-2 filter-label">SSH Port</label>
            <div class="col-md-3">
                <input name="txtSshPort" class="form-input form-control" type="text" data-bind="value:ExedCommandStop.SshPort" />
            </div>
          </div>

          <div class="row" data-bind="visible:ExedCommandStop.Type() == 'CommandType_SSH'">
            <label class="col-md-3 filter-label">Auth Type</label>
            <div class="col-md-8 dd-ping">
              <input name="ddsshauthtype" style="width:100%" data-bind="kendoDropDownList:{data:Home.SshAuthType, value: ExedCommandStop.SshAuthType}"/>
            </div>
          </div>

          <div class="row" data-bind="visible:ExedCommandStop.Type() == 'CommandType_SSH'">
            <label class="col-md-3 filter-label">SSH User</label>
            <div class="col-md-3">
                <input name="txtsshuser" class="form-input form-control" type="text" data-bind="value:ExedCommandStop.SshUser" />
            </div>
            <label class="col-md-2 filter-label" data-bind="visible:ExedCommandStop.SshAuthType() == 'SSHAuthType_Password'">SSH Pass</label>
            <div class="col-md-3" data-bind="visible:ExedCommandStop.SshAuthType() == 'SSHAuthType_Password'">
                <input name="txtsshpassword" class="form-input form-control" type="password" data-bind="value:ExedCommandStop.SshPassword" />
            </div>
            <label class="col-md-2 filter-label" data-bind="visible:ExedCommandStop.SshAuthType() != 'SSHAuthType_Password'">Key Loc</label>
            <div class="col-md-3" data-bind="visible:ExedCommandStop.SshAuthType() != 'SSHAuthType_Password'">
                <input name="txtsshkeyloc" class="form-input form-control" type="text" data-bind="value:ExedCommandStop.SshKeyLocation" />
            </div>
          </div>

          <div class="row">
            <label class="col-md-3 filter-label">Validation Type</label>
            <div class="col-md-8 dd-ping">
              <input name="ddvaltype" style="width:100%" data-bind="kendoDropDownList:{data:Home.ValidationType, value: ExedCommandStop.ValidationType}"/>
            </div>
          </div>

          <div class="row">
            <label class="col-md-3 filter-label">Value Stop</label>
            <div class="col-md-8">
              <input name="txtvalidationvalue" class="form-input form-control" data-bind="value:ExedCommandStop.ValidationValue"/>
            </div>
          </div>
        </div>
        <div style="clear:both"></div>
        <!-- End Content Add Service -->
      </div>
      <div class="modal-footer">
        <button type="button" class="btn btn-default" data-dismiss="modal"><span class="glyphicon glyphicon-floppy-remove"></span> Cancel</button>
        <button type="button" class="btn btn-warning" data-bind="visible:Home.ModeAdd() == 'Ping' || Home.ModeAdd() != 'Service' || Home.ModeAdd() == 'EXEC COMMAND START',click:function(){NextBackAdd('Back')}"><span class="glyphicon glyphicon-chevron-left"></span> Back</button>
        <button type="button" class="btn btn-primary" data-bind="visible:Home.ModeAdd() == 'Ping' || Home.ModeAdd() == 'Service' || Home.ModeAdd() == 'EXEC COMMAND START', click:function(){NextBackAdd('Next')}"><span class="glyphicon glyphicon-chevron-right"></span> Next</button>
        <button type="button" class="btn btn-primary" data-bind="visible:Home.ModeAdd() != 'Ping' && Home.ModeAdd() != 'Service' && Home.ModeAdd() != 'EXEC COMMAND START',click:SaveService"><span class="glyphicon glyphicon-floppy-saved"></span> <span class="titleSave" data-bind="text:Home.ModeSave">Save</span></button>
      </div>
    </div>
  </div>
</div>

<div class="modal fade" id="modalDetailService" tabindex="-1" role="dialog" aria-labelledby="myModalLabel">
  <div class="modal-dialog" role="document">
    <div class="modal-content" data-bind="with:Home">
      <div class="modal-header">
        <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
        <h4 class="modal-title" id="myModalLabel">Add Service</h4>
      </div>
      <div class="modal-body" data-bind="with:RecordServiceNew">
        <!-- Content Detail Service -->
        <div class="row">
          <div class="col-md-12">
            <button class="btn btn-sm btn-success space-right" data-bind="click:function(){Home.DetailService(Service.ID(),'Start', Home.IndexService())}"><span class="glyphicon glyphicon-play"></span> Start</button>
            <button class="btn btn-sm btn-danger space-right" data-bind="click:function(){Home.ServiceStop(Service.ID())}"><span class="glyphicon glyphicon-stop"></span> Stop</button>
            <button class="btn btn-sm btn-danger space-right" data-bind="click:function(){Home.StopServer(Service.ID())}"><span class="glyphicon glyphicon-stop"></span> Stop Server</button>
            <button class="btn btn-sm btn-warning space-right" data-bind="click:Home.Log"><span class="glyphicon glyphicon-list-alt"></span> Log</button>
            <button class="btn btn-sm btn-info space-right" data-bind="click:function(){Home.EditService('Detail')}"><span class="glyphicon glyphicon-pencil"></span> Update</button>
            <button class="btn btn-sm btn-danger space-right" data-bind="click:function(){Home.RemoveService(Service.ID())}"><span class="glyphicon glyphicon-trash"></span> Remove</button>
          </div>
          <div class="col-md-12">
            <div class="col-md-12 subCenter">
              <h3 data-bind="text:Service.Title"></h3>
            </div>
            <div class="col-md-12">
              <img class="meta__avatar" src="/static/img/symbol_check.png" data-bind="attr:{src:Service.StatusService() == 'Start' ? '/static/img/symbol_check.png' : '/static/img/stop1normalred.png'}" alt="author01" width="60" height="60" />
            </div>
            <div class="col-md-12 subCenter">
              <label data-bind="text:Service.Description() + ' ' + Ping.Host() + ' : ' + Ping.Port()"></label><br/>
              <label data-bind="text:Home.DateStatus"></label>
            </div>
          </div>
        </div>
        <div style="clear:both"></div>
        <!-- End Content Detail Service -->
      </div>
    </div>
  </div>
</div>

<script src="/static/gridlib/js/classie.js"></script>
<!-- <script src="/static/gridlib/js/main.js"></script> -->

<script>
  Home.validationField = function(val){
    if (val === 'Save'){
      var boolVal = true;
      if (Home.RecordServiceNew.Service.Title() == "")
        boolVal = false;
      return boolVal;
    } else {
      // Service
      var ServiceVal = new Array(), PingVal = new Array(), ExecStartVal = new Array(), ExecStopVal = new Array(), typePing = Home.RecordServiceNew.Ping.Type(), typeExecStart = Home.RecordServiceNew.ExedCommandStart.Type(), typeExecStop = Home.RecordServiceNew.ExedCommandStop.Type(), StringValidation = "";
      $.each( Home.RecordServiceNew.Service, function( key, value ) {
        if (key !== 'EmailWarning' && key !== 'EmailError' && key !== 'RestartAfterNCritical' && key !== 'Interval' && key !== 'PathLog' && key !== 'StatusService'){
          if (value() === '')
            ServiceVal.push(key);
        } else if (key === 'EmailWarning' || key == 'EmailError'){
          if (value().length == 0)
            ServiceVal.push(key);
        } else if (key === 'RestartAfterNCritical' || key === 'Interval'){
          if (value() == 0)
            ServiceVal = 0;
        }
      });
      // Ping
      $.each( Home.RecordServiceNew.Ping, function( key, value ) {
        if (typePing === 'PingType_Network' || typePing === 'PingType_HttpStatus' && key != 'Type'){
          if (key === 'Host' || key === 'Port'){
            if (value() === '')
              PingVal.push(key);
          }
        } else if (typePing === 'PingType_HttpBody' && key != 'Type'){
          if (key === 'Host' || key === 'Port' || key === 'HttpBodySearch'){
            if (value() === '')
              PingVal.push(key);
          }
        } else if (typePing === 'PingType_Command' && key != 'Type'){
          if (key === 'Command' || key === 'ResponseValue'){
            if (value() === '')
              PingVal.push(key);
          }
          else if (key === 'CommandParm' && value().length === 0)
            PingVal.push(key);
        }
      });
      // Exec Start
      $.each( Home.RecordServiceNew.ExedCommandStart, function( key, value ) {
        if (typeExecStart === 'CommandType_Local' && key != 'Type'){
          if (key === 'CommandText' && value() === '')
            ExecStartVal.push(key);
          else if (key === 'CommandParm' && value().length === 0)
            ExecStartVal.push(key);
        } else if (typeExecStart === 'CommandType_SSH' && key != 'Type'){
          if (key === 'CommandTextSsh' || key === 'SshHost' || key === 'SshPort'){
            if (value() === '')
              ExecStartVal.push(key);
          }
          else if (key === 'SshAuthType' && value() === 'SSHAuthType_Password'){
            if (key === 'SshPassword' && value() === '')
              ExecStartVal.push(key);
          } else if (key === 'SshAuthType' && value() === 'SSHAuthType_Certificate'){
            if (key === 'SshKeyLocation' && value() === '')
              ExecStartVal.push(key);
          }
        } else if (typeExecStart === 'CommandType_REST' && key != 'Type'){
          if (key === 'RestUrl' || key === 'RestMenthod'){
            if (value() === '')
              ExecStartVal.push(key);
          }
          else if (key === 'RestAuthType' && value() !== 'RESTAuthType_None'){
            if (key === 'RestUser' || key === 'RestPassword'){
              if (value() === '')
                ExecStartVal.push(key);
            }
          }
        }
        if (key === 'ValidationValue' && value() === '')
          ExecStartVal.push(key);
      });
      // Exec Stop
      $.each( Home.RecordServiceNew.ExedCommandStop, function( key, value ) {
        if (typeExecStart === 'CommandType_Local' && key != 'Type'){
          if (key === 'CommandText' && value() === '')
            ExecStopVal.push(key);
          else if (key === 'CommandParm' && value().length === 0)
            ExecStopVal.push(key);
        } else if (typeExecStart === 'CommandType_SSH' && key != 'Type'){
          if (key === 'CommandTextSsh' || key === 'SshHost' || key === 'SshPort'){
            if (value() === '')
              ExecStopVal.push(key);
          }
          else if (key === 'SshAuthType' && value() === 'SSHAuthType_Password'){
            if (key === 'SshPassword' && value() === '')
              ExecStopVal.push(key);
          } else if (key === 'SshAuthType' && value() === 'SSHAuthType_Certificate'){
            if (key === 'SshKeyLocation' && value() === '')
              ExecStopVal.push(key);
          }
        } else if (typeExecStart === 'CommandType_REST' && key != 'Type'){
          if (key === 'RestUrl' || key === 'RestMenthod'){
            if (value() === '')
              ExecStopVal.push(key);
          }
          else if (key === 'RestAuthType' && value() !== 'RESTAuthType_None'){
            if (key === 'RestUser' || key === 'RestPassword'){
              if (value() === '')
                ExecStopVal.push(key);
            }
          }
        }
        if (key === 'ValidationValue' && value() === '')
          ExecStopVal.push(key);
      });

      if (ServiceVal.length > 0){
        StringValidation += "Service : " +ServiceVal.join(",") + "\n";
      }
      if (PingVal.length > 0){
        StringValidation += "Ping : " + PingVal.join(",") + "\n";
      }
      if (ExecStartVal.length > 0){
        StringValidation += "Exec Start : " + ExecStartVal.join(",") + "\n";
      }
      if (ExecStopVal.length > 0){
        StringValidation += "Exec Stop : " + ExecStopVal.join(",") + "\n";
      }
      return StringValidation;
    }
  }
  Home.AddService = function(){
    Home.ModeSave('Save');
    ko.mapping.fromJS(modelServiceNew, Home.RecordServiceNew);
    $("#txtcommandparmping").tokenInput('clear');
    $("#txtcommandparmexedstart").tokenInput('clear');
    $("#txtcommandparmexedstop").tokenInput('clear');
    $("#txtemailerror").tokenInput('clear');
    $("#txtemailwarning").tokenInput('clear');
    Home.ModeAdd('Service');
    $('#modalAddService').modal('show');
  }
  Home.SaveService = function(){
    $('#modalAddService').modal('hide');
    var validationSave = Home.validationField('Save');
    if(Home.RecordServiceNew.Service.StatusService() != 'Start' && validationSave == true){
      Home.Processing(true);
      var url = "/home/addservice";
      if(Home.ModeSave() === 'Update')
        url = "/home/updateservice";
      // Home.RecordServiceNew.Ping.CommandParmString(Home.RecordServiceNew.Ping.CommandParm().join(","));
      // Home.RecordServiceNew.ExedCommand.CommandParmString(Home.RecordServiceNew.ExedCommand.CommandParms().join(","));
      if (Home.RecordServiceNew.Ping.Type()==='PingType_Command'){
        Home.RecordServiceNew.Ping.Host('localhost');
        Home.RecordServiceNew.Ping.Port('-');
      }
      for(var key in $("#txtcommandparmping").tokenInput('get')){
        Home.RecordServiceNew.Ping.CommandParm.push($("#txtcommandparmping").tokenInput('get')[key].name);
      }
      for(var key in $("#txtcommandparmexedstart").tokenInput('get')){
        Home.RecordServiceNew.ExedCommandStart.CommandParm.push($("#txtcommandparmexedstart").tokenInput('get')[key].name);
      }
      for(var key in $("#txtcommandparmexedstop").tokenInput('get')){
        Home.RecordServiceNew.ExedCommandStop.CommandParm.push($("#txtcommandparmexedstop").tokenInput('get')[key].name);
      }
      for(var key in $("#txtemailwarning").tokenInput('get')){
        Home.RecordServiceNew.Service.EmailWarning.push($("#txtemailwarning").tokenInput('get')[key].name);
      }
      for(var key in $("#txtemailerror").tokenInput('get')){
        Home.RecordServiceNew.Service.EmailError.push($("#txtemailerror").tokenInput('get')[key].name);
      }
      $.ajax({
        url: url,
        type: 'post',
        // dataType: 'json',
        contentType: "application/json; charset=utf-8",
        data : ko.mapping.toJSON(Home.RecordServiceNew),
        success : function(res) {
          if(res.success){
            Home.Processing(false);
            Home.GetService();
          }else{
            alert(res.message);
            Home.Processing(false);
          }
        },
      });
    } else {
      if (validationSave == false)
        alert('Please Input Service Title !');
      else
        alert('You must stop service before update service !');
    }
  }
  Home.GetService = function(){
    Home.Processing(true);
    var url = "/home/getservice";
    $.ajax({
      url: url,
      type: 'post',
      dataType: 'json',
      data : {Statuslive:"Start"},
      success : function(res) {
        if(res.success){
          Home.Processing(false);
          Home.RecordServices(_.map(res.data, function (r) { return ko.mapping.fromJS(r); }));
          // for (var key in res.data){
          //   if (res.data[key].Service.StatusService === 'Start'){
          //     if (res.data[key].Service.LogStatus !== 'Success' || res.data[key].Service.LogStatus !== 'OK' || res.data[key].Service.LogStatus() !== 'Fail' || res.data[key].Service.LogStatus !== 'Error'){  
          //       // console.log(Home.RecordServices()[key].Service.LogStatus());
          //       setTimeout(function() { Home.ServiceStart(res.data[key].Service.ID,'Live', key, 'Grid'); }, 1000);
          //     }
          //   }
          // }
        }else{
          alert(res.message);
          Home.Processing(false);
        }
      },
    });
  }
  Home.RemoveService = function(idService){
    if (confirm("Are you sure remove this !") == true) {
        Home.Processing(true);
        var url = "/home/removeservice";
        $.ajax({
          url: url,
          type: 'post',
          dataType: 'json',
          data : {ID: idService},
          success : function(res) {
            if(res.success){
              Home.Processing(false);
              Home.GetService();
            }else{
              alert(res.message);
              Home.Processing(false);
            }
          },
        });
    }
  }
  Home.DetailService = function(idService, valview, indexSer){
    // console.log(idService);
    Home.IndexService(indexSer);
    var url = "/home/getdetailservice";
    $.ajax({
      url: url,
      type: 'post',
      dataType: 'json',
      data : {ID: idService},
      success : function(res) {
        if(res.success){
          ko.mapping.fromJS(res.data, Home.RecordServiceNew);
          Home.DateStatus(diffDateTime(res.data.Service.DateStatus));
          Home.DateStatusString(res.data.Service.DateStatus);
          if(valview == 'Grid'){
            Home.EditService(valview);
          } else if(valview == 'Log'){
            Home.Log();
          } else if (valview == 'Start') {
            Home.ServiceStart(idService,'Start', Home.IndexService(), 'Detail');
          } else {
            $('#modalDetailService').modal('show');
          }
        }else{
          alert(res.message);
        }
      },
    });
  }
  Home.Log = function(){
    $('#modalDetailService').modal('hide');
    Home.Mode(false);
    Home.DateFilter(moment().format('DD MMM YYYY'));
    Home.GetLogService(Home.RecordServiceNew.Service.ID());
    Home.DateStatus(diffDateTime(Home.DateStatusString()));
  }
  Home.BackGrid = function(){
    Home.Mode(true);
  }
  Home.NextBackAdd = function(data){
    if(data === 'Next' && Home.ModeAdd() === 'Ping')
      Home.ModeAdd('EXEC COMMAND START');
    else if (data === 'Next' && Home.ModeAdd() === 'Service')
      Home.ModeAdd('Ping');
    else if (data === 'Next' && Home.ModeAdd() === 'EXEC COMMAND START')
      Home.ModeAdd('EXEC COMMAND STOP');
    else if (data === 'Back' && Home.ModeAdd() === 'Ping')
      Home.ModeAdd('Service');
    else if (data === 'Back' && Home.ModeAdd() === 'EXEC COMMAND STOP')
      Home.ModeAdd('EXEC COMMAND START');
    else
      Home.ModeAdd('Ping');
  }
  Home.ServiceStart = function(idService, statusLive, indexSer, statusCheck){
    // console.log(indexSer);
    var validationStart = Home.validationField('Start');
    if (validationStart == '' || statusCheck == 'Grid'){
      var url = "/home/startservice";
      $.ajax({
        url: url,
        type: 'post',
        dataType: 'json',
        data : {Status: 'Start', ID: idService, Statuslive : statusLive},
        success : function(res) {
            Home.ArrService()[indexSer].Service.LogStatus(res.data);
            if (res.data === 'OK' || res.data === 'Fail' || res.data === 'Error'){
              $('#modalDetailService').modal('hide');
              if (statusCheck != 'Grid')
                Home.GetService();
            } else {
              var StatusService = statusCheck;
              setTimeout(function() { Home.ServiceStart(idService,'Live', indexSer, StatusService); }, 1000);
            }
        },
      });
    } else {
      alert("Field Can't be Null !! \n" + validationStart);
    }
  }
  Home.ServiceStop = function(idService){
    var url = "/home/stopservice";
    $.ajax({
      url: url,
      type: 'post',
      dataType: 'json',
      data : {ID: idService, IndexService: Home.IndexService()},
      success : function(res) {
          $('#modalDetailService').modal('hide');
          Home.GetService();
      },
    });
  }
  Home.EditService = function(valview){
    Home.ModeSave('Update');
    $("#txtcommandparmping").tokenInput('clear');
    $("#txtcommandparmexedstart").tokenInput('clear');
    $("#txtcommandparmexedstop").tokenInput('clear');
    $("#txtemailerror").tokenInput('clear');
    $("#txtemailwarning").tokenInput('clear');

    if(valview === 'Detail')
      $('#modalDetailService').modal('hide');
    for (var key in Home.RecordServiceNew.Ping.CommandParm()){
      $("#txtcommandparmping").tokenInput("add", {id: Home.RecordServiceNew.Ping.CommandParm()[key], name: Home.RecordServiceNew.Ping.CommandParm()[key]});
    }
    for (var key in Home.RecordServiceNew.ExedCommandStart.CommandParm()){
      $("#txtcommandparmexedstart").tokenInput("add", {id: Home.RecordServiceNew.ExedCommandStart.CommandParm()[key], name: Home.RecordServiceNew.ExedCommandStart.CommandParm()[key]});
    }
    for (var key in Home.RecordServiceNew.ExedCommandStop.CommandParm()){
      $("#txtcommandparmexedstop").tokenInput("add", {id: Home.RecordServiceNew.ExedCommandStop.CommandParm()[key], name: Home.RecordServiceNew.ExedCommandStop.CommandParm()[key]});
    }
    for (var key in Home.RecordServiceNew.Service.EmailWarning()){
      $("#txtemailwarning").tokenInput("add", {id: Home.RecordServiceNew.Service.EmailWarning()[key], name: Home.RecordServiceNew.Service.EmailWarning()[key]});
    }
    for (var key in Home.RecordServiceNew.Service.EmailError()){
      $("#txtemailerror").tokenInput("add", {id: Home.RecordServiceNew.Service.EmailError()[key], name: Home.RecordServiceNew.Service.EmailError()[key]});
    }
    Home.ModeAdd('Service');
    $('#modalAddService').modal('show');
  }
  Home.GetLogService = function(idService){
    var url = "/home/getlogservice";
    $.ajax({
      url: url,
      type: 'post',
      dataType: 'json',
      data : {ID: idService, DateFilter: moment($('input[name=txtfilterdate]').val()).format('YYYYMMDD')},
      success : function(res) {
          var ks = res.data.split("\n"), arrLog = new Array();
          $.each(ks, function(k){
            if (ks[k] != ""){
              var a = ks[k].split(' ');
              b = ks[k].match(/\[(.*)\]/).pop();
              arrLog.push({
                Date:a[1] + ' ' + a[2],
                Type:a[0],
                Log:b,
              });
            }
          });
          var ds = new kendo.data.DataSource({
            data: arrLog,
            refresh: true,
            pageSize: 20
          });
          if ($("#gridlog").data("kendoGrid") != undefined)
              $("#gridlog").data("kendoGrid").setDataSource(ds);
      },
    });
  }
  Home.StopServer = function(){
    var url = "/home/stopserver";
    $.ajax({
      url: url,
      type: 'post',
      dataType: 'json',
      data : {ID: idService},
      success : function(res) {
          $('#modalDetailService').modal('hide');
          Home.GetService();
      },
    });
  }

  Home.ArrService = ko.computed(function () {
    var search = Home.filterKeyword();
    return ko.utils.arrayFilter(Home.RecordServices(), function (c) {
        return c.Service.Title().toLowerCase().indexOf(search.toLowerCase()) >= 0;
    }).sort(function (a, b) {
        return a.Service.Title().toLowerCase() < b.Service.Title().toLowerCase() ? -1 : a.Service.Title().toLowerCase() > b.Service.Title().toLowerCase() ? 1 : a.Service.Title().toLowerCase() == b.Service.Title().toLowerCase() ? 0 : 0;
    });
  });

	$(document).ready(function(){
    Home.GetService();
    $("#txtcommandparmping").tokenInput([], { 
      noResultsText: "Add New Command",
      theme: "facebook",
      zindex: 9999,
      allowFreeTagging: true,
    });
    $("#txtcommandparmexedstart").tokenInput([], { 
      noResultsText: "Add New Command",
      theme: "facebook",
      zindex: 9999,
      allowFreeTagging: true,
    });
    $("#txtcommandparmexedstop").tokenInput([], { 
      noResultsText: "Add New Command",
      theme: "facebook",
      zindex: 9999,
      allowFreeTagging: true,
    });

    $("#txtemailwarning").tokenInput([], { 
      noResultsText: "Add New Command",
      theme: "facebook",
      zindex: 9999,
      allowFreeTagging: true,
    });
    $("#txtemailerror").tokenInput([], { 
      noResultsText: "Add New Command",
      theme: "facebook",
      zindex: 9999,
      allowFreeTagging: true,
    });
    $(".token-input-dropdown").css({"z-index":"9999","width":"40%"});
    $(".token-input-list").addClass("form-control");
    $(".token-input-list").css("width","100%");
    Home.DateFilter(moment().format('DD MMM YYYY'));
	});
</script>