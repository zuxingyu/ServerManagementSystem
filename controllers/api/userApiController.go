package api

import "github.com/astaxie/beego"

type UserApiController struct {
	beego.Controller
}

func (c *UserApiController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}


