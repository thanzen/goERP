package controllers

import (
	"fmt"
	"pms/models/base"
	"pms/utils"
	"time"
)

type LoginController struct {
	BaseController
}

func (this *LoginController) Get() {
	action := this.GetString(":action")
	if action == "out" {
		this.Logout()
	}
	this.Redirect("/", 302)
}
func (this *LoginController) Post() {
	fmt.Println(1232)
	this.Layout = "base/base.html"
	this.TplName = "base/base.html"
	loginName := this.GetString("loginName")
	password := this.GetString("password")
	rememberMe := this.GetString("remember")

	if loginName == "" {
		this.Redirect("/", 302)
	}
	if password == "" {
		this.Redirect("/", 302)
	}

	var (
		user base.User
		err  error
	)
	if user, err = base.GetUserByName(loginName); err != nil {
		this.Redirect("/", 302)
	}

	//判断密码是否正确，若正确设置session
	if utils.PasswordMD5(password, user.Email) != user.Password {
		return
	} else {
		this.SetSession("UserId", user.Id)
		this.SetSession("UserName", user.Name)
		this.SetSession("LastLogin", user.LastLogin)
		this.SetSession("IsAdmin", user.IsAdmin)
		user.LastLogin = time.Now()
		if rememberMe != "" {
			this.Ctx.SetCookie("Remember", "on", 31536000, "/")
		} else {
			this.Ctx.SetCookie("Remember", "off", 31536000, "/")
		}
	}
	//通过验证跳转到主界面
	this.Redirect("/", 302)
}

//登出
func (this *LoginController) Logout() {

	this.SetSession("UserName", nil)
	this.SetSession("LastLogin", nil)
	this.SetSession("IsAdmin", nil)
	// this.SetSession("Group", nil)

	this.DelSession("UserName")
	this.DelSession("LastLogin")
	this.DelSession("IsAdmin")
	// this.DelSession("Group")
}
