package base

import (
	"fmt"
	"pms/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type City struct {
	Base
	Name      string      //城市名称
	Province  *Province   `orm:"rel(fk)"`       //国家
	Districts []*District `orm:"reverse(many)"` //城市
}

//添加城市
func AddCity(obj City, user User) (int64, error) {
	o := orm.NewOrm()
	o.Using("default")
	city := new(City)
	city.Name = obj.Name
	city.CreateUser = &user
	city.UpdateUser = &user
	city.Province = obj.Province
	id, err := o.Insert(city)
	return id, err
}

//获得某一个城市信息
func GetCityByID(id int64) (City, error) {
	o := orm.NewOrm()
	o.Using("default")
	city := City{Base: Base{Id: id}}
	err := o.Read(&city)
	return city, err
}

//根据名称查询城市
func GetCityByName(name string) (City, error) {
	o := orm.NewOrm()
	o.Using("default")
	var (
		city City
		err  error
	)
	cond := orm.NewCondition()
	qs := o.QueryTable(new(City))

	if name != "" {
		cond = cond.And("name", name)
		qs = qs.SetCond(cond)
		qs = qs.RelatedSel()
		err = qs.One(&city)
	} else {
		err = fmt.Errorf("%s", "查询条件不成立")
	}

	return city, err
}

//列出记录
func ListCity(condArr map[string]interface{}, page, offset int64) (utils.Paginator, []City, error) {

	if page < 1 {
		page = 1
	}

	if offset < 1 {
		offset, _ = beego.AppConfig.Int64("pageoffset")
	}

	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(new(City))
	// qs = qs.RelatedSel()
	cond := orm.NewCondition()
	if name, ok := condArr["name"]; ok {
		cond = cond.And("name_icontains", name)
	}
	var (
		citys []City
		num   int64
		err   error
	)
	var paginator utils.Paginator

	//后面再考虑查看权限的问题
	qs = qs.SetCond(cond)
	qs = qs.RelatedSel()
	if cnt, err := qs.Count(); err == nil {
		paginator = utils.GenPaginator(page, offset, cnt)
	}

	start := (page - 1) * offset
	if num, err = qs.OrderBy("-id").Limit(offset, start).All(&citys); err == nil {
		paginator.CurrentPageSize = num
	}

	return paginator, citys, err
}
