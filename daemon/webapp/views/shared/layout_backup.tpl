<!DOCTYPE html>
<html>
<head>
    <title>EACIIT SmartView</title>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link rel="shortcut icon" href="/static/favicon.ico" type="image/x-icon" />
    <link rel="icon" href="/static/favicon.ico" type="image/x-icon" />

    <script src="/static/js/jquery-2.1.0.min.js"></script>
    <script src="/static/js/jquery-ui.min.js"></script>
    <script src="/static/js/knockout-3.1.0.js"></script>
    <script src="/static/js/knockout.mapping-latest.js"></script>
    <script src="/static/js/knockout.validation.js"></script>
    <script src="/static/js/jszip.min.js"></script>
    <script src="/static/kendoui/js/kendo.all.min.js"></script>
    <script src="/static/js/knockout-kendo.min.js"></script>
    <script src="/static/js/bootstrap.min.js"></script>
    <script src="/static/js/underscore.min.js"></script>
    <script src="/static/js/moment.min.js"></script>
    <script src="/static/js/color.js"></script>
    <script src="/static/js/tools.js"></script>
    <script src="/static/js/linq.js"></script>

    <link href="/static/css/bootstrap.css" type="text/css" rel="stylesheet" />
    <link href="/static/css/bootstrap-theme.css" type="text/css" rel="stylesheet" />

    <link rel="stylesheet" href="/static/kendoui/styles/kendo.material.min.css" />
    <link rel="stylesheet" href="/static/kendoui/styles/kendo.common-bootstrap.min.css" />
    <link rel="stylesheet" href="/static/kendoui/styles/kendo.dataviz.min.css" />
    <link rel="stylesheet" href="/static/kendoui/styles/kendo.dataviz.bootstrap.min.css" />

    <script src="/static/js/underscore.min.js"></script>
    <script src="/static/js/ecis_config.js"></script>
    <script src="/static/js/main.js"></script>
    <script src="/static/js/ecis_start.js"></script>
    <script src="/static/js/jquery.fullscreen.min.js"></script>

    <link rel="stylesheet" href="/static/css/font-awesome.css" />
    <link rel="stylesheet" href="/static/css/Site.css" />
    <link rel="stylesheet" href="/static/css/custom.css" />
    <script>
        var Now = new Date();
        function getQuarter(d) {
          d = d || new Date(); // If no date supplied, use today
          var q = [1,2,3,4];
          return q[Math.floor(d.getMonth() / 3)];
        }
        function getFinanceQuarter(d) {
          d = d || new Date(); // If no date supplied, use today
          var q = [4,1,2,3];
          return q[Math.floor(d.getMonth() / 3)];
        }
        var model = {
            Processing: ko.observable(true)
        }
    </script>

    <style>
        .k-multiselect{
            background:none;
        }
        html, body {
            max-width: 100%;
            overflow-x: hidden;
        }
        .menu-header {
        height: 30px;
        background-color: #333;
        }

        .nav-bar-header li {
        float: left;
        }

        .nav-bar-header li a {
        font-size: 12px;
        margin: 0px;
        padding: 5px;
        background-color: #000;
        }

        .nav-bar-header li.selected a {
        background-color: #D33;
        }

        .nav-bar-header li a:hover {
        font-size: 12px;
        margin: 0px;
        padding: 5px;
        background-color: #D33;
        }

        .form-group label {
        text-align: right;
        padding-right: 10px;
        }

        .form-group input {
        border: solid 1px #ccc;
        padding: 2px;
        }

        .form-group input[type='number'] {
        text-align: right;
        }


        #logo > #triangle1{
            width: 0px;
            height: 0px;
            border-top: 75px solid transparent;
            border-right: 80px solid #035882;
            margin-left: 220px;
        }
        #logo > #triangle2{
            width: 0px;
            height: 0px;
            border-bottom: 75px solid transparent;
            border-left: 20px solid #035882;
            margin-top: -75px;
            margin-left: 300px;
        }
        #logo > #triangle3{
            width: 0px;
            height: 0px;
            border-top: 75px solid transparent;
            border-right: 40px solid #F79E44;
            margin-top: -75px;
            margin-left: 279px
        }
        #logo > #triangle4{
            width: 0px;
            height: 0px;
            border-bottom: 75px solid transparent;
            border-left: 75px solid #F79E44;
            margin-top: -75px;
            margin-left: 319px;
        }
        #HypothesisCategory{
            color: #E67401;
            font-size: 15px;
            font-style: italic;
            padding: 3px;
            font-weight: bold;
        }
        #HypothesisCategory:after{
            content: ' Analysis';
            font-weight: normal;
        }
    </style>


    <script id="navbarTemplate" type="text/html">
        <!-- ko if: Submenus().length==0 -->
        <li data-bind="css:{selected:$root.PageId()==_id()}">
            <a data-bind="text: Title, attr:{href:Url}"></a>
        </li>
        <!-- /ko -->
        <!-- ko if: Submenus().length>0 -->
        <li data-bind="css:{selected:$root.PageId()==_id()}" class="dropdown">
            <a data-bind="text: Title" style="cursor:pointer" class="dropdown-toggle" data-toggle="dropdown"></a>
            <ul class="dropdown-menu" role="menu" data-bind="template:{name:'navbarSubTemplate', foreach:Submenus}"></ul>
        </li>
        <!-- /ko -->
    </script>

    <script id="navbarSubTemplate" type="text/html">
        <!-- ko if: Submenus().length==0 -->
        <li data-bind="css:{selected:$root.PageId()==_id()}">
            <a data-bind="text: Title, attr:{href:Url}"></a>
        </li>
        <!-- /ko -->
        <!-- ko if: Submenus().length>0 -->
        <li data-bind="css:{selected:$root.PageId()==_id()}" class="dropdown-submenu">
            <a data-bind="text: Title" style="cursor:pointer" class="dropdown-toggle" data-toggle="dropdown"></a>
            <ul class="dropdown-menu" role="menu" data-bind="template:{name:'navbarSubTemplate', foreach:Submenus}"></ul>
        </li>
        <!-- /ko -->
    </script>

    <script id="userTemplate" type="text/html">
        <!-- ko if: Submenus().length==0 -->
        <li data-bind="css:{selected:$root.PageId()==_id()}">
            <a data-bind="text: Title, attr:{href:Url}"></a>
        </li>
        <!-- /ko -->
        <!-- ko if: Submenus().length>0 -->
        <li data-bind="css:{selected:$root.PageId()==_id()}" class="dropdown-submenu-left">
            <a data-bind="text: Title" style="cursor:pointer" class="dropdown-toggle" data-toggle="dropdown"></a>
            <ul class="dropdown-menu" role="menu" data-bind="template:{name:'userSubTemplate', foreach:Submenus}"></ul>
        </li>
        <!-- /ko -->
    </script>

    <script id="userSubTemplate" type="text/html">
        <!-- ko if: Submenus().length==0 -->
        <li data-bind="css:{selected:$root.PageId()==_id()}">
            <a data-bind="text: Title, attr:{href:Url}"></a>
        </li>
        <!-- /ko -->
        <!-- ko if: Submenus().length>0 -->
        <li data-bind="css:{selected:$root.PageId()==_id()}" class="dropdown-submenu-left">
            <a data-bind="text: Title" style="cursor:pointer" class="dropdown-toggle" data-toggle="dropdown"></a>
            <ul class="dropdown-menu" role="menu" data-bind="template:{name:'userSubTemplate', foreach:Submenus}"></ul>
        </li>
        <!-- /ko -->
    </script>

    <script id="breadcrumbTemplate" type="text/html">
        <li data-bind="attr:{class:CssClass}"><a data-bind="text: Title, attr:{href:Url, class:CssClass, onclick: (Action!=undefined?Action:'')}"></a></li>
    </script>

</head>
<body>

    <div data-bind="visible:Processing()">
        {{template "shared/processing.tpl"}}
    </div>

    <!-- wrapper starts here -->
    <div class="container-fluid" style="display:none" data-bind="visible:!Processing()">
        <!-- header starts here -->
        <header id="page-header">
            <div style="float:left;width:20px;background:white;height:75px;"></div>
            <section id="logo">
                <a href="res/" id="logo-link"><span>RM</span></a>
                <div id="triangle1"></div>
                <div id="triangle2"></div>
                <div id="triangle3"></div>
                <div id="triangle4"></div>
            </section>
            <section id="user-nav">
                <div class="user-nav-wrapper">
                    <a href="#" class="user-nav-info" data-toggle="dropdown">Welcome <span>ADMIN<!-- {{.UserID}} --></span></a>
                    <!-- ko if:UserMenus().length==0 -->
                    <ul class="dropdown-menu" role="menu">
                        <!--li><a href="res/webmenu/menu">Menu Configuration</a></li-->
                        <li><a onClick="UserManagementMenu()" class="logout-link">User Management</a></li>
                        <li><a onClick="Logout()" class="logout-link">Logout</a></li><!--href="" onClick="localStorage.clear();"-->
                    </ul>
                    <!-- /ko -->
                    <!-- ko if:UserMenus().length>0 -->
                    <ul class="dropdown-menu" role="menu" data-bind="template:{name:'userTemplate', foreach:UserMenus}"></ul>
                    <!-- /ko -->
                </div>
            </section>

        </header>

        <script>
            function Logout(){
                localStorage.clear();
                window.location.href = "/Login/Logout";
            }
            function UserManagementMenu(){
                localStorage.clear();
                window.location.href = "/Administration?PageId=UserManagement&PageTitle=User Management";
            }

            function MenuItem(id, url, title, submenus, baseURL) {
                var obj = {
                    _id: ko.observable(id),
                    Title: ko.observable(title == undefined ? id : title),
                    Url: ko.observable(url.replace("~/",baseURL)),
                    Submenus: ko.observableArray([])
                };

                var arr = submenus;
                for(var i in arr){
                    obj.Submenus.push(
                        new MenuItem(
                            arr[i]._id,
                            arr[i].Url,
                            arr[i].Title,
                            arr[i].Submenus,
                            baseURL
                        )
                    );
                }
                return obj;
            };

            function BreadCrumb(id, title, url, cssClass, action) {
                var obj = {
                    _id: ko.observable(id),
                    Title: ko.observable(title == undefined ? id : title),
                    Url: url,
                    Action: action,
                    CssClass: cssClass
                };

                return obj;
            }


            model.PageId = ko.observable("{{.PageId}}");
            model.PageTitle = ko.observable("{{.PageTitle}}");
            model.HypothesisId = ko.observable("{{.HypothesisId}}");
            model.HypothesisCategory = ko.observable("{{.HypothesisCategory}}");
            model.MainMenus = ko.observableArray([]);
            model.MenuList  = [
                {_id:"Home",Url:"/Home",Title:"Home",Submenus:[]},
                // {_id:"FunctionalLocation",Url:"/FunctionalLocation",Title:"Functional Location",Submenus:[]},
                {_id:"ComparisonHypothesis",Url:"/Hypothesis?PageId=ComparisonHypothesis&PageTitle=Comparison Hypothesis",Title:"Comparison Hypothesis",Submenus:[]},
                {_id:"ValueHypothesis",Url:"/Hypothesis?PageId=ValueHypothesis&PageTitle=Value Hypothesis",Title:"Value Hypothesis",Submenus:[]},
                {_id:"FailureHypothesis",Url:"/Hypothesis?PageId=FailureHypothesis&PageTitle=Failure Hypothesis",Title:"Failure Hypothesis",Submenus:[]},
                 // {_id:"Administration",Url:"#",Title:"Administration",Submenus:[
                 //    {_id:"UserManagement",Url:"/Administration?PageId=UserManagement&PageTitle=User Management",Title:"User Management",Submenus:[]},
                 //    // {_id:"RoleManagement",Url:"RoleManagement?user={{.Data}}",Title:"Role Management",Submenus:[]},
                 // ]},
            ]
            model.UserMenus = ko.observableArray([]);

            model.BreadCrumbs = ko.observableArray([]);

            model.getMainMenu = function(){
                var baseURL = "res";
                var arr = model.MenuList;
                for(var i in arr){
                    model.MainMenus.push(
                        new MenuItem(
                            arr[i]._id,
                            arr[i].Url,
                            arr[i].Title,
                            arr[i].Submenus,
                            baseURL
                        )
                    );
                }
            }


            model.getUserMenu = function(){
                var url = "res/webmenu/menu/getmenu";
                var baseURL = "res";
                model.UserMenus([]);
                ajaxPost(url,{collection_name:"User_Menu"},function(respondse){
                    var arr = respondse.Data;
                    for(var i in arr){
                        model.UserMenus.push(
                            new MenuItem(
                                arr[i]._id,
                                arr[i].Url,
                                arr[i].Title,
                                arr[i].Submenus,
                                baseURL
                            )
                        );
                    }
                });
            }
        </script>

        <nav class="navbar">
            <div class="container-fluid">
                <div class="navbar-header">
                    <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#navbar" aria-expanded="false" aria-controls="navbar">
                        <span class="sr-only">Toggle navigation</span>
                        <i class="fa fa-bars fa-2x"></i>
                    </button>
                </div>
                <div id="navbar" class="collapse navbar-collapse">
                    <ul class="nav navbar-nav" data-bind="template:{name:'navbarTemplate', foreach:MainMenus}"></ul>
                </div>
            </div>
        </nav>


        <script>
            model.BreadCrumbs.push(new BreadCrumb("youare", "You are here : ", "#", "youare"));
        </script>
        <section class="breadcrumb-section">
            <ol class="breadcrumb" data-bind="template:{name:'breadcrumbTemplate', foreach:BreadCrumbs}"></ol>
        </section>
    </div>
    <!-- header ends here -->
    <!-- section starts here -->
    <div class="container-fluid content-main" data-bind="visible:!Processing()">

        <section class="content">
            <div class="content-header">
                <div class="row">
                    <div class="col-md-9">
                        <h2 class="content-title">
                            <span id="HypothesisCategory" data-bind="visible:HypothesisCategory()!=='',text:HypothesisCategory()"></span>
                            <span id="HypothesisId" data-bind="visible:HypothesisId()!=='',text:HypothesisId"></span>
                            <span data-bind="text:PageTitle,attr:HypothesisId()!==''?{style:'font-size:15px;'}:{}"></span>
                        </h2>
                    </div>
                </div>
            </div>
            <!-- Main panel starts here -->
            {{.LayoutContent}}
        </section>

    </div>
    <!-- section ends here -->
    <!-- wrapper ends here -->
    <script>
        $(function(){
            $('#logo-link').on('click', function() {
                $('body').fullscreen({ overflow: 'auto' });
                return false;
            });
        });

        function setPageTitle(s) {
            $("#pageTitle").text(s);
        }

        ko.applyBindings(model);

        $(document).ready(function () {
            model.Processing(false);
            model.getMainMenu();
            //model.getUserMenu();
            if (typeof PageUpdate == "function") {
                PageUpdate();
            }
        });
    </script>

</body>
</html>