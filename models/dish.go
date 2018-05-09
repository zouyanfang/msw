package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
)

type Dish struct {
	Id             int
	Uid            int    //用户id
	DishName       string //菜名
	DishImg        string //菜成品图
	ReleaseDate    string //发布时间
	MainMaterial   string //主料
	SecondMaterial string //辅料
	ReleaseRole    int    //发布角色 0 官方 1用户
	Tasty          string //口味
	DishSystem     string //菜系
	DishDescribe   string //菜谱描述
	CollectCount   int    //收藏人数
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

