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

func (ctl *ProductAttributeValueController) Post() {
	action := ctl.Input().Get("action")
	switch action {
	case "validator":
		ctl.Validator()
	case "table": //bootstrap table的post请求
		ctl.PostList()
	default:
		ctl.PostList()
	}
}
func (ctl *ProductAttributeValueController) Get() {
	ctl.GetList()

	ctl.URL = "/product/attribute"
	ctl.Data["URL"] = ctl.URL
	ctl.Layout = "base/base.html"
	ctl.Data["MenuProductAttributeActive"] = "active"
}
func (ctl *ProductAttributeValueController) Validator() {
	name := ctl.GetString("name")
	name = strings.TrimSpace(name)
	result := make(map[string]bool)
	if _, err := mp.GetProductAttributeValueByName(name); err != nil {
		result["valid"] = true
	} else {
		result["valid"] = false
	}
	ctl.Data["json"] = result
	ctl.ServeJSON()
}

// 获得符合要求的城市数据
func (ctl *ProductAttributeValueController) productAttributeList(start, length int64, condArr map[string]interface{}) (map[string]interface{}, error) {

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
func (ctl *ProductAttributeValueController) PostList() {
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
	if result, err := ctl.productAttributeList(startInt64, lengthInt64, condArr); err == nil {
		ctl.Data["json"] = result
	}
	ctl.ServeJSON()

}

func (ctl *ProductAttributeValueController) GetList() {
	ctl.Data["tableId"] = "table-product-attributevalue"
	ctl.TplName = "base/table_base.html"
}
