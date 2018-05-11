package models

import (
	"github.com/astaxie/beego/orm"
)

//菜谱
type Dish struct {
	Id             int
	Uid            int    `description:"用户id"`
	DishName       string `description:"菜名"`
	DishImg        string `description:"菜成品图"`
	IsSlideshow    int    `description:"是否是轮播图"`
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

//获取菜谱列表
func GetDishList(pageIndex, pageSize int, condition string, paras []string) (d []Dish, err error) {
	sql := `SELECT * FROM dish WHERE 1=1 `
	if condition != "" {
		sql += condition
	}
	sql += " LIMIT ?,?"
	_, err = orm.NewOrm().Raw(sql, paras, pageIndex, pageSize).QueryRows(&d)
	return
}

type DishInfo struct {
	Uid            int
	UserImg        string `description:"用户头像"`
	Name           string `description:"用户昵称"`
	DishImg        string `description:"菜谱成品图"`
	MainMaterial   string `description:"主料"`
	SecondMaterial string `description:"辅料"`
	Tasty          string `description:"口味"`
	DishSystem     string `description:"菜系"`
	DishId         int    `description:"菜谱id"`
	Step           int    `description:"步骤"`
	StepImg        string `description:"步骤图"`
	StepDescribe   string `description:"步骤描述"`
}

func GetDishInfo(uid,dishId int)(dishInfo []DishInfo,err error)  {
	sql := `SELECT d.uid,u.user_img,u.name,d.dish_img,d.main_material,d.second_material,
			d.tasty,d.dish_system,ds.dish_id,ds.step,ds.step_img,ds.step_describe
			FROM dish d
			LEFT JOIN users u ON d.uid = u.id
			LEFT JOIN dish_step ds ON d.id = ds.dish_id
			WHERE d.uid = ? AND ds.dish_id = ? `
	_,err = orm.NewOrm().Raw(sql,uid,dishId).QueryRows(&dishInfo)
	return
}
