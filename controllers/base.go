package controllers

import (
	"github.com/astaxie/beego"
	"msw/models"
)

type BaseController struct {
	beego.Controller
	User *models.User
} 

//验证用户登录信息
func (this *BaseController)Prepare(){
	name := this.Ctx.GetCookie("uid")
	pwd := this.Ctx.GetCookie("pid")
	if name != "" && pwd != ""{
		user,_ := models.Login(name ,pwd)
		this.Data["user"] = &user
	}
	this.Data["user"] = models.User{}
}

//TODO
//是否需要模板
func (this *BaseController)IsNeddTemplate(){
	this.Layout = "..."
}

