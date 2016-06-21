package models

import "github.com/astaxie/beego/orm"

type Profile struct {
	Id         int
	Invitecode string //设置的邀请码
	User       *User  `orm:"reverse(one)"` // 设置一对一反向关系(可选)
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(Profile))
}

func CheckCode(code string) bool {
	obj := Profile{Invitecode: code}
	o := orm.NewOrm()
	err := o.Read(&obj, "Invitecode")
	if err == orm.ErrNoRows {
		return false
	}
	return true
}
