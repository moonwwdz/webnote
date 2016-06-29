package models

import "github.com/astaxie/beego/orm"

func AllList() []*Tag {
	var tags []*Tag
	_, err := orm.NewOrm().QueryTable("Tag").All(&tags)
	if err == nil {
		for _, elem := range tags {
			var notes []*Notebook
			orm.NewOrm().QueryTable("Notebook").Filter("Tag__Tag__Tag", elem.Tag).All(&notes)
			elem.Notebook = notes
		}
	}
	return tags
}
