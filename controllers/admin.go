package controllers

import (
	m "webnote/models"

	"github.com/astaxie/beego"
)

type AdminController struct {
	beego.Controller
}

func (c *AdminController) Login() {
	var isLogin bool = true
	flash := beego.NewFlash()
	email := c.GetString("uemail")
	pwd := c.GetString("upasswd")

	if len(email) < 0 || len(pwd) < 0 {
		flash.Error("登录信息不完整！")
		isLogin = false
	}

	if !m.CheckUser(email, pwd) {
		flash.Error("登录信息不正确，请重新输入！")
		isLogin = false
	}

	flash.Store(&c.Controller)
	if isLogin {
		session := m.GetUserByEmail(email)
		c.SetSession("userInfo", session)
	}
	c.Ctx.Redirect(302, "/")
}

func (c *AdminController) Logout() {
	c.DelSession("userInfo")
	c.Ctx.Redirect(302, "/")
}
