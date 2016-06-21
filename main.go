package main

import (
	_ "webnote/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.SetStaticPath("/static", "static")
	beego.Run()
}
