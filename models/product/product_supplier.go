package product

import (
	"pms/models/base"
	"pms/models/partner"
	"time"
)

type ProductSupplier struct {
	base.Base
	Sequence        int32            //序列号
	Supplier        *partner.Partner `orm:"rel(fk)"` //供应商
	ProductName     string           //供应商产品名称
	ProductCode     string           //供应商产品编码
	FirstMinQty     float32          //第一单位采购最小数量
	SecondMinQty    float32          `orm:"default(0)"` //第二单位采购最小数量
	FirstPrice      float64          //第一单位采购价格
	SecondPrice     float64          `orm:"default(0)"`     //第二单位采购价格
	DateStart       time.Time        `orm:"type(datetime)"` //价格有效开始时间
	DateEnd         time.Time        `orm:"type(datetime)"` //价格有效截止时间
	DelayHour       int32            //下单到交货所需时间(小时)
	ProductTemplate *ProductTemplate `orm:"rel(fk);null"` //产品款式
	ProductProduct  *ProductProduct  `orm:"rel(fk);null"` //产品规格
}
