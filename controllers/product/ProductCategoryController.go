package product

import (
	"encoding/json"
	"pms/controllers/base"
	mp "pms/models/product"
	"strconv"
	"strings"
)

type ProductCategoryController struct {
	base.BaseController
}

func (this *ProductCategoryController) Post() {
	action := this.Input().Get("action")
	switch action {
	case "validator":
		this.Validator()
	case "table": //bootstrap table的post请求
		this.PostList()
	case "create":
		this.PostCreate()
	default:
		this.PostList()
	}
}

func (this *ProductCategoryController) Get() {
	this.GetList()
	action := this.Input().Get("action")
	switch action {
	case "create":
		this.Create()
	case "edit":
		this.Edit()
	case "detail":
		this.Detail()
	default:
		this.GetList()

	}
	this.URL = "/product/category"
	this.Data["URL"] = this.URL
	this.Layout = "base/base.html"
	this.Data["MenuProductCategoryActive"] = "active"
}
func (this *ProductCategoryController) Edit() {
	id := this.Ctx.Input.Param(":id")
	categoryInfo := make(map[string]interface{})
	if id != "" {
		if idInt64, e := strconv.ParseInt(id, 10, 64); e == nil {

			if category, err := mp.GetProductCategoryByID(idInt64); err == nil {

				categoryInfo["name"] = category.Name
				parent := make(map[string]interface{})
				if category.Parent != nil {
					parent["id"] = category.Parent.Id
					parent["name"] = category.Parent.Name
				}
				categoryInfo["parent"] = parent

			}
		}

	}
	this.Data["Action"] = "edit"
	this.Data["RecordId"] = id
	this.Data["Category"] = categoryInfo

	this.TplName = "product/product_category_form.html"
}

func (this *ProductCategoryController) Detail() {
	//获取信息一样，直接调用Edit
	this.Edit()
	this.Data["Readonly"] = true
	this.Data["Action"] = "detial"
}

func (this *ProductCategoryController) PostCreate() {

	category := new(mp.ProductCategory)

	if err := this.ParseForm(category); err == nil {

		if parentId, err := this.GetInt64("parent"); err == nil {
			if parent, err := mp.GetProductCategoryByID(parentId); err == nil {
				category.Parent = &parent

			}
		}

		if id, err := mp.AddProductCategory(category, this.User); err == nil {
			this.Redirect("/product/category/"+strconv.FormatInt(id, 10), 302)
		} else {
			this.PostList()
		}
	} else {
		this.PostList()
	}

}
func (this *ProductCategoryController) Create() {
	method := strings.ToUpper(this.Ctx.Request.Method)
	if method == "GET" {
		this.Data["Action"] = "create"
		this.Data["Readonly"] = false
		this.Data["listName"] = "创建类别"
		this.TplName = "product/product_category_form.html"

	}
}
func (this *ProductCategoryController) Validator() {
	name := this.GetString("name")
	recordId := this.GetString("recordId")
	name = strings.TrimSpace(name)
	result := make(map[string]bool)
	obj, err := mp.GetProductCategoryByName(name)
	if err != nil {
		result["valid"] = true
	} else {
		if obj.Name == name {
			if recordId != "" {
				result["valid"] = true
			} else {
				result["valid"] = false
			}

		} else {
			result["valid"] = true
		}

	}
	this.Data["json"] = result
	this.ServeJSON()
}

// 获得符合要求的城市数据
func (this *ProductCategoryController) productCategoryList(start, length int64, condArr map[string]interface{}) (map[string]interface{}, error) {

	var arrs []mp.ProductCategory
	paginator, arrs, err := mp.ListProductCategory(condArr, start, length)
	result := make(map[string]interface{})
	if err == nil {

		// result["recordsFiltered"] = paginator.TotalCount
		tableLines := make([]interface{}, 0, 4)
		for _, line := range arrs {
			oneLine := make(map[string]interface{})
			oneLine["name"] = line.Name
			if line.Parent != nil {
				oneLine["parent"] = line.Parent.Name
			} else {
				oneLine["parent"] = "-"
			}
			oneLine["path"] = line.ParentFullPath
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
func (this *ProductCategoryController) PostList() {
	condArr := make(map[string]interface{})
	start := this.Input().Get("offset")
	length := this.Input().Get("limit")
	name := this.Input().Get("name")
	name = strings.TrimSpace(name)
	if name != "" {
		condArr["name"] = name
	}
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
	if result, err := this.productCategoryList(startInt64, lengthInt64, condArr); err == nil {
		this.Data["json"] = result
	}
	this.ServeJSON()

}

func (this *ProductCategoryController) GetList() {
	this.Data["tableId"] = "table-product-category"
	this.TplName = "base/table_base.html"
}
