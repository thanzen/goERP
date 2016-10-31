package product

import (
	"pms/controllers/base"
	"pms/models/product"
	"pms/utils"
	"strconv"
)

const (
	productAttributeValueListCellLength = 2
)

type ProductAttributeValueController struct {
	base.BaseController
}

func (this *ProductAttributeValueController) Get() {
	action := this.GetString(":action")
	viewType := this.Input().Get("view_type")
	switch action {
	case "list":
		switch viewType {
		case "list":
			this.List()
		default:
			this.List()
		}
	default:
		this.List()
	}
	this.Data["searchKeyWords"] = "属性值"
}
func (this *ProductAttributeValueController) List() {
	this.Data["listName"] = "产品属性值"
	this.Layout = "base/base.html"
	this.TplName = "user/record_list.html"
	this.Data["productRootActive"] = "active"
	this.Data["productAttributeActive"] = "active"
	condArr := make(map[string]interface{})
	page := this.Input().Get("page")
	offset := this.Input().Get("offset")
	var (
		err         error
		pageInt64   int64
		offsetInt64 int64
	)
	if pageInt, ok := strconv.Atoi(page); ok == nil {
		pageInt64 = int64(pageInt)
	}
	if offsetInt, ok := strconv.Atoi(offset); ok == nil {
		offsetInt64 = int64(offsetInt)
	}
	var productAttributeValues []product.ProductAttributeValue
	paginator, err, productAttributeValues := product.ListProductAttributeValue(condArr, pageInt64, offsetInt64)
	paginator.Url = "/city"
	this.Data["Paginator"] = paginator
	tableInfo := new(utils.TableInfo)
	tableInfo.Url = "/city"
	tableTitle := make(map[string]interface{})
	tableTitle["titleName"] = [productAttributeListCellLength]string{"属性", "操作"}
	tableInfo.Title = tableTitle
	tableBody := make(map[string]interface{})
	bodyLines := make([]interface{}, 0, 20)
	if err == nil {
		for _, productAttributeValue := range productAttributeValues {
			oneLine := make([]interface{}, productAttributeListCellLength, productAttributeListCellLength)
			lineInfo := make(map[string]interface{})
			action := map[string]map[string]string{}
			edit := make(map[string]string)
			detail := make(map[string]string)
			id := int(productAttributeValue.Id)

			lineInfo["id"] = id
			oneLine[0] = productAttributeValue.Name

			edit["name"] = "编辑"
			edit["url"] = tableInfo.Url + "/edit/" + strconv.Itoa(id)
			detail["name"] = "详情"
			detail["url"] = tableInfo.Url + "/detail/" + strconv.Itoa(id)
			action["edit"] = edit
			action["detail"] = detail

			oneLine[2] = action

			lineData := make(map[string]interface{})
			lineData["oneLine"] = oneLine
			lineData["lineInfo"] = lineInfo
			bodyLines = append(bodyLines, lineData)
		}
		tableBody["bodyLines"] = bodyLines
		tableInfo.Body = tableBody
		tableInfo.TitleLen = productAttributeListCellLength
		tableInfo.TitleIndexLen = productAttributeListCellLength - 1
		tableInfo.BodyLen = paginator.CurrentPageSize
		this.Data["tableInfo"] = tableInfo
	}
}
