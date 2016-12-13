package web

import (
	"ServerManagementSystem/controllers"
	"fmt"
	"net/http"
	"errors"
	"ServerManagementSystem/logs"
	"github.com/astaxie/beego"
	"strings"
)

type BaseWebController struct {
	controllers.BaseController
}



func (this *BaseWebController) MustString(key string) string {
	v := this.GetString(key)
	if v == "" {
		// 400 Error, Parameter error
		this.ErrorCode = http.StatusBadRequest
		this.Err = errors.New(fmt.Sprintf("必填参数为空: %s", key))
		logs.Logger.Error(this.Err)
		//this.Failed()
	}
	return v
}


func (this *BaseWebController) Prepare() {
	userName := this.GetSession("userName")
	userId := this.GetSession("userId")

	if userName == nil || userId == nil {
		flash := beego.NewFlash()

		flash.Error("登陆信息已过期或用户名不存在")
		logs.Logger.Info(this.Err.Error())
		flash.Store(&this.Controller)

		this.Redirect("/", 302)
		return
	} else {
		this.Data["LoginUserName"] = userName.(string)
		this.Data["LoginUserId"] = userId.(int64)
	}
	this.Data["isAdmin"] = this.GetSession("isAdmin").(bool)

	// 如果没有头像，使用默认头像
	userAvatar := this.GetSession("avatar")
	var avatarStr string
	if userAvatar == nil || strings.EqualFold("", userAvatar.(string)) {
		avatarStr = beego.AppConfig.String("defaultAvatar")
	} else {
		avatarStr = userAvatar.(string)
	}

	this.Data["LoginAvatar"] = avatarStr

	// 如果没有昵称，用户名代替昵称
	nickName := this.GetSession("nickName")
	var nickNameStr string
	if nickName == nil || strings.EqualFold("", nickName.(string)) {
		nickNameStr = userName.(string)
	} else {
		nickNameStr = nickName.(string)
	}
	this.Data["LoginNickName"] = nickNameStr
}






