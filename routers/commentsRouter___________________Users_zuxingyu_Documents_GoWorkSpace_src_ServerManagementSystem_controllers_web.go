package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["ServerManagementSystem/controllers/web:PublicController"] = append(beego.GlobalControllerRouter["ServerManagementSystem/controllers/web:PublicController"],
		beego.ControllerComments{
			Method: "Index",
			Router: `/index`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ServerManagementSystem/controllers/web:UserWebController"] = append(beego.GlobalControllerRouter["ServerManagementSystem/controllers/web:UserWebController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}
