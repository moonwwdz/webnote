package routers

import (
	"webnote/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.UserController{}, "get:Login")
	beego.Router("/signup", &controllers.UserController{}, "post:SignUp")
}
