package api

import "ServerManagementSystem/controllers"

type BaseApiController struct {
	controllers.BaseController
	Err        error
	ErrCode    int
	ReturnData interface{}
}




