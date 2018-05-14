package routers

import (
	"msw/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/",&controllers.IndexController{},"get:Index")
	beego.Router("/login",&controllers.AccountController{},"get:Login;post:VerifyPassword")
	beego.Router("/loginout",&controllers.AccountController{},"get:LoginOut")
	beego.Router("/register",&controllers.AccountController{},"get:Register")
	beego.Router("/toregister",&controllers.AccountController{},"get:ToRegister")
	beego.AutoRouter(&controllers.IndexController{}) //首页列表展示
	beego.AutoRouter(&controllers.DishController{}) //菜谱大全
	beego.AutoRouter(&controllers.UserController{}) //用户中心
}
