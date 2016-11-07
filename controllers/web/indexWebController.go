package web

import (
	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}



func (c *IndexController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "zuxingyu@gmail.com"
	c.TplName = "index.tpl"
}

