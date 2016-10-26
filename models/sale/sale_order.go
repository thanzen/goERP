package purchase

import (
	"pms/models/base"
	"pms/models/partner"
)

type PurchaseOrder struct {
	base.Base
	Name      string               `orm:"unique"`             //采购订单号
	Origin    string               `orm:"default(\"\")"`      //订单来源
	Partner   *partner.Partner     `orm:"rel(fk)"`            //供应商
	Comment   string               `orm:"type(text)"`         //说明
	State     string               `orm:"default(\"state\")"` //订单状态
	OrderLine []*PurchaseOrderLine `orm:"reverse(many)"`      //订单明细
}
