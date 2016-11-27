package address

import (
	"pms/controllers/base"
	mb "pms/models/base"
	"strconv"
	"strings"
)

type ProvinceController struct {
	base.BaseController
}

func (this *ProvinceController) Post() {
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
func (this *ProvinceController) Get() {

	this.GetList()

	this.URL = "/city"
	this.Data["URL"] = this.URL
	this.Layout = "base/base.html"
	this.Data["MenuProvinceActive"] = "active"
}
func (this *ProvinceController) Validator() {
	name := this.GetString("name")
	name = strings.TrimSpace(name)
	result := make(map[string]bool)
	if _, err := mb.GetProvinceByName(name); err != nil {
		result["valid"] = true
	} else {
		result["valid"] = false
	}
	this.Data["json"] = result
	this.ServeJSON()
}
func (this *ProvinceController) Table() {
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
	var provinces []mb.Province
	paginator, provinces, err := mb.ListProvince(condArr, startInt64, lengthInt64)
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
		for _, province := range provinces {
			oneLine := make(map[string]interface{})
			oneLine["name"] = province.Name
			oneLine["country"] = province.Country.Name
			oneLine["id"] = province.Id
			tableLines = append(tableLines, oneLine)
		}
		result["data"] = tableLines
	}
	this.Data["json"] = result
	this.ServeJSON()
}

func (this *ProvinceController) GetList() {
	this.TplName = "address/table_province.html"
}
