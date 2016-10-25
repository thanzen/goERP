package db

import (
	"pms/models/base"

	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(base.User), new(base.LoginLog))
	orm.RegisterModel(new(base.Group), new(base.Department))
	orm.RegisterModel(new(base.Country), new(base.Province), new(base.City), new(base.District))
}
