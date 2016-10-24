package product

import (
	"pms/models/base"
)

type ProductUom struct {
	base.Base
	Name      string           `orm:"unique"`        //计量单位名称
	Active    bool             `orm:"default(true)"` //有效
	Category  *ProductUomCateg `orm:"rel(fk)"`       //计量单位类别
	Factor    float64          //比率
	FactorInv float64          //更大比率
	Rounding  float64          //舍入精度
}
