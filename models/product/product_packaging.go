package product

import (
	"pms/models/base"
	"pms/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type ProductPackaging struct {
	base.Base
	Name            string
	sequence        int32            //序列号
	ProductTemplate *ProductTemplate `orm:"rel(fk);null"` //产品款式
	ProductProduct  *ProductProduct  `orm:"rel(fk);null"` //产品规格
	FirstQty        float64          //第一单位最大数量
	SecondQty       float64          //第二单位最大数量

}

//列出记录
func ListProductPackaging(condArr map[string]interface{}, page, offset int64) (utils.Paginator, error, []ProductPackaging) {

	if page < 1 {
		page = 1
	}

	if offset < 1 {
		offset, _ = beego.AppConfig.Int64("pageoffset")
	}

	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(new(ProductPackaging))
	// qs = qs.RelatedSel()
	cond := orm.NewCondition()

	var (
		productPackagings []ProductPackaging
		num               int64
		err               error
	)
	var paginator utils.Paginator

	//后面再考虑查看权限的问题
	qs = qs.SetCond(cond)
	qs = qs.RelatedSel()
	if cnt, err := qs.Count(); err == nil {
		paginator = utils.GenPaginator(page, offset, cnt)
	}

	start := (page - 1) * offset
	if num, err = qs.OrderBy("-id").Limit(offset, start).All(&productPackagings); err == nil {
		paginator.CurrentPageSize = num
	}

	return paginator, err, productPackagings
}

//添加属性
func AddProductPackaging(obj ProductPackaging, user base.User) (int64, error) {
	o := orm.NewOrm()
	o.Using("default")
	productPackaging := new(ProductPackaging)
	productPackaging.CreateUser = &user
	productPackaging.UpdateUser = &user

	id, err := o.Insert(productPackaging)
	return id, err
}

//获得某一个属性信息
func GetProductPackaging(id int64) (ProductPackaging, error) {
	o := orm.NewOrm()
	o.Using("default")
	productPackaging := ProductPackaging{Base: base.Base{Id: id}}
	err := o.Read(&productPackaging)
	return productPackaging, err
}
