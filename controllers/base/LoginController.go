package base

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

	loginName := this.GetString("loginName")
	password := this.GetString("password")
	rememberMe := this.GetString("remember")

	if loginName == "" && password == "" {
		this.Redirect("/login/in", 302)
	}

	var (
		user   base.User
		err    error
		record *base.Record
		ok     bool
	)
	if user, err, ok = base.CheckUserByName(loginName, password); ok != true {
		this.Redirect("/login/in", 302)
	} else {
		if record, err = base.GetRecord(user); err == nil {
			this.SetSession("LastLogin", record.CreateDate)
			this.SetSession("LastIp", record.Ip)
		}
		base.AddRecord(user, this.Ctx.Input.IP())
		this.SetSession("User", user)

		this.Ctx.SetCookie("Remember", rememberMe, 31536000, "/")
		//通过验证跳转到主界面
		this.Redirect("/", 302)
	}
}

//登出
func (this *LoginController) Logout() {
	base.UpdateRecord(this.User.Id, this.Ctx.Input.IP())
	this.SetSession("User", nil)
	this.DelSession("User")

}
