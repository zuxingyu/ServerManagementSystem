package web

import (
	"ServerManagementSystem/controllers"
	"fmt"
	"net/http"
	"errors"
	"ServerManagementSystem/logs"
	//"ServerManagementSystem/models"
)

type BaseWebController struct {
	controllers.BaseController
}

func (this *BaseWebController) MustString(key string) string {
	v := this.GetString(key)
	if v == "" {
		// 400 Error, Parameter error
		this.ErrorCode = http.StatusBadRequest
		this.Err = errors.New(fmt.Sprintf("必填参数为空: %s", key))
		logs.Logger.Error(this.Err)
		//this.Failed()
	}
	return v
}





//func (this *BaseWebController) Failed() {
//	if this.ErrorCode == 0 {
//		this.ErrorCode = models.ServerfailUndefinedCode
//	}
//
//	if this.ErrorCode < model.ServerSuccessCode {
//		// Http Error. 401 400 500
//		//http.Error(this.Ctx.ResponseWriter, this.Err.Error(), this.ErrCode)
//		this.Ctx.Output.SetStatus(this.ErrorCode)
//		this.Data["json"] = map[string]string{
//			controllers.HeaderDesKey : this.Err.Error(),
//		}
//	} else {
//		this.Data["json"] = map[string]interface{}{
//			"header" : map[string]string{
//				controllers.HeaderCodeKey : fmt.Sprintf("%d",this.ErrorCode),
//				controllers.HeaderDesKey : this.Err.Error(),
//			},
//		}
//	}
//	this.Redirect("/",302)
//	this.StopRun()
//}
//
//func (this *BaseWebController) Finsh() {
//	if this.Err != nil {
//		this.Failed()
//	}
//
//
//}







