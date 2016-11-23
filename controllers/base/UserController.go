package base

import (
	"fmt"
	"pms/models/base"
	"pms/utils"
	"strconv"
	"strings"
)

//列表视图列数-1，第一列为checkbox
const (
	userListCellLength = 11
)

type UserController struct {
	BaseController
}

func (this *UserController) Post() {
	action := this.GetString(":action")
	switch action {
	case "create":
		this.Create()
	case "search":
		this.Search()
	default:
		this.List()
	}
}
func (this *UserController) Get() {

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
	case "show":
		this.Show()
	case "create":
		this.Create()
	case "edit":
		this.Edit()
	case "changepasswd":
		this.ChangePwd()
	default:
		this.List()
	}
	this.Data["searchKeyWords"] = "邮箱/手机号码"
	this.URL = "/user"
	this.Data["URL"] = this.URL
	this.Layout = "base/base.html"
	this.Data["settingRootActive"] = "active"

}
func (this *UserController) ChangePwd() {
	this.Data["personInfoActive"] = "active"
	this.Data["settingPassword"] = "active"
	this.TplName = "user/user_change_password_form.html"
}
func (this *UserController) Search() {
	username := this.GetString("username")
	username = strings.TrimSpace(username)
	result := make(map[string]bool)
	if _, err := base.GetUserByName(username); err != nil {
		result["valid"] = true
	} else {
		result["valid"] = false
	}
	this.Data["json"] = result
	this.ServeJSON()
}
func (this *UserController) Create() {
	method := strings.ToUpper(this.Ctx.Request.Method)
	if method == "GET" {
		this.Data["Readonly"] = false
		this.Data["listName"] = "创建用户"
		this.TplName = "user/user_form.html"

	} else if method == "POST" {
		name := this.GetString("name")
		mobile := this.GetString("mobile")
		fmt.Print(name, mobile)

	}

}
func (this *UserController) Edit() {
	id, _ := this.GetInt64(":id")
	user, _ := base.GetUserByID(id)
	fmt.Println(user)
	this.TplName = "user/user_form.html"
}
func (this *UserController) Show() {
	id, _ := this.GetInt64(":id")
	fmt.Print(id)
	this.TplName = "user/user_form.html"
}
func (this *UserController) List() {
	this.Data["listName"] = "用户信息"
	this.Data["userListActive"] = "active"
	this.TplName = "user/user_list.html"
	this.URL = "/user"
	this.Data["URL"] = this.URL
	condArr := make(map[string]interface{})
	condArr["active"] = true
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

	paginator, users, err := base.ListUser(condArr, this.User, pageInt64, offsetInt64)
	this.Data["Paginator"] = paginator
	tableInfo := new(utils.TableInfo)

	tableTitle := make(map[string]interface{})
	tableTitle["titleName"] = [userListCellLength]string{"用户名", "中文用户名", "部门", "邮箱", "手机号码", "固定号码", "超级用户", "有效", "QQ", "微信", "操作"}
	tableInfo.Title = tableTitle
	tableBody := make(map[string]interface{})
	bodyLines := make([]interface{}, 0, ListNum)
	if err == nil {
		for _, user := range users {
			oneLine := make([]interface{}, userListCellLength, userListCellLength)
			lineInfo := make(map[string]interface{})
			action := map[string]map[string]string{}
			edit := make(map[string]string)
			delete := make(map[string]string)
			disable := make(map[string]string)
			detail := make(map[string]string)
			id := int(user.Id)

			lineInfo["id"] = id
			oneLine[0] = user.Name
			oneLine[1] = user.NameZh
			if user.Department != nil {
				oneLine[2] = user.Department.Name
			} else {
				oneLine[2] = "-"
			}

			oneLine[3] = user.Email
			oneLine[4] = user.Mobile
			oneLine[5] = user.Tel
			if user.IsAdmin {
				oneLine[6] = "是"
			} else {
				oneLine[6] = "否"
			}
			if user.Active {
				oneLine[7] = "有效"
			} else {
				oneLine[7] = "无效"
			}
			oneLine[9] = user.Qq
			oneLine[9] = user.WeChat
			edit["name"] = "编辑"
			edit["url"] = this.URL + "/edit/" + strconv.Itoa(id)
			delete["name"] = "删除"
			delete["url"] = this.URL + "/delete/" + strconv.Itoa(id)
			detail["name"] = "详情"
			detail["url"] = this.URL + "/show/" + strconv.Itoa(id)
			disable["name"] = "无效"
			disable["url"] = this.URL + "/disable/" + strconv.Itoa(id)
			action["edit"] = edit
			action["delete"] = delete
			action["detail"] = detail
			action["disable"] = disable
			oneLine[10] = action
			lineData := make(map[string]interface{})
			lineData["oneLine"] = oneLine
			lineData["lineInfo"] = lineInfo
			bodyLines = append(bodyLines, lineData)
		}
		tableBody["bodyLines"] = bodyLines
		tableInfo.Body = tableBody
		tableInfo.TitleLen = userListCellLength
		tableInfo.TitleIndexLen = userListCellLength - 1
		tableInfo.BodyLen = paginator.CurrentPageSize
		this.Data["tableInfo"] = tableInfo
	}

}
