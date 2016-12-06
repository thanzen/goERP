package product

import (
	"encoding/json"
	"pms/controllers/base"
	mp "pms/models/product"
	"strconv"
	"strings"
)

const (
	productAttributeValueListCellLength = 2
)

type ProductAttributeValueController struct {
	base.BaseController
}

func (this *ProductAttributeValueController) Post() {
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
func (this *ProductAttributeValueController) Get() {
	this.GetList()

	this.URL = "/product/attribute"
	this.Data["URL"] = this.URL
	this.Layout = "base/base.html"
	this.Data["MenuProductAttributeActive"] = "active"
}
func (this *ProductAttributeValueController) Validator() {
	name := this.GetString("name")
	name = strings.TrimSpace(name)
	result := make(map[string]bool)
	if _, err := mp.GetProductAttributeValueByName(name); err != nil {
		result["valid"] = true
	} else {
		result["valid"] = false
	}
	this.Data["json"] = result
	this.ServeJSON()
}

// 获得符合要求的城市数据
func (this *ProductAttributeValueController) productAttributeList(start, length int64, condArr map[string]interface{}) (map[string]interface{}, error) {

	var arrs []mp.ProductAttributeValue
	paginator, arrs, err := mp.ListProductAttributeValue(condArr, start, length)
	result := make(map[string]interface{})
	if err == nil {

		//使用多线程来处理数据，待修改
		tableLines := make([]interface{}, 0, 4)
		for _, line := range arrs {
			oneLine := make(map[string]interface{})
			oneLine["value"] = line.Name
			oneLine["attribute"] = line.Attribute.Name
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
func (this *ProductAttributeValueController) PostList() {
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
	if result, err := this.productAttributeList(startInt64, lengthInt64, condArr); err == nil {
		this.Data["json"] = result
	}
	this.ServeJSON()

}

func (this *ProductAttributeValueController) GetList() {
	this.Data["tableId"] = "table-product-attributevalue"
	this.TplName = "base/table_base.html"
}
