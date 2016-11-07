package web

import "github.com/astaxie/beego"

type UserWebController struct {
	beego.Controller
}

// @router / [get]
func (c *UserWebController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "web/index.tpl"
}

