package main

import (
	"github.com/astaxie/beego"
	_ "ServerManagementSystem/routers"
	_ "ServerManagementSystem/monitor"
)

func main() {


	beego.Run()
}
