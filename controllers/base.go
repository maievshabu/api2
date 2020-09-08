package controllers

import (
	"github.com/astaxie/beego"
)

type BaseController struct{
	beego.Controller
}

func (b *BaseController)ReturnJson(msg string, code int)map[string]interface{}{
	return map[string]interface{}{"msg": code, "info": msg}
}
