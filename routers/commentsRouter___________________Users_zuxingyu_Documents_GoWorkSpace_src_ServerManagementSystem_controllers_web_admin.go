package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["ServerManagementSystem/controllers/web/admin:AdminUserController"] = append(beego.GlobalControllerRouter["ServerManagementSystem/controllers/web/admin:AdminUserController"],
		beego.ControllerComments{
			Method: "UserList",
			Router: `/list`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ServerManagementSystem/controllers/web/admin:AdminUserController"] = append(beego.GlobalControllerRouter["ServerManagementSystem/controllers/web/admin:AdminUserController"],
		beego.ControllerComments{
			Method: "UserListData",
			Router: `/listData`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}
