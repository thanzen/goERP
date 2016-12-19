package product

import "pms/controllers/base"

type ProductTagController struct {
	base.BaseController
}

func (ctl *ProductTagController) Get() {
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
	ctl.Data["searchKeyWords"] = "产品标签"
}
func (ctl *ProductTagController) List() {

}
