package product

import (
	"pms/controllers/base"
	"pms/models/product"
	"pms/utils"
	"strconv"
)

const (
	productProductListCellLength = 2
)

type ProductProductController struct {
	base.BaseController
}

func (this *ProductProductController) Get() {
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
	this.Data["searchKeyWords"] = "产品规格名"
	this.URL = "/product/product"
	this.Data["URL"] = this.URL
	this.Data["productRootActive"] = "active"
	this.Data["productProductActive"] = "active"
	this.Layout = "base/base.html"
}
func (this *ProductProductController) List() {
	this.Data["listName"] = "产品规格"

	this.TplName = "user/record_list.html"

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
	var productProducts []product.ProductProduct
	paginator, err, productProducts := product.ListProductProduct(condArr, pageInt64, offsetInt64)

	this.Data["Paginator"] = paginator
	tableInfo := new(utils.TableInfo)
	tableTitle := make(map[string]interface{})
	tableTitle["titleName"] = [productProductListCellLength]string{"属性", "操作"}
	tableInfo.Title = tableTitle
	tableBody := make(map[string]interface{})
	bodyLines := make([]interface{}, 0, 20)
	if err == nil {
		for _, productProduct := range productProducts {
			oneLine := make([]interface{}, productProductListCellLength, productProductListCellLength)
			lineInfo := make(map[string]interface{})
			action := map[string]map[string]string{}
			edit := make(map[string]string)
			detail := make(map[string]string)
			id := int(productProduct.Id)

			lineInfo["id"] = id
			oneLine[0] = productProduct.Name

			edit["name"] = "编辑"
			edit["url"] = this.URL + "/edit/" + strconv.Itoa(id)
			detail["name"] = "详情"
			detail["url"] = this.URL + "/detail/" + strconv.Itoa(id)
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
		tableInfo.TitleLen = productProductListCellLength
		tableInfo.TitleIndexLen = productProductListCellLength - 1
		tableInfo.BodyLen = paginator.CurrentPageSize
		this.Data["tableInfo"] = tableInfo
	}
}
