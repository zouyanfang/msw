package controllers

import "github.com/astaxie/beego"

type BaseController struct {
	beego.Controller
} 

func (this *BaseController)Prepare(){
/*	name := this.Ctx.GetCookie("username")
	pwd := this.Ctx.GetCookie("password")*/

}

func (this *BaseController)IsNeddTemplate(){
	this.Layout = "..."
}

