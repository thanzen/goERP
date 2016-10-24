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
	cond := orm.NewCondition()
	cond = cond.And("user_id", user.Id)
	qs := o.QueryTable(&login)
	qs = qs.SetCond(cond).OrderBy("-id").Limit(1)
	err = qs.One(&login)
	return login, err
}

//更新
func UpdateLoginLog(userId int64) {
	o := orm.NewOrm()
	var (
		login LoginLog
	)
	o.Using("default")
	cond := orm.NewCondition()
	cond = cond.And("user_id", userId)
	qs := o.QueryTable(&login)
	qs = qs.SetCond(cond).OrderBy("-id").Limit(1)
	if qs.One(&login) != nil {
		login.Logout = time.Now()
		if user, err := GetUser(userId); err == nil {
			login.UpdateUser = &user
		}
		o.Insert(login)
	}

}
