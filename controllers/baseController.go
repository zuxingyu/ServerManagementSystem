package controllers

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
	Err        error
	ErrorCode  int
	ReturnData interface{}
}



