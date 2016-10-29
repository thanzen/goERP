package base

type Position struct {
	Base
	Name string `orm:"unique"` //职位名称
}
