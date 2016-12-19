package base

import (
	"pms/utils"

	"github.com/astaxie/beego/orm"
)

type District struct {
	Base
	Name string //区县名称
	City *City  `orm:"rel(fk)"` //城市
}

//添加区县
func CreateDistrict(obj District, user User) (int64, error) {
	o := orm.NewOrm()
	o.Using("default")
	district := new(District)
	district.Name = obj.Name
	district.City = obj.City
	district.CreateUser = &user
	district.UpdateUser = &user
	id, err := o.Insert(district)
	return id, err
}

//获得某一个区县信息
func GetDistrictByID(id int64) (District, error) {
	o := orm.NewOrm()
	o.Using("default")
	district := District{Base: Base{Id: id}}

	err := o.Read(&district)
	_, err = o.LoadRelated(&district, "City")
	return district, err
}

//根据名称查询区县
func GetDistrictByName(name string) (District, error) {
	o := orm.NewOrm()
	o.Using("default")
	district := District{Name: name}

	err := o.Read(&district)
	_, err = o.LoadRelated(&district, "City")
	return district, err

}

//列出记录
func ListDistrict(condArr map[string]interface{}, start, length int64) (utils.Paginator, []District, error) {

	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(new(District))
	// qs = qs.RelatedSel()
	cond := orm.NewCondition()
	if name, ok := condArr["name"]; ok {
		cond = cond.And("name__icontains", name)
	}
	var (
		districts []District
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

	if num, err = qs.OrderBy("id").Limit(length, start).All(&districts); err == nil {
		paginator.CurrentPageSize = num
	}

	return paginator, districts, err
}
