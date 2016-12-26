package sale

import "pms/models/base"

type SaleState struct {
	base.Base
	SaleConfig *SaleConfig `orm:"rel(fk)"`                     //订单配置
	Name       string      `orm:" null"  form:"name"`          //订单状态名称
	Active     bool        `orm:"default(true)" form:"active"` //有效
}
