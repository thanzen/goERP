//访问用户登录日志信息
package controllers

import (
	"pms/models/base"
	"pms/utils"
	"strconv"
)

//列表视图列数-1，第一列为checkbox

const (
	recordListCellLength = 7
)

type RecordController struct {
	BaseController
}

func (this *RecordController) Get() {
	action := this.GetString(":action")
	switch action {
	case "list":
		this.List()
	default:
		this.List()

	}
}
func (this *RecordController) List() {

	this.Data["listName"] = "登录日志"
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
	var records []base.Record
	paginator, err, records := base.ListRecord(condArr, this.User.Id, pageInt64, offsetInt64)
	paginator.Url = "/record"
	this.Data["Paginator"] = paginator
	tableInfo := new(utils.TableInfo)
	tableInfo.Url = "/record"
	tableTitle := make(map[string]interface{})
	tableTitle["titleName"] = [recordListCellLength]string{"用户名", "中文用户名", "开始时间", "结束时间", "持续时间", "登录IP", "操作"}
	tableInfo.Title = tableTitle
	tableBody := make(map[string]interface{})
	bodyLines := make([]interface{}, 0, 20)
	if err == nil {
		for _, record := range records {
			oneLine := make([]interface{}, recordListCellLength, recordListCellLength)
			lineInfo := make(map[string]interface{})
			action := map[string]map[string]string{}
			// edit := make(map[string]string)
			// remove := make(map[string]string)
			// disable := make(map[string]string)
			// detail := make(map[string]string)
			id := int(record.Id)

			lineInfo["id"] = id
			oneLine[0] = record.User.Name
			oneLine[1] = record.User.NameZh

			oneLine[2] = record.CreateDate
			oneLine[3] = record.UpdateDate

			oneLine[4] = record.UpdateDate.Sub(record.CreateDate)
			oneLine[5] = record.Ip

			// edit["name"] = "编辑"
			// edit["url"] = tableInfo.Url + "/edit/" + strconv.Itoa(id)
			// remove["name"] = "删除"
			// remove["url"] = tableInfo.Url + "/remove/" + strconv.Itoa(id)
			// detail["name"] = "详情"
			// detail["url"] = tableInfo.Url + "/detail/" + strconv.Itoa(id)
			// disable["name"] = "无效"
			// disable["url"] = tableInfo.Url + "/disable/" + strconv.Itoa(id)
			// action["edit"] = edit
			// action["remove"] = remove
			// action["detail"] = detail
			// action["disable"] = disable
			oneLine[6] = action
			lineData := make(map[string]interface{})
			lineData["oneLine"] = oneLine
			lineData["lineInfo"] = lineInfo
			bodyLines = append(bodyLines, lineData)
		}
		tableBody["bodyLines"] = bodyLines
		tableInfo.Body = tableBody
		tableInfo.TitleLen = recordListCellLength
		tableInfo.TitleIndexLen = recordListCellLength - 1
		tableInfo.BodyLen = paginator.CurrentPageSize
		this.Data["tableInfo"] = tableInfo
	}
}
