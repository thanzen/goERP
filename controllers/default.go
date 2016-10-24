package controllers

type MainController struct {
	BaseController
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	// 基础布局页面
	// c.Layout = "base/base.html"
	c.TplName = "user/login.html"

}
