package product

import (
	"pms/models/base"
)

type ProductAttributeLine struct {
	base.Base
	Name            string                   `orm:"unique"`   //产品属性名称
	Attribute       *ProductAttribute        `orm:"rel(fk)"`  //属性
	ProductTemplate *ProductTemplate         `orm:"rel(fk)"`  //产品模版
	AttributeValues []*ProductAttributeValue `orm:"rel(m2m)"` //属性值

}
