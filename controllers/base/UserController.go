package base

import (
	"encoding/json"
	"fmt"
	mb "pms/models/base"
	"strconv"
	"strings"
)

type UserController struct {
	BaseController
}

func (ctl *UserController) Put() {
	id := ctl.Ctx.Input.Param(":id")
	ctl.URL = "/user/"
	if idInt64, e := strconv.ParseInt(id, 10, 64); e == nil {
		if user, err := mb.GetUserByID(idInt64); err == nil {
			if err := ctl.ParseForm(&user); err == nil {
				var upateField []string
				if departmentId, err := ctl.GetInt64("department"); err == nil {
					if department, err := mb.GetDepartmentByID(departmentId); err == nil {
						user.Department = &department
						upateField = append(upateField, "Department")
					}
				}
				groupIds := ctl.GetStrings("group")
				if len(groupIds) > 0 {
					var groups []*mb.Group
					var err error
					for groupId := range groupIds {
						var group mb.Group
						if group, err = mb.GetGroupByID(int64(groupId)); err == nil {
							groups = append(groups, &group)
						}
					}
					fmt.Println(groups[0])
					fmt.Println(groups[1])

					user.Groups = groups
					upateField = append(upateField, "Groups")
				}
				if positionId, err := ctl.GetInt64("position"); err == nil {
					if position, err := mb.GetPositionByID(positionId); err == nil {
						user.Position = &position
						upateField = append(upateField, "Position")
					}
				}

				if _, err := mb.UpdateUser(&user, ctl.User, upateField); err == nil {
					ctl.Redirect(ctl.URL+id+"?action=detail", 302)
				}
			}
		}
	}
	ctl.Redirect(ctl.URL+id+"?action=edit", 302)
}
func (ctl *UserController) Get() {
	ctl.URL = "/user/"
	ctl.Data["URL"] = ctl.URL

	action := ctl.Input().Get("action")
	switch action {
	case "create":
		ctl.Create()
	case "edit":
		ctl.Edit()
	case "detail":
		ctl.Detail()
	case "changepasswd":
		ctl.ChangePwd()
	default:
		ctl.GetList()
	}

}
func (ctl *UserController) Post() {
	action := ctl.Input().Get("action")
	ctl.URL = "/user/"
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
func (ctl *UserController) Create() {
	ctl.Data["Action"] = "create"
	ctl.Data["Readonly"] = false
	ctl.Data["listName"] = "创建用户"
	ctl.Layout = "base/base.html"
	ctl.TplName = "user/user_form.html"
}
func (ctl *UserController) Detail() {
	//获取信息一样，直接调用Edit
	ctl.Edit()
	ctl.Data["Readonly"] = true
	ctl.Data["MenuSelfInfoActive"] = "active"
	ctl.Data["Action"] = "detail"
}
func (ctl *UserController) GetList() {
	viewType := ctl.Input().Get("view")
	if viewType == "" || viewType == "table" {
		ctl.Data["ViewType"] = "table"
	}
	ctl.Data["listName"] = "用户管理"
	ctl.Data["tableId"] = "table-user"
	ctl.Data["MenuUserActive"] = "active"
	ctl.Layout = "base/base_list_view.html"
	ctl.TplName = "user/user_list_search.html"
}
func (ctl *UserController) Validator() {
	name := ctl.GetString("name")
	recordId := ctl.GetString("recordId")
	name = strings.TrimSpace(name)
	result := make(map[string]bool)
	obj, err := mb.GetUserByName(name)
	if err != nil {
		result["valid"] = true
	} else {
		if obj.Name == name {
			if recordId != "" {
				result["valid"] = true
			} else {
				result["valid"] = false
			}

		} else {
			result["valid"] = true
		}

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
	filter := ctl.GetString("filter")
	sortName := ctl.GetString("sort")
	orderSymbol := ctl.GetString("order")
	if orderSymbol == "desc" {
		orderSymbol = "-"
	} else {
		orderSymbol = ""
	}
	orderBy := make(map[string]string)
	if sortName != "" {
		orderBy[sortName] = orderSymbol
	} else {
		orderBy["id"] = "-"
	}
	if filter != "" {
		json.Unmarshal([]byte(filter), &condArr)
	}
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
	if result, err := ctl.userList(startInt64, lengthInt64, condArr, orderBy); err == nil {
		ctl.Data["json"] = result
	}
	ctl.ServeJSON()

}
func (ctl *UserController) userList(start, length int64, condArr map[string]interface{}, orderBy map[string]string) (map[string]interface{}, error) {

	var users []mb.User
	paginator, users, err := mb.ListUser(condArr, orderBy, start, length)
	result := make(map[string]interface{})
	if err == nil {

		// result["recordsFiltered"] = paginator.TotalCount
		tableLines := make([]interface{}, 0, 4)
		for _, user := range users {

			oneLine := make(map[string]interface{})
			oneLine["name"] = user.Name
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
	ctl.Layout = "base/base.html"
	ctl.TplName = "user/user_change_password_form.html"
}

func (ctl *UserController) GetCreate() {
	ctl.Data["Readonly"] = false
	ctl.Data["listName"] = "创建用户"
	ctl.Layout = "base/base.html"
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
	userInfo := make(map[string]interface{})
	if id != "" {
		if idInt64, e := strconv.ParseInt(id, 10, 64); e == nil {
			if user, err := mb.GetUserByID(idInt64); err == nil {
				userInfo["Id"] = user.Id
				userInfo["Name"] = user.Name
				userInfo["NameZh"] = user.NameZh
				userInfo["Email"] = user.Email
				userInfo["Mobile"] = user.Mobile
				userInfo["Tel"] = user.Tel
				department := make(map[string]string)
				if user.Department != nil {
					department["Id"] = strconv.FormatInt(user.Department.Id, 10)
					department["Name"] = user.Department.Name
					userInfo["Department"] = department
				}

				position := make(map[string]string)
				if user.Position != nil {
					position["Id"] = strconv.FormatInt(user.Position.Id, 10)
					position["Name"] = user.Position.Name
					userInfo["Position"] = position
				}
			}
		}
	}
	ctl.Data["RecordId"] = id
	ctl.Data["Action"] = "edit"
	ctl.Data["User"] = userInfo
	ctl.Layout = "base/base.html"
	ctl.TplName = "user/user_form.html"
}
func (ctl *UserController) Show() {
	ctl.Data["MenuSelfInfoActive"] = "active"
	ctl.Layout = "base/base.html"
	ctl.TplName = "user/user_form.html"
}
