package partner

import (
	"pms/models/base"
)

type Partner struct {
	base.Base
	Name       string         //合作伙伴名称
	IsCompany  bool           `orm:"default(true)"`  //是公司
	IsSupplier bool           `orm:"default(false)"` //是供应商
	IsCustomer bool           `orm:"default(true)"`  //是客户
	Active     bool           `orm:"default(true)"`  //有效
	Country    *base.Country  `orm:"rel(fk);null"`   //国家
	Province   *base.Province `orm:"rel(fk);null"`   //身份
	City       *base.City     `orm:"rel(fk);null"`   //城市
	District   *base.District `orm:"rel(fk);null"`   //区县
	Street     string         `orm:"default(\"\")"`  //街道
	Parent     *Partner       `orm:"rel(fk);null"`   //母公司
	Childs     []*Partner     `orm:"reverse(many)"`  //下级
	Mobile     string         `orm:"default(\"\")"`  //电话号码
	Tel        string         `orm:"default(\"\")"`  //座机
	Email      string         `orm:"default(\"\")"`  //邮箱
	Comment    string         `orm:"type(text)"`     //备注
}
