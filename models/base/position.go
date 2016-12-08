package base

import (
	"pms/utils"

	"github.com/astaxie/beego/orm"
)

type Position struct {
	Base
	Name string `orm:"unique"` //职位名称
}

func AddPosition(obj Position, user User) (int64, error) {
	o := orm.NewOrm()
	o.Using("default")
	position := new(Position)
	position.Name = obj.Name
	position.CreateUser = &user
	position.UpdateUser = &user
	id, err := o.Insert(position)
	return id, err
}

func GetPositionByID(id int64) (Position, error) {
	o := orm.NewOrm()
	o.Using("default")
	position := Position{Base: Base{Id: id}}

	err := o.Read(&position)

	return position, err
}

//根据名称查询部门
func GetPositionByName(name string) (Position, error) {
	o := orm.NewOrm()
	o.Using("default")
	position := Position{Name: name}

	err := o.Read(&position)

	return position, err

}

func ListPosition(condArr map[string]interface{}, start, length int64) (utils.Paginator, []Position, error) {

	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(new(Position))
	// qs = qs.RelatedSel()
	cond := orm.NewCondition()
	if name, ok := condArr["name"]; ok {
		cond = cond.And("name_icontains", name)
	}
	var (
		positions []Position
		num       int64
		err       error
	)
	var paginator utils.Paginator

	//后面再考虑查看权限的问题
	qs = qs.SetCond(cond)
	qs = qs.RelatedSel()
	if cnt, err := qs.Count(); err == nil {
		paginator = utils.GenPaginator(start, length, cnt)
	}

	if num, err = qs.OrderBy("id").Limit(length, start).All(&positions); err == nil {
		paginator.CurrentPageSize = num
	}

	return paginator, positions, err
}
