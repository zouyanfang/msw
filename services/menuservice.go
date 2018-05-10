package services

import (
	"msw/models"
	"msw/utils"
)

//菜单业务

//获取菜单列表
func GetMenuList(page,pageSize int,condition string,paras []interface{}) (resp models.BaseMsgResp) {
	resp.Ret = 403
	menuList,err := models.GetMenuList(utils.StartIndex(page,pageSize),pageSize,condition,paras)
	if err != nil {
		if err.Error() != utils.ERRROWS{
			resp.Msg = "服务器bug"
			return
		}
		resp.Msg = "查不到数据"
		resp.Ret = 200
		return
	}
	resp.Ret = 200
	resp.Object = menuList
	return
}
