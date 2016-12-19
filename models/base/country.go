package base

import (
	"pms/utils"

	"github.com/astaxie/beego/orm"
)

type Country struct {
	Base
	Name      string      `xml:"name"`          //国家名称
	Provinces []*Province `orm:"reverse(many)"` //省份
}

//添加国家
func CreateCountry(obj Country, user User) (int64, error) {
	o := orm.NewOrm()
	o.Using("default")
	country := new(Country)
	country.Name = obj.Name
	country.CreateUser = &user
	country.UpdateUser = &user
	id, err := o.Insert(country)
	return id, err
}

//获得某一个国家信息
func GetCountryByID(id int64) (Country, error) {
	o := orm.NewOrm()
	o.Using("default")
	country := Country{Base: Base{Id: id}}

	err := o.Read(&country)
	return country, err
}

//根据名称查询城市
func GetCountryByName(name string) (Country, error) {
	o := orm.NewOrm()
	o.Using("default")
	country := Country{Name: name}

	err := o.Read(&country)
	return country, err

}

//列出记录
func ListCountry(condArr map[string]interface{}, start, length int64) (utils.Paginator, []Country, error) {

	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(new(Country))
	// qs = qs.RelatedSel()
	cond := orm.NewCondition()
	if name, ok := condArr["name"]; ok {
		cond = cond.And("name__icontains", name)
	}
	var (
		countrys []Country
		num      int64
		err      error
	)
	var paginator utils.Paginator

	//后面再考虑查看权限的问题
	qs = qs.SetCond(cond)
	qs = qs.RelatedSel()
	if cnt, err := qs.Count(); err == nil {
		paginator = utils.GenPaginator(start, length, cnt)
	}

	if num, err = qs.OrderBy("id").Limit(length, start).All(&countrys); err == nil {
		paginator.CurrentPageSize = num
	}

	return paginator, countrys, err
}
