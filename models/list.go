package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type List struct {
	Tagid   int
	Tagname string
	Count   int
	Notes   []Note
}

type Note struct {
	Noteid   int
	Notename string
}

func AllList() int {
	var lists []orm.ParamsList
	o := orm.NewOrm()
	num, err := o.Raw("select tag.id,tag.tag,count(notebook_tags.notebook_id) as ctn from notebook_tags left join tag on notebook_tags.tag_id = tag.id group by notebook_tags.tag_id ").ValuesList(&lists)
	if err == nil && num > 0 {
		for id, arr := range lists {
			var note []orm.ParamsList
			_, _ = o.Raw("select notebook.id,notebook.title from notebook_tags left join notebook on notebook_tags.notebook_id = notebook.id where notebook_tags.tag_id = ?", arr[0]).ValuesList(&note)
			lists[id] = append(lists[id], note)
			beego.Debug(note)
		}
	}
	beego.Debug(lists)
	return 1
}
