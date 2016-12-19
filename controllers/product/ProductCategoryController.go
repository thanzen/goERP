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

func (ctl *ProductCategoryController) Post() {
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
func (ctl *ProductCategoryController) Put() {
	id := ctl.Ctx.Input.Param(":id")
	ctl.URL = "/product/category"
	if idInt64, e := strconv.ParseInt(id, 10, 64); e == nil {
		if category, err := mp.GetProductCategoryByID(idInt64); err == nil {
			if err := ctl.ParseForm(&category); err == nil {
				if parentId, err := ctl.GetInt64("parent"); err == nil {
					if parent, err := mp.GetProductCategoryByID(parentId); err == nil {
						category.Parent = &parent
					}
				}
				if _, err := mp.UpdateProductCategory(&category, ctl.User); err == nil {
					ctl.Redirect(ctl.URL+"/"+id+"?action=detail", 302)
				}
			}
		}
	}
	ctl.Redirect(ctl.URL+"/"+id+"?action=edit", 302)

}
func (ctl *ProductCategoryController) Get() {

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
	ctl.URL = "/product/category"
	ctl.Data["URL"] = ctl.URL
	ctl.Layout = "base/base.html"
	ctl.Data["MenuProductCategoryActive"] = "active"
}
func (ctl *ProductCategoryController) Edit() {
	id := ctl.Ctx.Input.Param(":id")
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
	ctl.Data["Action"] = "edit"
	ctl.Data["RecordId"] = id
	ctl.Data["Category"] = categoryInfo
	ctl.TplName = "product/product_category_form.html"
}

func (ctl *ProductCategoryController) Detail() {
	//获取信息一样，直接调用Edit
	ctl.Edit()
	ctl.Data["Readonly"] = true
	ctl.Data["Action"] = "detail"
}

//post请求创建产品分类
func (ctl *ProductCategoryController) PostCreate() {
	category := new(mp.ProductCategory)
	if err := ctl.ParseForm(category); err == nil {
		if parentID, err := ctl.GetInt64("parent"); err == nil {
			if parent, err := mp.GetProductCategoryByID(parentID); err == nil {
				category.Parent = &parent
			}
		}
		if id, err := mp.CreateProductCategory(category, ctl.User); err == nil {
			ctl.Redirect("/product/category/"+strconv.FormatInt(id, 10)+"?action=detail", 302)
		} else {
			ctl.Get()
		}
	} else {
		ctl.Get()
	}

}
func (ctl *ProductCategoryController) Create() {
	ctl.Data["Action"] = "create"
	ctl.Data["Readonly"] = false
	ctl.Data["listName"] = "创建类别"
	ctl.TplName = "product/product_category_form.html"

}
func (ctl *ProductCategoryController) Validator() {
	name := ctl.GetString("name")
	recordId := ctl.GetString("recordId")
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
	ctl.Data["json"] = result
	ctl.ServeJSON()
}

// 获得符合要求的城市数据
func (ctl *ProductCategoryController) productCategoryList(start, length int64, condArr map[string]interface{}) (map[string]interface{}, error) {

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
func (ctl *ProductCategoryController) PostList() {
	condArr := make(map[string]interface{})
	start := ctl.Input().Get("offset")
	length := ctl.Input().Get("limit")
	name := ctl.Input().Get("name")
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
	if result, err := ctl.productCategoryList(startInt64, lengthInt64, condArr); err == nil {
		ctl.Data["json"] = result
	}
	ctl.ServeJSON()

}

func (ctl *ProductCategoryController) GetList() {
	ctl.Data["tableId"] = "table-product-category"
	ctl.TplName = "base/table_base.html"
}
