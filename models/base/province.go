package base

import (
	"fmt"
	"pms/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Province struct {
	Base
	Name    string   `xml:"ProvinceName,attr"` //省份名称
	Country *Country `orm:"rel(fk)"`           //国家
	Citys   []*City  `orm:"reverse(many)"`     //城市

}

//添加省份
func AddProvince(obj Province, user User) (int64, error) {
	o := orm.NewOrm()
	o.Using("default")
	province := new(Province)
	province.Name = obj.Name
	province.Country = obj.Country
	province.CreateUser = &user
	province.UpdateUser = &user
	id, err := o.Insert(province)
	return id, err
}

//获得某一个省份信息
func GetProvinceByID(id int64) (Province, error) {
	o := orm.NewOrm()
	o.Using("default")
	var (
		province Province
		err      error
	)
	cond := orm.NewCondition()
	cond = cond.And("id", id)
	qs := o.QueryTable(new(Province))
	qs = qs.RelatedSel()
	err = qs.One(&province)
	return province, err
}
func GetProvinceByName(name string) (Province, error) {
	o := orm.NewOrm()
	o.Using("default")
	var (
		province Province
		err      error
	)
	cond := orm.NewCondition()
	qs := o.QueryTable(new(Province))

	if name != "" {
		cond = cond.And("name", name)
		qs = qs.SetCond(cond)
		qs = qs.RelatedSel()
		err = qs.One(&province)
	} else {
		err = fmt.Errorf("%s", "查询条件不成立")
	}

	return province, err
}

//列出记录
func ListProvince(condArr map[string]interface{}, page, offset int64) (utils.Paginator, []Province, error) {

	if page < 1 {
		page = 1
	}

	if offset < 1 {
		offset, _ = beego.AppConfig.Int64("pageoffset")
	}

	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(new(Province))
	// qs = qs.RelatedSel()
	cond := orm.NewCondition()

	var (
		provinces []Province
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
	if num, err = qs.OrderBy("id").Limit(offset, start).All(&provinces); err == nil {
		paginator.CurrentPageSize = num
	}

	return paginator, provinces, err
}
