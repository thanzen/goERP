package product

import "pms/controllers/base"

type ProductAttributePriceController struct {
	base.BaseController
}

func (this *ProductAttributePriceController) Get() {
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
	this.Data["searchKeyWords"] = "属性值价格"
	this.Layout = "base/base.html"
}
func (this *ProductAttributePriceController) List() {

}
