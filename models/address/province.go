package address

import (
	. "pms/models/base"

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
func GetProvince(id int64) (Province, error) {
	o := orm.NewOrm()
	o.Using("default")
	province := Province{Base: Base{Id: id}}
	err := o.Read(&province)
	return province, err
}
