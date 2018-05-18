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
	fmt.Println(page)
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

	}else if page > pages{
		resp.Page = 1
	}else {
		resp.Page = page
	}
	if page >= pages {
		resp.Next = false
	}else {
		resp.Next = true
	}
	fmt.Println("当前页",page)
	if page == 1 {
		resp.Pref = false
	}else {
		resp.Pref = true
	}

	resp.Ret = 200
	resp.Object = allDish
	return
}

//获取菜谱信息
func GetDishInfo(uid,dishId int) (dishDetail models.DishDetailResp) {
	models.AddPopular(dishId)
	dishInfo,err := models.GetDishInfo(uid,dishId)
	mainMaterial := strings.Split(dishInfo.MainMaterial,"/")
	secondMaterial := strings.Split(dishInfo.SecondMaterial,"/")
	stepinfo ,err := models.GetStepByDish(dishId)
	if err != nil {
		fmt.Println(err.Error())
	}
	dishDetail.DishDetail = dishInfo
	dishDetail.StepDetail = stepinfo
	dishDetail.MainMaterial = mainMaterial
	dishDetail.SecondMaterial = secondMaterial

	return
}

//获取菜谱评论
func GetDishComment(page,pageSize ,dishId int) (resp models.DishResp)  {
	resp.Ret = 403
	dishComment,err := models.GetDishComment(utils.StartIndex(page,pageSize),utils.PAFESIZE5,dishId)
	if err != nil  {
		if err.Error() != utils.ERRROWS{
			resp.Msg = "查询菜谱评论失败!"
			return
		}
		resp.Msg = " 查询不到数据"
		resp.Ret = 200
		return
	}
	count,err := models.GetDishCommentCount(dishId)
	resp.Count = count
	if err != nil{
		resp.Msg = "查询总条数失败!"
		return
	}
	pages := utils.PageCount(utils.PAFESIZE5,count)
	if utils.IsHaveNext(page,pages){
		resp.Page = page
	}else {
		resp.Page = 0
	}
	resp.Object = dishComment
	resp.Ret = 200
	return
}


