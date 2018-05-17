package controllers

import (
	"msw/models"
	"msw/services"
	)

type UserController struct {
	BaseController
}

//用户菜谱页面


func (this *UserController)UserCenter(){
	this.Data["type"] = 1

	this.IsNeedTemplate()
	this.TplName = "site/user/dishcenter.html"
}

//用户菜单页面
func (this *UserController)UserMenu(){
	this.Data["type"] = 2
	this.IsNeedTemplate()
	this.TplName = "site/user/menucenter.html"
}

//用户个人信息设置页面
func (this *UserController)UserConfig(){
	this.Data["type"] = 3
	this.IsNeedTemplate()
	this.TplName = "site/user/userconfig.html"
}


//用户模板
func (this *UserController)IsNeedTemplate(){
	this.IsHaveUser()
	this.Layout = "site/user/template.html"
}

//判断是否有用户
func (this *UserController)IsHaveUser(){
	if this.User == nil {
		this.Redirect("/login",302)
		return
	}
}

func (this *UserController)Demo(){
	this.TplName = "site/user/1.html"
}

func (this *UserController)GetImg(){
	//_,s,_ := this.GetFile("file")
	this.ServeJSON()
}

func (this *UserController)ReleaseTalk(){
	var resp models.BaseMsgResp
	resp.Ret = 403
	defer func() {
		this.Data["json"] = resp
		this.ServeJSON()
	}()
	if (this.User == nil ){
		resp.Msg = "用户还未登入，请先登入"
		return
	}
	content := this.GetString("content")
	dishid,err := this.GetInt("dishid")
	if err != nil {
		resp.Msg = "获取参数失败"
		return
	}
	resp = services.ReleaseDishTalk(dishid,this.User.Id,content)
	return
}

//用户收藏菜谱
func (this *UserController)UserCollection(){
	var resp models.BaseMsgResp
	resp.Ret = 403
	defer func() {
		this.Data["json"] = resp
		this.ServeJSON()
	}()
	if this.User == nil {
		resp.Msg = "请先登录！"
		return
	}
	dishid,err := this.GetInt("dishid")
	if err != nil {
		resp.Msg = err.Error()
		return
	}
	status,err := this.GetInt("status")
	if err != nil {
		resp.Msg = err.Error()
		return
	}
	resp  = services.UserDishCollect(dishid,this.User.Id,status)
}

