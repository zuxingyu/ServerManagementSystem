package web

import (
	"ServerManagementSystem/models"
	"github.com/astaxie/beego"
	"wenshixiongServer/util"
	"ServerManagementSystem/logs"
	"strings"
)

type LoginController struct {
	BaseWebController
}

func (this *LoginController) Prepare() {

}

func (this *LoginController) Get() {
	flash := beego.ReadFromRequest(&this.Controller)
	if n, ok := flash.Data["error"]; ok {
		logs.Logger.Error(n)
		this.Data["error"] = flash.Data["error"]
	}

	this.TplName = "login.tpl"
}

func (this *LoginController) Post() {
	userName := this.MustString("username")
	userPassword := this.MustString("password")

	var user models.User

	this.Err, this.ErrorCode = models.User_Login(userName, userPassword, &user)

	if this.Err == nil {
		logs.Logger.Info(user)
		var isAdmin bool
		if strings.EqualFold("admin", user.UserName) {
			isAdmin = true
		} else {
			isAdmin = false
		}

		this.SetSession("isAdmin", isAdmin)
		this.SetSession("userName", user.UserName)
		this.SetSession("nickName", user.NickName)
		this.SetSession("userId", user.Id)
		this.SetSession("avatar", user.Avatar)
		this.Redirect("/web/public/index", 302)
		return
	} else {
		flash := beego.NewFlash()

		flash.Error(this.Err.Error())
		util.Logger.Info(this.Err.Error())
		flash.Store(&this.Controller)
		this.Redirect("/", 302)
		return
	}
}



