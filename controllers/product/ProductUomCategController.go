package product

import "pms/controllers/base"

type ProductUomCategController struct {
	base.BaseController
}

func (this *ProductUomCategController) Get() {
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
	this.Data["searchKeyWords"] = "计量分类"
}
func (this *ProductUomCategController) List() {

}
