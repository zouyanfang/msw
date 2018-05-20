package controllers

import (
	"msw/services"

	"msw/utils"
	"fmt"
	"msw/models"
)

//TODO
//菜单中心
type MenuCenterController struct {
	BaseController
}

//菜单中心
func (this *MenuCenterController)MenuCenter(){
	this.Data["type"] = 3
	this.IsNeddTemplate()
	desc,_ := this.GetInt("desc")
	fmt.Println(desc)
	condition := ""
	if desc == 0 {
		condition += " ORDER BY collect_count DESC "
	}else {
		condition += "ORDER BY popular_count DESC "
	}
	resp := services.GetMenuList(1,utils.PAGESIZE6,condition,nil)
	this.Data["menu"] = resp.Object
	this.Data["page"] = resp.Page
	this.Data["desc"] = desc
	this.TplName = "site/menu.html"
}

func (this *MenuCenterController)LoadMoreMenu(){
	var resp models.MenuResp
	defer func() {
		this.Data["json"] = resp
		this.ServeJSON()
	}()
	page,_ := this.GetInt("page")
	desc ,_:= this.GetInt("desc")
	condition := ""
	if desc == 0 {
		condition += " ORDER BY collect_count DESC "
	}else {
		condition += "ORDER BY popular_count DESC "
	}
	resp = services.GetMenuList(page,utils.PAGESIZE6,condition,nil)
}

//菜单详情
func (this *MenuCenterController)MenuDetail(){
	fmt.Print("woshiqdknlnkjlkjvnkja")
	this.Data["type"] = 3
	menuid ,_ := this.GetInt("menuid")
	resp := services.GetMenuDetatil(menuid)
	this.Data["dish"] = resp.Dish
	this.Data["menu"] = resp.Object
	this.Data["name"] = resp.Name
	this.Data["img"] = resp.Img
	fmt.Println("我去是")
	if this.User == nil {
		this.Data["status"] = 0
		fmt.Println("123123")
	}else {
		status ,_:= models.GetMenuStatus(menuid,this.User.Id)
		this.Data["status"] = status

	}
	this.IsNeddTemplate()
	this.TplName = "site/menudetail.html"
}





