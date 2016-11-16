package base

import "pms/models/base"

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
	case "detail":
		this.Detail()
	case "create":
		this.Create()
	case "edit":
		this.Edit()
	case "search":
		this.Search()
	default:
		this.List()
	}
	this.Data["searchKeyWords"] = "部门名称"
	this.URL = "/department"
	this.Data["URL"] = this.URL
	this.Layout = "base/base.html"
}
func (this *DepartmentController) List() {

}
func (this *DepartmentController) Detail() {

}
func (this *DepartmentController) Create() {

}
func (this *DepartmentController) Edit() {

}
func (this *DepartmentController) Search() {
	name := this.GetString("name")
	this.Data["json"] = ""
	if _, departments, err := base.GetDepartmentByName(name, false); err == nil {
		data := make([]interface{}, 0)
		for _, department := range departments {
			line := make(map[string]interface{})
			line["id"] = department.Id
			line["name"] = department.Name
			if department.Leader != nil {
				line["leader"] = department.Leader.Name
			} else {
				line["leader"] = "-"
			}

			data = append(data, line)
		}
		this.Data["json"] = data
	}

}
