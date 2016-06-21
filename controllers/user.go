package controllers

import (
	_ "webnote/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type UserController struct {
	beego.Controller
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	mysqlU := beego.AppConfig.String("mysqlname")
	mysqlP := beego.AppConfig.String("mysqlpwd")
	mysqlDB := beego.AppConfig.String("dbname")
	handle := mysqlU + ":" + mysqlP + "@/" + mysqlDB + "?charset=utf8"
	orm.RegisterDataBase("default", "mysql", handle)
	orm.RunSyncdb("default", false, true)
}

func (this *UserController) Post() {
	email := this.GetString("uemail")
	this.Ctx.SetCookie("ThisKad", email, "/")
	this.Ctx.Redirect(302, "/")
}
