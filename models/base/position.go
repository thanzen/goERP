package base

import (
	"pms/utils"

	"github.com/astaxie/beego"
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

func ListPosition(condArr map[string]interface{}, page, offset int64) (utils.Paginator, []Position, error) {
	if page < 1 {
		page = 1
	}

	if offset < 1 {
		offset, _ = beego.AppConfig.Int64("pageoffset")
	}

	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(new(Position))
	// qs = qs.RelatedSel()
	cond := orm.NewCondition()

	if name, ok := condArr["name"]; ok {
		cond = cond.And("name__icontains", name)
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
		paginator = utils.GenPaginator(page, offset, cnt)
	}
	start := (page - 1) * offset
	if num, err = qs.Limit(offset, start).All(&positions); err == nil {
		paginator.CurrentPageSize = num
	}

	return paginator, positions, err
}
