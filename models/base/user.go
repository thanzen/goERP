package base

import (
	"pms/utils"
	"time"

	"github.com/astaxie/beego/orm"
)

type User struct {
	Base
	Name      string    `orm:"size(20)" xml:"name"`                 //用户名
	NameZh    string    `orm:"size(20)" `                           //中文用户名
	Email     string    `orm:"size(20)" xml:"email"`                //邮箱
	Mobile    string    `orm:"size(20);default(\"\")" xml:"mobile"` //手机号码
	Tel       string    `orm:"size(20);default(\"\")"`              //固定号码
	Password  string    `xml:"password"`                            //密码
	Group     []*Group  `orm:"rel(m2m);rel_table(user_groups)"`
	IsAdmin   bool      `orm:"default(false)" xml:"isAdmin"` //是否为超级用户
	LastLogin time.Time `orm:"null"`                         //上次登录时间
	Active    bool      `orm:"default(true)" xml:"active"`   //有效
	Ip        string    //上次登录IP
}

//多字段唯一
func (u *User) TableUnique() [][]string {
	return [][]string{
		[]string{"Email", "Mobile"},
	}
}

// 多字段索引
func (u *User) TableIndex() [][]string {
	return [][]string{
		[]string{"Id", "Email", "Mobile"},
	}
}
func (u *User) TableName() string {
	return "auth_user"
}

//添加用户
func AddUser(obj User, cUser User) (int64, error) {
	o := orm.NewOrm()
	o.Using("default")
	user := new(User)
	user.Name = obj.Name
	user.Email = obj.Email
	user.Mobile = obj.Mobile
	user.Tel = obj.Tel
	user.IsAdmin = obj.IsAdmin
	user.Active = obj.Active
	user.CreateUser = &cUser
	user.UpdateUser = &cUser
	user.Password = utils.PasswordMD5(obj.Password, obj.Mobile)
	id, err := o.Insert(user)
	return id, err
}

//获得某一个用户信息
func GetUser(id int64) (User, error) {
	o := orm.NewOrm()
	o.Using("default")
	user := User{Base: Base{Id: id}}
	err := o.Read(&user)
	return user, err
}
func GetUserByName(name string) (User, error) {
	o := orm.NewOrm()
	var user User
	//7LR8ZC-855575-64657756081974692
	o.Using("default")
	cond := orm.NewCondition()
	cond = cond.And("mobile", name).Or("email", name)
	qs := o.QueryTable(&user)
	qs = qs.SetCond(cond)
	err := qs.One(&user)
	return user, err
}
func CheckUserByName(name, password string) (User, error) {
	o := orm.NewOrm()
	var (
		user User
		err  error
	)

	//7LR8ZC-855575-64657756081974692
	o.Using("default")
	cond := orm.NewCondition()
	cond = cond.And("active", true).And("mobile", name).Or("email", name)
	qs := o.QueryTable(&user)
	qs = qs.SetCond(cond)
	if err = qs.One(&user); err == nil {
		if user.Password == utils.PasswordMD5(password, user.Mobile) {
			return user, err
		}
	}
	return user, err

}
func UpdateUser(user User) (int64, error) {
	o := orm.NewOrm()
	return o.Update(&user)
}
