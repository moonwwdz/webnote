package controllers

import "github.com/astaxie/beego"

type AdminController struct {
	beego.Controller
}

func (c *AdminController) Login() {
	c.Ctx.Redirect(302, "/")
}

func (c *AdminController) Logout() {
	c.Ctx.Redirect(302, "/")
}
