package main

import (
	_ "webnote/routers"

	"webnote/helper"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	mysqlU := beego.AppConfig.String("mysqlname")
	mysqlP := beego.AppConfig.String("mysqlpwd")
	mysqlDB := beego.AppConfig.String("dbname")
	handle := mysqlU + ":" + mysqlP + "@/" + mysqlDB + "?charset=utf8"
	orm.RegisterDataBase("default", "mysql", handle)
	orm.RunSyncdb("default", false, true)
}

func main() {
	beego.SetStaticPath("/static", "static")
	beego.AddFuncMap("md", helper.Markdown)
	beego.Run()
}
