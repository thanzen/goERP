package product

import (
	"pms/models/base"
	"pms/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type ProductAttributePrice struct {
	base.Base
	ProductTemplate *ProductTemplate       `orm:"rel(fk)"`    //产品款式
	AttributeValue  *ProductAttributeValue `orm:"rel(fk)"`    //属性值
	PriceExtra      float64                `orm:"default(0)"` //属性价格
}

//列出记录
func ListProductAttributePrice(condArr map[string]interface{}, page, offset int64) (utils.Paginator, error, []ProductAttributePrice) {

	if page < 1 {
		page = 1
	}

	if offset < 1 {
		offset, _ = beego.AppConfig.Int64("pageoffset")
	}

	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(new(ProductAttributePrice))
	// qs = qs.RelatedSel()
	cond := orm.NewCondition()

	var (
		productAttributePrices []ProductAttributePrice
		num                    int64
		err                    error
	)
	var paginator utils.Paginator

	//后面再考虑查看权限的问题
	qs = qs.SetCond(cond)
	qs = qs.RelatedSel()
	if cnt, err := qs.Count(); err == nil {
		paginator = utils.GenPaginator(page, offset, cnt)
	}

	start := (page - 1) * offset
	if num, err = qs.OrderBy("-id").Limit(offset, start).All(&productAttributePrices); err == nil {
		paginator.CurrentPageSize = num
	}

	return paginator, err, productAttributePrices
}

//添加属性
func CreateProductAttributePrice(obj ProductAttributePrice, user base.User) (int64, error) {
	o := orm.NewOrm()
	o.Using("default")
	productAttributePrice := new(ProductAttributePrice)
	productAttributePrice.CreateUser = &user
	productAttributePrice.UpdateUser = &user

	id, err := o.Insert(productAttributePrice)
	return id, err
}

//获得某一个属性信息
func GetProductAttributePrice(id int64) (ProductAttributePrice, error) {
	o := orm.NewOrm()
	o.Using("default")
	productAttributePrice := ProductAttributePrice{Base: base.Base{Id: id}}
	err := o.Read(&productAttributePrice)
	return productAttributePrice, err
}
