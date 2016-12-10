package web

import (
	"ServerManagementSystem/controllers"
)

type UserWebController struct {
	controllers.BaseController
}

// @router / [get]
func (this *UserWebController) Get() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplName = "web/index.tpl"
}

