package product

import (
	"pms/models/base"
)

type ProductPackaging struct {
	base.Base
	Name            string
	sequence        int32            //序列号
	ProductTemplate *ProductTemplate `orm:"rel(fk);null"` //产品款式
	ProductProduct  *ProductProduct  `orm:"rel(fk);null"` //产品规格
	FirstQty        float64          //第一单位最大数量
	SecondQty       float64          //第二单位最大数量

}
