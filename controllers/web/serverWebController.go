package web

import (
	"strconv"
	"ServerManagementSystem/util"
	"ServerManagementSystem/models"
	"github.com/astaxie/beego"
	"ServerManagementSystem/logs"
	"fmt"
)

type ServerWebController struct {
	BaseWebController
}



// @router /list [get]
func (this *ServerWebController) ServerList() {
	this.TplName = "web/serverList.tpl"
}

// @router /listData [get]
func (this *ServerWebController) ServerDataList() {
	aColumns := []string{
		"id",
		"alias_name",
		"server_o_s",
		"server_state",
		"c_p_u_rate",
		"memory_rate",
		"Created",
	}

	userId := this.GetSession("userId")
	var where = make(map[string]interface{})
	where["user_id="] = userId.(int64)


	/*
	 * Paging
	 * 分页请求
	 */
	Input := this.Ctx.Input
	iDisplayStart, _ := strconv.Atoi(Input.Query("iDisplayStart"))
	iDisplayLength, _ := strconv.Atoi(Input.Query("iDisplayLength"))

	orderStr, args, whereStr := util.Datatables("server", aColumns, Input, where)

	counts, err := models.Orm.Table("server").Where(whereStr, args...).Count(new(models.UserDTO))
	if err != nil {
		flash := beego.NewFlash()

		flash.Error(this.Err.Error())
		logs.Logger.Info(this.Err.Error())
		flash.Store(&this.Controller)
		this.Redirect("/", 302)
		return
	}

	fmt.Println(counts)


	var serverList []models.ServerDTO
	models.Orm.Table("server").Select("*,(select server_memory_log.memory_rate FROM server_memory_log WHERE server_memory_log.server_id = server.id ORDER BY server_memory_log.created DESC LIMIT 1 ) AS memory_rate , (select server_c_p_u_log.c_p_u_rate FROM server_c_p_u_log WHERE server_c_p_u_log.server_id = server.id ORDER BY server_c_p_u_log.created DESC LIMIT 1 ) AS c_p_u_rate ").Where(whereStr, args...).OrderBy(orderStr).Limit(iDisplayLength, iDisplayStart).Find(&serverList)
	data := make(map[string]interface{}, iDisplayLength)
	data["sEcho"] = this.GetString("sEcho")
	data["iTotalRecords"] = counts
	data["iTotalDisplayRecords"] = counts

	fmt.Println(serverList)
	if len(serverList) == 0{
		data["aaData"] = ""
	} else {
		data["aaData"] = serverList
	}

	this.Data["json"] = data
	this.ServeJSON()
}

