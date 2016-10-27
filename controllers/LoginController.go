package controllers

import "pms/models/base"

type LoginController struct {
	BaseController
}

func (this *LoginController) Get() {
	action := this.GetString(":action")
	if action == "out" {
		this.Logout()
		this.Redirect("/login/in", 302)
	} else if action == "in" {
		user := this.GetSession("User")
		if user != nil {
			this.Redirect("/", 302)
		}
		this.TplName = "login.html"
	}

}
func (this *LoginController) Post() {

	this.Layout = "base/base.html"
	this.TplName = "test.html"
	loginName := this.GetString("loginName")
	password := this.GetString("password")
	rememberMe := this.GetString("remember")

	if loginName == "" && password == "" {
		this.Redirect("/login/in", 302)
	}

	var (
		user  base.User
		err   error
		login base.LoginLog
		ok    bool
	)
	if user, err, ok = base.CheckUserByName(loginName, password); ok != true {
		this.Redirect("/login/in", 302)
	} else {

		this.Data["user"] = user
		if login, err = base.GetLoginLog(user); err == nil {
			this.Data["LastLogin"] = login.CreateDate
			this.Data["LastIp"] = login.Ip
			this.SetSession("LastLogin", login.CreateDate)
			this.SetSession("LastIp", login.Ip)
		}
		base.AddLoginLog(user, this.Ctx.Input.IP())
		this.SetSession("User", user)

		this.Ctx.SetCookie("Remember", rememberMe, 31536000, "/")
		//通过验证跳转到主界面
		this.Redirect("/", 302)
	}
}

//登出
func (this *LoginController) Logout() {
	base.UpdateLoginLog(this.User.Id)
	this.SetSession("User", nil)
	this.DelSession("User")

}
