package models

import (
	"github.com/astaxie/beego/orm"
	_"github.com/go-sql-driver/mysql"
	"msw/utils"
	"fmt"
)

func init(){
	err :=orm.RegisterDataBase("default","mysql",utils.MYSQL_URL)
	if err != nil {
		fmt.Print(err.Error())
	}
}
