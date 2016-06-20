package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id      int
	Name    string
	Passwd  string
	Profile *Profile `orm:"rel(one)"`      // OneToOne relation
	Post    []*Post  `orm:"reverse(many)"` // 设置一对多的反向关系
}

type Profile struct {
	Id         int
	Invitecode int16 //设置的邀请码
	User       *User `orm:"reverse(one)"` // 设置一对一反向关系(可选)
}

type Post struct {
	Id    int
	Title string
	User  *User `orm:"rel(fk)"` //设置一对多关系
	// Tags  []*Tag `orm:"rel(m2m)"`
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(User), new(Profile))
}
