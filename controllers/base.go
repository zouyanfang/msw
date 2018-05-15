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
		this.User,_ = models.Login(name ,pwd)
		this.Data["user"] = this.User
	}else {
		this.Data["user"] = nil
	}
}

//TODO
//是否需要模板
func (this *BaseController)IsNeddTemplate(){
	this.Layout = "site/index.html"
}

