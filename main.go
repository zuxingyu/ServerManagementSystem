package main

import (
	"github.com/astaxie/beego"
	_ "ServerManagementSystem/model"
	_ "ServerManagementSystem/routers"
	_ "ServerManagementSystem/ServerMonitor"
	"ServerManagementSystem/logs"
)

func main() {
	defer logs.FlushLog()
	beego.Run()
}
