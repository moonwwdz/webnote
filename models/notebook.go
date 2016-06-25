package models

import (
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
	UpdateTime time.Time `orm:"auto_now;type(datetime)`
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(Notebook))
}

//添加一条笔记
func AddNote(note *Notebook, tags *[]string) bool {
	var tagObjs []*Tag
	tagObj := Tag{}
	o := orm.NewOrm()
	for _, tag := range *tags {
		tagObj.Tag = tag
		if _ = o.Read(&tagObj); tagObj.Id != 0 {
			tagObjs = append(tagObjs, &tagObj)
		} else {
			_, err := o.Insert(&tagObj)
			if err == nil {
				tagObjs = append(tagObjs, &tagObj)
			} else {
				beego.Debug(err)
			}
		}
	}
	if _, nerr := o.Insert(note); nerr != nil {
		beego.Debug(nerr)
		return false
	}
	return true
}

//删除一条笔记
func DelNote(id int) bool {
	return true
}

//更新一条笔记
func UpdateNote(id int) bool {
	return true
}
