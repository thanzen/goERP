package product

import (
	"encoding/json"
	"fmt"
	"pms/controllers/base"
	mp "pms/models/product"
	"strconv"
	"strings"
)

type ProductCategoryController struct {
	base.BaseController
}

func (this *ProductCategoryController) Post() {
	action := this.Input().Get("action")
	switch action {
	case "validator":
		this.Validator()
	case "table": //bootstrap table的post请求
		this.PostList()
	default:
		this.PostList()
	}
}
func (this *ProductCategoryController) Get() {
	this.GetList()
	action := this.Input().Get("action")
	switch action {
	case "create":
		this.Create()
	default:
		this.GetList()

	}
	this.URL = "/product/category"
	this.Data["URL"] = this.URL
	this.Layout = "base/base.html"
	this.Data["MenuProductCategoryActive"] = "active"
}
func (this *ProductCategoryController) Create() {
	method := strings.ToUpper(this.Ctx.Request.Method)
	if method == "GET" {
		this.Data["Readonly"] = false
		this.Data["listName"] = "创建类别"
		this.TplName = "product/product_category_form.html"

	}
}
func (this *ProductCategoryController) Validator() {
	name := this.GetString("name")
	name = strings.TrimSpace(name)
	result := make(map[string]bool)
	if _, err := mp.GetProductCategoryByName(name); err != nil {
		result["valid"] = true
	} else {
		result["valid"] = false
	}
	this.Data["json"] = result
	this.ServeJSON()
}

// 获得符合要求的城市数据
func (this *ProductCategoryController) productCategoryList(start, length int64, condArr map[string]interface{}) (map[string]interface{}, error) {

	var arrs []mp.ProductCategory
	paginator, arrs, err := mp.ListProductCategory(condArr, start, length)
	result := make(map[string]interface{})
	if err == nil {

		// result["recordsFiltered"] = paginator.TotalCount
		tableLines := make([]interface{}, 0, 4)
		for _, line := range arrs {
			oneLine := make(map[string]interface{})
			oneLine["name"] = line.Name
			if line.Parent != nil {
				oneLine["parent"] = line.Parent.Name
			} else {
				oneLine["parent"] = "-"
			}
			oneLine["path"] = line.ParentFullPath
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
func (this *ProductCategoryController) PostList() {
	condArr := make(map[string]interface{})
	start := this.Input().Get("offset")
	length := this.Input().Get("limit")
	fmt.Println(start)
	fmt.Println(length)

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
	if result, err := this.productCategoryList(startInt64, lengthInt64, condArr); err == nil {
		this.Data["json"] = result
	}
	this.ServeJSON()

}

func (this *ProductCategoryController) GetList() {
	this.Data["tableId"] = "table-product-category"
	this.TplName = "base/table_base.html"
}
