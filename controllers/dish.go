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
	paras:= []interface{}{}
	conditon := ""
	page,_ := this.GetInt("page",0)
	uid,_ := this.GetInt("uid")
	if uid != 0 {
		conditon += " AND  "
		paras = append(paras,uid)
	}
	dishId,_ := this.GetInt("dishId")
	if dishId != 0  {
		conditon += " AND "
		paras = append(paras,dishId)
	}
	dishInfoResp := services.GetDishInfo(uid,dishId)
	dishCommnet := services.GetDishComment(page,utils.PAFESIZE5," ORDER BY ds.comment_date DESC",paras)

	this.Data["dishInfo"] = dishInfoResp.DishDetail
	this.Data["stepInfo"] = dishInfoResp.StepDetail
	this.Data["mainMaterial"] = dishInfoResp.MainMaterial
	this.Data["secondMaterial"] = dishInfoResp.SecondMaterial
	this.Data["dishComment"] = dishCommnet.Object
	this.Data["type"] = 2
	this.IsNeddTemplate()
	this.TplName = "site/foodetail.html"
}
