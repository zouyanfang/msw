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
	page,_ := this.GetInt("page",1)
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
	if this.User == nil {
		this.Data["status"] = 0
	}else {
		s ,_:= models.FindUserCollection(dishId,this.User.Id)
		if s != nil {
			this.Data["status"] = s.Status
		}else {
			this.Data["status"] = 0
		}

	}
	dishInfoResp := services.GetDishInfo(uid,dishId)
	dishCommnet := services.GetDishComment(page,utils.PAFESIZE5,dishId)
	this.Data["dishInfo"] = dishInfoResp.DishDetail
	this.Data["stepInfo"] = dishInfoResp.StepDetail
	this.Data["mainMaterial"] = dishInfoResp.MainMaterial
	this.Data["secondMaterial"] = dishInfoResp.SecondMaterial
	this.Data["dishComment"] = dishCommnet.Object
	this.Data["page"] = dishCommnet.Page
	this.Data["dishid"] = dishId
	this.Data["type"] = 2
	this.Data["count"] = dishCommnet.Count
	this.IsNeddTemplate()
	this.TplName = "site/foodetail.html"
}


//获取更多评论
func (this *DishController)GetTalk(){
	var resp models.DishResp
	resp.Ret = 403
	defer func(){
		this.Data["json"] = resp
		this.ServeJSON()
	}()
	page,err := this.GetInt("page")
	if err != nil {
		resp.Msg = "获取参数失败"
		return
	}
	dishid,err := this.GetInt("dishid")
	if err != nil {
		resp.Msg = "获取参数失败"
		return
	}
	resp = services.GetDishComment(page,utils.PAFESIZE5,dishid)
	return
}