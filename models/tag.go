package models

import "github.com/astaxie/beego/orm"

type Tag struct {
	Id       int
	Tag      string
	Notebook []*Notebook `orm:"reverse(many)"`
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(Tag))
}
