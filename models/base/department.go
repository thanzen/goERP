package base

import (
	"pms/utils"

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

//添加部门
func AddDepartment(obj Department, user User) (int64, error) {
	o := orm.NewOrm()
	o.Using("default")
	department := new(Department)
	department.Name = obj.Name
	department.CreateUser = &user
	department.UpdateUser = &user
	department.Company = obj.Company
	department.Leader = obj.Leader
	id, err := o.Insert(department)
	return id, err
}

//获得某一个部门信息
func GetDepartmentByID(id int64) (Department, error) {
	o := orm.NewOrm()
	o.Using("default")
	department := Department{Base: Base{Id: id}}

	err := o.Read(&department)
	if department.Leader != nil {
		o.Read(department.Leader)
	}

	return department, err
}

//根据名称查询部门
func GetDepartmentByName(name string) (Department, error) {
	o := orm.NewOrm()
	o.Using("default")
	department := Department{Name: name}

	err := o.Read(&department)
	if department.Leader != nil {
		o.Read(department.Leader)
	}
	return department, err

}
func ListDepartment(condArr map[string]interface{}, page, offset int64) (utils.Paginator, []Department, error) {

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
	if name, ok := condArr["name"]; ok {
		cond = cond.And("name__icontains", name)
	}
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

	return paginator, departments, err
}
