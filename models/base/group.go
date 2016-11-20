package base

import (
	"fmt"
	"pms/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//权限组
type Group struct {
	Base
	Name          string  `orm:"unique" xml:"name"` //组名称
	Members       []*User `orm:"reverse(many)"`     //组员
	GlobalLoation string  `orm:"unique" `           //全局定位
}

func AddGroup(obj Group, user User) (int64, error) {
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
	group := Group{Base: Base{Id: id}}
	err := o.Read(&group)
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
func ListGroup(condArr map[string]interface{}, page, offset int64) (utils.Paginator, []Group, error) {

	if page < 1 {
		page = 1
	}

	if offset < 1 {
		offset, _ = beego.AppConfig.Int64("pageoffset")
	}

	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(new(Group))
	// qs = qs.RelatedSel()
	cond := orm.NewCondition()

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
		paginator = utils.GenPaginator(page, offset, cnt)
	}
	start := (page - 1) * offset
	if num, err = qs.Limit(offset, start).All(&groups); err == nil {
		paginator.CurrentPageSize = num
	}

	return paginator, groups, err
}
