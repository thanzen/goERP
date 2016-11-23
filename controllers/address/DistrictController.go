package address

import (
	"pms/controllers/base"
	mb "pms/models/base"
	"pms/utils"

	"strconv"
)

const (
	districtListCellLength = 4
)

type DistrictController struct {
	base.BaseController
}

func (this *DistrictController) Get() {
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
	default:
		this.List()
	}
	this.Data["searchKeyWords"] = "国家/省份/城市/区县"
}
func (this *DistrictController) List() {
	this.Data["listName"] = "区县信息"
	this.Layout = "base/base.html"
	this.TplName = "user/record_list.html"
	this.Data["settingRootActive"] = "active"
	this.Data["addressManageActive"] = "active"
	this.Data["addressDistrictActive"] = "active"
	condArr := make(map[string]interface{})
	page := this.Input().Get("page")
	offset := this.Input().Get("offset")
	var (
		err         error
		pageInt64   int64
		offsetInt64 int64
	)
	if pageInt, ok := strconv.Atoi(page); ok == nil {
		pageInt64 = int64(pageInt)
	}
	if offsetInt, ok := strconv.Atoi(offset); ok == nil {
		offsetInt64 = int64(offsetInt)
	}
	var districts []mb.District
	paginator, districts, err := mb.ListDistrict(condArr, pageInt64, offsetInt64)
	URL := "/district"
	this.Data["URL"] = URL
	this.Data["Paginator"] = paginator
	tableInfo := new(utils.TableInfo)
	tableTitle := make(map[string]interface{})
	tableTitle["titleName"] = [districtListCellLength]string{"区县", "城市", "省份", "操作"}
	tableInfo.Title = tableTitle
	tableBody := make(map[string]interface{})
	bodyLines := make([]interface{}, 0, base.ListNum)
	if err == nil {
		for _, district := range districts {

			//
			oneLine := make([]interface{}, districtListCellLength, districtListCellLength)
			lineInfo := make(map[string]interface{})
			action := map[string]map[string]string{}
			edit := make(map[string]string)
			detail := make(map[string]string)
			id := int(district.Id)

			lineInfo["id"] = id
			oneLine[0] = district.Name

			oneLine[1] = district.City.Province.Name
			oneLine[2] = district.City.Name

			edit["name"] = "编辑"
			edit["url"] = URL + "/edit/" + strconv.Itoa(id)
			detail["name"] = "详情"
			detail["url"] = URL + "/show/" + strconv.Itoa(id)
			action["edit"] = edit
			action["detail"] = detail

			oneLine[3] = action
			lineData := make(map[string]interface{})
			lineData["oneLine"] = oneLine
			lineData["lineInfo"] = lineInfo
			bodyLines = append(bodyLines, lineData)
		}
		tableBody["bodyLines"] = bodyLines
		tableInfo.Body = tableBody
		tableInfo.TitleLen = districtListCellLength
		tableInfo.TitleIndexLen = districtListCellLength - 1
		tableInfo.BodyLen = paginator.CurrentPageSize
		this.Data["tableInfo"] = tableInfo
	}
}
