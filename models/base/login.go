package base

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type LoginLog struct {
	Base
	User   *User     `orm:"rel(fk)"`
	Logout time.Time `orm:"type(datetime);null"` //登出时间
	Ip     string    //上次登录IP
}

//添加记录
func AddLoginLog(user User, IP string) (int64, error) {
	o := orm.NewOrm()
	o.Using("default")
	login := new(LoginLog)
	login.User = &user
	login.CreateUser = &user
	login.Ip = IP
	id, err := o.Insert(login)
	return id, err
}

//获得某一个用户记录信息
func GetLoginLog(user User) (LoginLog, error) {
	o := orm.NewOrm()
	var (
		login LoginLog
		err   error
	)

	o.Using("default")
	err = o.QueryTable(&login).Filter("User", user.Id).RelatedSel().OrderBy("-id").Limit(1).One(&login)
	return login, err
}

//更新
func UpdateLoginLog(userId int64) {
	o := orm.NewOrm()
	var (
		login LoginLog
		err   error
	)
	o.Using("default")
	err = o.QueryTable(&login).Filter("User", userId).RelatedSel().OrderBy("-id").Limit(1).One(&login)

	if err == nil {
		login.Logout = time.Now()
		if user, err := GetUser(userId); err == nil {
			login.UpdateUser = &user
			o.Update(&login)
		}
	}
}
