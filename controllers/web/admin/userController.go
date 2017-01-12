package admin

import (
	"strconv"
	"ServerManagementSystem/util"
	"ServerManagementSystem/models"
	"github.com/astaxie/beego"
	"ServerManagementSystem/logs"
)

type AdminUserController struct {
	BaseAdminController
}

// @router /list [get]
func (this *AdminUserController) UserList() {
	this.TplName = "web/admin/userList.tpl"
}

// @router /listData [get]
func (this *AdminUserController) UserListData() {
	aColumns := []string{
		"id",
		"nick_name",
		"user_name",
		"mobile",
		"created",
		"user_state",
	}

	var where = make(map[string]interface{})
	where["id!="] = 1


	/*
	 * Paging
	 * 分页请求
	 */
	Input := this.Ctx.Input
	iDisplayStart, _ := strconv.Atoi(Input.Query("iDisplayStart"))
	iDisplayLength, _ := strconv.Atoi(Input.Query("iDisplayLength"))

	orderStr, args, whereStr := util.Datatables(aColumns, Input, where)



	counts, err := models.Orm.Table("user").Where(whereStr, args...).Count(new(models.UserDTO))
	if err != nil {
		flash := beego.NewFlash()

		flash.Error(this.Err.Error())
		logs.Logger.Info(this.Err.Error())
		flash.Store(&this.Controller)
		this.Redirect("/", 302)
		return
	}

	var userList []models.UserDTO
	models.Orm.Table("user").Where(whereStr, args...).OrderBy(orderStr).Limit(iDisplayLength, iDisplayStart).Find(&userList)
	data := make(map[string]interface{}, iDisplayLength)
	data["sEcho"] = this.GetString("sEcho")
	data["iTotalRecords"] = counts
	data["iTotalDisplayRecords"] = counts
	if len(userList) == 0{
		data["aaData"] = ""
	} else {
		data["aaData"] = userList
	}

	this.Data["json"] = data
	this.ServeJSON()
}

