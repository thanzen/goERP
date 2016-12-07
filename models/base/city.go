package base

import (
	"pms/utils"

	"github.com/astaxie/beego/orm"
)

type City struct {
	Base
	Name      string      `json:"name"`                          //城市名称
	Province  *Province   `orm:"rel(fk)" json:"province"`        //国家
	Districts []*District `orm:"reverse(many)" json:"districts"` //城市
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
	_, err = o.LoadRelated(&city, "Province")
	return city, err
}

//根据名称查询城市
func GetCityByName(name string) (City, error) {
	o := orm.NewOrm()
	o.Using("default")
	city := City{Name: name}

	err := o.Read(&city)
	_, err = o.LoadRelated(&city, "Province")
	return city, err

}

//列出记录
func ListCity(condArr map[string]interface{}, start, length int64) (utils.Paginator, []City, error) {

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
		paginator = utils.GenPaginator(start, length, cnt)
	}

	if num, err = qs.OrderBy("id").Limit(length, start).All(&citys); err == nil {
		paginator.CurrentPageSize = num
	}

	return paginator, citys, err
}
