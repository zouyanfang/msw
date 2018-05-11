package controllers

import (
	"msw/services"
)

//菜谱大全
type DishController struct {
	BaseController
}

type S struct {
	s interface{}

}
func (this *DishController) Dish()  {
	var s S
	defer func (){
		this.Data["json"] = s
		this.ServeJSON()
	}()
	uid,_ := this.GetInt("uid")
	dishId,_ := this.GetInt("dishId")
	dishInfo := services.GetDishInfo(uid,dishId)
	s.s = dishInfo.Object
}