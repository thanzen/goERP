package product

import "pms/controllers/base"

type ProductUomController struct {
	base.BaseController
}

func (ctl *ProductUomController) Get() {
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
	ctl.Data["searchKeyWords"] = "计量单位"
}
func (ctl *ProductUomController) List() {

}
