package models

import "github.com/astaxie/beego/orm"

type User struct {
	Id       int
	Name     string
	Passwd   string
	Profile  *Profile    `orm:"rel(one)"`      // OneToOne relation
	Notebook []*Notebook `orm:"reverse(many)"` // 设置一对多的反向关系
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(User))
}

func AddUser(u string, p string, c string) (id int64, err error) {
	o := orm.NewOrm()
	pro := Profile{Invitecode: c}
	pe := o.Read(&pro)
	if pe != nil {
		return 0, pe
	}
	user := User{Name: u, Passwd: p, Profile: &pro}
	id, ue := o.Insert(&user)
	if ue != nil {
		return 0, ue
	}
	return id, nil
}