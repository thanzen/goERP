package base

import (
	"encoding/json"
	mb "pms/models/base"
	"strconv"
	"strings"
)

type DepartmentController struct {
	BaseController
}

func (this *DepartmentController) Post() {

	action := this.Input().Get("action")
	switch action {
	case "validator":
		this.Validator()
	case "table": //bootstrap table的post请求
		this.PostList()
	case "selectSearch":
		this.PostList()
	default:
		this.PostList()
	}
}

func (this *DepartmentController) Get() {
	this.GetList()

	this.URL = "/department"
	this.Data["URL"] = this.URL
	this.Layout = "base/base.html"
	this.Data["MenuDepartmentActive"] = "active"
}

func (this *DepartmentController) Validator() {
	name := this.GetString("name")
	name = strings.TrimSpace(name)
	result := make(map[string]bool)
	if _, err := mb.GetDepartmentByName(name); err != nil {
		result["valid"] = true
	} else {
		result["valid"] = false
	}
	this.Data["json"] = result
	this.ServeJSON()
}

// 获得符合要求的城市数据
func (this *DepartmentController) departmentList(start, length int64, condArr map[string]interface{}) (map[string]interface{}, error) {

	var departments []mb.Department
	paginator, departments, err := mb.ListDepartment(condArr, start, length)

	result := make(map[string]interface{})
	if err == nil {

		// result["recordsFiltered"] = paginator.TotalCount
		tableLines := make([]interface{}, 0, 4)
		for _, department := range departments {
			oneLine := make(map[string]interface{})

			oneLine["Id"] = department.Id
			oneLine["id"] = department.Id
			oneLine["name"] = department.Name

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
func (this *DepartmentController) PostList() {
	condArr := make(map[string]interface{})
	start := this.Input().Get("offset")
	length := this.Input().Get("limit")
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
	if result, err := this.departmentList(startInt64, lengthInt64, condArr); err == nil {
		this.Data["json"] = result
	}
	this.ServeJSON()

}

func (this *DepartmentController) GetList() {
	this.Data["tableId"] = "table-department"
	this.TplName = "base/table_base.html"
}
