package models

import "github.com/astaxie/beego/orm"

type Tag struct {
	Id       int
	Tag      string      `valid:"Required"`
	Notebook []*Notebook `orm:"reverse(many)"`
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(Tag))
}

//添加tag
func AddTag(tag string) bool {
	return true
}

//获取tag列表
func GetTagList() map[int]Tag {
	tags := make(map[int]Tag)
	return tags
}
