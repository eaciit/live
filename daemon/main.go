package main

import (
	"github.com/astaxie/beego"
	_ "github.com/eaciit/live/daemon/routers"
)

func main() {
	beego.Run()
}
