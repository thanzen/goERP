package db

import (
	. "pms/models/base"

	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(User))
	orm.RegisterModel(new(Group))
	orm.RegisterModel(new(Country), new(Province), new(City), new(District))
}
