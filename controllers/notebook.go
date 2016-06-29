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
		return
	}

	note := models.Notebook{}
	if nerr := c.ParseForm(&note); nerr != nil {

	}
	if s, ok := session.(models.User); ok {
		note.User = &s
	}
	tags := c.GetString("ntag")
	tagList := strings.Split(tags, ",")
	if !models.AddNote(&note, &tagList) {
		flash.Error("添加失败！")
		flash.Store(&c.Controller)
	}
	c.Ctx.Redirect(302, "/")
}

func (c *NoteController) GetNote() {
	c.Ctx.Redirect(302, "/")
}
