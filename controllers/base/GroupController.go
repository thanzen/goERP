package base

import (
	"encoding/json"
	mb "pms/models/base"
	"strconv"
	"strings"
)

type GroupController struct {
	BaseController
}

func (ctl *GroupController) Post() {
	action := ctl.Input().Get("action")
	ctl.URL = "/group/"
	switch action {
	case "validator":
		ctl.Validator()
	case "table": //bootstrap table的post请求
		ctl.PostList()
	case "create":
		ctl.PostCreate()
	default:
		ctl.PostList()
	}
}
func (ctl *GroupController) Get() {
	action := ctl.Input().Get("action")
	switch action {
	case "create":
		ctl.Create()
	case "edit":
		ctl.Edit()
	case "detail":
		ctl.Detail()
	default:
		ctl.GetList()
	}
	ctl.URL = "/group/"
	ctl.Data["URL"] = ctl.URL
	ctl.Layout = "base/base.html"
	ctl.Data["MenuGroupActive"] = "active"
}
func (ctl *GroupController) PostCreate() {
	group := new(mb.Group)
	if err := ctl.ParseForm(group); err == nil {

		if id, err := mb.CreateGroup(group, ctl.User); err == nil {
			ctl.Redirect(ctl.URL+strconv.FormatInt(id, 10)+"?action=detail", 302)
		} else {
			ctl.Get()
		}
	} else {
		ctl.Get()
	}
}
func (ctl *GroupController) GetList() {
	ctl.Data["tableId"] = "table-group"
	ctl.TplName = "base/table_base.html"
}

func (ctl *GroupController) Detail() {
	//获取信息一样，直接调用Edit
	ctl.Edit()
	ctl.Data["Readonly"] = true
	ctl.Data["Action"] = "detail"
}
func (ctl *GroupController) Create() {
	ctl.Data["Action"] = "create"
	ctl.Data["Readonly"] = false
	ctl.Data["listName"] = "创建权限"
	ctl.TplName = "user/group_form.html"
}
func (ctl *GroupController) Edit() {
	id := ctl.Ctx.Input.Param(":id")
	groupInfo := make(map[string]interface{})
	if id != "" {
		if idInt64, e := strconv.ParseInt(id, 10, 64); e == nil {

			if group, err := mb.GetGroupByID(idInt64); err == nil {

				groupInfo["name"] = group.Name

			}
		}
	}
	ctl.Data["Action"] = "edit"
	ctl.Data["RecordId"] = id
	ctl.Data["Group"] = groupInfo
	ctl.TplName = "user/group_form.html"
}
func (ctl *GroupController) Validator() {
	name := ctl.GetString("name")
	name = strings.TrimSpace(name)
	result := make(map[string]bool)
	if _, err := mb.GetGroupByName(name); err != nil {
		result["valid"] = true
	} else {
		result["valid"] = false
	}
	ctl.Data["json"] = result
	ctl.ServeJSON()
}
func (ctl *GroupController) PostList() {
	condArr := make(map[string]interface{})
	start := ctl.Input().Get("offset")
	length := ctl.Input().Get("limit")
	name := ctl.Input().Get("name")
	name = strings.TrimSpace(name)
	if name != "" {
		condArr["name"] = name
	}
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
	if result, err := ctl.groupList(startInt64, lengthInt64, condArr); err == nil {
		ctl.Data["json"] = result
	}
	ctl.ServeJSON()

}
func (ctl *GroupController) groupList(start, length int64, condArr map[string]interface{}) (map[string]interface{}, error) {

	var groups []mb.Group
	paginator, groups, err := mb.ListGroup(condArr, start, length)

	result := make(map[string]interface{})
	if err == nil {

		// result["recordsFiltered"] = paginator.TotalCount
		tableLines := make([]interface{}, 0, 4)
		for _, group := range groups {
			oneLine := make(map[string]interface{})
			oneLine["name"] = group.Name

			oneLine["Id"] = group.Id
			oneLine["id"] = group.Id
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
