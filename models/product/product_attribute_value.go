package product

import (
	"fmt"
	"pms/models/base"
	"pms/utils"

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
func ListProductAttributeValue(condArr map[string]interface{}, start, length int64) (utils.Paginator, []ProductAttributeValue, error) {

	o := orm.NewOrm()

	o.Using("default")
	qs := o.QueryTable(new(ProductAttributeValue))
	// qs = qs.RelatedSel()
	cond := orm.NewCondition()

	var (
		arrs []ProductAttributeValue
		num  int64
		err  error
	)
	var paginator utils.Paginator

	//后面再考虑查看权限的问题
	qs = qs.SetCond(cond)
	qs = qs.RelatedSel()
	if cnt, err := qs.Count(); err == nil {
		paginator = utils.GenPaginator(start, length, cnt)
	}

	if num, err = qs.OrderBy("-id").Limit(length, start).All(&arrs); err == nil {
		paginator.CurrentPageSize = num
	}

	return paginator, arrs, err
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
func GetProductAttributeValueByID(id int64) (ProductAttributeValue, error) {
	o := orm.NewOrm()
	o.Using("default")
	productProduct := ProductAttributeValue{Base: base.Base{Id: id}}
	err := o.Read(&productProduct)
	return productProduct, err
}

func GetProductAttributeValueByName(name string) (ProductAttributeValue, error) {
	o := orm.NewOrm()
	o.Using("default")
	var (
		obj ProductAttributeValue
		err error
	)
	cond := orm.NewCondition()
	qs := o.QueryTable(new(ProductAttributeValue))

	if name != "" {
		cond = cond.And("name", name)
		qs = qs.SetCond(cond)
		qs = qs.RelatedSel()
		err = qs.One(&obj)
	} else {
		err = fmt.Errorf("%s", "查询条件不成立")
	}

	return obj, err
}
