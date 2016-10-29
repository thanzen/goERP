package routers

import (
	"pms/controllers/base"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &base.IndexController{})
	//登录
	beego.Router("/login/:action([A-Za-z]+)/", &base.LoginController{})
	//用户
	beego.Router("/user/:action([A-Za-z]+)/", &base.UserController{})
	//登录日志
	beego.Router("/record/:action([A-Za-z]+)/", &base.RecordController{})

}
