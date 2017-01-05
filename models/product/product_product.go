package product

import (
	"fmt"
	"pms/models/base"
	"pms/utils"

	"github.com/astaxie/beego/orm"
)

type ProductProduct struct {
	base.Base
	Name                string                   `orm:"unique"`        //产品属性名称
	IsProductVariant    bool                     `orm:"default(true)"` //是变形产品
	ProductTags         []*ProductTag            `orm:"reverse(many)"` //产品标签
	Categ               *ProductCategory         `orm:"rel(fk)"`       //产品类别
	Active              bool                     `orm:"default(true)"` //有效
	Barcode             string                   //条码,如ean13
	DefaultCode         string                   `orm:"unique"`        //产品编码
	ProductTemplate     *ProductTemplate         `orm:"rel(fk)"`       //产品款式
	AttributeValues     []*ProductAttributeValue `orm:"reverse(many)"` //产品属性
	FirstSaleUom        *ProductUom              `orm:"rel(fk)"`       //第一销售单位
	SecondSaleUom       *ProductUom              `orm:"rel(fk)"`       //第二销售单位
	FirstPurchaseUom    *ProductUom              `orm:"rel(fk)"`       //第一采购单位
	SecondPurchaseUom   *ProductUom              `orm:"rel(fk)"`       //第二采购单位
	ProductPackagings   []*ProductPackaging      `orm:"reverse(many)"` //打包方式
	PackagingDependTemp bool                     `orm:"default(true)"` //根据款式打包
	PurchaseDependTemp  bool                     `orm:"default(true)"` //根据款式采购，ture一个供应商可以供应所有的款式

}

//列出记录
func ListProductProduct(condArr map[string]interface{}, start, length int64) (utils.Paginator, []ProductProduct, error) {

	o := orm.NewOrm()

	o.Using("default")
	qs := o.QueryTable(new(ProductProduct))
	// qs = qs.RelatedSel()
	cond := orm.NewCondition()

	var (
		arrs []ProductProduct
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
func UpdateProductProduct(obj *ProductProduct, user base.User) (int64, error) {
	o := orm.NewOrm()
	o.Using("default")
	updateObj := ProductProduct{Base: base.Base{Id: obj.Id}}
	updateObj.UpdateUser = &user
	updateObj.Name = obj.Name
	return o.Update(&updateObj, "Name", "UpdateUser", "UpdateDate")
}

//添加属性
func CreateProductProduct(obj ProductProduct, user base.User) (int64, error) {
	o := orm.NewOrm()
	o.Using("default")
	productProduct := new(ProductProduct)
	productProduct.Name = obj.Name
	productProduct.CreateUser = &user
	productProduct.UpdateUser = &user
	id, err := o.Insert(productProduct)
	return id, err
}

//获得某一个产品模版信息
func GetProductProductByID(id int64) (ProductProduct, error) {
	o := orm.NewOrm()
	o.Using("default")
	productProduct := ProductProduct{Base: base.Base{Id: id}}
	err := o.Read(&productProduct)
	return productProduct, err
}

func GetProductProductByName(name string) (ProductProduct, error) {
	o := orm.NewOrm()
	o.Using("default")
	var (
		obj ProductProduct
		err error
	)
	cond := orm.NewCondition()
	qs := o.QueryTable(new(ProductProduct))

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
