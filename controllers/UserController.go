package controllers

import (
	"fmt"
	"pms/models/base"
	"pms/utils"
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
	var users []base.User
	paginator, err, users := base.ListUser(condArr, this.User, pageInt64, offsetInt64)
	tableLineInfo := new(utils.TableLineInfo)
	tableLineInfo.Url = "/user/list"
	// tableTitle := make(map[string]interface{})
	if err == nil {
		for i, user := range users {
			fmt.Println(i)
			fmt.Println(user)
		}
	}
	this.Data["Paginator"] = paginator

}
