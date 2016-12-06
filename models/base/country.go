package base

import (
	"fmt"
	"pms/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Country struct {
	Base
	Name      string      `xml:"name"`          //国家名称
	Provinces []*Province `orm:"reverse(many)"` //省份
}

//添加国家
func AddCountry(obj Country, user User) (int64, error) {
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
	var (
		country Country
		err     error
	)
	cond := orm.NewCondition()
	cond = cond.And("id", id)
	qs := o.QueryTable(new(Country))
	qs = qs.RelatedSel()
	err = qs.One(&country)
	return country, err
}

//根据名称查询国家
func GetCountryByName(name string) (Country, error) {
	o := orm.NewOrm()
	o.Using("default")
	var (
		country Country
		err     error
	)
	cond := orm.NewCondition()
	qs := o.QueryTable(new(Country))

	if name != "" {
		cond = cond.And("name", name)
		qs = qs.SetCond(cond)
		qs = qs.RelatedSel()
		err = qs.One(&country)
	} else {
		err = fmt.Errorf("%s", "查询条件不成立")
	}

	return country, err
}

//列出记录
func ListCountry(condArr map[string]interface{}, page, offset int64) (utils.Paginator, []Country, error) {

	if page < 1 {
		page = 1
	}

	if offset < 1 {
		offset, _ = beego.AppConfig.Int64("pageoffset")
	}

	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(new(Country))
	// qs = qs.RelatedSel()
	cond := orm.NewCondition()
	if name, ok := condArr["name"]; ok {
		cond = cond.And("name_icontains", name)
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
		paginator = utils.GenPaginator(page, offset, cnt)
	}

	start := (page - 1) * offset
	if num, err = qs.OrderBy("id").Limit(offset, start).All(&countrys); err == nil {
		paginator.CurrentPageSize = num
	}

	return paginator, countrys, err
}
