package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
)

type Menu struct {
	Id           int
	Uid          int    `description:"用户id"`
	MenuName     string `description:"菜单名"`
	MenuImg      string `description:"菜单图片"`
	CreateDate   string `description:"创建日期"`
	MenuDescribe string `description:"菜单描述"`
	CollectCount int    `description:"收藏人数"`
	PopularCount int    `description:"人气"`
	Counts        int    `description:"菜谱的总数"`
}

//获取菜单列表
func GetMenuList(startIndex,pageSize int,condition string,paras []interface{}) (menuList []Menu,err error)  {
	sql := `SELECT * FROM menu WHERE 1 = 1 `
	if condition != "" {
		sql += condition
	}
	sql += " LIMIT ?,?"
	_,err = orm.NewOrm().Raw(sql,paras,startIndex,pageSize).QueryRows(&menuList)
	fmt.Println(sql,menuList)
	return
}

//获取用户menu
func GetUserMenu(uid int)(m []Menu,err error){
	sql := `SELECT * FROM menu WHERE uid = ?`
	o := orm.NewOrm()
	_,err = o.Raw(sql,uid).QueryRows(&m)
	fmt.Println(m)
	return
}

//查看是否在菜单内
func FindMenu(dish_id,menu_id int)(count int,err error){
	sql := `SELECT COUNT(1) FROM menu_dish WHERE dish_id = ? AND menu_id = ? `
	o  := orm.NewOrm()
	err = o.Raw(sql,dish_id,menu_id).QueryRow(&count)
	return
}

func AddDishNumber(menu_id int)(err error){
	sql := `UPDATE menu SET counts = counts+1 WHERE id = ?`
	o := orm.NewOrm()
	_,err = o.Raw(sql,menu_id).Exec()
	return
}
//添加至菜单
func AddToUserMenu(dish_id int,menu_id int)(err error){
	sql := `INSERT INTO menu_dish (menu_id,dish_id) VALUES (?,?) `
	o := orm.NewOrm()
	_,err = o.Raw(sql,menu_id,dish_id).Exec()
	return
}

//新建菜单
func InsertMenu(uid int,menuname string,menuimg string,describe string)(err error){
	sql := `INSERT INTO menu (uid,menu_name,menu_img,menu_describe,create_date) VALUES (?,?,?,?,NOW())`
	o := orm.NewOrm()
	_,err = o.Raw(sql,uid,menuname,menuimg,describe).Exec()
	return
}

//检查菜单名
func CheckMenuName(uid int,menuname string)(count int,err error){
	sql := `SELECT COUNT(1) FROM menu WHERE uid = ? AND menu_name = ?`
	o := orm.NewOrm()
	err = o.Raw(sql,uid,menuname).QueryRow(&count)
	return
}

func CountMenu()(count int,err error){
	sql := `SELECT COUNT(1) FROM menu`
	o := orm.NewOrm()
	err = o.Raw(sql).QueryRow(&count)
	return
}

//获取菜单信息
func GetMenuInfo(menuid int)(m Menu,err error){
	sql := `SELECT * FROM menu WHERE id = ?`
	o := orm.NewOrm()
	err = o.Raw(sql,menuid).QueryRow(&m)
	return
}

//获取用户名
func GetUserNameByMenu(uid int)(name string,err error){
	sql := `SELECT name FROM users WHERE id = ?`
	o := orm.NewOrm()
	err = o.Raw(sql,uid).QueryRow(&name)
	return
}

func GetUserImgByMenu(uid int)(menuimg string,err error){
	sql := `SELECT user_img FROM users WHERE id = ?`
	o := orm.NewOrm()
	err = o.Raw(sql,uid).QueryRow(&menuimg)
	fmt.Println(sql,uid)
	return
}

//获取菜谱的
func GetDishInfoByMenuId(menuid int)(dish []Dish,err error){
	sql := `SELECT d.* FROM dish d
			LEFT JOIN menu_dish m on m.dish_id = d.id
			WHERE m.menu_id = ?`
	o := orm.NewOrm()
	_,err = o.Raw(sql,menuid).QueryRows(&dish)
	return
}

//模糊查询菜谱
func GetDishByVague(condition string)(dish []Dish,err error){
	sql := `SELECT * FROM dish WHERE dish_name like `+condition
	o := orm.NewOrm()
	_, err = o.Raw(sql).QueryRows(&dish)
	return
}

func GetMenuByVague(condition string)(menu []Menu,err error){
	sql := `SELECT * FROM menu WHERE menu_name like `+condition
	o := orm.NewOrm()
	_, err = o.Raw(sql).QueryRows(&menu)
	return
}