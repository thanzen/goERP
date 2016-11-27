package base

import (
	"pms/models/base"
	"strings"
)

//列表视图列数-1，第一列为checkbox
const (
	departmentListCellLength = 4
)

type DepartmentController struct {
	BaseController
}

func (this *DepartmentController) Post() {

	action := this.GetString(":action")
	switch action {
	case "search":
		this.Search()
	}
	this.ServeJSON()
}

func (this *DepartmentController) Get() {
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
	case "show":
		this.Show()
	case "create":
		this.Create()
	case "edit":
		this.Edit()
	default:
		this.List()
	}
	this.Data["searchKeyWords"] = "部门名称"
	this.URL = "/department"
	this.Data["URL"] = this.URL
	this.Layout = "base/base.html"
}
func (this *DepartmentController) Validator() {
	name := this.GetString("name")
	name = strings.TrimSpace(name)
	result := make(map[string]bool)
	if _, err := base.GetDepartmentByName(name); err != nil {
		result["valid"] = true
	} else {
		result["valid"] = false
	}
	this.Data["json"] = result
	this.ServeJSON()
}
func (this *DepartmentController) List() {

}
func (this *DepartmentController) Show() {

}
func (this *DepartmentController) Create() {

}

//get请求
func (this *DepartmentController) Edit() {

}

//post请求
func (this *DepartmentController) Update() {

}
func (this *DepartmentController) Search() {
	name := this.GetString("name")

	page, _ := this.GetInt64("page")
	offset, _ := this.GetInt64("offset")
	var condArr = make(map[string]interface{})
	name = strings.TrimSpace(name)
	if name != "" {
		condArr["name"] = name
	}
	paginator, departments, err := base.ListDepartment(condArr, page, offset)
	data := make(map[string]interface{})
	if err == nil {

		items := make([]interface{}, 0, 5)
		for _, department := range departments {
			line := make(map[string]interface{})
			line["id"] = department.Id
			line["name"] = department.Name
			if department.Leader != nil {
				line["leader"] = department.Leader.Name
			} else {
				line["leader"] = "-"
			}
			items = append(items, line)
		}
		data["items"] = items
		data["total_count"] = paginator.TotalCount
		data["pageSize"] = paginator.PageSize
		data["page"] = 2 //paginator.CurrentPage
	} else {
		data["msg"] = "failed"
	}
	this.Data["json"] = data
}
