package base

import "github.com/astaxie/beego/orm"

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
func GetCountry(id int64) (Country, error) {
	o := orm.NewOrm()
	o.Using("default")
	country := Country{Base: Base{Id: id}}
	err := o.Read(&country)
	return country, err
}
