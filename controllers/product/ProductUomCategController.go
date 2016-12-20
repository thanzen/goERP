package product

import (
	"encoding/json"
	"pms/controllers/base"
	mp "pms/models/product"
	"strconv"
	"strings"
)

type ProductUomCategController struct {
	base.BaseController
}

func (ctl *ProductUomCategController) Post() {
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
func (ctl *ProductUomCategController) Get() {
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
	ctl.URL = "/product/uomcateg"
	ctl.Data["URL"] = ctl.URL
	ctl.Layout = "base/base.html"
	ctl.Data["MenuProductUomCategActive"] = "active"
}
func (ctl *ProductUomCategController) Put() {
	id := ctl.Ctx.Input.Param(":id")
	ctl.URL = "/product/uomcateg"
	if idInt64, e := strconv.ParseInt(id, 10, 64); e == nil {
		if uomCateg, err := mp.GetProductUomCategByID(idInt64); err == nil {
			if err := ctl.ParseForm(&uomCateg); err == nil {

				if _, err := mp.UpdateProductUomCateg(&uomCateg, ctl.User); err == nil {
					ctl.Redirect(ctl.URL+"/"+id+"?action=detail", 302)
				}
			}
		}
	}
	ctl.Redirect(ctl.URL+"/"+id+"?action=edit", 302)
}
func (ctl *ProductUomCategController) Validator() {
	name := ctl.GetString("name")
	recordId := ctl.GetString("recordId")
	name = strings.TrimSpace(name)
	result := make(map[string]bool)
	obj, err := mp.GetProductUomCategByName(name)
	if err != nil {
		result["valid"] = true
	} else {
		if obj.Name == name {
			if recordId != "" {
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
func (ctl *ProductUomCategController) PostCreate() {
	uom := new(mp.ProductUomCateg)
	if err := ctl.ParseForm(uom); err == nil {
		if id, err := mp.CreateProductUomCateg(uom, ctl.User); err == nil {
			ctl.Redirect("/product/uomcateg/"+strconv.FormatInt(id, 10)+"?action=detail", 302)
		} else {
			ctl.Get()
		}
	} else {
		ctl.Get()
	}
}
func (ctl *ProductUomCategController) productUomCategList(start, length int64, condArr map[string]interface{}) (map[string]interface{}, error) {
	var arrs []mp.ProductUomCateg
	paginator, arrs, err := mp.ListProductUomCateg(condArr, start, length)
	result := make(map[string]interface{})
	if err == nil {

		// result["recordsFiltered"] = paginator.TotalCount
		tableLines := make([]interface{}, 0, 4)
		for _, line := range arrs {
			oneLine := make(map[string]interface{})
			oneLine["name"] = line.Name
			oneLine["Id"] = line.Id
			oneLine["id"] = line.Id
			uoms := line.Uoms
			mapValues := make(map[int64]string)
			for _, line := range uoms {
				mapValues[line.Id] = line.Name
			}
			oneLine["uoms"] = mapValues
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
func (ctl *ProductUomCategController) PostList() {
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
	if result, err := ctl.productUomCategList(startInt64, lengthInt64, condArr); err == nil {
		ctl.Data["json"] = result
	}
	ctl.ServeJSON()
}
func (ctl *ProductUomCategController) Edit() {
	id := ctl.Ctx.Input.Param(":id")
	categInfo := make(map[string]interface{})
	if id != "" {
		if idInt64, e := strconv.ParseInt(id, 10, 64); e == nil {
			if categ, err := mp.GetProductUomCategByID(idInt64); err == nil {
				categInfo["name"] = categ.Name
			}
		}
	}
	ctl.Data["Action"] = "edit"
	ctl.Data["RecordId"] = id
	ctl.Data["UomCateg"] = categInfo
	ctl.TplName = "product/product_uom_categ_form.html"
}
func (ctl *ProductUomCategController) Detail() {
	//获取信息一样，直接调用Edit
	ctl.Edit()
	ctl.Data["Readonly"] = true
	ctl.Data["Action"] = "detail"
}
func (ctl *ProductUomCategController) GetList() {
	ctl.Data["tableId"] = "table-product-uom-categ"
	ctl.TplName = "base/table_base.html"
}
func (ctl *ProductUomCategController) Create() {
	ctl.Data["Action"] = "create"
	ctl.Data["Readonly"] = false
	ctl.Data["listName"] = "创建单位类别"
	ctl.TplName = "product/product_uom_categ_form.html"
}
