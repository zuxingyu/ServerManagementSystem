package api

import (
	"ServerManagementSystem/controllers"
)

type UserApiController struct {
	controllers.BaseController
}

func (c *UserApiController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "zuxingyu@gmail.com"
	c.TplName = "index.tpl"
}


