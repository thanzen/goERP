package product

import "pms/controllers/base"

const (
	productProductListCellLength = 2
)

type ProductProductController struct {
	base.BaseController
}

func (this *ProductProductController) Get() {
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
	this.Data["searchKeyWords"] = "产品规格名"
	this.URL = "/product/product"
	this.Data["URL"] = this.URL
	this.Data["productRootActive"] = "active"
	this.Data["productProductActive"] = "active"
	this.Layout = "base/base.html"
}
func (this *ProductProductController) List() {

}
