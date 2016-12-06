package product

import "pms/controllers/base"

const (
	productTemplateListCellLength = 2
)

type ProductTemplateController struct {
	base.BaseController
}

func (this *ProductTemplateController) Get() {
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
	this.Data["searchKeyWords"] = "产品款式名"
	this.URL = "/product/template"
	this.Data["URL"] = this.URL
	this.Layout = "base/base.html"
	this.Data["productRootActive"] = "active"
	this.Data["productTemplateActive"] = "active"
}
func (this *ProductTemplateController) List() {

}
