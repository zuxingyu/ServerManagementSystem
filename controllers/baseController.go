package controllers

import (
	"github.com/astaxie/beego"
	"strings"
)

type BaseController struct {
	beego.Controller
	Err        error
	ErrorCode  int
	ReturnData interface{}
}

func (this *BaseController) Prepare() {
	userName := this.GetSession("userName")
	userId := this.GetSession("userId")

	if userName == nil || userId == nil {
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



