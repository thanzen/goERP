package product

import (
	"encoding/json"
	"pms/controllers/base"
	mp "pms/models/product"
	"strconv"
	"strings"
)

type ProductProductController struct {
	base.BaseController
}

func (this *ProductProductController) Post() {
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
func (this *ProductProductController) Get() {
	this.GetList()

	this.URL = "/product/product"
	this.Data["URL"] = this.URL
	this.Layout = "base/base.html"
	this.Data["MenuProductProductActive"] = "active"
}
func (this *ProductProductController) Validator() {
	name := this.GetString("name")
	name = strings.TrimSpace(name)
	result := make(map[string]bool)
	if _, err := mp.GetProductProductByName(name); err != nil {
		result["valid"] = true
	} else {
		result["valid"] = false
	}
	this.Data["json"] = result
	this.ServeJSON()
}

// 获得符合要求的城市数据
func (this *ProductProductController) productProductList(start, length int64, condArr map[string]interface{}) (map[string]interface{}, error) {

	var arrs []mp.ProductProduct
	paginator, arrs, err := mp.ListProductProduct(condArr, start, length)
	result := make(map[string]interface{})
	if err == nil {

		//使用多线程来处理数据，待修改
		tableLines := make([]interface{}, 0, 4)
		for _, line := range arrs {
			oneLine := make(map[string]interface{})
			oneLine["name"] = line.Name
			oneLine["id"] = line.Id
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
func (this *ProductProductController) PostList() {
	condArr := make(map[string]interface{})
	start := this.Input().Get("offset")
	length := this.Input().Get("limit")
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
	if result, err := this.productProductList(startInt64, lengthInt64, condArr); err == nil {
		this.Data["json"] = result
	}
	this.ServeJSON()

}

func (this *ProductProductController) GetList() {
	this.Data["tableId"] = "table-product-product"
	this.TplName = "base/table_base.html"
}
