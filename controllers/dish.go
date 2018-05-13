package controllers

import (
	"msw/services"
	"msw/utils"
	"msw/models"
)

//菜谱大全
type DishController struct {
	BaseController
}

//获取菜谱大全列表
func (this *DishController) GetAllDishList()  {
	page := 1
	allDish := services.GetAllDishList(page,utils.PAFESIZE9,"ORDER BY d.popular_count DESC",nil)
	this.Data["allDish"] = allDish.Object
	this.Data["type"] = 2
	this.IsNeddTemplate()
	this.TplName = "site/food.html"
}

func (this *DishController) GetConditionDishList()  {
	var resp models.DishResp
	resp.Ret = 403
	defer func() {
		this.Data["json"] = resp
		this.ServeJSON()
	}()
	condition := ""
	paras := []interface{}{}
	page,_:= this.GetInt("page",0)
	taste:= this.GetString("taste")
	if taste != "" {
		condition += " tasty = ?"
		paras = append(paras,taste)
	}
	area := this.GetString("area")
	if area != "" {
		condition += " dish_system = ?"
		paras = append(paras,area)
	}
	resp = services.GetAllDishList(page,utils.PAFESIZE9,condition,paras)

}

func (this *DishController) GetDishInfo()  {
	uid,_ := this.GetInt("uid")
	dishId,_ := this.GetInt("dishId")
	dishInfo := services.GetDishInfo(uid,dishId)
	this.Data["dishInfo"] = dishInfo
}