package controllers

import (
	"github.com/astaxie/beego"
	"msw/models"
)

type AccountController struct {
	beego.Controller
}

//登录
func (c *AccountController) Login()  {
	c.Data["state"] = 0
	c.TplName = "site/login.html"
}

//验证密码
func (c *AccountController) VerifyPassword()  {
	resultMap := make(map[string]interface{})
	resultMap["ret"] = 304
	defer func() {
		c.Data["json"] = resultMap
		c.ServeJSON()
	}()
	account := c.GetString("account")
	password := c.GetString("password")
	//password = utils.MD5(password)
	user,_ := models.Login(account,password)
	if user == nil{
		resultMap["msg"] ="登录失败！用户名或密码不正确"
		return
	}else {
		//保存用户cookie
		c.Ctx.SetCookie("uid",user.Account)
		c.Ctx.SetCookie("pid",user.Pwd)
		resultMap["ret"] = 200
		resultMap["msg"] = "登录成功"
		return
	}

}

//退出登录
func (c *AccountController) LogOut ()  {
	uid := c.Ctx.GetCookie("uid")
	if uid != ""{
		//清除cookie
		c.Ctx.SetCookie("uid","0",-1)
		c.Ctx.SetCookie("pid","0",-1)
	}
	c.Ctx.Redirect(302,"/account")
}

//注册
func (c *AccountController) Register()  {
	resultMap := make(map[string]interface{})
	resultMap["ret"] = 304
	defer func() {
		c.Data["json"] = resultMap
		c.ServeJSON()
	}()
	account := c.GetString("account")
	password := c.GetString("password")
	count,_ := models.Register(account)
	if count > 0 {
		resultMap["msg"] = "该用户名已被注册过"
		return
	}
	//添加用户注册信息
	err := models.AddUser(account,password)
	if err != nil {
		resultMap["msg"] = "添加注册用户失败"
		return
	}
	resultMap["msg"] = "注册成功"
	resultMap["ret"] = 200
	return
}

func (this *AccountController)ToRegister(){
	this.Data["state"] = 1
	this.TplName = "site/login.html"
}
