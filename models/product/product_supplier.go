package product

import (
	"pms/models/base"
	"pms/models/partner"
	"pms/utils"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type ProductSupplier struct {
	base.Base
	Sequence    int32            //序列号
	Supplier    *partner.Partner `orm:"rel(fk)"` //供应商
	ProductName string           //供应商产品名称
	ProductCode string           //供应商产品编码
	FirstMinQty float32          //第一单位采购最小数量
	// SecondMinQty    float32          `orm:"default(0)"` //第二单位采购最小数量
	FirstPrice float64 //第一单位采购价格
	// SecondPrice     float64          `orm:"default(0)"`     //第二单位采购价格
	DateStart       time.Time        `orm:"type(datetime)"` //价格有效开始时间
	DateEnd         time.Time        `orm:"type(datetime)"` //价格有效截止时间
	DelayHour       int32            //下单到交货所需时间(小时)
	ProductTemplate *ProductTemplate `orm:"rel(fk);null"` //产品款式
	ProductProduct  *ProductProduct  `orm:"rel(fk);null"` //产品规格
}

//列出记录
func ListProductSupplier(condArr map[string]interface{}, page, offset int64) (utils.Paginator, error, []ProductSupplier) {

	if page < 1 {
		page = 1
	}

	if offset < 1 {
		offset, _ = beego.AppConfig.Int64("pageoffset")
	}

	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(new(ProductSupplier))
	// qs = qs.RelatedSel()
	cond := orm.NewCondition()

	var (
		productSuppliers []ProductSupplier
		num              int64
		err              error
	)
	var paginator utils.Paginator

	//后面再考虑查看权限的问题
	qs = qs.SetCond(cond)
	qs = qs.RelatedSel()
	if cnt, err := qs.Count(); err == nil {
		paginator = utils.GenPaginator(page, offset, cnt)
	}

	start := (page - 1) * offset
	if num, err = qs.OrderBy("-id").Limit(offset, start).All(&productSuppliers); err == nil {
		paginator.CurrentPageSize = num
	}

	return paginator, err, productSuppliers
}

//添加属性
func CreateProductSupplier(obj ProductSupplier, user base.User) (int64, error) {
	o := orm.NewOrm()
	o.Using("default")
	productSupplier := new(ProductSupplier)
	productSupplier.CreateUser = &user
	productSupplier.UpdateUser = &user

	id, err := o.Insert(productSupplier)
	return id, err
}

//获得某一个属性信息
func GetProductSupplier(id int64) (ProductSupplier, error) {
	o := orm.NewOrm()
	o.Using("default")
	productSupplier := ProductSupplier{Base: base.Base{Id: id}}
	err := o.Read(&productSupplier)
	return productSupplier, err
}
