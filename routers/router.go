package routers

import (
	"github.com/astaxie/beego"
	"ServerManagementSystem/controllers/api"
	"ServerManagementSystem/controllers/web"
	"github.com/astaxie/beego/context"
	"net/http"
	"html/template"
	"ServerManagementSystem/controllers/web/admin"
)

func init() {

	nsApi := beego.NewNamespace("/api",
		beego.NSNamespace("/user",
			beego.NSInclude(&api.UserApiController{}),
		),
	)

	nsWeb := beego.NewNamespace("/web",
		beego.NSNamespace("/public",
			beego.NSInclude(&web.PublicController{}),
		),
		beego.NSNamespace("/userManager",
			beego.NSInclude(&admin.AdminUserController{}),
		),
	)


	// 设置首页
	beego.Router("/", &web.LoginController{})

	// 登陆
	beego.Router("/login", &web.LoginController{})

	// 登出
	beego.Router("/logout", &web.LogOutController{})


	// 设置404，500页面
	beego.ErrorHandler("404", page_not_found)

	// 设置其他
	beego.AddNamespace(nsApi, nsWeb)

	// 增加过滤器判断
	beego.InsertFilter("/web/*",beego.BeforeRouter,FilterUser)

}

var FilterUser = func(ctx *context.Context) {
	_, ok := ctx.Input.Session("userId").(int64)
	if !ok && ctx.Request.RequestURI != "/login" {
		ctx.Redirect(302, "/")
	}
}

func page_not_found(rw http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles(beego.AppConfig.String("ViewsPath") + "/errorpage/page404.tpl")
	data := make(map[string]interface{})
	t.Execute(rw, data)
}
