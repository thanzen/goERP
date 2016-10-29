package routers

import (
	"pms/controllers/address"
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
	//国家
	beego.Router("/country/:action([A-Za-z]+)/", &address.CountryController{})
	//省份
	beego.Router("/province/:action([A-Za-z]+)/", &address.ProvinceController{})
	//城市
	beego.Router("/city/:action([A-Za-z]+)/", &address.CityController{})
	//区县
	beego.Router("/district/:action([A-Za-z]+)/", &address.DistrictController{})

}
