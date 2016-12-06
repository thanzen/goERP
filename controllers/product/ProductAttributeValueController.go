package product

import "pms/controllers/base"

const (
	productAttributeValueListCellLength = 2
)

type ProductAttributeValueController struct {
	base.BaseController
}

func (this *ProductAttributeValueController) Get() {
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
	this.Data["searchKeyWords"] = "属性值"
	this.URL = "/product/attributevalue"
	this.Data["URL"] = this.URL
	this.Data["productRootActive"] = "active"
	this.Data["productAttributeActive"] = "active"
	this.Layout = "base/base.html"
}
func (this *ProductAttributeValueController) List() {

}
