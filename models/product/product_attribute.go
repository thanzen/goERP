package product

import (
	"pms/models/base"
	"pms/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type ProductAttribute struct {
	base.Base
	Name           string                   `orm:"unique"`        //产品属性名称
	Code           string                   `orm:"default(\"\")"` //产品属性编码
	Sequence       int32                    //序列
	ValueIds       []*ProductAttributeValue `orm:"reverse(many)"` //属性值
	AttributeLines []*ProductAttributeLine  `orm:"reverse(many)"`
}

//列出记录
func ListProductAttribute(condArr map[string]interface{}, page, offset int64) (utils.Paginator, error, []ProductAttribute) {

	if page < 1 {
		page = 1
	}

	if offset < 1 {
		offset, _ = beego.AppConfig.Int64("pageoffset")
	}

	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(new(ProductAttribute))
	// qs = qs.RelatedSel()
	cond := orm.NewCondition()

	var (
		productAttributes []ProductAttribute
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
	if num, err = qs.OrderBy("-id").Limit(offset, start).All(&productAttributes); err == nil {
		paginator.CurrentPageSize = num
	}

	return paginator, err, productAttributes
}

//添加属性
func AddProductAttribute(obj ProductAttribute, user base.User) (int64, error) {
	o := orm.NewOrm()
	o.Using("default")
	productAttribute := new(ProductAttribute)
	productAttribute.Name = obj.Name
	productAttribute.CreateUser = &user
	productAttribute.UpdateUser = &user
	productAttribute.Code = obj.Code
	productAttribute.Sequence = obj.Sequence
	id, err := o.Insert(productAttribute)
	return id, err
}

//获得某一个属性信息
func GetProductAttribute(id int64) (ProductAttribute, error) {
	o := orm.NewOrm()
	o.Using("default")
	productAttribute := ProductAttribute{Base: base.Base{Id: id}}
	err := o.Read(&productAttribute)
	return productAttribute, err
}
