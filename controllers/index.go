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
	page := 1
	resp := services.GetDishList(page,utils.PAGESIZE3,"ORDER BY release_date DESC",[]string{})
	this.Data["dish"] = resp.Object
	fmt.Print(resp.Object)
	fmt.Print(resp.Object)
	this.TplName = "site/index.html"
}
