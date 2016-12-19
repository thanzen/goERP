package product

import "pms/controllers/base"

type ProductPackagingController struct {
	base.BaseController
}

func (ctl *ProductPackagingController) Get() {
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
	ctl.Data["searchKeyWords"] = "产品包装"
}
func (ctl *ProductPackagingController) List() {

}
