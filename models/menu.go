package models

import (
	"github.com/astaxie/beego/orm"
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
}

//获取菜单列表
func GetMenuList(startIndex,pageSize int,condition string,paras []interface{}) (menuList []Menu,err error)  {
	sql := `SELECT * FROM menu WHERE 1 = 1 `
	if condition != "" {
		sql += condition
	}
	sql += " LIMIT ?,?"
	_,err = orm.NewOrm().Raw(sql,paras,startIndex,pageSize).QueryRows(&menuList)
	return
}
