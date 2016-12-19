package product

import "pms/controllers/base"

type ProductUomCategController struct {
	base.BaseController
}

func (ctl *ProductUomCategController) Get() {
	action := ctl.GetString(":action")
	viewType := ctl.Input().Get("view_type")
	switch action {
	case "list":
		switch viewType {
		case "list":
			ctl.List()
		default:
			ctl.List()
		}
	default:
		ctl.List()
	}
	ctl.Data["searchKeyWords"] = "计量分类"
}
func (ctl *ProductUomCategController) List() {

}
