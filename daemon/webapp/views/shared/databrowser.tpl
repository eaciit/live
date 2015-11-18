
<div data-bind="with:DataBrowser">
 <div class="row" data-bind="visible:Processing()">
          <div class="col-md-12 align-center">
              {{template "shared/processing.tpl"}}
          </div>
      </div>
<div class="row" data-bind="visible:!Processing()">
    <div class="col-md-2">Fields :</div>
    <div class="col-md-10" > 
     
        <div  id="MSFields" class="min-top" data-bind="kendoMultiSelect:{value:SelectedFields,data:Fields,dataValueField:'field',dataTextField:'alias',placeholder:'Choose Fields..',filter: 'contains'}"></div>
    </div>
</div>
<div class="row" data-bind="visible:!Processing()">
    <div class="col-md-12 align-center marg-top">
        <div id="databrowsergrid"></div>
    </div>
</div>

</div>

<style type="text/css">
    .sum-header{
        color: #a8a8a8;
        background-color: #f9f9f9;
    }
    .marg-top{
        margin-top: 10px;
    }
    .min-top{
        margin-top: -5px;
    }
</style>

<script type="text/javascript">

 

function FillField(){
    var fieldx =DBFie;
    var selectedx = SelectedFie;

    if($("#HypothesisId").html()=="H16"){       
        DataBrowser().Fields([]);
        DataBrowser().SelectedFields([]); 

        for(var i in fieldx){
            DataBrowser().Fields().push(fieldx[i]);
        }

        for(var i in selectedx){
            DataBrowser().SelectedFields().push(selectedx[i]);
        }

        for(var i in CompleteBearing){
            var d = {};
            d["tipe"]= "double";
            d["alias"]= CompleteBearing[i];
            d["field"] = CompleteBearing[i];
            DataBrowser().Fields().push(d);
            DataBrowser().SelectedFields().push(CompleteBearing[i]);
        }
    }
    else
    {
        DataBrowser().Fields(fieldx);
        DataBrowser().SelectedFields(selectedx); 
    }
}


// function GenerateSUmmary(datas,typelist){
//             var dt = datas;
//                         if(dt.length>0){
//                         setTimeout(function(){
                           

//                             var strhtml = "<tr role='row'><td class='sum-header' colspan='"+ typelist.length +"' > Average </td></tr><tr>";
//                                 for(var i in typelist)
//                                 {


//                                     var td = "";
//                                     if(typelist[i].tipe=="string"||typelist[i].tipe=="date")
//                                         td+="<td role='gridcell'>-";
//                                     else
//                                     {
//                                          var sum = 0;
//                                         for(var x in dt){
//                                             if(dt[x][typelist[i].field]!=undefined)
//                                             sum+=dt[x][typelist[i].field]
//                                         }

//                                         var avg = sum/dt.length;

//                                         td+= "<td role='gridcell' class='align-right'>"+ kendo.toString(avg,'N2');
//                                     }
//                                     td+="</td>"
//                                     strhtml+=td;
//                                 }


//                             strhtml+="</tr>"
//                              $('#databrowsergrid').find('tbody').append(strhtml);

//                              var strhtml = "<tr role='row'><td class='sum-header' colspan='"+ typelist.length +"' > Total </td></tr><tr>";
//                                 for(var i in typelist)
//                                 {
//                                     var td = "";
//                                     if(typelist[i].tipe=="string"||typelist[i].tipe=="date")
//                                        td+="<td role='gridcell'>-";
//                                     else
//                                     {
//                                          var sum = 0;
//                                         for(var x in dt){
//                                             if(dt[x][typelist[i].field]!=undefined)
//                                             sum+=dt[x][typelist[i].field]
//                                         }
//                                         td+= "<td role='gridcell' class='align-right'>"+ kendo.toString(sum ,'N2');
//                                     }
//                                     td+="</td>"
//                                     strhtml+=td;
//                                 }


//                             strhtml+="</tr>"
//                              $('#databrowsergrid').find('tbody').append(strhtml);


//                         },300);
//                 }
//     }


function GenerateSumH16(dt,typelist){

    if(dt.length>0){
                        setTimeout(function(){

                           

                            var strhtml = "<tr role='row'><td class='sum-header' colspan='"+ typelist.length +"' > Average </td></tr><tr>";
                                for(var i in typelist)
                                {
                                    var td = "";
                                    if(typelist[i].tipe=="string"||typelist[i].tipe=="date")
                                        td+="<td role='gridcell'>-";
                                    else
                                    {
                                        td+= "<td role='gridcell' class='align-right'>"+ kendo.toString(dt[0][typelist[i].field.replace(".","")+"avg"] ,'N2');
                                    }
                                    td+="</td>"
                                    strhtml+=td;
                                }


                            strhtml+="</tr>"
                             $('#databrowsergrid').find('tbody').append(strhtml);

                             var strhtml = "<tr role='row'><td class='sum-header' colspan='"+ typelist.length +"' > Total </td></tr><tr>";
                                for(var i in typelist)
                                {
                                    var td = "";
                                    if(typelist[i].tipe=="string"||typelist[i].tipe=="date")
                                       td+="<td role='gridcell'>-";
                                    else
                                    {
                                        td+= "<td role='gridcell' class='align-right'>"+ kendo.toString(dt[0][typelist[i].field.replace(".","")+"sum"] ,'N2');
                                    }
                                    td+="</td>"
                                    strhtml+=td;
                                }


                            strhtml+="</tr>"
                             $('#databrowsergrid').find('tbody').append(strhtml);


                        },300);
                }
}

</script>

<script type="text/javascript">
   
    var DataBrowser = ko.observable({
        Fields : ko.observableArray([]),
        SelectedFields :  ko.observableArray([]),
        Columns :  ko.observableArray([]),
        SelectedColumns :  ko.observableArray([]),
        Processing:ko.observable(true),
    });


	function GenerateDataBrowser(){
        if(DataBrowser().Fields().length==0 || $("#HypothesisId").html() =="H16" ){
            FillField();
        }

        var fields = DataBrowser().SelectedFields();
        var paramdouble = [];
        var allfields = [];
        var schema = {};    
        var columns = [];
        var typelist = [];

        for(var i in fields){
            var f = _.find(DataBrowser().Fields(), function(num){ return num.field == fields[i] ; });

            //Set Parameter
            if(f.tipe=="double"){
            paramdouble.push(f.field);}

            allfields.push(f.field);

            //build Schema
            if(f.tipe=="string")
            schema[f.field] = { type : "string"  }
            else if(f.tipe=="double")
            schema[f.field] = { type : "number"  }
            else
            schema[f.field] = { type : "date"  }

            //build Columns
            if(f.tipe=="string")
            columns.push({
                "field":f.field,
                "title":f.alias,
                "attributes":{ class:"align-left"},
                width:150
            });
            else if(f.tipe == "double"){
                columns.push({
                    "field":f.field,
                    "title":f.alias,
                    "format":"{0:N2}",
                    "attributes":{ class:"align-right"},
                    width:100
                });
            }
            else if(f.tipe == "date")
            columns.push({
                "field":f.field,
                "title":f.alias,
                template: "#= moment("+f.field+").utc().format('DD MMM YYYY') #", 
                attributes: { "style": "text-align:center" },
                filterable:false,
                width:100
            });

            //settype
            typelist.push({
                "field":f.field,
                "tipe":f.tipe,
            });

        }

    var Filter =  $('#HypothesisId').html() =="H18"?{PeriodFrom:model.Filter().PeriodFrom,PeriodTo :model.Filter().PeriodFrom  }:  Hypothesis.GetFilter();
	var param = {fields:allfields,fieldsdouble:paramdouble,hypoid:$('#HypothesisId').html()};

    for (var property in Filter) {
        if (Filter.hasOwnProperty(property)) {
            param[property]  = Filter[property]
        }
    }
     var hypo = $("#HypothesisId").html();
    var url = "/databrowser/getgriddb";

    if(hypo =="H16")
    {
        $("#databrowsergrid").html("");
        $('#databrowsergrid').kendoGrid({
            dataSource:{
                data : DataGrid,
                pageSize : 10,
             },
            resizable: true,
            scrollable: true,
            sortable: true,
            pageable: {
                refresh: false,
                pageSizes: 10,
                buttonCount: 5
            },
            dataBound:function(e){
                     var dt = DataGridSum;
                     GenerateSumH16(dt,typelist);
            },
            columns:columns,

        });
            // var dt = DataGridSum;
            // GenerateSumH16(dt);
                
     DataBrowser().Processing(false);
    }
    else{
      DataBrowser().Processing(true);
      $("#databrowsergrid").html("");
      $('#databrowsergrid').kendoGrid({
         dataSource: {
            transport: {
                    read: {
                       url: url,
                       dataType: "json",
                       data: param,
                       type: "POST",
                       complete: function (datas) {
                        DataBrowser().Processing(false);
                        var data = JSON.parse(datas.responseText);
                        var dt = data.Data.Summary;
                        if(dt.length>0){
                        setTimeout(function(){

                           

                            var strhtml = "<tr role='row'><td class='sum-header' colspan='"+ typelist.length +"' > Average </td></tr><tr>";
                                for(var i in typelist)
                                {
                                    var td = "";
                                    if(typelist[i].tipe=="string"||typelist[i].tipe=="date")
                                        td+="<td role='gridcell'>-";
                                    else
                                    {
                                        td+= "<td role='gridcell' class='align-right'>"+ kendo.toString(dt[0][typelist[i].field.replace(".","")+"avg"] ,'N2');
                                    }
                                    td+="</td>"
                                    strhtml+=td;
                                }


                            strhtml+="</tr>"
                             $('#databrowsergrid').find('tbody').append(strhtml);

                             var strhtml = "<tr role='row'><td class='sum-header' colspan='"+ typelist.length +"' > Total </td></tr><tr>";
                                for(var i in typelist)
                                {
                                    var td = "";
                                    if(typelist[i].tipe=="string"||typelist[i].tipe=="date")
                                       td+="<td role='gridcell'>-";
                                    else
                                    {
                                        td+= "<td role='gridcell' class='align-right'>"+ kendo.toString(dt[0][typelist[i].field.replace(".","")+"sum"] ,'N2');
                                    }
                                    td+="</td>"
                                    strhtml+=td;
                                }


                            strhtml+="</tr>"
                             $('#databrowsergrid').find('tbody').append(strhtml);


                        },300);
                }
                       }
                    }
                },
                schema: {
                    data: "Data.Datas",
                    total: "Data.Total",
                    model:{
                        fields: schema
                    }
                },
                
                pageSize: 10,
                serverPaging: true, // enable server paging
                serverSorting: true,
            },
            resizable: true,
            scrollable: true,
            sortable: true,
            pageable: {
                refresh: false,
                pageSizes: 10,
                buttonCount: 5
            },
            columns:columns
    });
}
   //        }else{
   //            alert(res.Message);
   //        }
   //    },
   // });

	}

  
</script>