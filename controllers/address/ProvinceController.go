package address

import (
	"encoding/json"
	"pms/controllers/base"
	mb "pms/models/base"
	"strconv"
	"strings"
)

type ProvinceController struct {
	base.BaseController
}

func (ctl *ProvinceController) Post() {
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
func (ctl *ProvinceController) Get() {

	ctl.URL = "/address/city/"
	ctl.Data["URL"] = ctl.URL
	ctl.Data["MenuProvinceActive"] = "active"
	ctl.GetList()
}
func (ctl *ProvinceController) PostList() {
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
	if result, err := ctl.provinceList(startInt64, lengthInt64, condArr); err == nil {
		ctl.Data["json"] = result
	}
	ctl.ServeJSON()

}
func (ctl *ProvinceController) Validator() {
	name := ctl.GetString("name")
	name = strings.TrimSpace(name)
	recordID, _ := ctl.GetInt64("recordId")
	result := make(map[string]bool)
	obj, err := mb.GetPositionByName(name)
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

// 获得符合要求的地区数据
func (ctl *ProvinceController) provinceList(start, length int64, condArr map[string]interface{}) (map[string]interface{}, error) {

	var provinces []mb.Province
	paginator, provinces, err := mb.ListProvince(condArr, start, length)
	result := make(map[string]interface{})
	if err == nil {

		// result["recordsFiltered"] = paginator.TotalCount
		tableLines := make([]interface{}, 0, 4)
		for _, province := range provinces {
			oneLine := make(map[string]interface{})
			oneLine["name"] = province.Name
			oneLine["country"] = province.Country.Name
			oneLine["Id"] = province.Id
			oneLine["id"] = province.Id

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

func (ctl *ProvinceController) GetList() {
	viewType := ctl.Input().Get("view")
	if viewType == "" || viewType == "table" {
		ctl.Data["ViewType"] = "table"
	}
	ctl.Data["tableId"] = "table-province"
	ctl.Layout = "base/base_list_view.html"
	ctl.TplName = "address/province_list_search.html"
}
