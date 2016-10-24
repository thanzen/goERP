package base

import "time"

type Base struct {
	Id         int64     `orm:"pk;auto"` //主键
	CreateUser *User     `orm:"rel(fk);null"`
	UpdateUser *User     `orm:"rel(fk);null"`
	CreateDate time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateDate time.Time `orm:"auto_now;type(datetime)"`
}
