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

func (this *UserController) Put() {

}
func (this *UserController) Get() {
	action := this.Input().Get("action")
	switch action {
	case "create":
		this.Create()
	default:
		this.GetList()

	}

	this.URL = "/user"
	this.Data["URL"] = this.URL
	this.Layout = "base/base.html"
	this.Data["MenuUserActive"] = "active"

}
func (this *UserController) Post() {
	action := this.Input().Get("action")
	switch action {
	case "validator":
		this.Validator()
	case "table":
		this.PostList()
	default:
		this.PostList()
	}
}
func (this *UserController) GetList() {
	this.Data["tableId"] = "table-user"
	this.TplName = "base/table_base.html"
}
func (this *UserController) Validator() {
	username := this.GetString("username")
	username = strings.TrimSpace(username)
	result := make(map[string]bool)
	if _, err := mb.GetUserByName(username); err != nil {
		result["valid"] = true
	} else {
		result["valid"] = false
	}
	this.Data["json"] = result
	this.ServeJSON()
}
func (this *UserController) PostList() {
	condArr := make(map[string]interface{})
	start := this.Input().Get("offset")
	length := this.Input().Get("limit")
	fmt.Println(start)
	fmt.Println(length)

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
	if result, err := this.userList(startInt64, lengthInt64, condArr); err == nil {
		this.Data["json"] = result
	}
	this.ServeJSON()

}
func (this *UserController) userList(start, length int64, condArr map[string]interface{}) (map[string]interface{}, error) {

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

			oneLine["email"] = user.Email
			oneLine["mobile"] = user.Mobile
			oneLine["tel"] = user.Tel
			if user.IsAdmin {
				oneLine["isadmin"] = "是"
			} else {
				oneLine["isadmin"] = "否"
			}
			if user.Active {
				oneLine["active"] = "有效"
			} else {
				oneLine["active"] = "无效"
			}
			oneLine["qq"] = user.Qq
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

func (this *UserController) ChangePwd() {
	this.Data["MenuChangePwdActive"] = "active"
	this.TplName = "user/user_change_password_form.html"
}

func (this *UserController) Create() {

	method := strings.ToUpper(this.Ctx.Request.Method)
	if method == "GET" {
		this.Data["Readonly"] = false
		this.Data["listName"] = "创建用户"
		this.TplName = "user/user_form.html"

	} else if method == "POST" {
		fmt.Println("enter create post")
		user := new(mb.User)
		if err := this.ParseForm(user); err == nil {

			department := this.Input().Get("department")
			fmt.Println("======================")
			fmt.Println(department)
			if id, err := mb.AddUser(user, this.User); err == nil {
				this.Redirect("/user/"+strconv.FormatInt(id, 10), 302)
			}
		} else {
			fmt.Print("%T", err)
		}

	}

}
func (this *UserController) Edit() {
	id, _ := this.GetInt64(":id")
	user, _ := mb.GetUserByID(id)
	fmt.Println(user)
	this.TplName = "user/user_form.html"
}
func (this *UserController) Show() {
	this.Data["MenuSelfInfoActive"] = "active"
	id, _ := this.GetInt64(":id")
	fmt.Print(id)
	this.TplName = "user/user_form.html"
}
