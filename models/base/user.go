package base

import (
	"pms/utils"

	"github.com/astaxie/beego/orm"
)

type User struct {
	Base
	Name       string      `orm:"size(20)" xml:"name" form:"username"`               //用户名
	NameZh     string      `orm:"size(20)"  form:"namezh"`                           //中文用户名
	Department *Department `orm:"rel(fk);null;" form:"department"`                   //部门
	Email      string      `orm:"size(20)" xml:"email" form:"email"`                 //邮箱
	Mobile     string      `orm:"size(20);default(\"\")" xml:"mobile" form:"mobile"` //手机号码
	Tel        string      `orm:"size(20);default(\"\")" form:"tel"`                 //固定号码
	Password   string      `xml:"password" form:"password"`                          //密码
	Group      []*Group    `orm:"rel(m2m);rel_table(user_groups)"`                   //权限组
	IsAdmin    bool        `orm:"default(false)" xml:"isAdmin" form:"isadmin"`       //是否为超级用户
	Active     bool        `orm:"default(true)" xml:"active" form:"active"`          //有效
	Qq         string      `orm:"default(\"\")" xml:"qq" form:"qq"`                  //QQ
	WeChat     string      `orm:"default(\"\")" xml:"wechat" form:"wechat"`          //微信
	Position   *Position   `orm:"rel(fk);null;" form:"position"`                     //职位

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
func ListUser(condArr map[string]interface{}, start, length int64) (utils.Paginator, []User, error) {

	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(new(User))
	// qs = qs.RelatedSel()
	cond := orm.NewCondition()
	if active, ok := condArr["active"]; ok {
		cond = cond.And("active", active)

	} else {
		cond = cond.And("active", true)
	}
	if name, ok := condArr["name"]; ok {
		cond = cond.And("name__icontains", name)
	}
	if departmentId, ok := condArr["departmentId"]; ok {
		cond = cond.And("department__id", departmentId)
	}
	var (
		users []User
		num   int64
		err   error
	)
	var paginator utils.Paginator

	//后面再考虑查看权限的问题
	qs = qs.SetCond(cond)
	qs = qs.RelatedSel()
	if cnt, err := qs.Count(); err == nil {
		paginator = utils.GenPaginator(start, length, cnt)
	}
	if num, err = qs.Limit(length, start).All(&users); err == nil {
		paginator.CurrentPageSize = num
	}

	return paginator, users, err
}

//添加用户
func CreateUser(insetUser *User, currentUser User) (int64, error) {
	o := orm.NewOrm()
	o.Using("default")

	insetUser.CreateUser = &currentUser
	insetUser.UpdateUser = &currentUser
	password := utils.PasswordMD5(insetUser.Password, insetUser.Mobile)
	insetUser.Password = password

	id, err := o.Insert(insetUser)
	return id, err
}

//获得某一个用户信息
func GetUserByID(id int64) (User, error) {
	o := orm.NewOrm()
	o.Using("default")
	user := User{Base: Base{Id: id}}
	err := o.Read(&user)
	if user.Department != nil {
		o.Read(user.Department)
	}

	if user.Position != nil {
		o.Read(user.Position)
	}
	return user, err
}
func GetUserByName(name string) (User, error) {
	o := orm.NewOrm()
	var user User
	//7LR8ZC-855575-64657756081974692
	o.Using("default")
	cond := orm.NewCondition()
	cond = cond.And("mobile", name).Or("email", name).Or("name", name)
	qs := o.QueryTable(&user)
	qs = qs.SetCond(cond)
	err := qs.One(&user)
	return user, err
}
func CheckUserByName(name, password string) (User, error, bool) {
	o := orm.NewOrm()
	var (
		user User
		err  error
		ok   bool
	)
	ok = false
	//7LR8ZC-855575-64657756081974692
	o.Using("default")
	cond := orm.NewCondition()
	cond = cond.And("active", true).And("mobile", name).Or("email", name)
	qs := o.QueryTable(&user)
	qs = qs.SetCond(cond)
	if err = qs.One(&user); err == nil {
		if user.Password == utils.PasswordMD5(password, user.Mobile) {
			ok = true
		}
	}
	return user, err, ok

}

func UpdateUser(obj *User, user User) (int64, error) {
	o := orm.NewOrm()
	o.Using("default")
	updateObj := User{Base: Base{Id: obj.Id}}
	updateObj.UpdateUser = &user
	return o.Update(&obj, "UpdateUser", "UpdateDate")

}
