package product

import (
	"pms/models/base"
)

type ProductAttributeValue struct {
	base.Base
	Name       string                 `orm:"unique"`     //产品属性名称
	Attribute  *ProductAttribute      `orm:"rel(fk)"`    //属性
	Products   []*ProductProduct      `orm:"rel(m2m)"`   //产品规格
	PriceExtra float64                `orm:"default(0)"` //额外价格
	// Prices     *ProductAttributePrice `orm:"reverse(many)"`
	Sequence   int32                  //序列
}
