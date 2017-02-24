package web

import (
	"ServerManagementSystem/models"
	"github.com/astaxie/beego"
	"ServerManagementSystem/logs"
	"encoding/json"
	"strconv"
	"time"
	"fmt"
)


// 登陆
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
		fmt.Println(user.UserType)
		fmt.Println(models.USER_TYPE_SUPER_ADMIN)
		if user.UserType == models.USER_TYPE_SUPER_ADMIN {
			isAdmin = true
		} else {
			isAdmin = false
		}
		fmt.Println(isAdmin)

		this.SetSession("isAdmin", isAdmin)
		this.SetSession("userName", user.UserName)
		this.SetSession("nickName", user.NickName)
		this.SetSession("userId", user.Id)
		this.SetSession("avatar", user.Avatar)

		// 将登陆信息存入redis
		jsonByte, _ := json.Marshal(user)
		models.RedisCache.Put(models.USER_INFO_REDIS + "_" + strconv.FormatInt(user.Id, 10), jsonByte, 60*60*24*time.Second)

		this.Redirect("/web/public/index", 302)
		return
	} else {
		flash := beego.NewFlash()

		flash.Error(this.Err.Error())
		logs.Logger.Info(this.Err.Error())
		flash.Store(&this.Controller)
		this.Redirect("/", 302)
		return
	}
}

// 登出
type LogOutController struct {
	BaseWebController
}

func (this *LogOutController) Prepare() {

}

func (this *LogOutController) Get() {
	this.DestroySession()
	this.Redirect("/", 302)
	return
}





