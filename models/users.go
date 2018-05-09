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
