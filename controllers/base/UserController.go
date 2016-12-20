package base

import (
	"encoding/json"
	mb "pms/models/base"
	"strconv"
	"strings"
)

type UserController struct {
	BaseController
}

func (ctl *UserController) Put() {

}
func (ctl *UserController) Get() {
	if id, err := ctl.GetInt64(":id"); err == nil {
		if user, err := mb.GetUserByID(id); err == nil {
			userMap := make(map[string]interface{})
			userMap["Id"] = user.Id
			userMap["Name"] = user.Name
			userMap["NameZh"] = user.NameZh
			userMap["Email"] = user.Email
			userMap["Mobile"] = user.Mobile
			userMap["Tel"] = user.Tel
			department := make(map[string]string)
			if user.Department != nil {
				department["Id"] = strconv.FormatInt(user.Department.Id, 10)
				department["Name"] = user.Department.Name
				userMap["Department"] = department
			}

			position := make(map[string]string)
			if user.Position != nil {
				position["Id"] = strconv.FormatInt(user.Position.Id, 10)
				position["Name"] = user.Position.Name
				userMap["Position"] = position

			}

			ctl.Data["User"] = userMap

			ctl.Data["Readonly"] = true
			ctl.TplName = "user/user_form.html"
		}
	} else {
		action := ctl.Input().Get("action")
		switch action {
		case "create":
			ctl.GetCreate()
		default:
			ctl.GetList()
		}
	}

	ctl.URL = "/user"
	ctl.Data["URL"] = ctl.URL
	ctl.Layout = "base/base.html"
	ctl.Data["MenuUserActive"] = "active"

}
func (ctl *UserController) Post() {
	action := ctl.Input().Get("action")
	switch action {
	case "validator":
		ctl.Validator()
	case "table":
		ctl.PostList()
	case "create":
		ctl.PostCreate()
	default:
		ctl.PostList()
	}
}
func (ctl *UserController) GetList() {
	ctl.Data["tableId"] = "table-user"
	ctl.TplName = "base/table_base.html"
}
func (ctl *UserController) Validator() {
	username := ctl.GetString("username")
	username = strings.TrimSpace(username)
	result := make(map[string]bool)
	if _, err := mb.GetUserByName(username); err != nil {
		result["valid"] = true
	} else {
		result["valid"] = false
	}
	ctl.Data["json"] = result
	ctl.ServeJSON()
}
func (ctl *UserController) PostList() {
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
	if result, err := ctl.userList(startInt64, lengthInt64, condArr); err == nil {
		ctl.Data["json"] = result
	}
	ctl.ServeJSON()

}
func (ctl *UserController) userList(start, length int64, condArr map[string]interface{}) (map[string]interface{}, error) {

	var users []mb.User
	paginator, users, err := mb.ListUser(condArr, start, length)
	result := make(map[string]interface{})
	if err == nil {

		// result["recordsFiltered"] = paginator.TotalCount
		tableLines := make([]interface{}, 0, 4)
		for _, user := range users {

			oneLine := make(map[string]interface{})
			oneLine["username"] = user.Name
			oneLine["namezh"] = user.NameZh
			if user.Department != nil {
				oneLine["department"] = user.Department.Name
			} else {
				oneLine["department"] = "-"
			}
			if user.Position != nil {
				oneLine["position"] = user.Position.Name
			} else {
				oneLine["position"] = "-"
			}

			oneLine["email"] = user.Email
			oneLine["mobile"] = user.Mobile
			oneLine["tel"] = user.Tel
			oneLine["isadmin"] = user.IsAdmin
			oneLine["active"] = user.Active
			oneLine["qq"] = user.Qq
			oneLine["Id"] = user.Id
			oneLine["id"] = user.Id
			oneLine["wechat"] = user.WeChat

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

func (ctl *UserController) ChangePwd() {
	ctl.Data["MenuChangePwdActive"] = "active"
	ctl.TplName = "user/user_change_password_form.html"
}

func (ctl *UserController) GetCreate() {
	ctl.Data["Readonly"] = false
	ctl.Data["listName"] = "创建用户"
	ctl.TplName = "user/user_form.html"
}
func (ctl *UserController) PostCreate() {

	user := new(mb.User)
	if err := ctl.ParseForm(user); err == nil {

		if deparentId, err := ctl.GetInt64("department"); err == nil {
			if department, err := mb.GetDepartmentByID(deparentId); err == nil {
				user.Department = &department
			}
		}

		if positionId, err := ctl.GetInt64("position"); err == nil {
			if position, err := mb.GetPositionByID(positionId); err == nil {
				user.Position = &position
			}
		}

		if id, err := mb.CreateUser(user, ctl.User); err == nil {
			ctl.Redirect("/user/"+strconv.FormatInt(id, 10), 302)
		}
	}

}
func (ctl *UserController) Edit() {
	id := ctl.Ctx.Input.Param(":id")
	ctl.Data["RecordId"] = id
	ctl.Data["Action"] = "edit"
	ctl.TplName = "user/user_form.html"
}
func (ctl *UserController) Show() {
	ctl.Data["MenuSelfInfoActive"] = "active"

	ctl.TplName = "user/user_form.html"
}
