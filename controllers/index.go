package controllers

import (
	"msw/services"
	"msw/utils"
	"fmt"
)

//首页

type IndexController struct {
	BaseController
}

func (this *IndexController)Index(){
	//首页轮播图展示
	this.Data["type"] = 1
	this.IsNeddTemplate()
	page := 1
	slideShow := services.GetDishList(page,utils.PAGESIZE3,"AND release_role = 2 AND is_slideshow = 1 ORDER BY release_date DESC",nil)
	this.Data["dish"] = slideShow.Object

	//首页菜单列表展示
	menuShow := services.GetMenuList(page,utils.PAGESIZE4,"ORDER BY popular_count DESC",nil)
	this.Data["menu"] = menuShow.Object
	//首页菜谱列表展示
	dishShow := services.GetDishList(page,utils.PAGESIZE6,"AND is_slideshow = 2 ORDER BY release_date DESC",nil)
	this.Data["dishShow"] = dishShow.Object
	//TODO
	//首页最新留言展示
	messageShow := services.GetMessageBoardList(page,utils.PAGESIZE3,"ORDER BY mb.create_time DESC",nil)
	this.Data["message"] = messageShow.Object
	/*fmt.Print(resp.Object)*/
	this.TplName = "site/index1.html"
}

/*//菜谱中心
func (this *IndexController)DishCenter (){
	page := 1
	allDish := services.GetAllDishList(page,utils.PAFESIZE9,"ORDER BY d.popular_count DESC",nil)
	this.Data["allDish"] = allDish.Object
	this.Data["type"] = 2
	this.IsNeddTemplate()
	this.TplName = "site/food.html"
}*/



/*//菜谱详情
	func (this *IndexController)FoodDetail(){
	this.Data["type"] = 2
	this.IsNeddTemplate()
	this.TplName = "site/foodetail.html"
}*/


//留言板页面
func (this *IndexController)Message(){
	this.Data["type"] = 4
	this.IsNeddTemplate()
	resp :=services.GetUserMsgList(1,utils.PAGESIZE9,"ORDER BY create_time DESC ",nil)
	this.Data["msg"] = resp.Object
	this.Data["count"] = resp.Count
	this.Data["page"] = resp.Page
	this.TplName = "site/message.html"
}


//查询菜单菜谱页面
func (this *IndexController)Search(){
	this.Data["type"] = 0
	name := this.GetString("Search")
	if name == ""{
		this.Abort("404")
	}
	condition := "'%"+name+"%'"
	resp := services.GetSearchResutl(condition)
	fmt.Println(resp)
	this.Data["status"] = resp.Status
	this.Data["menu"] = resp.Menu
	this.Data["dish"] = resp.Dish
	this.IsNeddTemplate()
	this.TplName = "site/search.html"
}

