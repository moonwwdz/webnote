package controllers

import (
	m "webnote/models"

	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) SignUp() {
	email := this.GetString("semail")
	pwd := this.GetString("spasswd")
	invcode := this.GetString("invitecode")
	if len(email) < 1 || len(pwd) < 1 || len(invcode) < 1 {
		return
	}
	if !m.CheckCode(invcode) {
		return
	}
	if _, err := m.AddUser(email, pwd, invcode); err != nil {
		this.Ctx.Output.JSON(err, true, true)
	}
	this.Ctx.Redirect(302, "/")
}
