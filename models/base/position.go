package base

import (
	"fmt"
	"pms/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Position struct {
	Base
	Name string `orm:"unique"` //职位名称
}

func ListPosition(condArr map[string]interface{}, page, offset int64) (utils.Paginator, error, []Position) {

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

	return paginator, err, positions
}

func GetPositionByName(name string, exact bool) (int64, []Position, error) {
	o := orm.NewOrm()
	o.Using("default")
	var (
		positions []Position
		err       error
		num       int64
	)
	cond := orm.NewCondition()
	qs := o.QueryTable(new(Position))

	if name != "" {
		cond = cond.And("name__icontains", name)
		qs = qs.SetCond(cond)
		qs = qs.RelatedSel()
		num, err = qs.All(&positions)
	} else {
		if exact == true {
			err = fmt.Errorf("%s", "查询条件不成立")
		} else {
			qs = qs.SetCond(cond)
			qs = qs.RelatedSel()
			num, err = qs.Limit(5, 0).All(&positions)
		}
	}

	return num, positions, err
}
func GetPositionById(id int64) (Position, error) {
	o := orm.NewOrm()
	o.Using("default")
	position := Position{Base: Base{Id: id}}
	err := o.Read(&position)

	return position, err
}
