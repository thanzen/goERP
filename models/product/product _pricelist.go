package product

import "pms/models/base"

type ProductPriceList struct {
	base.Base
	Name   string                  //价格表名称
	Active bool                    `orm:"default(true)"` //有效
	Items  []*ProductPricelistItem `orm:"reverse(many)"`
}
