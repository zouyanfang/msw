package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
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

func Register(account string) (count int, err error) {
	o := orm.NewOrm()
	sql := `SELECT COUNT(1) FROM users WHERE account = ?`
	err = o.Raw(sql,account).QueryRow(&count)
	return
}

//添加用户注册信息
func AddUser(account, pwd,userimg string) (err error) {
	o := orm.NewOrm()
	sql := `INSERT INTO users (account,pwd,register_date,user_img)
			VALUES(?,?,NOW(),?)`
	_, err = o.Raw(sql, account, pwd,userimg).Exec()
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

func GetUserCollectDish(uid int)(dish []Dish,err error){
	sql := 	`SELECT d.* FROM dish d  LEFT JOIN  user_collection uc ON d.id = uc.dish_id WHERE uc.uid = ? AND status = 1`
	o := orm.NewOrm()
	_,err =o.Raw(sql,uid).QueryRows(&dish)
	fmt.Println(sql)
	return
}

func GetMyDish(uid int)(dish []Dish,err error){
	sql := `SELECT * FROM dish WHERE uid = ? `
	o := orm.NewOrm()
	_,err = o.Raw(sql,uid).QueryRows(&dish)
	return
}

func DeleteMenu(menuid int)(err error){
	o := orm.NewOrm()
	defer func (){
		if err != nil {
			o.Rollback()
			return
		}
		o.Commit()
	}()
	sql := `DELETE FROM menu WHERE id = ?`
	_,err = o.Raw(sql,menuid).Exec()
	if err != nil {
		return
	}
	sql = "DELETE FROM menu_dish WHERE menu_id = ? "
	_,err = o.Raw(sql,menuid).Exec()
	if err != nil{
		return
	}
	sql = "DELETE FROM user_collection WHERE menu_id = ?"
	_,err = o.Raw(sql,menuid).Exec()
	if err != nil {
		return
	}
	return
}

func GetDishCountByUid(uid int)(count int){
	sql := `SELECT COUNT(1) FROM dish WHERE uid = ?`
	o := orm.NewOrm()
	o.Raw(sql,uid).QueryRow(&count)
	return
}

func GetMenuCountByUid(uid int)(count int){
	sql := `SELECT COUNT(1) FROM menu WHERE uid = ?`
	o := orm.NewOrm()
	o.Raw(sql,uid).QueryRow(&count)
	return
}

func ModifyUserBaseMsg(uid int,condition string,paras interface{})(err error){
	sql := `UPDATE users `
	if condition != ""{
		sql += condition
	}
	sql += "WHERE id = ?"
	o := orm.NewOrm()
	_,err = o.Raw(sql,paras,uid).Exec()
	fmt.Println(sql,uid)
	return
}

func CreateDish(uid int,dishname,dishimg,describe string)(err error){
	sql := `INSERT INTO dish (uid,dish_name,dish_img,release_date,release_role,dish_describe) VALUES (?,?,?,NOW(),2,?)`
	o := orm.NewOrm()
	_,err = o.Raw(sql,uid,dishname,dishimg,describe).Exec()
	fmt.Println(sql,uid,dishname,dishimg,describe,err)
	return
}

func GetLastDish()(id int,err error){
	sql := `SELECT id FROM dish ORDER BY id DESC LIMIT 0,1`
	o := orm.NewOrm()
	err = o.Raw(sql).QueryRow(&id)
	return
}

func UpdateDish(dishid int,taste,system string,main ,second string)(err error){
	sql := `UPDATE dish SET main_material = ?,second_material = ?,tasty = ?,dish_system = ? WHERE id = ?`
	o := orm.NewOrm()
	_,err = o.Raw(sql,main,second,taste,system,dishid).Exec()
	return
}

func CountStep(dishid int,)(count int,err error){
	sql := `SELECT COUNT(1) FROM dish_step  WHERE dish_id = ? AND step_img is NULL `
	o := orm.NewOrm()
	err = o.Raw(sql,dishid).QueryRow(&count)
	return
}

