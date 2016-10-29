package address

import (
	. "pms/models/base"
	"pms/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Country struct {
	Base
	Name      string      `xml:"name"`          //国家名称
	Provinces []*Province `orm:"reverse(many)"` //省份
}

//列出记录
func ListCountry(condArr map[string]interface{}, page, offset int64) (utils.Paginator, error, []Country) {

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
	if num, err = qs.OrderBy("-id").Limit(offset, start).All(&countrys); err == nil {
		paginator.CurrentPageSize = num
	}

	return paginator, err, countrys
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
func GetCountry(id int64) (Country, error) {
	o := orm.NewOrm()
	o.Using("default")
	country := Country{Base: Base{Id: id}}
	err := o.Read(&country)
	return country, err
}
