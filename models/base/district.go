package base

import "github.com/astaxie/beego/orm"

type District struct {
	Base
	Name string //区县名称
	City *City  `orm:"rel(fk)"` //城市
}

//添加区县
func AddDistrict(obj District, user User) (int64, error) {
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
func GetDistrict(id int64) (District, error) {
	o := orm.NewOrm()
	o.Using("default")
	district := District{Base: Base{Id: id}}
	err := o.Read(&district)
	return district, err
}
