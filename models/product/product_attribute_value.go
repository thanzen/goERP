package product

import (
	"pms/models/base"
	"pms/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type ProductAttributeValue struct {
	base.Base
	Name       string            `orm:"unique"`     //产品属性名称
	Attribute  *ProductAttribute `orm:"rel(fk)"`    //属性
	Products   []*ProductProduct `orm:"rel(m2m)"`   //产品规格
	PriceExtra float64           `orm:"default(0)"` //额外价格
	// Prices     *ProductAttributePrice `orm:"reverse(many)"`
	Sequence int32 //序列
}

//列出记录
func ListProductAttributeValue(condArr map[string]interface{}, page, offset int64) (utils.Paginator, error, []ProductAttributeValue) {

	if page < 1 {
		page = 1
	}

	if offset < 1 {
		offset, _ = beego.AppConfig.Int64("pageoffset")
	}

	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(new(ProductAttributeValue))
	// qs = qs.RelatedSel()
	cond := orm.NewCondition()

	var (
		productAttributeValues []ProductAttributeValue
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
	if num, err = qs.OrderBy("-id").Limit(offset, start).All(&productAttributeValues); err == nil {
		paginator.CurrentPageSize = num
	}

	return paginator, err, productAttributeValues
}

//添加属性
func AddProductAttributeValue(obj ProductAttributeValue, user base.User) (int64, error) {
	o := orm.NewOrm()
	o.Using("default")
	productAttributeValue := new(ProductAttributeValue)
	productAttributeValue.Name = obj.Name
	productAttributeValue.CreateUser = &user
	productAttributeValue.UpdateUser = &user
	id, err := o.Insert(productAttributeValue)
	return id, err
}

//获得某一个属性信息
func GetProductAttributeValue(id int64) (ProductAttributeValue, error) {
	o := orm.NewOrm()
	o.Using("default")
	productAttributeValue := ProductAttributeValue{Base: base.Base{Id: id}}
	err := o.Read(&productAttributeValue)
	return productAttributeValue, err
}
