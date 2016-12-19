package base

import "pms/models/base"

type LoginController struct {
	BaseController
}

func (ctl *LoginController) Get() {
	action := ctl.GetString(":action")
	if action == "out" {
		ctl.Logout()
		ctl.Redirect("/login/in", 302)
	} else if action == "in" {
		user := ctl.GetSession("User")
		if user != nil {
			ctl.Redirect("/", 302)
		}
		ctl.TplName = "login.html"
	}

}
func (ctl *LoginController) Post() {

	loginName := ctl.GetString("loginName")
	password := ctl.GetString("password")
	rememberMe := ctl.GetString("remember")

	if loginName == "" && password == "" {
		ctl.Redirect("/login/in", 302)
	}

	var (
		user   base.User
		err    error
		record base.Record
		ok     bool
	)
	if user, err, ok = base.CheckUserByName(loginName, password); ok != true {
		ctl.Redirect("/login/in", 302)
	} else {
		if record, err = base.GetLastRecordByUserID(user.Id); err == nil {

			ctl.SetSession("LastLogin", record.CreateDate)
			ctl.SetSession("LastIp", record.Ip)
		}
		base.CreateRecord(user, ctl.Ctx.Input.IP(), ctl.Ctx.Request.UserAgent())
		ctl.SetSession("User", user)

		ctl.Ctx.SetCookie("Remember", rememberMe, 31536000, "/")
		//通过验证跳转到主界面
		ctl.Redirect("/", 302)
	}
}

//登出
func (ctl *LoginController) Logout() {
	base.UpdateRecord(ctl.User.Id, ctl.Ctx.Input.IP())
	ctl.SetSession("User", nil)
	ctl.DelSession("User")

}
