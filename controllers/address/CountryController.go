package address

import (
	"encoding/json"
	"pms/controllers/base"
	mb "pms/models/base"
	"strconv"
	"strings"
)

type CountryController struct {
	base.BaseController
}

func (ctl *CountryController) Post() {
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
func (ctl *CountryController) Get() {

	ctl.GetList()

	ctl.URL = "/address/city"
	ctl.Data["URL"] = ctl.URL
	ctl.Layout = "base/base.html"
	ctl.Data["MenuCountryActive"] = "active"
}

func (ctl *CountryController) PostList() {
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
	if result, err := ctl.countryList(startInt64, lengthInt64, condArr); err == nil {
		ctl.Data["json"] = result
	}
	ctl.ServeJSON()

}

// 获得符合要求的国家数据
func (ctl *CountryController) countryList(start, length int64, condArr map[string]interface{}) (map[string]interface{}, error) {

	var countries []mb.Country
	paginator, countries, err := mb.ListCountry(condArr, start, length)
	result := make(map[string]interface{})
	if err == nil {

		// result["recordsFiltered"] = paginator.TotalCount
		tableLines := make([]interface{}, 0, 4)
		for _, country := range countries {
			oneLine := make(map[string]interface{})
			oneLine["name"] = country.Name
			oneLine["Id"] = country.Id
			oneLine["id"] = country.Id
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

func (ctl *CountryController) Validator() {
	name := ctl.GetString("name")
	name = strings.TrimSpace(name)
	result := make(map[string]bool)
	if _, err := mb.GetCountryByName(name); err != nil {
		result["valid"] = true
	} else {
		result["valid"] = false
	}
	ctl.Data["json"] = result
	ctl.ServeJSON()
}

func (ctl *CountryController) GetList() {
	ctl.Data["tableId"] = "table-country"
	ctl.TplName = "base/table_base.html"
}
