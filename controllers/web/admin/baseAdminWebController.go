package admin

import (
	"ServerManagementSystem/controllers/web"
	"strings"
	"github.com/astaxie/beego"
	"ServerManagementSystem/logs"
	"ServerManagementSystem/models"
	"errors"
)

type BaseAdminController struct {
	web.BaseWebController
}


func (this *BaseAdminController) Prepare() {
	userName := this.GetSession("userName")
	userId := this.GetSession("userId")

	if userName == nil || userId == nil {
		flash := beego.NewFlash()

		flash.Error("登陆信息已过期或用户名不存在")
		logs.Logger.Error(this.Err.Error())
		flash.Store(&this.Controller)

		this.Redirect("/", 302)
		return
	}
	// 判断用户是否是管理员
	user, thErr := models.User_GetByUid(userId.(int64), true)
	if thErr != nil {
		flash := beego.NewFlash()
		flash.Error(thErr.Error())
		logs.Logger.Error(this.Err.Error())
		flash.Store(&this.Controller)
		this.Redirect("/", 302)
		return
	} else {
		if models.USER_TYPE_SUPER_ADMIN != user.UserType {
			flash := beego.NewFlash()
			err := errors.New("您没有此操作权限")
			flash.Error(err.Error())
			logs.Logger.Error(err.Error())
			flash.Store(&this.Controller)
			this.Redirect("/", 302)
			return
		}
	}

	this.Data["LoginUserName"] = userName.(string)
	this.Data["LoginUserId"] = userId.(int64)
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
