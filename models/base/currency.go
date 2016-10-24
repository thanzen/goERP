package base

type Currncy struct {
	Base
	Active         bool   `orm:"default(true)"` //有效
	Name           string `orm:"unique"`        //货币代码
	Symbol         string //货币符号
	PositionBefore bool   `orm:"default(true)"` //符号位于金额前面
}
