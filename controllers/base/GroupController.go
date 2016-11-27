package base

import (
	"pms/models/base"
	"strconv"
)

type GroupController struct {
	BaseController
}

func (this *GroupController) Post() {

	this.PostList()
}

func (this *GroupController) PostList() {
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
	var groups []base.Group
	paginator, groups, err := base.ListGroup(condArr, startInt64, lengthInt64)
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
		for _, group := range groups {
			oneLine := make(map[string]interface{})
			oneLine["name"] = group.Name

			oneLine["id"] = group.Id
			tableLines = append(tableLines, oneLine)
		}
		result["data"] = tableLines
	}
	this.Data["json"] = result
	this.ServeJSON()
}
