package base

import (
	"html/template"
	. "pms/init"
	"pms/models/base"
	. "pms/utils"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/beego/i18n"
)

var (
	AppVer string
	IsPro  bool
)

type BaseController struct {
	beego.Controller
	IsAdmin   bool
	UserName  string
	LastLogin time.Time
	User      base.User
	i18n.Locale
}

// Prepare implemented Prepare method for baseRouter.
func (this *BaseController) Prepare() {
	// Setting properties.
	this.StartSession()
	this.Data["AppVer"] = AppVer
	this.Data["IsPro"] = IsPro
	this.Data["xsrf"] = template.HTML(this.XSRFFormHTML())
	this.Data["PageStartTime"] = time.Now()
	// Redirect to make URL clean.
	if this.setLangVer() {
		i := strings.Index(this.Ctx.Request.RequestURI, "?")
		this.Redirect(this.Ctx.Request.RequestURI[:i], 302)
		return
	}

	user := this.GetSession("User")
	if user != nil {
		this.User = user.(base.User)
		this.Data["user"] = user
		this.Data["LastLogin"] = this.GetSession("LastLogin")
	} else {
		if this.Ctx.Request.RequestURI != "/login/in" {
			this.Redirect("/login/in", 302)
		}

		this.Data["LastLogin"] = this.GetSession("LastLogin")
		this.Data["LastIp"] = this.GetSession("LastIp")
	}

}

// setLangVer sets site language version.
func (this *BaseController) setLangVer() bool {
	isNeedRedir := false
	hasCookie := false

	// 1. Check URL arguments.
	lang := this.Input().Get("lang")

	// 2. Get language information from cookies.
	if len(lang) == 0 {
		lang = this.Ctx.GetCookie("lang")
		hasCookie = true
	} else {
		isNeedRedir = true
	}

	// Check again in case someone modify by purpose.
	if !i18n.IsExist(lang) {
		lang = ""
		isNeedRedir = false
		hasCookie = false
	}

	// 3. Get language information from 'Accept-Language'.
	if len(lang) == 0 {
		al := this.Ctx.Request.Header.Get("Accept-Language")
		if len(al) > 4 {
			al = al[:5] // Only compare first 5 letters.
			if i18n.IsExist(al) {
				lang = al
			}
		}
	}

	// 4. Default language is English.
	if len(lang) == 0 {
		lang = "en-US"
		isNeedRedir = false
	}

	curLang := LangType{
		Lang: lang,
	}

	// Save language information in cookies.
	if !hasCookie {
		this.Ctx.SetCookie("lang", curLang.Lang, 1<<31-1, "/")
	}

	restLangs := make([]*LangType, 0, len(LangTypes)-1)
	for _, v := range LangTypes {
		if lang != v.Lang {
			restLangs = append(restLangs, v)
		} else {
			curLang.Name = v.Name
		}
	}

	// Set language properties.
	this.Lang = lang
	this.Data["Lang"] = curLang.Lang
	this.Data["CurLang"] = curLang.Name
	this.Data["RestLangs"] = restLangs

	return isNeedRedir
}
