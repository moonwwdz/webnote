package models

import "github.com/astaxie/beego/orm"

type User struct {
	Id       int
	Name     string      `form:"sname" valid:"Required"`
	Email    string      `form:"semail" valid:"Required;Email"`
	Passwd   string      `form:"spasswd" valid:"Required"`
	Profile  *Profile    `orm:"rel(one)"`      // OneToOne relation
	Notebook []*Notebook `orm:"reverse(many)"` // 设置一对多的反向关系
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(User))
}

func AddUser(u *User, c string) (id int64, err error) {
	o := orm.NewOrm()
	pro := Profile{Invitecode: c}
	pe := o.Read(&pro, "Invitecode")
	if pe != nil {
		return 0, pe
	}
	user := User{Name: u.Name, Email: u.Email, Passwd: u.Passwd, Profile: &pro}
	id, ue := o.Insert(&user)
	return id, ue
}
