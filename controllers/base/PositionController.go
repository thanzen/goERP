package base

import (
	"pms/models/base"
	"strings"
)

//列表视图列数-1，第一列为checkbox
const (
	positionListCellLength = 2
)

type PositionController struct {
	BaseController
}

func (this *PositionController) Post() {

	action := this.GetString(":action")
	switch action {
	case "search":
		this.Search()
	}
	this.ServeJSON()
}

func (this *PositionController) Get() {

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
	case "detail":
		this.Detail()
	case "create":
		this.Create()
	case "edit":
		this.Edit()
	default:
		this.List()
	}
	this.Data["searchKeyWords"] = "部门名称"
	this.URL = "/postion"
	this.Data["URL"] = this.URL
	this.Layout = "base/base.html"
}
func (this *PositionController) List() {

}
func (this *PositionController) Detail() {

}
func (this *PositionController) Create() {

}
func (this *PositionController) Edit() {

}
func (this *PositionController) Search() {
	name := this.GetString("name")
	page, _ := this.GetInt64("page")
	offset, _ := this.GetInt64("offset")
	var condArr = make(map[string]interface{})
	name = strings.TrimSpace(name)
	if name != "" {
		condArr["name"] = name
	}
	paginator, positions, err := base.ListPosition(condArr, page, offset)
	data := make(map[string]interface{})
	if err == nil {

		items := make([]interface{}, 0, 5)
		for _, postion := range positions {
			line := make(map[string]interface{})
			line["id"] = postion.Id
			line["name"] = postion.Name
			items = append(items, line)
		}
		data["items"] = items
		data["total"] = paginator.TotalCount
		data["pageSize"] = paginator.PageSize
		data["page"] = paginator.CurrentPage

	} else {
		data["msg"] = "failed"
	}
	this.Data["json"] = data

}
