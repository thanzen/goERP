package product

import (
	"encoding/json"
	"pms/controllers/base"
	mp "pms/models/product"
	"strconv"
	"strings"
)

type ProductAttributeController struct {
	base.BaseController
}

func (this *ProductAttributeController) Post() {
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
func (this *ProductAttributeController) Get() {
	this.GetList()

	this.URL = "/product/attribute"
	this.Data["URL"] = this.URL
	this.Layout = "base/base.html"
	this.Data["MenuProductAttributeActive"] = "active"
}
func (this *ProductAttributeController) Validator() {
	name := this.GetString("name")
	name = strings.TrimSpace(name)
	result := make(map[string]bool)
	if _, err := mp.GetProductAttributeByName(name); err != nil {
		result["valid"] = true
	} else {
		result["valid"] = false
	}
	this.Data["json"] = result
	this.ServeJSON()
}

// 获得符合要求的城市数据
func (this *ProductAttributeController) productAttributeList(start, length int64, condArr map[string]interface{}) (map[string]interface{}, error) {

	var arrs []mp.ProductAttribute
	paginator, arrs, err := mp.ListProductAttribute(condArr, start, length)
	result := make(map[string]interface{})
	if err == nil {

		//使用多线程来处理数据，待修改
		tableLines := make([]interface{}, 0, 4)
		for _, line := range arrs {
			oneLine := make(map[string]interface{})
			oneLine["name"] = line.Name
			oneLine["code"] = line.Code
			oneLine["sequence"] = line.Sequence
			mapValues := make(map[int64]string)
			oneLine["id"] = line.Id
			values := line.ValueIds
			for _, line := range values {
				mapValues[line.Id] = line.Name
			}
			//测试代码
			mapValues[12] = "11231232"
			mapValues[2] = "21231230"
			mapValues[3] = "121321"
			mapValues[4] = "20123"
			mapValues[52] = "12"
			mapValues[72] = "20"
			mapValues[21] = "12"
			mapValues[37] = "20"
			oneLine["values"] = mapValues
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
func (this *ProductAttributeController) PostList() {
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

func (this *ProductAttributeController) GetList() {
	this.Data["tableId"] = "table-product-attribute"
	this.TplName = "base/table_base.html"
}
