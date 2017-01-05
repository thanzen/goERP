package product

import (
	"pms/controllers/base"
	mp "pms/models/product"
	"strings"
)

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
func (ctl *ProductTagController) Validator() {
	name := ctl.GetString("name")
	name = strings.TrimSpace(name)
	recordID, _ := ctl.GetInt64("recordId")
	result := make(map[string]bool)
	obj, err := mp.GetProductTagByName(name)
	if err != nil {
		result["valid"] = true
	} else {
		if obj.Name == name {
			if recordID == obj.Id {
				result["valid"] = true
			} else {
				result["valid"] = false
			}

		} else {
			result["valid"] = true
		}

	}
	ctl.Data["json"] = result
	ctl.ServeJSON()
}
