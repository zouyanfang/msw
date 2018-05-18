package services

import (
	"msw/models"
	"github.com/astaxie/beego/orm"
)

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

//用户收藏菜单
func UserDishCollect(dishid,uid ,status int)(resp models.BaseMsgResp){
	o := orm.NewOrm()
	defer func (){
		if resp.Msg != "" {
			o.Rollback()
			return
		}
		o.Commit()
	}()
	resp.Ret = 403
	/*判断是否用该收藏字段*/
	exist,_ := models.FindUserCollection(dishid,uid)
	if exist == nil {
		err := models.InsertUserCollection(o,dishid,uid)
		if err != nil {
			resp.Msg = err.Error()
			return
		}
		resp.Ret = 200
		return
	}
	err := models.UpdateCollect(o,dishid,uid,status)
	if err != nil {
		resp.Msg = err.Error()
		return
	}
	err = models.UpdateDishNum(o,dishid,status)
	if err != nil {
		resp.Msg = "更新菜谱收藏数量失败"
		return
	}
	resp.Ret = 200
	return
}

//发布一条信息
func ReleaseUserMsg(uid int,content string)(resp models.BaseMsgResp){
	err := models.InsertUserMsg(uid,content)
	if err != nil {
		resp.Msg = err.Error()
		return
	}
	resp.Ret = 200
	return
}

//func GetUserInfoByIndex(uid int)(resp models.BaseMsgResp){
//	u,err := models.GetUserInfoByIndex(uid)
//	if err != nil {
//		resp.Msg = err.Error()
//		return
//	}
//	dish,err := models.GetUserCollectDish(uid)
//	if err != nil {
//		resp.Msg = err.Error()
//		return
//	}
//	resp.
//	return
//}
