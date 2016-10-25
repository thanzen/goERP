package controllers

import (
	"fmt"
	"pms/models/base"
	"strconv"
)

type UserController struct {
	BaseController
}

func (this *UserController) Get() {
	action := this.GetString(":action")
	switch action {
	case "list":
		this.List()
	default:
		this.List()

	}
}
func (this *UserController) List() {
	this.Layout = "base/base.html"
	this.TplName = "user/user_list.html"
	condArr := make(map[string]interface{})
	condArr["active"] = true
	page := this.Input().Get("page")
	offset := this.Input().Get("offset")
	var err error
	pageInt, _ := strconv.Atoi(page)
	offsetInt, _ := strconv.Atoi(offset)
	var users []base.User
	num, err, users := base.ListUser(condArr, this.User, pageInt, offsetInt)
	fmt.Println(num)
	fmt.Println(err)
	fmt.Println(users)

}
