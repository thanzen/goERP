package db

import (
	"pms/models/base"
	"pms/models/partner"

	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(base.User), new(base.Record))
	orm.RegisterModel(new(base.Company), new(base.Group), new(base.Department))
	orm.RegisterModel(new(base.Country), new(base.Province), new(base.City), new(base.District))
	orm.RegisterModel(new(partner.Partner))

}
