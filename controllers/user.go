package controllers



type UserController struct {
	BaseController
}

func (this *UserController)UserCenter(){
	this.Data["type"] = 1
	this.IsNeedTemplate()
	this.TplName = "site/user/dishcenter.html"
}

func (this *UserController)UserMenu(){
	this.Data["type"] = 2
	this.IsNeedTemplate()
	this.TplName = "site/user/menucenter.html"
}

func (this *UserController)IsNeedTemplate(){

	this.Layout = "site/user/template.html"
}

