package db

import (
	"pms/models/address"
	"pms/models/base"

	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(base.User), new(base.Record))
	orm.RegisterModel(new(base.Group), new(base.Department))
	orm.RegisterModel(new(address.Country), new(address.Province), new(address.City), new(address.District))
}
