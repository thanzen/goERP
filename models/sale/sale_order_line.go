package sale

import (
	"pms/models/base"
	mp "pms/models/product"
)

type SaleOrderLine struct {
	base.Base
	Name          string             `orm:"unique"`  //销售订单明细号
	SaleOrder     *SaleOrder         `orm:"rel(fk)"` //销售订单
	Product       *mp.ProductProduct `orm:"rel(fk)"` //产品
	ProductName   string             //产品名称
	FirstSaleUom  *mp.ProductUom     `orm:"rel(fk)"` //第一销售单位
	SecondSaleUom *mp.ProductUom     `orm:"rel(fk)"` //第二销售单位
	FirstQty      float64            //第一单位数量
	SecondQty     float64            //第二单位数量
	PriceTotal    float64            //合计
}
