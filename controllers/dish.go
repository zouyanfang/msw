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
	tasty:= this.GetString("tasty")
	if tasty != "" {
		condition += " AND d.tasty = ?"
		paras = append(paras,tasty)
	}
	area := this.GetString("area")
	if area != "" {
		condition += " AND d.dish_system = ?"
		paras = append(paras,area)
	}
	resp = services.GetAllDishList(page,utils.PAFESIZE9,condition,paras)

}

//菜谱详情
func (this *DishController)GetDishDetail(){
	/*var resp models.BaseMsgResp
	resp.Ret = 403
	defer func() {
		this.Data["json"] = resp
		this.ServeJSON()
	}()*/

	uid,_ := this.GetInt("uid")
	dishId,_ := this.GetInt("dishId")
	dishInfo,stepInfo,mainMaterial,secondMaterial := services.GetDishInfo(uid,dishId)
	/*foodInfo := map[string]interface{}{"dishInfo":dishInfo,"stepInfo":stepInfo,"material":material}
	resp.Object = foodInfo
	resp.Ret = 200*/
	this.Data["dishInfo"] = dishInfo
	this.Data["stepInfo"] = stepInfo
	this.Data["mainMaterial"] = mainMaterial
	this.Data["secondMaterial"] = secondMaterial
	this.Data["type"] = 2
	this.IsNeddTemplate()
	this.TplName = "site/foodetail.html"
}
