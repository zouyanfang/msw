package routers

import (
	"msw/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/",&controllers.IndexController{},"get:Index")
	beego.Router("/login",&controllers.AccountController{},"get:Login;post:VerifyPassword")
	beego.Router("/logout",&controllers.AccountController{},"get:LogOut")
	beego.Router("/register",&controllers.AccountController{},"get:Register")
	beego.Router("/toregister",&controllers.AccountController{},"get:ToRegister")
	beego.AutoRouter(&controllers.IndexController{})
}
