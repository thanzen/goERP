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

func (ctl *DepartmentController) Post() {

	action := ctl.Input().Get("action")
	switch action {
	case "validator":
		ctl.Validator()
	case "table": //bootstrap table的post请求
		ctl.PostList()
	case "selectSearch":
		ctl.PostList()
	default:
		ctl.PostList()
	}
}

func (ctl *DepartmentController) Get() {
	ctl.GetList()

	ctl.URL = "/department/"
	ctl.Data["URL"] = ctl.URL
	ctl.Layout = "base/base.html"
	ctl.Data["MenuDepartmentActive"] = "active"
}

func (ctl *DepartmentController) Validator() {
	name := ctl.GetString("name")
	name = strings.TrimSpace(name)
	result := make(map[string]bool)
	if _, err := mb.GetDepartmentByName(name); err != nil {
		result["valid"] = true
	} else {
		result["valid"] = false
	}
	ctl.Data["json"] = result
	ctl.ServeJSON()
}

// 获得符合要求的城市数据
func (ctl *DepartmentController) departmentList(start, length int64, condArr map[string]interface{}) (map[string]interface{}, error) {

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
func (ctl *DepartmentController) PostList() {
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
	if result, err := ctl.departmentList(startInt64, lengthInt64, condArr); err == nil {
		ctl.Data["json"] = result
	}
	ctl.ServeJSON()

}

func (ctl *DepartmentController) GetList() {
	ctl.Data["tableId"] = "table-department"
	ctl.TplName = "base/table_base.html"
}
