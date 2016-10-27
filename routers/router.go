package routers

import (
	"pms/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.IndexController{})
	//登录
	beego.Router("/login/:action([A-Za-z]+)/", &controllers.LoginController{})
	//用户
	beego.Router("/user/:action([A-Za-z]+)/", &controllers.UserController{})
	//登录日志
	beego.Router("/record/:action([A-Za-z]+)/", &controllers.RecordController{})

}
