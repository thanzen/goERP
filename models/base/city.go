package base

import "github.com/astaxie/beego/orm"

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
func GetCity(id int64) (City, error) {
	o := orm.NewOrm()
	o.Using("default")
	city := City{Base: Base{Id: id}}
	err := o.Read(&city)
	return city, err
}
