package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
)

//菜谱
type Dish struct {
	Id             int
	Uid            int    `description:"用户id"`
	DishName       string `description:"菜名"`
	DishImg        string `description:"菜成品图"`
	ReleaseDate    string `description:"发布时间"`
	MainMaterial   string `description:"主料"`
	SecondMaterial string `description:"辅料"`
	ReleaseRole    int    `description:"发布角色 0 官方 1用户"`
	Tasty          string `description:"口味"`
	DishSystem     string `description:"菜系"`
	DishDescribe   string `description:"菜谱描述"`
	CollectCount   int    `description:"收藏人数"`
	PopularCount   int    `description:"人气"`
}


func GetDishList(pageIndex,pageSize int,condition string,paras []string)(d []Dish,err error){
	sql := `SELECT * FROM dish WHERE 1=1 `
	if condition != ""{
		sql += condition
	}
	sql += " LIMIT ?,?"
	fmt.Print(sql,paras,pageIndex,pageSize)
	_,err = orm.NewOrm().Raw(sql,paras,pageIndex,pageSize).QueryRows(&d)
	return
}

//获取菜谱的总条数
func GetDishListCount(condition string,paras []string)(count int,err error){
	sql := `SELECT COUNT(1) FROM dish WHERE 1=1 `
	if condition != ""{
		sql += condition
	}
	err = orm.NewOrm().Raw(sql,paras).QueryRow(&count)
	return
}

