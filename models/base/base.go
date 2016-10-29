package base

import "time"

type Base struct {
	Id         int64     `orm:"pk;auto"`                     //主键
	CreateUser *User     `orm:"rel(fk);null"`                //创建者
	UpdateUser *User     `orm:"rel(fk);null"`                //最后更新者
	CreateDate time.Time `orm:"auto_now_add;type(datetime)"` //创建时间
	UpdateDate time.Time `orm:"auto_now;type(datetime)"`     //最后更新时间
}
