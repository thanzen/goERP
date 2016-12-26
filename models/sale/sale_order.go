package sale

import "pms/models/base"

type SaleOrder struct {
	base.Base

	Name      string           `orm:"unique"`        //销售订单号
	Company   *base.Company    `orm:"rel(fk)"`       //公司
	Comment   string           `orm:"type(text)"`    //说明
	State     *SaleState       `orm:"rel(fk)"`       //订单状态
	OrderLine []*SaleOrderLine `orm:"reverse(many)"` //订单明细
	Note      string           //说明
}
