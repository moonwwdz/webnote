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

	//取note的数据
	note := models.GetNote(intid)

	//取关联的title
	c.Data["title"] = models.GetNoteTags(note)

	//是否登录
	if session := c.GetSession("userInfo"); session != nil {
		c.Data["isLogin"] = true
		c.Data["info"] = session
	}
	c.Data["note"] = note
	c.TplName = "note.html"
}

func (c *NoteController) ChangeNote() {
	flash := beego.NewFlash()
	session := c.GetSession("userInfo")
	o := orm.NewOrm()
	if session == nil {
		flash.Error("请先登录！")
		flash.Store(&c.Controller)
		c.Ctx.Redirect(302, "/")
		return
	}

	note_id := c.GetString("note_id")
	intid, err := strconv.Atoi(note_id)
	if err != nil {
		beego.Debug(err)
	}
	note := models.Notebook{Id: intid}
	if err := o.Read(&note); err != nil {
		beego.Debug(err)
		flash.Error("修改失败！")
		flash.Store(&c.Controller)
		c.Ctx.Redirect(302, "/")
	}
	if nerr := c.ParseForm(&note); nerr != nil {
		beego.Debug(nerr)
	}
	if s, ok := session.(models.User); ok {
		note.User = &s
	}

	o.Begin()
	if !models.UpdateNote(&note) {
		o.Rollback()
		flash.Error("添加失败！")
		flash.Store(&c.Controller)
		c.Ctx.Redirect(302, "/")
	}

	tags := c.GetString("ntag")
	tagList := strings.Split(tags, ",")
	tagObjs := models.GetOrCreateTags(&tagList)
	m2m := o.QueryM2M(&note, "Tag")
	m2m.Clear()
	for _, tag := range tagObjs {
		m2m.Add(tag)
	}
	o.Commit()
	c.Ctx.Redirect(302, "/")
}
