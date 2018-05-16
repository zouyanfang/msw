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


