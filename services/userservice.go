package services

import "msw/models"

//发布菜谱评论

func ReleaseDishTalk(dish_id,uid int,content string)(resp models.BaseMsgResp){
	resp.Ret = 403
	if content == ""{
		resp.Msg = "没有输入内容，请重新输入"
		return
	}
	err := models.InsertTalk(dish_id,uid,content)
	if err != nil {
		resp.Msg = "发布评论失败！请稍后重试"
		return
	}
	resp.Ret = 200
	return
}
