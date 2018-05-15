package services

import (
	"msw/models"
	"msw/utils"
	"fmt"
	"strings"
)

//菜谱业务

//首页菜谱列表
func GetDishList(page ,pageSize int,condition string ,paras []string)(resp models.BaseMsgResp){
	resp.Ret = 403
	dishs,err := models.GetDishList(utils.StartIndex(page,pageSize),pageSize,condition,paras)
	if err != nil{
		if err.Error()!=utils.ERRROWS{
			resp.Msg = "查询菜谱列表失败!"
			return
		}
		resp.Msg = "查不到数据"
		resp.Ret = 200
		return
	}
	resp.Object = dishs
	return
}


//获取菜谱大全
func GetAllDishList(page,pageSize int,condition string,paras []interface{}) (resp models.DishResp)  {
	resp.Ret = 403
	allDish,err := models.GetAllDishList(utils.StartIndex(page,pageSize),pageSize,condition,paras)
	if err != nil{
		if err.Error() != utils.ERRROWS {
			resp.Msg = "获取菜谱大全失败!"
			return
		}
		resp.Msg = "查不到数据"
		resp.Ret = 200
		return
	}
	count,err := models.GetDishCount(condition,paras)
	if err != nil {
		resp.Msg = "查询总条数失败！"
		return
	}
	pages := utils.PageCount(pageSize,count)
	if utils.IsHaveNext(page,pages) {
		resp.Page = page
	}else {
		resp.Page = 0
	}
	resp.Ret = 200
	resp.Object = allDish
	return
}

//获取菜谱信息
func GetDishInfo(uid,dishId int) (dish *models.DishInfo,step []models.StepInfo,mainMaterial []string,secondMaterial []string) {
	dishInfo,err := models.GetDishInfo(uid,dishId)
	stepInfo,err := models.GetDishStep(uid,dishId)
	mainMaterial = strings.Split(dishInfo.MainMaterial,"/")
	secondMaterial = strings.Split(dishInfo.SecondMaterial,"/")
	if err != nil {
		fmt.Println(err.Error())
	}
	return dishInfo,stepInfo,mainMaterial,secondMaterial
}


