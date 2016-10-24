package product

import (
	"pms/models/base"
)

type ProductCatgory struct {
	base.Base
	Name     string            `orm:"unique"`        //产品属性名称
	Parent   *ProductCatgory   `orm:"ref(fk);null"`  //上级分类
	Childs   []*ProductCatgory `orm:"reverse(many)"` //下级分类
	Sequence int32             //序列
}
