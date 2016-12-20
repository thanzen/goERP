package product

import (
	"fmt"
	"pms/models/base"
	"pms/utils"

	"github.com/astaxie/beego/orm"
)

type ProductAttribute struct {
	base.Base
	Name           string                   `orm:"unique" form:"name"`        //产品属性名称
	Code           string                   `orm:"default(\"\")" form:"code"` //产品属性编码
	Sequence       int32                    `form:"sequence"`                 //序列
	ValueIds       []*ProductAttributeValue `orm:"reverse(many)"`             //属性值
	AttributeLines []*ProductAttributeLine  `orm:"reverse(many)"`
}

//列出记录
func ListProductAttribute(condArr map[string]interface{}, start, length int64) (utils.Paginator, []ProductAttribute, error) {

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
		paginator = utils.GenPaginator(start, length, cnt)
	}

	if num, err = qs.OrderBy("-id").Limit(length, start).All(&productAttributes); err == nil {
		paginator.CurrentPageSize = num
	}
	//后期需要改成线程池来获得关联数据,下面为线程池两种实现
	//https://github.com/Jeffail/tunny
	//https://github.com/jolestar/go-commons-pool
	for i, _ := range productAttributes {
		o.LoadRelated(&productAttributes[i], "ValueIds")
	}

	return paginator, productAttributes, err
}

//添加属性
func CreateProductAttribute(obj *ProductAttribute, user base.User) (int64, error) {
	o := orm.NewOrm()
	o.Using("default")
	obj.CreateUser = &user
	obj.UpdateUser = &user
	id, err := o.Insert(obj)
	return id, err
}

//更新产品属性
func UpdateProductAttribute(obj *ProductAttribute, user base.User) (int64, error) {
	o := orm.NewOrm()
	o.Using("default")
	updateObj := ProductAttribute{Base: base.Base{Id: obj.Id}}
	updateObj.UpdateUser = &user
	updateObj.Name = obj.Name
	updateObj.Code = obj.Code
	updateObj.Sequence = obj.Sequence

	return o.Update(&updateObj, "Name", "Code", "Sequence", "UpdateUser", "UpdateDate")

}

//获得某一个属性信息
func GetProductAttributeByID(id int64) (ProductAttribute, error) {
	o := orm.NewOrm()
	o.Using("default")
	obj := ProductAttribute{Base: base.Base{Id: id}}
	err := o.Read(&obj)
	o.LoadRelated(&obj, "ValueIds")
	return obj, err
}
func GetProductAttributeByName(name string) (ProductAttribute, error) {
	o := orm.NewOrm()
	o.Using("default")
	var (
		obj ProductAttribute
		err error
	)
	cond := orm.NewCondition()
	qs := o.QueryTable(new(ProductAttribute))

	if name != "" {
		cond = cond.And("name", name)
		qs = qs.SetCond(cond)
		qs = qs.RelatedSel()
		err = qs.One(&obj)
		o.LoadRelated(&obj, "ValueIds")
	} else {
		err = fmt.Errorf("%s", "查询条件不成立")
	}

	return obj, err
}
