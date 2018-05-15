package controllers



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
		this.Redirect("/",302)
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


