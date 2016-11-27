package base

import (
	"fmt"
	"pms/models/base"
	"strconv"
	"strings"
)

type UserController struct {
	BaseController
}

func (this *UserController) Put() {

}
func (this *UserController) Post() {
	action := this.Input().Get("action")
	switch action {
	case "validator":
		this.Validator()
	case "table":
		this.Table()
	default:
		this.Create()
	}
}
func (this *UserController) Validator() {
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
func (this *UserController) Table() {
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
	var users []base.User
	paginator, users, err := base.ListUser(condArr, this.User, startInt64, lengthInt64)
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
	}
	this.Data["json"] = result
	this.ServeJSON()

}
func (this *UserController) Get() {
	action := this.Input().Get("action")
	switch action {
	case "create":
		this.Create()
	default:
		this.List()

	}

	this.URL = "/user"
	this.Data["URL"] = this.URL
	this.Layout = "base/base.html"
	this.Data["MenuUserActive"] = "active"

}
func (this *UserController) List() {

	this.TplName = "user/table_user.html"
}
func (this *UserController) ChangePwd() {
	this.Data["MenuChangePwdActive"] = "active"
	this.TplName = "user/user_change_password_form.html"
}

func (this *UserController) Create() {
	fmt.Println("enter create")
	method := strings.ToUpper(this.Ctx.Request.Method)
	if method == "GET" {
		this.Data["Readonly"] = false
		this.Data["listName"] = "创建用户"
		this.TplName = "user/user_form.html"

	} else if method == "POST" {
		fmt.Println("enter create post")
		user := new(base.User)
		if err := this.ParseForm(user); err == nil {

			department := this.Input().Get("department")
			fmt.Println("======================")
			fmt.Println(department)
			if id, err := base.AddUser(user, this.User); err == nil {
				this.Redirect("/user/"+strconv.FormatInt(id, 10), 302)
			}
		} else {
			fmt.Print("%T", err)
		}

	}

}
func (this *UserController) Edit() {
	id, _ := this.GetInt64(":id")
	user, _ := base.GetUserByID(id)
	fmt.Println(user)
	this.TplName = "user/user_form.html"
}
func (this *UserController) Show() {
	this.Data["MenuSelfInfoActive"] = "active"
	id, _ := this.GetInt64(":id")
	fmt.Print(id)
	this.TplName = "user/user_form.html"
}
