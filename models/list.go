package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type List struct {
	T     Tag
	Count int
	N     []*Notebook
}

func AllList() []List {
	var tags []Tag
	tagCtn := make(map[string]int)
	tagObj := make(map[string]Tag)
	var lists []List
	if ctn, _ := orm.NewOrm().QueryTable("Tag").All(&tags); ctn > 0 {
		for _, tag := range tags {
			beego.Debug(tag.Tag)
			tagCtn[tag.Tag]++
			tagObj[tag.Tag] = tag
		}
	}
	beego.Debug(tagCtn)
	beego.Debug(tagObj)
	for tag, ctn := range tagCtn {
		list := List{T: tagObj[tag], Count: ctn, N: tagObj[tag].Notebook}
		lists = append(lists, list)
	}
	return lists
}
