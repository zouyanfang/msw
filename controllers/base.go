package controllers

import "github.com/astaxie/beego"

type BaseController struct {
	beego.Controller
} 

func (this *BaseController)Prepare(){
	this.Ctx.GetCookie("")
}

func (this *BaseController)IsNeddTemplate(){
	this.Layout = "..."
}

