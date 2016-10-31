package product

import (
	"pms/models/base"
	"pms/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type ProductPriceList struct {
	base.Base
	Name   string                  //价格表名称
	Active bool                    `orm:"default(true)"` //有效
	Items  []*ProductPricelistItem `orm:"reverse(many)"`
}

//列出记录
func ListProductPriceList(condArr map[string]interface{}, page, offset int64) (utils.Paginator, error, []ProductPriceList) {

	if page < 1 {
		page = 1
	}

	if offset < 1 {
		offset, _ = beego.AppConfig.Int64("pageoffset")
	}

	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(new(ProductPriceList))
	// qs = qs.RelatedSel()
	cond := orm.NewCondition()

	var (
		productPriceLists []ProductPriceList
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
	if num, err = qs.OrderBy("-id").Limit(offset, start).All(&productPriceLists); err == nil {
		paginator.CurrentPageSize = num
	}

	return paginator, err, productPriceLists
}

//添加属性
func AddProductPriceList(obj ProductPriceList, user base.User) (int64, error) {
	o := orm.NewOrm()
	o.Using("default")
	productPriceList := new(ProductPriceList)
	productPriceList.Name = obj.Name
	productPriceList.CreateUser = &user
	productPriceList.UpdateUser = &user

	id, err := o.Insert(productPriceList)
	return id, err
}

//获得某一个属性信息
func GetProductPriceList(id int64) (ProductPriceList, error) {
	o := orm.NewOrm()
	o.Using("default")
	productPriceList := ProductPriceList{Base: base.Base{Id: id}}
	err := o.Read(&productPriceList)
	return productPriceList, err
}
