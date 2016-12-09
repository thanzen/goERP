package base

import (
	"pms/utils"
	"time"

	"github.com/astaxie/beego/orm"
)

type Record struct {
	Base
	User      *User     `orm:"rel(fk)"`
	Logout    time.Time `orm:"type(datetime);null"` //登出时间
	UserAgent string    `orm:"null"`                //用户代理
	Ip        string    //上次登录IP
}

//列出记录
func ListRecord(condArr map[string]interface{}, userId, start, length int64) (utils.Paginator, []Record, error) {

	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(new(Record))
	// qs = qs.RelatedSel()
	cond := orm.NewCondition()
	if start_date, ok := condArr["start_date"]; ok {
		cond = cond.And("create_date__gte", start_date)

	}
	if end_date, ok := condArr["end_date"]; ok {
		cond = cond.And("create_date__lte", end_date)

	}
	var (
		records []Record
		num     int64
		err     error
	)
	var paginator utils.Paginator

	//后面再考虑查看权限的问题
	qs = qs.SetCond(cond)
	qs = qs.RelatedSel()
	if cnt, err := qs.Count(); err == nil {
		paginator = utils.GenPaginator(start, length, cnt)
	}

	if num, err = qs.OrderBy("-id").Limit(length, start).All(&records); err == nil {
		paginator.CurrentPageSize = num
	}

	return paginator, records, err
}

//添加记录
func AddRecord(user User, IP, UserAgent string) (int64, error) {
	o := orm.NewOrm()
	o.Using("default")
	record := new(Record)
	record.User = &user
	record.CreateUser = &user
	record.Ip = IP
	record.UserAgent = UserAgent
	id, err := o.Insert(record)
	return id, err
}

//获得某一个用户记录信息
func GetLastRecordByUserID(userId int64) (Record, error) {
	o := orm.NewOrm()
	var (
		record Record
		err    error
	)

	o.Using("default")
	err = o.QueryTable(&record).Filter("User", userId).RelatedSel().OrderBy("-id").Limit(1).One(&record)
	return record, err
}

//更新
func UpdateRecord(userId int64, Ip string) {
	o := orm.NewOrm()
	var (
		record Record
		err    error
	)
	o.Using("default")
	err = o.QueryTable(&record).Filter("User", userId).Filter("ip", Ip).RelatedSel().OrderBy("-id").Limit(1).One(&record)

	if err == nil {
		record.Logout = time.Now()
		if user, err := GetUserByID(userId); err == nil {
			record.UpdateUser = &user
			o.Update(&record)
		}
	}
}
