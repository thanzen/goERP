package product

import (
	"pms/models/base"
	"pms/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type ProductPricelistItem struct {
	base.Base
}

//列出记录
func ListProductPricelistItem(condArr map[string]interface{}, page, offset int64) (utils.Paginator, error, []ProductPricelistItem) {

	if page < 1 {
		page = 1
	}

	if offset < 1 {
		offset, _ = beego.AppConfig.Int64("pageoffset")
	}

	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(new(ProductPricelistItem))
	// qs = qs.RelatedSel()
	cond := orm.NewCondition()

	var (
		productPricelistItems []ProductPricelistItem
		num                   int64
		err                   error
	)
	var paginator utils.Paginator

	//后面再考虑查看权限的问题
	qs = qs.SetCond(cond)
	qs = qs.RelatedSel()
	if cnt, err := qs.Count(); err == nil {
		paginator = utils.GenPaginator(page, offset, cnt)
	}

	start := (page - 1) * offset
	if num, err = qs.OrderBy("-id").Limit(offset, start).All(&productPricelistItems); err == nil {
		paginator.CurrentPageSize = num
	}

	return paginator, err, productPricelistItems
}

//添加属性
func AddProductPricelistItem(obj ProductPricelistItem, user base.User) (int64, error) {
	o := orm.NewOrm()
	o.Using("default")
	productPricelistItem := new(ProductPricelistItem)
	productPricelistItem.CreateUser = &user
	productPricelistItem.UpdateUser = &user

	id, err := o.Insert(productPricelistItem)
	return id, err
}

//获得某一个属性信息
func GetProductPricelistItem(id int64) (ProductPricelistItem, error) {
	o := orm.NewOrm()
	o.Using("default")
	productPricelistItem := ProductPricelistItem{Base: base.Base{Id: id}}
	err := o.Read(&productPricelistItem)
	return productPricelistItem, err
}
