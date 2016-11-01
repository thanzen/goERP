package product

import (
	"pms/controllers/base"
	"pms/models/product"
	"pms/utils"
	"strconv"
)

const (
	productAttributeListCellLength = 2
)

type ProductAttributeController struct {
	base.BaseController
}

func (this *ProductAttributeController) Get() {
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
	case "create":
		this.Create()
	default:
		this.List()
	}
	this.Data["searchKeyWords"] = "属性值"
}
func (this *ProductAttributeController) Create() {
	this.Layout = "base/base.html"
	this.TplName = "product/product_attribute_form.html"
	this.Data["Url"] = "/product/category"
	this.Data["Action"] = "create"
}
func (this *ProductAttributeController) List() {
	this.Data["listName"] = "产品属性信息"
	this.Layout = "base/base.html"
	this.TplName = "product/product_attribute_list.html"
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
	var productAttributes []product.ProductAttribute
	paginator, err, productAttributes := product.ListProductAttribute(condArr, pageInt64, offsetInt64)
	URL := "/product/attribute"
	this.Data["URL"] = URL

	this.Data["Paginator"] = paginator
	tableInfo := new(utils.TableInfo)
	tableTitle := make(map[string]interface{})
	tableTitle["titleName"] = [productAttributeListCellLength]string{"属性", "操作"}
	tableInfo.Title = tableTitle
	tableBody := make(map[string]interface{})
	bodyLines := make([]interface{}, 0, 20)
	if err == nil {
		for _, productAttribute := range productAttributes {
			oneLine := make([]interface{}, productAttributeListCellLength, productAttributeListCellLength)
			lineInfo := make(map[string]interface{})
			action := map[string]map[string]string{}
			edit := make(map[string]string)
			detail := make(map[string]string)
			id := int(productAttribute.Id)

			lineInfo["id"] = id
			oneLine[0] = productAttribute.Name

			edit["name"] = "编辑"
			edit["url"] = URL + "/edit/" + strconv.Itoa(id)
			detail["name"] = "详情"
			detail["url"] = URL + "/detail/" + strconv.Itoa(id)
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
