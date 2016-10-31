package product

import (
	"fmt"
	"pms/controllers/base"
	"pms/models/product"
	"pms/utils"
	"strconv"
)

const (
	productCategoryListCellLength = 3
)

type ProductCategoryController struct {
	base.BaseController
}

func (this *ProductCategoryController) Get() {
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
	this.Data["searchKeyWords"] = "产品类别"
}
func (this *ProductCategoryController) List() {
	this.Data["listName"] = "产品类别"
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
	var productcategories []product.ProductCategory
	paginator, err, productcategories := product.ListProductCategory(condArr, pageInt64, offsetInt64)
	paginator.Url = "/city"
	this.Data["Paginator"] = paginator
	tableInfo := new(utils.TableInfo)
	tableInfo.Url = "/city"
	tableTitle := make(map[string]interface{})
	tableTitle["titleName"] = [productCategoryListCellLength]string{"类别", "上级类别", "操作"}
	tableInfo.Title = tableTitle
	tableBody := make(map[string]interface{})
	bodyLines := make([]interface{}, 0, 20)
	if err == nil {
		for _, productcategory := range productcategories {
			oneLine := make([]interface{}, productCategoryListCellLength, productCategoryListCellLength)
			lineInfo := make(map[string]interface{})
			action := map[string]map[string]string{}
			edit := make(map[string]string)
			detail := make(map[string]string)
			id := int(productcategory.Id)

			lineInfo["id"] = id
			oneLine[0] = productcategory.Name

			edit["name"] = "编辑"
			edit["url"] = tableInfo.Url + "/edit/" + strconv.Itoa(id)
			detail["name"] = "详情"
			detail["url"] = tableInfo.Url + "/detail/" + strconv.Itoa(id)
			action["edit"] = edit
			action["detail"] = detail
			// if productcategory.Parent != nil {
			// 	oneLine[1] = product.GetFullPathCategory(productcategory)
			// } else {
			// 	oneLine[1] = "-"
			// }
			oneLine[1] = product.GetFullPathCategory(productcategory)
			kko, _ := product.GetProductCategory(productcategory.Id)
			fmt.Println(kko)
			oneLine[2] = action

			lineData := make(map[string]interface{})
			lineData["oneLine"] = oneLine
			lineData["lineInfo"] = lineInfo
			bodyLines = append(bodyLines, lineData)
		}
		tableBody["bodyLines"] = bodyLines
		tableInfo.Body = tableBody
		tableInfo.TitleLen = productCategoryListCellLength
		tableInfo.TitleIndexLen = productCategoryListCellLength - 1
		tableInfo.BodyLen = paginator.CurrentPageSize
		this.Data["tableInfo"] = tableInfo
	}
}
