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

func (ctl *ProductProductController) Post() {
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
func (ctl *ProductProductController) Get() {
	ctl.URL = "/product/product/"
	ctl.Data["URL"] = ctl.URL
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
	ctl.URL = "/product/product/"
	ctl.Data["URL"] = ctl.URL
	ctl.Data["MenuProductProductActive"] = "active"

}
func (ctl *ProductProductController) Put() {
	id := ctl.Ctx.Input.Param(":id")
	ctl.URL = "/product/product/"
	//需要判断文件上传时页面不用跳转的情况
	if idInt64, e := strconv.ParseInt(id, 10, 64); e == nil {
		if product, err := mp.GetProductProductByID(idInt64); err == nil {
			if err := ctl.ParseForm(&product); err == nil {

				if _, err := mp.UpdateProductProduct(&product, ctl.User); err == nil {
					ctl.Redirect(ctl.URL+id+"?action=detail", 302)
				}
			}
		}
	}
	ctl.Redirect(ctl.URL+id+"?action=edit", 302)
}
func (ctl *ProductProductController) Create() {
	ctl.Data["Action"] = "create"
	ctl.Data["Readonly"] = false
	ctl.Data["listName"] = "创建规格"
	ctl.Layout = "base/base.html"
	ctl.TplName = "product/product_product_form.html"
}
func (ctl *ProductProductController) Edit() {
	id := ctl.Ctx.Input.Param(":id")
	productInfo := make(map[string]interface{})
	if id != "" {
		if idInt64, e := strconv.ParseInt(id, 10, 64); e == nil {
			if product, err := mp.GetProductProductByID(idInt64); err == nil {
				productInfo["name"] = product.Name
				productInfo["defaultCode"] = product.DefaultCode
				productInfo["standardPrice"] = product.DefaultCode

				// 款式类别
				categ := product.Categ
				categValues := make(map[string]string)
				if categ != nil {
					categValues["id"] = strconv.FormatInt(categ.Id, 10)
					categValues["name"] = categ.Name
				}
				productInfo["category"] = categValues
				// 销售第一单位
				firstSaleUom := product.FirstSaleUom
				firstSaleUomValues := make(map[string]string)
				if firstSaleUom != nil {
					firstSaleUomValues["id"] = strconv.FormatInt(firstSaleUom.Id, 10)
					firstSaleUomValues["name"] = firstSaleUom.Name
				}
				productInfo["firstSaleUom"] = firstSaleUomValues
				// 销售第二单位
				secondSaleUom := product.SecondSaleUom
				secondSaleUomValues := make(map[string]string)
				if secondSaleUom != nil {
					secondSaleUomValues["id"] = strconv.FormatInt(secondSaleUom.Id, 10)
					secondSaleUomValues["name"] = secondSaleUom.Name
				}
				productInfo["secondSaleUom"] = secondSaleUomValues
				// 采购第一单位
				firstPurchaseUom := product.FirstPurchaseUom
				firstPurchaseUomValues := make(map[string]string)
				if firstPurchaseUom != nil {
					firstPurchaseUomValues["id"] = strconv.FormatInt(firstPurchaseUom.Id, 10)
					firstPurchaseUomValues["name"] = firstPurchaseUom.Name
				}
				productInfo["firstPurchaseUom"] = firstSaleUomValues
				// 采购第二单位
				secondPurchaseUom := product.SecondPurchaseUom
				secondPurchaseUomValues := make(map[string]string)
				if secondSaleUom != nil {
					secondPurchaseUomValues["id"] = strconv.FormatInt(secondPurchaseUom.Id, 10)
					secondPurchaseUomValues["name"] = secondPurchaseUom.Name
				}
				productInfo["secondPurchaseUom"] = secondPurchaseUomValues
			}
		}
	}
	ctl.Data["Action"] = "edit"
	ctl.Data["RecordId"] = id
	ctl.Data["Product"] = productInfo
	ctl.Layout = "base/base.html"
	ctl.TplName = "product/product_product_form.html"
}
func (ctl *ProductProductController) Detail() {
	ctl.Edit()
	ctl.Data["Readonly"] = true
	ctl.Data["Action"] = "detail"
}
func (ctl *ProductProductController) Validator() {
	name := ctl.GetString("name")
	name = strings.TrimSpace(name)
	recordID, _ := ctl.GetInt64("recordId")
	result := make(map[string]bool)
	obj, err := mp.GetProductProductByName(name)
	if err != nil {
		result["valid"] = true
	} else {
		if obj.Name == name {
			if recordID == obj.Id {
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

// 获得符合要求的城市数据
func (ctl *ProductProductController) productProductList(start, length int64, condArr map[string]interface{}) (map[string]interface{}, error) {

	var arrs []mp.ProductProduct
	paginator, arrs, err := mp.ListProductProduct(condArr, start, length)
	result := make(map[string]interface{})
	if err == nil {

		//使用多线程来处理数据，待修改
		tableLines := make([]interface{}, 0, 4)
		for _, line := range arrs {
			oneLine := make(map[string]interface{})
			oneLine["name"] = line.Name
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
func (ctl *ProductProductController) PostList() {
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
	if result, err := ctl.productProductList(startInt64, lengthInt64, condArr); err == nil {
		ctl.Data["json"] = result
	}
	ctl.ServeJSON()

}

func (ctl *ProductProductController) GetList() {
	viewType := ctl.Input().Get("view")
	if viewType == "" || viewType == "table" {
		ctl.Data["ViewType"] = "table"
	}
	ctl.Data["listName"] = "产品规格管理"
	ctl.Data["tableId"] = "table-product-product"
	ctl.Layout = "base/base_list_view.html"
	ctl.TplName = "product/product_product_list_search.html"
}
