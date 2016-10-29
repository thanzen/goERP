package address

import (
	"pms/controllers/base"
	"pms/models/address"
	"pms/utils"
	"strconv"
)

const (
	districtListCellLength = 5
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
	var districts []address.District
	paginator, err, districts := address.ListDistrict(condArr, pageInt64, offsetInt64)
	paginator.Url = "/district"
	this.Data["Paginator"] = paginator
	tableInfo := new(utils.TableInfo)
	tableInfo.Url = "/district"
	tableTitle := make(map[string]interface{})
	tableTitle["titleName"] = [districtListCellLength]string{"区县", "国家", "省份", "城市", "操作"}
	tableInfo.Title = tableTitle
	tableBody := make(map[string]interface{})
	bodyLines := make([]interface{}, 0, 20)
	if err == nil {
		for _, district := range districts {
			oneLine := make([]interface{}, districtListCellLength, districtListCellLength)
			lineInfo := make(map[string]interface{})
			action := map[string]map[string]string{}
			edit := make(map[string]string)
			remove := make(map[string]string)
			detail := make(map[string]string)
			id := int(district.Id)

			lineInfo["id"] = id
			oneLine[0] = district.Name
			oneLine[1] = district.City.Province.Country.Name
			oneLine[2] = district.City.Province.Name
			oneLine[3] = district.City.Name
			edit["name"] = "编辑"
			edit["url"] = tableInfo.Url + "/edit/" + strconv.Itoa(id)
			remove["name"] = "删除"
			remove["url"] = tableInfo.Url + "/remove/" + strconv.Itoa(id)
			detail["name"] = "详情"
			detail["url"] = tableInfo.Url + "/detail/" + strconv.Itoa(id)
			action["edit"] = edit
			action["remove"] = remove
			action["detail"] = detail

			oneLine[4] = action
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
