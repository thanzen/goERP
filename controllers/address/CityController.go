package address

import (
	"encoding/json"
	"pms/controllers/base"
	mb "pms/models/base"
	"strconv"
	"strings"
)

type CityController struct {
	base.BaseController
}

func (this *CityController) Post() {
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
func (this *CityController) Get() {
	this.GetList()

	this.URL = "/address/city"
	this.Data["URL"] = this.URL
	this.Layout = "base/base.html"
	this.Data["MenuCityActive"] = "active"
}
func (this *CityController) Validator() {
	name := this.GetString("name")
	name = strings.TrimSpace(name)
	result := make(map[string]bool)
	if _, err := mb.GetCityByName(name); err != nil {
		result["valid"] = true
	} else {
		result["valid"] = false
	}
	this.Data["json"] = result
	this.ServeJSON()
}

// 获得符合要求的城市数据
func (this *CityController) cityList(start, length int64, condArr map[string]interface{}) (map[string]interface{}, error) {

	var cities []mb.City
	paginator, cities, err := mb.ListCity(condArr, start, length)
	result := make(map[string]interface{})
	if err == nil {

		// result["recordsFiltered"] = paginator.TotalCount
		tableLines := make([]interface{}, 0, 4)
		for _, city := range cities {
			oneLine := make(map[string]interface{})
			oneLine["name"] = city.Name
			oneLine["province"] = city.Province.Name
			oneLine["country"] = city.Province.Country.Name
			oneLine["Id"] = city.Id
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
func (this *CityController) PostList() {
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
	if result, err := this.cityList(startInt64, lengthInt64, condArr); err == nil {
		this.Data["json"] = result
	}
	this.ServeJSON()

}

func (this *CityController) GetList() {
	this.Data["tableId"] = "table-city"
	this.TplName = "base/table_base.html"
}
