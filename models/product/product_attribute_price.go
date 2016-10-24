package product

import (
	"pms/models/base"
)

type ProductAttributePrice struct {
	base.Base
	ProductTemplate *ProductTemplate       `orm:"rel(fk)"`    //产品款式
	AttributeValue  *ProductAttributeValue `orm:"rel(fk)"`    //属性值
	PriceExtra      float64                `orm:"default(0)"` //属性价格
}
