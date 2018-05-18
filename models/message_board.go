package models

import (
	"github.com/astaxie/beego/orm"
)

//留言板
type MessageBoard struct {
	Id         int
	Title      string `description:"留言标题"`
	Content    string `description:"留言内容"`
	CreateTime string `description:"创建时间"`
	Name       string `description:"用户昵称"`
	UserImg    string `description:"用户图片"`
}

//获取留言信息
func GetMessageList(startIndex, pageSize int, condition string, paras []interface{}) (list []MessageBoard, err error) {
	sql := `SELECT * FROM message_board mb
			LEFT JOIN users u ON u.id = mb.uid
			WHERE 1 = 1 `
	if condition != "" {
		sql += condition
	}
	sql += " LIMIT ?,?"
	_, err = orm.NewOrm().Raw(sql, paras, startIndex, pageSize).QueryRows(&list)
	return
}

//获取留言信息
func GetUserMessage(startIndex, pageSize int, condition string, paras []interface{}) (list []MessageBoard, err error) {
	sql := `SELECT um.*,u.Name,u.user_img FROM user_message um
			LEFT JOIN users u ON u.id = um.uid
			WHERE 1 = 1 `
	if condition != "" {
		sql += condition
	}
	sql += " LIMIT ?,?"
	_, err = orm.NewOrm().Raw(sql, paras, startIndex, pageSize).QueryRows(&list)
	return
}

func GetUserMessageCount()(count int,err error){
	sql := `SELECT COUNT(1) FROM user_message `
	o := orm.NewOrm()
	err = o.Raw(sql).QueryRow(&count)
	return
}
