package services

import (
	"msw/models"
	"github.com/astaxie/beego/orm"
	"fmt"
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

//用户收藏菜谱
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

//用户收藏菜单
func  UserMenuCollect(menu_id,uid ,status int)(resp models.BaseMsgResp){
	fmt.Println("menuid",menu_id)
	o := orm.NewOrm()
	defer func (){
		fmt.Println(resp.Msg)
		if resp.Msg != ""{
			o.Rollback()
			return
		}
		o.Commit()
	}()
	exist,err:= models.FindMenuIsExits(menu_id,uid,o)
	if err != nil {
		resp.Msg = err.Error()+"123123"
		return
	}
	if exist == 0{
		fmt.Println("2131313")
		err := models.CollectMenu(menu_id,uid,o)
		if err != nil {
			resp.Msg = err.Error()+"456"
			return
		}
		resp.Ret = 200
		return
	}
	err = models.UPDATEMenuCollection(uid,menu_id,status,o)
	if err != nil {
		resp.Msg = err.Error()+"789"
		return
	}
	resp.Ret = 200
	return
}

func GetMenuListByUid(uid int)(resp models.UserMenuResp){
	//获取我的菜单
	m ,err := models.GetMenuListByUid(uid)
	if err != nil {
		resp.Msg = err.Error()
		return
	}
	fmt.Println("12312313:",m)
	mc ,err := models.GetMenuCollectListByUid(uid)
	if err != nil {
		resp.Msg = err.Error()
		return
	}
	resp.CollectMenu = mc
	resp.Menu = m
	resp.Ret = 200
	return
}

func DeleteMenu(menuid int,uid int)(resp models.BaseMsgResp){
	count ,err := models.IsUserMenu(menuid,uid)
	if err != nil {
		resp.Msg = err.Error()
		return
	}
	if count == 0 {
		resp.Msg = "你没有该权限删除菜单"
		return
	}
	err = models.DeleteMenu(menuid)
	if err != nil {
		resp.Msg = err.Error()
		return
	}
	resp.Ret = 200
	return
}

func ModifyUserByUid(uid int, condition string,paras interface{} )(resp models.BaseMsgResp){
	if condition == ""{
		resp.Msg = "没有更新的内容"
		return
	}
	err := models.ModifyUserBaseMsg(uid,condition,paras)
	if err != nil{
		resp.Msg = err.Error()
		return
	}
	resp.Ret = 200
	return
}

func GetUserDish(uid int)(resp models.UserDishResp){
	collectdish,err := models.GetUserCollectDish(uid)
	if err != nil {
		resp.Msg = err.Error()
		return
	}
	resp.CollectDish = collectdish
	dish ,err := models.GetMyDish(uid)
	if err != nil {
		resp.Msg = err.Error()
		return
	}
	resp.Dish = dish
	resp.Ret = 200
	return
}

func UpdateDishDetail(dishid int,taste,system string,main ,second string)(resp models.BaseMsgResp){
	err := models.UpdateDish(dishid,taste,system,main,second)
	if err != nil {
		resp.Msg = err.Error()
		return
	}
	resp.Ret = 200
	return
}
