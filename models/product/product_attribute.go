package product

import (
	"pms/models/base"
)

type ProductAttribute struct {
	base.Base
	Name           string                   `orm:"unique"`        //产品属性名称
	Code           string                   `orm:"default(\"\")"` //产品属性编码
	Sequence       int32                    //序列
	ValueIds       []*ProductAttributeValue `orm:"reverse(many)"` //属性值
	AttributeLines []*ProductAttributeLine  `orm:"reverse(many)"`
}
