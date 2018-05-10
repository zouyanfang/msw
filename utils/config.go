package utils

import "github.com/astaxie/beego"

const (
	PAGESIZE3 = 3
	PAGESIZE4 = 4
	PAGESIZE6 = 6
	ERRROWS   = "<QuerySeter> no row found"
)

var (
	MYSQL_URL string
)

func init() {
	runmode := beego.AppConfig.String("runmode")
	config, err := beego.AppConfig.GetSection(runmode)
	if err != nil {
		panic("配置文件初始化失败")
	}
	MYSQL_URL = config["mysql_url"]
}
