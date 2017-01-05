package product

import (
	"encoding/json"
	"pms/controllers/base"
	mp "pms/models/product"
	"strconv"
	"strings"
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
	ctl.URL = "/product/uom/"
	ctl.Data["URL"] = ctl.URL
	ctl.Data["MenuProductUomActive"] = "active"
}
func (ctl *ProductUomController) Put() {
	id := ctl.Ctx.Input.Param(":id")
	ctl.URL = "/product/uom/"
	if idInt64, e := strconv.ParseInt(id, 10, 64); e == nil {
		if uom, err := mp.GetProductUomByID(idInt64); err == nil {
			if err := ctl.ParseForm(&uom); err == nil {

				if _, err := mp.UpdateProductUom(&uom, ctl.User); err == nil {
					ctl.Redirect(ctl.URL+id+"?action=detail", 302)
				}
			}
		}
	}
	ctl.Redirect(ctl.URL+id+"?action=edit", 302)
}

func (ctl *ProductUomController) Validator() {
	name := ctl.GetString("name")
	recordID, _ := ctl.GetInt64("recordId")
	name = strings.TrimSpace(name)
	result := make(map[string]bool)
	obj, err := mp.GetProductUomByName(name)
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
func (ctl *ProductUomController) PostCreate() {
	uom := new(mp.ProductUom)
	if err := ctl.ParseForm(uom); err == nil {
		if uomCategID, err := ctl.GetInt64("category"); err == nil {
			if category, err := mp.GetProductUomCategByID(uomCategID); err == nil {
				uom.Category = &category
				if id, err := mp.CreateProductUom(uom, ctl.User); err == nil {
					ctl.Redirect("/product/uom/"+strconv.FormatInt(id, 10)+"?action=detail", 302)
				}
			}
		}
	}
	ctl.Get()

}
func (ctl *ProductUomController) productUomList(start, length int64, condArr map[string]interface{}) (map[string]interface{}, error) {

	var arrs []mp.ProductUom
	paginator, arrs, err := mp.ListProductUom(condArr, start, length)
	result := make(map[string]interface{})
	if err == nil {

		// result["recordsFiltered"] = paginator.TotalCount
		tableLines := make([]interface{}, 0, 4)
		for _, line := range arrs {
			oneLine := make(map[string]interface{})
			oneLine["name"] = line.Name
			oneLine["Id"] = line.Id
			oneLine["id"] = line.Id
			oneLine["active"] = line.Active
			oneLine["rounding"] = line.Rounding
			oneLine["symbol"] = line.Symbol
			switch line.Type {
			case 1:
				oneLine["type"] = "小于参考计量单位"
				oneLine["factor"] = line.Factor
			case 2:
				oneLine["type"] = "参考计量单位"
			case 3:
				oneLine["type"] = "大约参考计量单位"
				oneLine["factorInv"] = line.FactorInv
			default:
				oneLine["type"] = "参考计量单位"
			}

			oneLine["category"] = line.Category.Name
			tableLines = append(tableLines, oneLine)
		}
		result["data"] = tableLines
		if jsonResult, er := json.Marshal(&paginator); er == nil {
			result["paginator"] = string(jsonResult)
			result["total"] = paginator.TotalCount
		}
	}
	return result, err
}
func (ctl *ProductUomController) PostList() {
	condArr := make(map[string]interface{})
	start := ctl.Input().Get("offset")
	length := ctl.Input().Get("limit")
	name := ctl.Input().Get("name")
	name = strings.TrimSpace(name)
	if name != "" {
		condArr["name"] = name
	}
	var (
		startInt64  int64
		lengthInt64 int64
	)
	if startInt, ok := strconv.Atoi(start); ok == nil {
		startInt64 = int64(startInt)
	}
	if lengthInt, ok := strconv.Atoi(length); ok == nil {
		lengthInt64 = int64(lengthInt)
	}
	if result, err := ctl.productUomList(startInt64, lengthInt64, condArr); err == nil {
		ctl.Data["json"] = result
	}
	ctl.ServeJSON()
}
func (ctl *ProductUomController) Edit() {
	id := ctl.Ctx.Input.Param(":id")
	uomInfo := make(map[string]interface{})
	if id != "" {
		if idInt64, e := strconv.ParseInt(id, 10, 64); e == nil {

			if uom, err := mp.GetProductUomByID(idInt64); err == nil {

				uomInfo["name"] = uom.Name
				uomInfo["factor"] = uom.Factor
				uomInfo["active"] = uom.Active
				uomInfo["factorInv"] = uom.FactorInv
				uomInfo["rounding"] = uom.Rounding
				typeUom := make(map[string]interface{})
				switch uom.Type {
				case 1:
					typeUom["id"] = 1
					typeUom["name"] = "小于参考计量单位"
				case 2:
					typeUom["id"] = 2
					typeUom["name"] = "参考计量单位"
				case 3:
					typeUom["id"] = 3
					typeUom["name"] = "大于参考计量单位"
				default:
					typeUom["id"] = 1
					typeUom["name"] = "参考计量单位"
				}
				uomInfo["type"] = typeUom
				category := make(map[string]interface{})
				category["id"] = uom.Category.Id
				category["name"] = uom.Category.Name
				uomInfo["category"] = category
			}
		}
	}
	ctl.Data["Action"] = "edit"
	ctl.Data["RecordId"] = id
	ctl.Data["Uom"] = uomInfo
	ctl.Layout = "base/base.html"

	ctl.TplName = "product/product_uom_form.html"
}
func (ctl *ProductUomController) Detail() {
	//获取信息一样，直接调用Edit
	ctl.Edit()
	ctl.Data["Readonly"] = true
	ctl.Data["Action"] = "detail"
}
func (ctl *ProductUomController) GetList() {
	viewType := ctl.Input().Get("view")
	if viewType == "" || viewType == "table" {
		ctl.Data["ViewType"] = "table"
	}
	ctl.Data["listName"] = "产品单位"
	ctl.Data["tableId"] = "table-product-uom"
	ctl.Layout = "base/base_list_view.html"
	ctl.TplName = "product/product_uom_list_search.html"
}
func (ctl *ProductUomController) Create() {
	ctl.Data["Action"] = "create"
	ctl.Data["Readonly"] = false
	ctl.Data["listName"] = "创建产品单位"
	ctl.Layout = "base/base.html"

	ctl.TplName = "product/product_uom_form.html"
}
