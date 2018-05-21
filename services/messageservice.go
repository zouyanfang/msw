package services

import (
	"msw/models"
	"msw/utils"
)

//留言板业务

//获取留言列表
func GetMessageBoardList(page,pageSize int,condition string,paras []interface{}) (resp models.BaseMsgResp) {
	resp.Ret = 403
	messageList,err := models.GetMessageList(utils.StartIndex(page,pageSize),pageSize,condition,paras)
	if err != nil {
		if err.Error() != utils.ERRROWS {
			resp.Msg = "服务器bug"
			return
		}
		resp.Msg = "查不到数据"
		resp.Ret = 200
	}
	resp.Object = messageList
	resp.Ret = 200
	return
}

func GetUserMsgList(page,pageSize int,condition string,paras []interface{}) (resp models.UserMsgResp) {
	resp.Ret = 403
	messageList,err := models.GetUserMessage(utils.StartIndex(page,pageSize),pageSize,condition,paras)
	if err != nil {
		if err.Error() != utils.ERRROWS {
			resp.Msg = "服务器bug"
			return
		}
		resp.Msg = "查不到数据"
		resp.Ret = 200
	}
	count,err := models.GetUserMessageCount()
	if err != nil {
		resp.Msg = err.Error()
	}
	pages := utils.PageCount(utils.PAGESIZE6,count)
	if utils.IsHaveNext(page,pages) {
		resp.Page = page
	}else {
		resp.Page = 0
	}
	resp.Count = count
	resp.Object = messageList
	resp.Ret = 200
	return
}