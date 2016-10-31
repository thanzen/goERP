package product

import "pms/controllers/base"

type ProductTagController struct {
	base.BaseController
}

func (this *ProductTagController) Get() {
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
	this.Data["searchKeyWords"] = "产品标签"
}
func (this *ProductTagController) List() {

}
