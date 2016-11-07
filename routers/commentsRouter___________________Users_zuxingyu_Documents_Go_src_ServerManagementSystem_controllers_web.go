package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["ServerManagementSystem/controllers/web:UserWebController"] = append(beego.GlobalControllerRouter["ServerManagementSystem/controllers/web:UserWebController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}
