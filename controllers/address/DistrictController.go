package address

import (
	"encoding/json"
	"pms/controllers/base"
	mb "pms/models/base"
	"strings"

	"strconv"
)

type DistrictController struct {
	base.BaseController
}

func (ctl *DistrictController) Post() {
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
func (ctl *DistrictController) PostList() {
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
	if result, err := ctl.districtList(startInt64, lengthInt64, condArr); err == nil {
		ctl.Data["json"] = result
	}
	ctl.ServeJSON()

}

// 获得符合要求的地区数据
func (ctl *DistrictController) districtList(start, length int64, condArr map[string]interface{}) (map[string]interface{}, error) {

	var districtes []mb.District
	paginator, districtes, err := mb.ListDistrict(condArr, start, length)
	result := make(map[string]interface{})
	if err == nil {
		provinceMap := make(map[int64]string)
		// result["recordsFiltered"] = paginator.TotalCount
		tableLines := make([]interface{}, 0, 4)
		for _, district := range districtes {
			oneLine := make(map[string]interface{})
			oneLine["name"] = district.Name
			oneLine["province"] = district.City.Province.Name

			provinceId := district.City.Province.Id
			if _, ok := provinceMap[provinceId]; ok != true {
				if province, e := mb.GetProvinceByID(district.City.Province.Id); e == nil {
					provinceMap[provinceId] = province.Country.Name
				}
			}
			if _, ok := provinceMap[provinceId]; ok {
				oneLine["country"] = provinceMap[provinceId]
			}
			oneLine["city"] = district.City.Name
			oneLine["Id"] = district.Id
			oneLine["id"] = district.Id
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
func (ctl *DistrictController) Validator() {
	name := ctl.GetString("name")
	name = strings.TrimSpace(name)
	recordID, _ := ctl.GetInt64("recordId")
	result := make(map[string]bool)
	obj, err := mb.GetDistrictByName(name)
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

func (ctl *DistrictController) Get() {

	ctl.Data["listName"] = "区县管理"
	ctl.URL = "/address/district/"
	ctl.Data["URL"] = ctl.URL
	ctl.Data["MenuDistrictActive"] = "active"
	ctl.GetList()

}
func (ctl *DistrictController) GetList() {
	viewType := ctl.Input().Get("view")
	if viewType == "" || viewType == "table" {
		ctl.Data["ViewType"] = "table"
	}
	ctl.Data["tableId"] = "table-district"
	ctl.Layout = "base/base_list_view.html"
	ctl.TplName = "address/district_list_search.html"
}
