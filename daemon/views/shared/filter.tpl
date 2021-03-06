<style>
#showHideFilter{
    float:right;
    cursor: pointer;
}
</style>
<script>
    var Now = new Date();
    var Filter = {
        dataSource:ko.observableArray([]),
        startMonthly:ko.observable(new Date(Now.getFullYear(),Now.getMonth()-1,1)),
        endMonthly:ko.observable(new Date(Now.getFullYear(),Now.getMonth()+1,0)),
        state:ko.observableArray([]),
        stateList:ko.observableArray([]),
        service:ko.observableArray([]),
        serviceList:ko.observableArray([
            {Id:"Electricity",Title:"Electricity"},
            {Id:"Gas",Title:"Gas"}
        ]),
        source:ko.observable(""),
        sourceList:ko.observableArray([
            {Id:"ABS",Title:"ABS"},
            {Id:"AEMO",Title:"AEMO"},
            {Id:"Competitors",Title:"Competitors"}
        ]),
    }
    Filter.Get = function(){
        var parm = {
            StartDate:moment(Filter.startMonthly()).toISOString(),
            EndDate:moment(Filter.endMonthly()).toISOString(),
            State:Filter.state(),
            Service:Filter.service(),
            Source:Filter.source(),
        }
        return parm;
    }
    Filter.Reset = function(){
        localStorage.clear();
        Filter.Initiate();
    }
    Filter.Initiate = function(){
        Filter.ApplyLocalStorageData();
        var url = "/filter/getdata";
         $.ajax({
            url: url,
            type: 'post',
            dataType: 'json',
            data : null,
            success : function(res) {
                if(res.success){
                    for(var i in res.data.State){
                        res.data.State[i].Title = res.data.State[i].Id;
                    }
                    Filter.dataSource(res.data);
                    Filter.stateList(res.data.State);
                    Filter.ApplyLocalStorageData();
                }else{
                    alert(res.message);
                }
            },
         });
    }
    Filter.ApplyLocalStorageData = function(){
        if(localStorage.length>0){
            Filter.startMonthly(new Date(localStorage["startMonthly"]));
            Filter.endMonthly(new Date(localStorage["endMonthly"]));
            Filter.source(localStorage["source"])
            //console.log(localStorage["state"]);
            if(localStorage["state"]!==""){
            Filter.state(localStorage["state"].split(","));    
            }
            if(localStorage["service"]!==""){
             Filter.service(localStorage["service"].split(","));   
            }
        }
    }
    Filter.SaveToLocalStorage = function(){
        localStorage["startMonthly"] = Filter.startMonthly();
        localStorage["endMonthly"] = Filter.endMonthly();
        localStorage["state"] = Filter.state();
        localStorage["service"] = Filter.service();
        localStorage["source"] = Filter.source();
    }

    Filter.source.subscribe(function(){
        if(Filter.source() == ""){
            Filter.source("ABS")
        }

        if(Filter.source() == "ABS"){
            $("#ABSContent").show();
            $("#AEMOContent").hide();
            $("#CompetitiveContent").hide();
        } else if(Filter.source() == "AEMO"){
            $("#ABSContent").hide();
            $("#AEMOContent").show();
            $("#CompetitiveContent").hide();
        } else if(Filter.source() == "Competitors"){
            $("#ABSContent").hide();
            $("#AEMOContent").hide();
            $("#CompetitiveContent").show();
        }
    });
</script>
<div class="panel panel-default" data-bind="with:Filter">
    <div class="panel-heading">
        FILTER
        <a class="align-right" id="showHideFilter" data-toggle="collapse" href="#contentFilter" aria-controls="contentFilter">show / hide</a>
    </div>
    <div class="panel-body collapse in" id="contentFilter">
        <div class="row form-group">
            <div class="col-md-3">
                <label class="col-md-4 filter-label">Data Source</label>
                <div class="col-md-8">
                    <input type="text" data-bind="kendoDropDownList:{value:source,data:sourceList,dataValueField:'Id',dataTextField:'Title'}"/>
                </div>
            </div>
            <div class="col-md-3">
                <label class="col-md-4 filter-label">Start</label>
                <div class="col-md-8">
                    <input type="text" data-bind="kendoDatePicker: {value: startMonthly, start: 'year', depth: 'year', format: 'MMM yyyy'}">
                </div>
            </div>
            <div class="col-md-3">
                <label class="col-md-4 filter-label">End</label>
                <div class="col-md-8">
                   <input type="text" data-bind="kendoDatePicker: {value: endMonthly, start: 'year', depth: 'year', format: 'MMM yyyy'}">
                </div>
            </div>
            <div class="col-md-3">
                <label class="col-md-4 filter-label">State</label>
                <div class="col-md-8">
                  <input type="text" data-bind="kendoMultiSelect:{value:state,data:stateList,dataValueField:'Id',dataTextField:'Title'}"/>
                </div>
            </div>
          </div>
          <div class="row form-group">
            <div class="col-md-3">
                <label class="col-md-4 filter-label">Service</label>
                <div class="col-md-8">
                  <input type="text" data-bind="kendoMultiSelect:{value:service,data:serviceList,dataValueField:'Id',dataTextField:'Title'}"/>
                </div>
            </div>
          </div>
          <div class="row form-group" id="filter-button">
            <div class="col-md-12 align-right">
                <button type="button" class="btn btn-primary btn-sm" data-bind="click:Refresh">
                  <span class="fa fa-refresh"></span>
                  Refresh
                </button>
                <button type="button" class="btn btn-warning btn-sm"  data-bind="click:Reset">
                  <span class="fa fa-share-square"></span>
                  Reset
                </button>
            </div>
          </div>
    </div>
</div>
<script>
    $(document).ready(function(){
        Filter.Initiate();
        $(window).bind('beforeunload', function(e){
          // Filter.SaveToLocalStorage();
        });
    });
</script>