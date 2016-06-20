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
	orm.RegisterDataBase("default", "mysql", "root:123456@/beego?charset=utf8")
	orm.RunSyncdb("default", false, true)
}

func (this *UserController) Post() {

}
