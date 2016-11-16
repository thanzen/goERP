package base

import (
	"pms/utils"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Record struct {
	Base
	User   *User     `orm:"rel(fk)"`
	Logout time.Time `orm:"type(datetime);null"` //登出时间
	Ip     string    //上次登录IP
}

//列出记录
func ListRecord(condArr map[string]interface{}, userId, page, offset int64) (utils.Paginator, error, []Record) {

	if page < 1 {
		page = 1
	}

	if offset < 1 {
		offset, _ = beego.AppConfig.Int64("pageoffset")
	}

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
		paginator = utils.GenPaginator(page, offset, cnt)
	}
	start := (page - 1) * offset
	if num, err = qs.OrderBy("-id").Limit(offset, start).All(&records); err == nil {
		paginator.CurrentPageSize = num
	}

	return paginator, err, records
}

//添加记录
func AddRecord(user User, IP string) (int64, error) {
	o := orm.NewOrm()
	o.Using("default")
	record := new(Record)
	record.User = &user
	record.CreateUser = &user
	record.Ip = IP
	id, err := o.Insert(record)
	return id, err
}

//获得某一个用户记录信息
func GetRecord(user User) (*Record, error) {
	o := orm.NewOrm()
	var (
		record Record
		err    error
	)

	o.Using("default")
	err = o.QueryTable(&record).Filter("User", user.Id).RelatedSel().OrderBy("-id").Limit(1).One(&record)
	return &record, err
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
		if user, err := GetUserById(userId); err == nil {
			record.UpdateUser = &user
			o.Update(&record)
		}
	}
}
