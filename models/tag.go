package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

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

//根据切片中的字符串，取得数据库中的对象，如果没有，新加一个对应的对象
//返回对象的切片
func GetOrCreateTags(tags *[]string) []*Tag {
	var tagObjs []*Tag
	o := orm.NewOrm()

	for _, tag := range *tags {
		tagObj := Tag{Tag: tag}
		if err := o.QueryTable("tag").Filter("Tag__Tag__Tag", tag).One(&tagObj); err != nil {
			beego.Debug(err)
			if _, err := o.Insert(&tagObj); err != nil {
				beego.Debug(err)
			}
		}
		tagObjs = append(tagObjs, &tagObj)
	}
	beego.Debug(tagObjs)
	return tagObjs
}
