package base

type IndexController struct {
	BaseController
}

func (ctl *IndexController) Get() {

	// 基础布局页面
	ctl.Layout = "base/base.html"
	ctl.TplName = "test.html"

}
