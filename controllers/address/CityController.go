package address

import (
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
	case "table":
		this.Table()
	default:
		this.Table()
	}
}
func (this *CityController) Get() {

	this.GetList()

	this.URL = "/city"
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
func (this *CityController) Table() {
	start := this.Input().Get("start")
	length := this.Input().Get("length")

	condArr := make(map[string]interface{})
	var (
		err         error
		startInt64  int64
		lengthInt64 int64
	)
	if startInt, ok := strconv.Atoi(start); ok == nil {
		startInt64 = int64(startInt)
	}
	if lengthInt, ok := strconv.Atoi(length); ok == nil {
		lengthInt64 = int64(lengthInt)
	}
	var cities []mb.City
	paginator, cities, err := mb.ListCity(condArr, startInt64, lengthInt64)
	result := make(map[string]interface{})
	if err == nil {
		result["draw"] = this.Input().Get("draw")
		result["recordsTotal"] = paginator.TotalCount
		result["recordsFiltered"] = paginator.TotalCount
		result["page"] = paginator.CurrentPage
		result["pages"] = paginator.TotalPage
		result["start"] = paginator.CurrentPage * paginator.PageSize
		result["length"] = length
		result["serverSide"] = true
		result["currentPageSize"] = paginator.CurrentPageSize

		// result["recordsFiltered"] = paginator.TotalCount
		tableLines := make([]interface{}, 0, 4)
		for _, city := range cities {
			oneLine := make(map[string]interface{})
			oneLine["name"] = city.Name
			oneLine["province"] = city.Province.Name
			oneLine["country"] = city.Province.Country.Name
			oneLine["id"] = city.Id
			tableLines = append(tableLines, oneLine)
		}
		result["data"] = tableLines
	}
	this.Data["json"] = result
	this.ServeJSON()
}

func (this *CityController) GetList() {
	this.TplName = "address/table_city.html"
}
