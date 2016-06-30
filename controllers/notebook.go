package controllers

import (
	"strconv"
	"strings"
	"webnote/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
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
	a := c.Ctx.Input.Params()
	intid, err := strconv.Atoi(a[":id"])
	if err != nil {
		beego.Debug(err)
	}
	note := new(models.Notebook)
	err = orm.NewOrm().QueryTable("notebook").Filter("id", intid).One(note)
	if err != nil {
		beego.Debug(err)
	}
	c.Data["note"] = note
	c.TplName = "note.html"
}
