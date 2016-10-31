package product

import "pms/controllers/base"

type ProductPackagingController struct {
	base.BaseController
}

func (this *ProductPackagingController) Get() {
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
	this.Data["searchKeyWords"] = "产品包装"
}
func (this *ProductPackagingController) List() {

}
