package product

import (
	"pms/models/base"
	"pms/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type ProductAttributeLine struct {
	base.Base
	Name            string                   `orm:"unique"`   //产品属性名称
	Attribute       *ProductAttribute        `orm:"rel(fk)"`  //属性
	ProductTemplate *ProductTemplate         `orm:"rel(fk)"`  //产品模版
	AttributeValues []*ProductAttributeValue `orm:"rel(m2m)"` //属性值

}

//列出记录
func ListProductAttributeLine(condArr map[string]interface{}, page, offset int64) (utils.Paginator, []ProductAttributeLine, error) {

	if page < 1 {
		page = 1
	}

	if offset < 1 {
		offset, _ = beego.AppConfig.Int64("pageoffset")
	}

	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(new(ProductAttributeLine))
	qs = qs.RelatedSel()
	cond := orm.NewCondition()
	if tmp_id, ok := condArr["Tmp_id"]; ok {
		cond = cond.And("ProductTemplate__id", tmp_id)
	}
	var (
		productAttributeLines []ProductAttributeLine
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
	if num, err = qs.OrderBy("-id").Limit(offset, start).All(&productAttributeLines); err == nil {
		paginator.CurrentPageSize = num
	}
	for i, _ := range productAttributeLines {
		o.QueryM2M(&productAttributeLines[i], "AttributeValues")
	}
	return paginator, productAttributeLines, err
}

//添加属性
func CreateProductAttributeLine(obj ProductAttributeLine, user base.User) (int64, error) {
	o := orm.NewOrm()
	o.Using("default")
	productAttributeLine := new(ProductAttributeLine)
	productAttributeLine.Name = obj.Name
	productAttributeLine.CreateUser = &user
	productAttributeLine.UpdateUser = &user

	id, err := o.Insert(productAttributeLine)
	return id, err
}

//获得某一个属性信息
func GetProductAttributeLine(id int64) (ProductAttributeLine, error) {
	o := orm.NewOrm()
	o.Using("default")
	productAttributeLine := ProductAttributeLine{Base: base.Base{Id: id}}
	err := o.Read(&productAttributeLine)
	return productAttributeLine, err
}
