package address

import (
	"pms/controllers/base"
	mb "pms/models/base"
	"pms/utils"
	"strconv"
)

const (
	provinceListCellLength = 3
)

type ProvinceController struct {
	base.BaseController
}

func (this *ProvinceController) Get() {
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
	this.Data["searchKeyWords"] = "国家/省份"
}
func (this *ProvinceController) List() {
	this.Data["listName"] = "省份信息"
	this.Layout = "base/base.html"
	this.TplName = "user/record_list.html"
	this.Data["settingRootActive"] = "active"
	this.Data["addressManageActive"] = "active"
	this.Data["addressProvinceActive"] = "active"
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
	var provinces []mb.Province
	paginator, provinces, err := mb.ListProvince(condArr, pageInt64, offsetInt64)

	URL := "/province"
	this.Data["URL"] = URL
	this.Data["Paginator"] = paginator
	tableInfo := new(utils.TableInfo)
	tableTitle := make(map[string]interface{})
	tableTitle["titleName"] = [provinceListCellLength]string{"省份", "国家", "操作"}
	tableInfo.Title = tableTitle
	tableBody := make(map[string]interface{})
	bodyLines := make([]interface{}, 0, base.ListNum)
	if err == nil {
		for _, province := range provinces {
			oneLine := make([]interface{}, provinceListCellLength, provinceListCellLength)
			lineInfo := make(map[string]interface{})
			action := map[string]map[string]string{}
			edit := make(map[string]string)
			detail := make(map[string]string)
			id := int(province.Id)

			lineInfo["id"] = id
			oneLine[0] = province.Name
			oneLine[1] = province.Country.Name
			edit["name"] = "编辑"
			edit["url"] = URL + "/edit/" + strconv.Itoa(id)
			detail["name"] = "详情"
			detail["url"] = URL + "/show/" + strconv.Itoa(id)
			action["edit"] = edit
			action["detail"] = detail

			oneLine[2] = action
			lineData := make(map[string]interface{})
			lineData["oneLine"] = oneLine
			lineData["lineInfo"] = lineInfo
			bodyLines = append(bodyLines, lineData)
		}
		tableBody["bodyLines"] = bodyLines
		tableInfo.Body = tableBody
		tableInfo.TitleLen = provinceListCellLength
		tableInfo.TitleIndexLen = provinceListCellLength - 1
		tableInfo.BodyLen = paginator.CurrentPageSize
		this.Data["tableInfo"] = tableInfo
	}
}
