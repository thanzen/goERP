package product

import (
	"pms/models/base"
)

type ProductUomCateg struct {
	base.Base
	Name string `orm:"unique"` //计量单位分类
}
