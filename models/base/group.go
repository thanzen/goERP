package base

import (
	"fmt"
	"pms/utils"

	"github.com/astaxie/beego/orm"
)

//权限组
type Group struct {
	Base
	Name          string  `orm:"unique" xml:"name"` //组名称
	Members       []*User `orm:"reverse(many)"`     //组员
	GlobalLoation string  `orm:"unique" `           //全局定位
	Active        bool    `orm:"default(true)"`     //是否有效
	Description   string  `orm:"default(\"\")"`     //描述
}

func CreateGroup(obj *Group, user User) (int64, error) {
	o := orm.NewOrm()
	o.Using("default")
	group := new(Group)
	group.Name = obj.Name
	group.CreateUser = &user
	group.UpdateUser = &user
	id, err := o.Insert(group)
	return id, err
}

//根据ID查询权限组
func GetGroupByID(id int64) (Group, error) {
	o := orm.NewOrm()
	o.Using("default")
	var (
		group Group
		err   error
	)
	cond := orm.NewCondition()
	cond = cond.And("id", id)
	qs := o.QueryTable(new(Group))
	qs = qs.RelatedSel()
	err = qs.One(&group)
	return group, err
}

//根据名称查询权限组
func GetGroupByName(name string) (Group, error) {
	o := orm.NewOrm()
	o.Using("default")
	var (
		group Group
		err   error
	)
	cond := orm.NewCondition()
	qs := o.QueryTable(new(Group))

	if name != "" {
		cond = cond.And("name", name)
		qs = qs.SetCond(cond)
		qs = qs.RelatedSel()
		err = qs.One(&group)
	} else {
		err = fmt.Errorf("%s", "查询条件不成立")
	}

	return group, err
}
func ListGroup(condArr map[string]interface{}, start, length int64) (utils.Paginator, []Group, error) {

	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(new(Group))
	// qs = qs.RelatedSel()
	cond := orm.NewCondition()
	if name, ok := condArr["name"]; ok {
		cond = cond.And("name__icontains", name)
	}
	var (
		groups []Group
		num    int64
		err    error
	)
	var paginator utils.Paginator

	//后面再考虑查看权限的问题
	qs = qs.SetCond(cond)
	qs = qs.RelatedSel()
	if cnt, err := qs.Count(); err == nil {
		paginator = utils.GenPaginator(start, length, cnt)
	}

	if num, err = qs.Limit(length, start).All(&groups); err == nil {
		paginator.CurrentPageSize = num
	}

	return paginator, groups, err
}
