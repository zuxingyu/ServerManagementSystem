package main

import (
	"github.com/astaxie/beego"
	_ "ServerManagementSystem/models"
	_ "ServerManagementSystem/routers"
	_ "ServerManagementSystem/ServerMonitor"
	_ "github.com/astaxie/beego/session/redis"
	"ServerManagementSystem/logs"
)

func main() {
	defer logs.FlushLog()
	beego.Run()
}