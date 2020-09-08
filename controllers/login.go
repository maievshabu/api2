package controllers

import (
	"github.com/astaxie/beego"
)

type LoginController struct{
	beego.Controller
}

func (t *LoginController) Get(){

	user :=make(map[string]string)
	user["age"] = "20"
	user["name"] = "maiev"

	t.Data["json"] = user
	t.ServeJSON()
}
