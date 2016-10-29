package address

import (
	"pms/controllers/base"
	. "pms/models/base"
	"pms/utils"
	"strconv"
)

const (
	cityListCellLength = 4
)

type CityController struct {
	base.BaseController
}

func (this *CityController) Get() {
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
	this.Data["searchKeyWords"] = "国家/省份/城市"
}
func (this *CityController) List() {
	this.Data["listName"] = "城市信息"
	this.Layout = "base/base.html"
	this.TplName = "user/record_list.html"
	this.Data["settingRootActive"] = "active"
	this.Data["addressManageActive"] = "active"
	this.Data["addressCityActive"] = "active"
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
	var citys []City
	paginator, err, citys := ListCity(condArr, pageInt64, offsetInt64)
	paginator.Url = "/city"
	this.Data["Paginator"] = paginator
	tableInfo := new(utils.TableInfo)
	tableInfo.Url = "/city"
	tableTitle := make(map[string]interface{})
	tableTitle["titleName"] = [cityListCellLength]string{"城市", "国家", "省份", "操作"}
	tableInfo.Title = tableTitle
	tableBody := make(map[string]interface{})
	bodyLines := make([]interface{}, 0, 20)
	if err == nil {
		for _, city := range citys {
			oneLine := make([]interface{}, cityListCellLength, cityListCellLength)
			lineInfo := make(map[string]interface{})
			action := map[string]map[string]string{}
			edit := make(map[string]string)
			detail := make(map[string]string)
			id := int(city.Id)

			lineInfo["id"] = id
			oneLine[0] = city.Name
			oneLine[1] = city.Province.Country.Name
			oneLine[2] = city.Province.Name
			edit["name"] = "编辑"
			edit["url"] = tableInfo.Url + "/edit/" + strconv.Itoa(id)
			detail["name"] = "详情"
			detail["url"] = tableInfo.Url + "/detail/" + strconv.Itoa(id)
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
		tableInfo.TitleLen = cityListCellLength
		tableInfo.TitleIndexLen = cityListCellLength - 1
		tableInfo.BodyLen = paginator.CurrentPageSize
		this.Data["tableInfo"] = tableInfo
	}
}
