package controllers

type IndexController struct {
	BaseController
}

func (this *IndexController) Get() {

	// 基础布局页面
	this.Layout = "base/base.html"
	this.TplName = "test.html"

}
