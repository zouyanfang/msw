package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id           int
	Name         string `description:"名字"`
	Account      string `description:"账号"`
	Pwd          string `description:"密码"`
	UserImg      string `description:"头像"`
	RegisterDate string `description:"注册日期"`
}

type UserCollection struct {
	Id          int
	Uid         int    `description:"用户名"`
	DishId      int    `description:"菜谱id"`
	MenuId      int    `description:"菜单id"`
	CollectTime string `description:"收藏时间"`
	Status      int    `description:"收藏状态"`
}

func Login(account, pwd string) (user *User, err error) {
	o := orm.NewOrm()
	sql := `SELECT * FROM users WHERE account=? AND pwd=? `
	err = o.Raw(sql, account, pwd).QueryRow(&user)
	return
}

func Register(name string) (count int, err error) {
	o := orm.NewOrm()
	sql := `SELECT COUNT(1) FROM users WHERE name = ?`
	err = o.Raw(sql, name).QueryRow(&count)
	return
}

//添加用户注册信息
func AddUser(account, pwd string) (err error) {
	o := orm.NewOrm()
	sql := `INSERT INTO users (account,pwd,register_date)
			VALUES(?,?,NOW())`
	_, err = o.Raw(sql, account, pwd).Exec()
	return
}

//初始化用户收藏表
func InsertUserCollection(o orm.Ormer,dishid int, uid int) (err error) {
	sql := `INSERT INTO user_collection (dish_id,uid,status,collect_time) VALUES (?,?,?,NOW())`
	_, err = o.Raw(sql, dishid, uid, 1).Exec()
	return
}

//更新收藏状态
func UpdateCollect(o orm.Ormer,dish_id, uid int, status int) (err error) {
	sql := "UPDATE user_collection SET status = ? WHERE dish_id = ? AND uid = ?"
	_, err = o.Raw(sql, status, dish_id, uid).Exec()
	return
}

//找到用户收藏表
func FindUserCollection(dish_id, uid int)(c *UserCollection,err error){
	o := orm.NewOrm()
	sql := `SELECT * FROM user_collection WHERE dish_id = ? AND uid = ?`
	err = o.Raw(sql,dish_id,uid).QueryRow(&c)
	return
}

//更新菜谱收藏量
func UpdateDishNum(o orm.Ormer,dishId,status int)(error){
	sql := ""
	if status == 0 {
		sql = `UPDATE dish SET collect_count = collect_count - 1 WHERE id = ? `
	}else if status == 1 {
		sql = `UPDATE dish SET collect_count = collect_count +1 WHERE id = ? `
	}
	_,err := o.Raw(sql,dishId).Exec()
	return err
}

//插入一条留言
func  InsertUserMsg(uid int , content string)(err error){
	sql := `INSERT INTO user_message (uid,content,create_time) VALUES (?,?,NOW())`
	o := orm.NewOrm()
	_,err = o.Raw(sql,uid,content).Exec()
	return
}

func GetUserInfoByIndex(uid int )(u User,err error){
 	sql := `SELECT * FROM users WHERE id = ?`
 	o := orm.NewOrm()
 	err = o.Raw(sql,uid).QueryRow(&u)
 	return
}

func GetUserCollectDish(uid int)(dish []Dish,err error){
	sql := 	`SELECT * FROM dish WHERE uid = ?`
	o := orm.NewOrm()
	_,err =o.Raw(sql,uid).QueryRows(&dish)
	return
}




