package sale

import (
	"pms/models/base"
)

type SaleConfig struct {
	base.Base
	Name    string        //配置名称
	Company *base.Company `orm:"rel(fk)"` //公司

}
