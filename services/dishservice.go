package services

import (
	"msw/models"
	"msw/utils"
)

//菜谱业务

//得到菜谱列表
func GetDishList(page ,pageSize int,condition string ,paras []string)(resp models.BaseMsgResp){
	resp.Ret = 403
	pageIndex := utils.StartIndex(page,pageSize)
	dishs,err := models.GetDishList(pageIndex,pageSize,condition,paras)
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


