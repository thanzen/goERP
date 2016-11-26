package address

import (
	"pms/controllers/base"
	mb "pms/models/base"

	"strconv"
)

const (
	districtListCellLength = 4
)

type DistrictController struct {
	base.BaseController
}

func (this *DistrictController) Post() {
	this.PostList()
}
func (this *DistrictController) PostList() {
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
	var districts []mb.District
	paginator, districts, err := mb.ListDistrict(condArr, startInt64, lengthInt64)
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
		for _, district := range districts {
			oneLine := make(map[string]interface{})
			oneLine["name"] = district.Name

			oneLine["province"] = district.City.Province.Name
			oneLine["city"] = district.City.Name
			oneLine["id"] = district.Id
			tableLines = append(tableLines, oneLine)
		}
		result["data"] = tableLines
	}
	this.Data["json"] = result
	this.ServeJSON()
}
func (this *DistrictController) Get() {
	this.List()
	this.List()

	this.URL = "/district"
	this.Data["URL"] = this.URL
	this.Layout = "base/base.html"
	this.Data["MenuDistrictActive"] = "active"

}
func (this *DistrictController) List() {
	this.Data["listName"] = "区县信息"
	this.Data["Readonly"] = true
	this.TplName = "address/table_district.html"
}
