package controllers

import (
	"strings"
	"webnote/models"

	"github.com/astaxie/beego"
)

type NoteController struct {
	beego.Controller
}

func (c *NoteController) NewNote() {
	flash := beego.NewFlash()
	session := c.GetSession("userInfo")
	if session == nil {
		flash.Error("请先登录！")
		flash.Store(&c.Controller)
		c.Ctx.Redirect(302, "/")
	}

	note := models.Notebook{User: models.User(session)}
	if nerr := c.ParseForm(&note); nerr != nil {

	}
	titles := c.GetString("ntitle")
	titleList := strings.Split(titles, ",")
	if !models.AddNote(&note, &titleList) {
		flash.Error("添加失败！")
		flash.Store(&c.Controller)
	}
	c.Ctx.Redirect(302, "/")
}

func (c *NoteController) GetNote() {
	c.Ctx.Redirect(302, "/")
}
