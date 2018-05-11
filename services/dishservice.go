package services

import (
	"msw/models"
	"msw/utils"
)

//菜谱业务

//首页菜谱列表
func GetDishList(page ,pageSize int,condition string ,paras []string)(resp models.BaseMsgResp){
	resp.Ret = 403
	dishs,err := models.GetDishList(utils.StartIndex(page,pageSize),pageSize,condition,paras)
	if err != nil{
		if err.Error()!=utils.ERRROWS{
			resp.Msg = "服务器bug"
			return
		}
		resp.Msg = "查不到数据"
		resp.Ret = 200
	}
	resp.Object = dishs
	return
}


//菜谱大全

//获取菜谱信息
func GetDishInfo(uid,dishId int) (resp models.BaseMsgResp) {
	resp.Ret = 403
	dishInfo,err := models.GetDishInfo(uid,dishId)
	if err != nil {
		if err.Error() != utils.ERRROWS {
			resp.Msg = "查询菜谱信息失败！"
			return
		}
		resp.Msg = "查不到数据"
		resp.Ret = 200
	}

	resp.Ret = 200
	resp.Object = dishInfo
	return
}
