package controllers

import (
	. "pms/models/base"
	"pms/utils"
	"time"

	"github.com/jinzhu/gorm"
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
	this.Layout = "base/base.html"
	this.TplName = "base/base.html"
	LoginName := this.GetString("LoginName")
	Password := this.GetString("PassWord")
	Remember := this.GetString("Remember")

	if LoginName == "" {
		return
	}
	if Password == "" {
		return
	}
	var (
		db *gorm.DB
		// err  error
		user User
	)
	//创建数据库的连接
	// if db, err = GormDbConnect(); err != nil {
	// 	return
	// }
	defer db.Close()
	if db.Where("name = ?", LoginName).Or("mobile = ?", LoginName).Or("email = ?", LoginName).First(&user).Error != nil {
		return
	}
	//判断密码是否正确，若正确设置session
	if utils.PasswordMD5(Password, user.Email) != user.Password {
		return
	} else {
		this.SetSession("UserId", user.Id)
		this.SetSession("UserName", user.Name)
		this.SetSession("LastLogin", user.LastLogin)
		this.SetSession("IsAdmin", user.IsAdmin)
		user.LastLogin = time.Now()
		db.Save(&user)
		if Remember != "" {
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
