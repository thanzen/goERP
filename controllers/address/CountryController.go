package address

import (
	"pms/controllers/base"
	. "pms/models/base"

	"pms/utils"
	"strconv"
)

const (
	countryListCellLength = 2
)

type CountryController struct {
	base.BaseController
}

func (this *CountryController) Get() {
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
	this.Data["searchKeyWords"] = "国家"
}
func (this *CountryController) List() {
	this.Data["listName"] = "国家信息"
	this.Layout = "base/base.html"
	this.TplName = "user/record_list.html"
	this.Data["settingRootActive"] = "active"
	this.Data["addressManageActive"] = "active"
	this.Data["addressCountryActive"] = "active"
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
	var countrys []Country
	paginator, err, countrys := ListCountry(condArr, pageInt64, offsetInt64)
	URL := "/country"
	this.Data["URL"] = URL
	this.Data["Paginator"] = paginator
	tableInfo := new(utils.TableInfo)

	tableTitle := make(map[string]interface{})
	tableTitle["titleName"] = [countryListCellLength]string{"国家", "操作"}
	tableInfo.Title = tableTitle
	tableBody := make(map[string]interface{})
	bodyLines := make([]interface{}, 0, base.ListNum)
	if err == nil {
		for _, country := range countrys {
			oneLine := make([]interface{}, countryListCellLength, countryListCellLength)
			lineInfo := make(map[string]interface{})
			action := map[string]map[string]string{}
			edit := make(map[string]string)
			detail := make(map[string]string)
			id := int(country.Id)

			lineInfo["id"] = id
			oneLine[0] = country.Name

			edit["name"] = "编辑"
			edit["url"] = URL + "/edit/" + strconv.Itoa(id)
			detail["name"] = "详情"
			detail["url"] = URL + "/detail/" + strconv.Itoa(id)
			action["edit"] = edit
			action["detail"] = detail

			oneLine[1] = action
			lineData := make(map[string]interface{})
			lineData["oneLine"] = oneLine
			lineData["lineInfo"] = lineInfo
			bodyLines = append(bodyLines, lineData)
		}
		tableBody["bodyLines"] = bodyLines
		tableInfo.Body = tableBody
		tableInfo.TitleLen = countryListCellLength
		tableInfo.TitleIndexLen = countryListCellLength - 1
		tableInfo.BodyLen = paginator.CurrentPageSize
		this.Data["tableInfo"] = tableInfo
	}
}
