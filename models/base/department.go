package base

import (
	"pms/utils"

	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Department struct {
	Base
	Name    string   `orm:"unique"`        //团队名称
	Leader  *User    `orm:"rel(fk);null"`  //团队领导者
	Members []*User  `orm:"reverse(many)"` //组员
	Company *Company `orm:"rel(fk);null"`  //公司
}

func ListDepartment(condArr map[string]interface{}, page, offset int64) (utils.Paginator, error, []Department) {

	if page < 1 {
		page = 1
	}

	if offset < 1 {
		offset, _ = beego.AppConfig.Int64("pageoffset")
	}

	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(new(Department))
	// qs = qs.RelatedSel()
	cond := orm.NewCondition()

	var (
		departments []Department
		num         int64
		err         error
	)
	var paginator utils.Paginator

	//后面再考虑查看权限的问题
	qs = qs.SetCond(cond)
	qs = qs.RelatedSel()
	if cnt, err := qs.Count(); err == nil {
		paginator = utils.GenPaginator(page, offset, cnt)
	}
	start := (page - 1) * offset
	if num, err = qs.Limit(offset, start).All(&departments); err == nil {
		paginator.CurrentPageSize = num
	}

	return paginator, err, departments
}
func GetDepartmentByName(name string, exact bool) (int64, []Department, error) {
	o := orm.NewOrm()
	o.Using("default")
	var (
		departments []Department
		err         error
		num         int64
	)
	cond := orm.NewCondition()
	qs := o.QueryTable(new(Department))

	if name != "" {
		cond = cond.And("name__icontains", name)
		qs = qs.SetCond(cond)
		qs = qs.RelatedSel()
		num, err = qs.All(&departments)
	} else {
		if exact == true {
			err = fmt.Errorf("%s", "查询条件不成立")
		} else {
			qs = qs.SetCond(cond)
			qs = qs.RelatedSel()
			num, err = qs.Limit(5, 0).All(&departments)
		}
	}

	return num, departments, err
}
func GetDepartmentById(id int64) (Department, error) {
	o := orm.NewOrm()
	o.Using("default")
	department := Department{Base: Base{Id: id}}
	err := o.Read(&department)

	return department, err
}
