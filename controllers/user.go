package controllers

import (
	m "webnote/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) SignUp() {
	flash := beego.NewFlash()
	invcode := this.GetString("invitecode")
	u := m.User{}
	//不为空
	if pErr := this.ParseForm(&u); pErr != nil || len(invcode) < 1 {
		flash.Error("输入完整信息！")
		flash.Store(&this.Controller)
	}
	//格式
	valid := validation.Validation{}
	if vResult, vErr := valid.Valid(&u); vErr != nil || !vResult {
		flash.Error("输入格式有错误！")
	}

	if !m.CheckCode(invcode) {
		flash.Error("Code不能用！")
	}

	if _, err := m.AddUser(&u, invcode); err != nil {
		flash.Error("数据库错误！换个邮箱试试。")
	}
	flash.Store(&this.Controller)
	this.Ctx.Redirect(302, "/")
}
