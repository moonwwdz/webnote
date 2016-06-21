package controllers

import (
	"fmt"
	m "webnote/models"

	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Login() {
	email := this.GetString("uemail")
	this.Ctx.SetCookie("ThisKad", email, "/")
	this.Ctx.Redirect(302, "/")
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
		fmt.Println(err)
		this.Ctx.SetCookie("warning", "xxxxx")
	}
	this.Ctx.Redirect(302, "/")
}
