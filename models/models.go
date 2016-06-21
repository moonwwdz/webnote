package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type User struct {
	Id       int
	Name     string
	Passwd   string
	Profile  *Profile    `orm:"rel(one)"`      // OneToOne relation
	Notebook []*Notebook `orm:"reverse(many)"` // 设置一对多的反向关系
}

type Profile struct {
	Id         int
	Invitecode int   //设置的邀请码
	User       *User `orm:"reverse(one)"` // 设置一对一反向关系(可选)
}

type Notebook struct {
	Id         int
	Title      string
	Content    string    `orm:"type(text)"`
	User       *User     `orm:"rel(fk)"`
	Tag        []*Tag    `orm:"rel(m2m)"`
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateTime time.Time `orm:"auto_now;type(datetime)`
}

type Tag struct {
	Id       int
	Tag      string
	Notebook []*Notebook `orm:"reverse(many)"`
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(User), new(Profile), new(Notebook), new(Tag))
}
