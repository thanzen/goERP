//访问用户登录日志信息
package base

import (
	"encoding/json"
	mb "pms/models/base"
	"strconv"
)

//列表视图列数-1，第一列为checkbox

type RecordController struct {
	BaseController
}

func (ctl *RecordController) Get() {

	ctl.URL = "/record/"
	ctl.Data["URL"] = ctl.URL
	ctl.Data["MenuRecordActive"] = "active"
	ctl.GetList()

}
func (ctl *RecordController) Post() {
	action := ctl.Input().Get("action")
	switch action {
	case "table":
		ctl.PostList()
	case "one":
		ctl.GetOneRecord()
	default:
		ctl.PostList()
	}
}
func (ctl *RecordController) GetOneRecord() {

}
func (ctl *RecordController) PostList() {
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
	if result, err := ctl.recordList(startInt64, lengthInt64, condArr); err == nil {
		ctl.Data["json"] = result
	}
	ctl.ServeJSON()

}
func (ctl *RecordController) recordList(start, length int64, condArr map[string]interface{}) (map[string]interface{}, error) {

	var records []mb.Record
	paginator, records, err := mb.ListRecord(condArr, ctl.User.Id, start, length)
	result := make(map[string]interface{})
	if err == nil {

		tableLines := make([]interface{}, 0, 4)
		for _, record := range records {
			oneLine := make(map[string]interface{})
			oneLine["email"] = record.User.Email
			oneLine["mobile"] = record.User.Mobile
			oneLine["username"] = record.User.Name
			oneLine["namezh"] = record.User.NameZh
			oneLine["UserAgent"] = record.UserAgent
			oneLine["start_time"] = record.CreateDate.Format("2006-01-02 15:04:05")
			oneLine["end_time"] = record.Logout.Format("2006-01-02 15:04:05")
			oneLine["ip"] = record.Ip
			oneLine["Id"] = record.Id
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
func (ctl *RecordController) GetList() {
	viewType := ctl.Input().Get("view")
	if viewType == "" || viewType == "table" {
		ctl.Data["ViewType"] = "table"
	}
	ctl.Data["tableId"] = "table-record"
	ctl.Layout = "base/base_list_view.html"
	ctl.TplName = "user/record_list_search.html"
}
