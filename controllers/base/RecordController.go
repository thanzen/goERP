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

func (this *RecordController) Get() {

	this.GetList()

	this.URL = "/user"
	this.Data["URL"] = this.URL
	this.Layout = "base/base.html"
	this.Data["MenuRecordActive"] = "active"

}
func (this *RecordController) Post() {
	action := this.Input().Get("action")
	switch action {

	case "table":
		this.PostList()
	default:
		this.PostList()
	}
}
func (this *RecordController) PostList() {
	condArr := make(map[string]interface{})
	start := this.Input().Get("offset")
	length := this.Input().Get("limit")

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
	if result, err := this.recordList(startInt64, lengthInt64, condArr); err == nil {
		this.Data["json"] = result
	}
	this.ServeJSON()

}
func (this *RecordController) recordList(start, length int64, condArr map[string]interface{}) (map[string]interface{}, error) {

	var records []mb.Record
	paginator, records, err := mb.ListRecord(condArr, this.User.Id, start, length)
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
			oneLine["id"] = record.Id
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
func (this *RecordController) GetList() {
	this.Data["tableId"] = "table-record"
	this.TplName = "base/table_base.html"
}
