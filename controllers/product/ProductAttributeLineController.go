package product

import "pms/controllers/base"

type ProductAttributeLineController struct {
	base.BaseController
}

func (this *ProductAttributeLineController) Get() {
	action := this.GetString(":action")
	viewType := this.Input().Get("view_type")
	switch action {
	case "list":
		switch viewType {
		case "list":
			this.List()
		default:
			this.List()
		}
	default:
		this.List()
	}
	this.Data["searchKeyWords"] = "属性明细"
	this.URL = "/product/attribute"
	this.Data["URL"] = this.URL
	this.Data["productRootActive"] = "active"
	this.Data["productAttributeLineActive"] = "active"
	this.Layout = "base/base.html"
}
func (this *ProductAttributeLineController) List() {

}
