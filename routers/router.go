package routers

import (
	"pms/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	//登录
	beego.Router("/login/:action([A-Za-z]+)/", &controllers.LoginController{})
}
