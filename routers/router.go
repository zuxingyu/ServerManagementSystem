package routers

import (
	"github.com/astaxie/beego"
	"ServerManagementSystem/controllers/api"
	"ServerManagementSystem/controllers/web"
)

func init() {
	nsApi := beego.NewNamespace("/api",
		beego.NSNamespace("/user",
			beego.NSInclude(&api.UserApiController{}),
		),
	)

	nsWeb := beego.NewNamespace("/web",
		beego.NSNamespace("/user",
			beego.NSInclude(&web.UserWebController{}),
		),
	)

	// 设置首页
	beego.Router("/", &web.IndexController{})

	// 设置其他
	beego.AddNamespace(nsApi, nsWeb)

}
