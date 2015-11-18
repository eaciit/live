package routers

import (
	"github.com/eaciit/live/daemon/webapp/controllers"

	"github.com/astaxie/beego"
)

func init() {

	// beego.Router("/", &controllers.HomeController{})
	beego.Router("/", &controllers.HomeController{})
	beego.AutoRouter(&controllers.HomeController{})
}
