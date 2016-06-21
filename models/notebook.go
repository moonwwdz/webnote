package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Notebook struct {
	Id         int
	Title      string
	Content    string    `orm:"type(text)"`
	User       *User     `orm:"rel(fk)"`
	Tag        []*Tag    `orm:"rel(m2m)"`
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateTime time.Time `orm:"auto_now;type(datetime)`
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(Notebook))
}
