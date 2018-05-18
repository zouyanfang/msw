package services

import (
	"msw/models"
	"msw/utils"
	"fmt"
)

//菜单业务

//获取菜单列表
func GetMenuList(page,pageSize int,condition string,paras []interface{}) (resp models.MenuResp) {
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
	count,err := models.CountMenu()
	if err != nil {
		resp.Msg = err.Error()
		return
	}
	pages := utils.PageCount(utils.PAGESIZE6,count)
	fmt.Println(pages,page)
	if utils.IsHaveNext(page,pages) {
		resp.Page = page
	}else {
		resp.Page = 0
	}
	resp.Ret = 200
	resp.Object = menuList
	return
}

//获取用户菜单
func GetUserMenu(uid int)(resp models.BaseMsgResp){
	resp.Ret = 403
	usermenu,err := models.GetUserMenu(uid)
	if err != nil {
		resp.Msg = err.Error()
		return
	}
	resp.Object = usermenu
	resp.Ret = 200
	return
}

func AddToMenu(menuid ,dish_id int)(resp models.BaseMsgResp){
	resp.Ret = 403
	count,err := models.FindMenu(dish_id,menuid)
	if err != nil {
		resp.Msg = err.Error()
		return
	}
	err = models.AddDishNumber(menuid)
	if err != nil {
		resp.Msg = err.Error()
		return
	}
	if count != 0 {
		resp.Msg = "该菜谱已经在该菜单内！不要重复添加"
		return
	}
	err = models.AddToUserMenu(dish_id,menuid)
	if err != nil {
		resp.Msg = err.Error()
		return
	}
	resp.Ret = 200
	return
}

func CreateNewMenu(uid int, menuname string,menuimg string,describe string)(resp models.BaseMsgResp){
	count,err := models.CheckMenuName(uid,menuname)
	if err != nil {
		resp.Msg = err.Error()
		return
	}
	if count != 0 {
		resp.Msg = "菜单重名了！请换一个名字"
	}
	err = models.InsertMenu(uid,menuname,menuimg,describe)
	if err != nil {
		resp.Msg = "创建失败"+err.Error()
		return
	}
	resp.Ret = 200
	return
}

func GetMenuDetatil(menuid int)(resp models.MenuDetailResp){
	m ,err := models.GetMenuInfo(menuid)
	if err != nil {
		resp.Msg = err.Error()
		return
	}
	resp.Object = m
	name ,err := models.GetUserNameByMenu(m.Uid)
	if err != nil {
		resp.Msg = err.Error()
		return
	}
	img,err := models.GetUserImgByMenu(m.Uid)
	if err != nil {
		resp.Msg = err.Error()
		return
	}
	resp.Img = img
	resp.Name = name
	dish ,err :=models.GetDishInfoByMenuId(menuid)
	if err != nil {
		resp.Msg = err.Error()
	}
	resp.Dish = dish
	resp.Ret = 200
	return
}

func GetSearchResutl(condition string)(resp models.SearchResp){
	dish,err := models.GetDishByVague(condition)
	if err != nil {
		resp.Msg = err.Error()
		return
	}

	menu,err := models.GetMenuByVague(condition)
	if err != nil {
		resp.Msg = err.Error()
		return
	}
	if dish == nil && menu == nil {
		resp.Status = false
		return
	}
	resp.Menu = menu
	resp.Dish = dish
	resp.Status = true
	return
}



