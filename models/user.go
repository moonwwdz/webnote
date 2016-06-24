package models

import (
	"webnote/helper"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id       int
	Name     string      `form:"sname" valid:"Required"`
	Email    string      `form:"semail" valid:"Required;Email" orm:"unique"`
	Passwd   string      `form:"spasswd" valid:"Required"`
	Photo    string      `orm:"null;size(1500)"`
	Profile  *Profile    `orm:"rel(one)"`      // OneToOne relation
	Notebook []*Notebook `orm:"reverse(many)"` // 设置一对多的反向关系
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(User))
}

func CheckUser(email, pwd string) bool {
	o := orm.NewOrm()
	qs := o.QueryTable("User").Filter("email", email).Filter("Passwd", helper.HashPw(pwd)).Exist()
	if !qs {
		beego.Debug(qs)
		return false
	}
	return true
}

func GetUserByEmail(email string) User {
	o := orm.NewOrm()
	user := User{}
	err := o.QueryTable("User").Filter("email", email).One(&user)
	if err != nil {
		return User{}
	}
	return user
}

func AddUser(u *User, c string) (id int64, err error) {
	o := orm.NewOrm()
	pro := Profile{Invitecode: c}
	pe := o.Read(&pro, "Invitecode")
	if pe != nil {
		return 0, pe
	}
	user := User{Name: u.Name, Email: u.Email, Passwd: helper.HashPw(u.Passwd), Photo: u.Photo, Profile: &pro}
	id, ue := o.Insert(&user)
	return id, ue
}
