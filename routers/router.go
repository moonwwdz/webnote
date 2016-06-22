package routers

import (
	"webnote/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/admin", &controllers.AdminController{}, "post:Login")
	beego.Router("/admin", &controllers.AdminController{}, "get:Logout")
	beego.Router("/signup", &controllers.UserController{}, "post:SignUp")
}
