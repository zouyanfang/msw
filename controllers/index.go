package controllers

import (
	"msw/services"
	"msw/utils"
)

//首页

type IndexController struct {
	BaseController
}

func (this *IndexController)Index(){
	//首页轮播图展示
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
	this.TplName = "site/index.html"
}

