package product

import (
	"pms/controllers/base"
	mp "pms/models/product"
	"strconv"
)

type ProductUomController struct {
	base.BaseController
}

func (ctl *ProductUomController) Post() {
	action := ctl.Input().Get("action")
	switch action {
	case "validator":
		ctl.Validator()
	case "table": //bootstrap table的post请求
		ctl.PostList()
	case "create":
		ctl.PostCreate()
	default:
		ctl.PostList()
	}
}
func (ctl *ProductUomController) Get() {
	action := ctl.Input().Get("action")
	switch action {
	case "create":
		ctl.Create()
	case "edit":
		ctl.Edit()
	case "detail":
		ctl.Detail()
	default:
		ctl.GetList()
	}
	ctl.URL = "/product/uom"
	ctl.Data["URL"] = ctl.URL
	ctl.Layout = "base/base.html"
	ctl.Data["MenuProductUomActive"] = "active"
}
func (ctl *ProductUomController) Put()       {}
func (ctl *ProductUomController) Validator() {}
func (ctl *ProductUomController) PostCreate() {
	uom := new(mp.ProductUom)
	if err := ctl.ParseForm(uom); err == nil {
		if id, err := mp.CreateProductUom(uom, ctl.User); err == nil {
			ctl.Redirect("/product/uom/"+strconv.FormatInt(id, 10)+"?action=detail", 302)
		} else {
			ctl.Get()
		}
	} else {
		ctl.Get()
	}
}
func (ctl *ProductUomController) PostList() {}
func (ctl *ProductUomController) Edit()     {}
func (ctl *ProductUomController) Detail() {
	//获取信息一样，直接调用Edit
	ctl.Edit()
	ctl.Data["Readonly"] = true
	ctl.Data["Action"] = "detail"
}
func (ctl *ProductUomController) GetList() {
	ctl.Data["tableId"] = "table-product-uom"
	ctl.TplName = "base/table_base.html"
}
func (ctl *ProductUomController) Create() {
	ctl.Data["Action"] = "create"
	ctl.Data["Readonly"] = false
	ctl.Data["listName"] = "创建类别"
	ctl.TplName = "product/product_uom_form.html"
}
