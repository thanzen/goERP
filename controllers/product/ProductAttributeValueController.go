package product

import (
	"encoding/json"
	"pms/controllers/base"
	mp "pms/models/product"
	"strconv"
	"strings"
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
	case "create":
		ctl.PostCreate()
	default:
		ctl.PostList()
	}
}
func (ctl *ProductAttributeValueController) Put() {
	id := ctl.Ctx.Input.Param(":id")
	ctl.URL = "/product/attributevalue/"
	if idInt64, e := strconv.ParseInt(id, 10, 64); e == nil {
		if attrValue, err := mp.GetProductAttributeValueByID(idInt64); err == nil {
			if err := ctl.ParseForm(&attrValue); err == nil {
				if attributeID, err := ctl.GetInt64("productAttributeID"); err == nil {
					if attribute, err := mp.GetProductAttributeByID(attributeID); err == nil {
						attrValue.Attribute = &attribute
					}
				}
				if _, err := mp.UpdateProductAttributeValue(&attrValue, ctl.User); err == nil {
					ctl.Redirect(ctl.URL+id+"?action=detail", 302)
				}
			}
		}
	}
	ctl.Redirect(ctl.URL+id+"?action=edit", 302)

}
func (ctl *ProductAttributeValueController) Get() {
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
	ctl.URL = "/product/attributevalue/"
	ctl.Data["URL"] = ctl.URL

	ctl.Data["MenuProductAttributeValueActive"] = "active"
}
func (ctl *ProductAttributeValueController) Create() {
	ctl.Data["Action"] = "create"
	ctl.Data["Readonly"] = false
	ctl.Data["listName"] = "创建属性值"
	ctl.Layout = "base/base.html"
	ctl.TplName = "product/product_attribute_value_form.html"
}
func (ctl *ProductAttributeValueController) Edit() {
	id := ctl.Ctx.Input.Param(":id")
	mapInfo := make(map[string]interface{})
	if id != "" {
		if idInt64, e := strconv.ParseInt(id, 10, 64); e == nil {

			if attributeValue, err := mp.GetProductAttributeValueByID(idInt64); err == nil {

				mapInfo["name"] = attributeValue.Name
				attribute := make(map[string]interface{})
				if attributeValue.Attribute != nil {
					attribute["id"] = attributeValue.Attribute.Id
					attribute["name"] = attributeValue.Attribute.Name
				}
				mapInfo["attribute"] = attribute
			}
		}
	}
	ctl.Data["Action"] = "edit"
	ctl.Data["RecordId"] = id
	ctl.Data["ProductAttValue"] = mapInfo
	ctl.Layout = "base/base.html"
	ctl.TplName = "product/product_attribute_value_form.html"
}
func (ctl *ProductAttributeValueController) Detail() {
	//获取信息一样，直接调用Edit
	ctl.Edit()
	ctl.Data["Readonly"] = true
	ctl.Data["Action"] = "detail"
}
func (ctl *ProductAttributeValueController) PostCreate() {
	attrValue := new(mp.ProductAttributeValue)
	if err := ctl.ParseForm(attrValue); err == nil {
		if attributeID, err := ctl.GetInt64("productAttributeID"); err == nil {
			if attribute, err := mp.GetProductAttributeByID(attributeID); err == nil {
				attrValue.Attribute = &attribute
			}
		}
		if id, err := mp.CreateProductAttributeValue(attrValue, ctl.User); err == nil {
			ctl.Redirect("/product/attributevalue/"+strconv.FormatInt(id, 10)+"?action=detail", 302)
		} else {
			ctl.Get()
		}
	} else {
		ctl.Get()
	}
}
func (ctl *ProductAttributeValueController) Validator() {
	name := ctl.GetString("name")
	name = strings.TrimSpace(name)
	recordID := ctl.GetString("recordId")
	result := make(map[string]bool)
	obj, err := mp.GetProductAttributeValueByName(name)
	if err != nil {
		result["valid"] = true
	} else {
		if obj.Name == name {
			if recordID != "" {
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

// 获得符合要求的数据
func (ctl *ProductAttributeValueController) productAttributeList(start, length int64, condArr map[string]interface{}) (map[string]interface{}, error) {

	var arrs []mp.ProductAttributeValue
	paginator, arrs, err := mp.ListProductAttributeValue(condArr, start, length)
	result := make(map[string]interface{})
	if err == nil {

		//使用多线程来处理数据，待修改
		tableLines := make([]interface{}, 0, 4)
		for _, line := range arrs {
			oneLine := make(map[string]interface{})
			oneLine["name"] = line.Name
			oneLine["attribute"] = line.Attribute.Name
			oneLine["Id"] = line.Id
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
	viewType := ctl.Input().Get("view")
	if viewType == "" || viewType == "table" {
		ctl.Data["ViewType"] = "table"
	}
	ctl.Data["listName"] = "产品属性值管理"
	ctl.Data["tableId"] = "table-product-attributevalue"
	ctl.Layout = "base/base_list_view.html"
	ctl.TplName = "product/product_attribute_value_list_search.html"
}
