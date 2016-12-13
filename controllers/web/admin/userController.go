package admin

type AdminUserController struct {
	BaseAdminController
}

// @router /list [get]
func (this *AdminUserController) UserList(){

	this.TplName = "web/admin/userList.tpl"
}



