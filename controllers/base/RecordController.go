//访问用户登录日志信息
package base

import (
	"pms/models/base"
	"strconv"
)

//列表视图列数-1，第一列为checkbox

const (
	recordListCellLength = 8
)

type RecordController struct {
	BaseController
}

func (this *RecordController) Post() {
	action := this.Input().Get("action")
	switch action {

	case "table":
		this.Table()
	default:
		this.Table()
	}

}
func (this *RecordController) Table() {
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
	var records []base.Record
	paginator, records, err := base.ListRecord(condArr, this.User.Id, startInt64, lengthInt64)
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
		tableLines := make([]interface{}, 0, ListNum)
		for _, record := range records {
			oneLine := make(map[string]interface{})
			oneLine["email"] = record.User.Email
			oneLine["mobile"] = record.User.Mobile
			oneLine["username"] = record.User.Name
			oneLine["namezh"] = record.User.NameZh

			oneLine["start_time"] = record.CreateDate.Format("2006-01-02 15:04:05")
			oneLine["end_time"] = record.Logout.Format("2006-01-02 15:04:05")
			oneLine["ip"] = record.Ip
			oneLine["id"] = record.Id
			tableLines = append(tableLines, oneLine)
		}
		result["data"] = tableLines
	}
	this.Data["json"] = result
	this.ServeJSON()

}
func (this *RecordController) Get() {

	this.List()

	this.URL = "/record"
	this.Data["URL"] = this.URL
	this.Layout = "base/base.html"

	this.Data["MenuRecordActive"] = "active"

}
func (this *RecordController) List() {
	this.TplName = "user/table_record.html"

}
