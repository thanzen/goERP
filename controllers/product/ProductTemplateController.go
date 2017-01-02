package product

import (
	"encoding/json"
	"fmt"
	"pms/controllers/base"
	mp "pms/models/product"
	"strconv"
	"strings"
)

type ProductTemplateController struct {
	base.BaseController
}

func (ctl *ProductTemplateController) Post() {
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
func (ctl *ProductTemplateController) Get() {
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
	ctl.URL = "/product/template/"
	ctl.Data["URL"] = ctl.URL
	ctl.Data["MenuProductTemplateActive"] = "active"
}
func (ctl *ProductTemplateController) Put() {
	id := ctl.Ctx.Input.Param(":id")
	ctl.URL = "/product/template/"
	if idInt64, e := strconv.ParseInt(id, 10, 64); e == nil {
		if template, err := mp.GetProductTemplateByID(idInt64); err == nil {
			if err := ctl.ParseForm(&template); err == nil {

				if _, err := mp.UpdateProductTemplate(&template, ctl.User); err == nil {
					ctl.Redirect(ctl.URL+id+"?action=detail", 302)
				}
			}
		}
	}
	ctl.Redirect(ctl.URL+id+"?action=edit", 302)

}
func (ctl *ProductTemplateController) PostCreate() {
	template := new(mp.ProductTemplate)
	if err := ctl.ParseForm(template); err == nil {

		if id, err := mp.CreateProductTemplate(template, ctl.User); err == nil {
			ctl.Redirect("/product/tempalte/"+strconv.FormatInt(id, 10)+"?action=detail", 302)
		} else {
			ctl.Get()
		}
	} else {
		ctl.Get()
	}
}
func (ctl *ProductTemplateController) Edit() {
	id := ctl.Ctx.Input.Param(":id")
	templateInfo := make(map[string]interface{})
	if id != "" {
		if idInt64, e := strconv.ParseInt(id, 10, 64); e == nil {
			if template, err := mp.GetProductTemplateByID(idInt64); err == nil {
				fmt.Println(template)
				templateInfo["name"] = template.Name
				templateInfo["defaultCode"] = template.DefaultCode
				templateInfo["sequence"] = template.Sequence
				templateInfo["description"] = template.Description
				templateInfo["descriptioPurchase"] = template.DescriptioPurchase
				templateInfo["descriptioSale"] = template.DescriptioSale
				templateInfo["productType"] = template.ProductType
				templateInfo["productMethod"] = template.ProductMethod
				categ := template.Categ
				fmt.Println(categ)
				categValues := make(map[string]string)
				if categ != nil {
					categValues["id"] = strconv.FormatInt(categ.Id, 10)
					categValues["name"] = categ.Name
				}
				fmt.Println(categValues)
				templateInfo["category"] = categValues
			}
		}
	}
	ctl.Data["Action"] = "edit"
	ctl.Data["RecordId"] = id
	ctl.Data["Tp"] = templateInfo
	fmt.Println(templateInfo)
	ctl.Layout = "base/base.html"
	ctl.TplName = "product/product_template_form.html"
}
func (ctl *ProductTemplateController) Detail() {
	ctl.Edit()
	ctl.Data["Readonly"] = true
	ctl.Data["Action"] = "detail"
}
func (ctl *ProductTemplateController) Create() {
	ctl.Data["Action"] = "create"
	ctl.Data["Readonly"] = false
	ctl.Data["listName"] = "创建款式"
	ctl.Layout = "base/base.html"

	ctl.TplName = "product/product_template_form.html"
}

func (ctl *ProductTemplateController) Validator() {
	name := ctl.GetString("name")
	name = strings.TrimSpace(name)
	result := make(map[string]bool)
	if _, err := mp.GetProductTemplateByName(name); err != nil {
		result["valid"] = true
	} else {
		result["valid"] = false
	}
	ctl.Data["json"] = result
	ctl.ServeJSON()
}

// 获得符合要求的城市数据
func (ctl *ProductTemplateController) productTemplateList(start, length int64, condArr map[string]interface{}) (map[string]interface{}, error) {

	var arrs []mp.ProductTemplate
	paginator, arrs, err := mp.ListProductTemplate(condArr, start, length)
	result := make(map[string]interface{})
	if err == nil {

		//使用多线程来处理数据，待修改
		tableLines := make([]interface{}, 0, 4)
		for _, line := range arrs {
			oneLine := make(map[string]interface{})
			oneLine["name"] = line.Name
			oneLine["sequence"] = line.Sequence
			oneLine["Id"] = line.Id
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
func (ctl *ProductTemplateController) PostList() {
	condArr := make(map[string]interface{})
	start := ctl.Input().Get("offset")
	length := ctl.Input().Get("limit")
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
	if result, err := ctl.productTemplateList(startInt64, lengthInt64, condArr); err == nil {
		ctl.Data["json"] = result
	}
	ctl.ServeJSON()

}

func (ctl *ProductTemplateController) GetList() {
	viewType := ctl.Input().Get("view")
	if viewType == "" || viewType == "table" {
		ctl.Data["ViewType"] = "table"
	}
	ctl.Data["listName"] = "产品款式管理"
	ctl.Data["tableId"] = "table-product-template"
	ctl.Layout = "base/base_list_view.html"
	ctl.TplName = "product/product_template_list_search.html"
}
