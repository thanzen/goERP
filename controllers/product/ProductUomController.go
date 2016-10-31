package product

import "pms/controllers/base"

type ProductUomController struct {
	base.BaseController
}

func (this *ProductUomController) Get() {
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
	this.Data["searchKeyWords"] = "计量单位"
}
func (this *ProductUomController) List() {

}
