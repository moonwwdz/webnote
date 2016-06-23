package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	flash := beego.ReadFromRequest(&c.Controller)
	if n, ok := flash.Data["error"]; ok {
		c.Data["isError"] = true
		c.Data["error"] = n
	}
	c.TplName = "index.html"
}
