package models

import (
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Notebook struct {
	Id         int
	Title      string    `form:"ntitle" valid:"Required"`
	Content    string    `form:"ncontent" valid:"Required" orm:"type(text)"`
	User       *User     `orm:"rel(fk)"`
	Tag        []*Tag    `orm:"rel(m2m)"`
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateTime time.Time `orm:"auto_now;type(datetime)"`
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(Notebook))
}

//添加一条笔记
func AddNote(note *Notebook, tags *[]string) bool {
	var tagObjs []*Tag
	o := orm.NewOrm()
	// beego.Debug(tags)
	// for _, tag := range *tags {
	// 	tagObj := Tag{}
	// 	tagObj.Tag = tag
	// 	tagObj.Notebook = []*Notebook{note}
	// 	if err := o.QueryTable("tag").Filter("Tag", tag).One(&tagObj); tagObj.Id != 0 {
	// 		beego.Debug(err)
	// 		tagObjs = append(tagObjs, &tagObj)
	// 	} else {
	// 		_, err := o.Insert(&tagObj)
	// 		if err == nil {
	// 			tagObjs = append(tagObjs, &tagObj)
	// 		} else {
	// 			beego.Debug(err)
	// 		}
	// 	}
	// 	beego.Debug(tagObj)
	// }
	tagObjs = GetOrCreateTags(tags)
	note.UpdateTime = time.Time{}
	note.Tag = tagObjs
	if _, nerr := o.Insert(note); nerr != nil {
		beego.Debug(nerr)
		return false
	}
	for _, tag := range tagObjs {
		if _, err := o.QueryM2M(note, "Tag").Add(tag); err != nil {
			beego.Debug(err)
		}
	}
	return true
}

//取一条笔记
func GetNote(id int) *Notebook {
	note := new(Notebook)
	err := orm.NewOrm().QueryTable("notebook").Filter("Id", id).One(note)
	if err != nil {
		beego.Debug(err)
	}
	return note
}

//取一条笔记里的tag
func GetNoteTags(note *Notebook) string {
	var title []string

	if _, terr := orm.NewOrm().LoadRelated(note, "Tag"); terr != nil {
		beego.Debug(terr)
	}
	for _, t := range note.Tag {
		title = append(title, t.Tag)
	}
	return strings.Join(title, ",")
}

//删除一条笔记
func DelNote(id int) bool {
	return true
}

//更新一条笔记
func UpdateNote(note *Notebook) bool {
	if _, err := orm.NewOrm().Update(note); err != nil {
		beego.Debug(err)
		return false
	}
	return true
}
