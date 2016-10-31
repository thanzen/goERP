package product

import (
	"pms/models/base"
	"pms/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type ProductProduct struct {
	base.Base
	Name                string                   `orm:"unique"`        //产品属性名称
	IsProductVariant    bool                     `orm:"default(true)"` //是变形产品
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
func ListProductProduct(condArr map[string]interface{}, page, offset int64) (utils.Paginator, error, []ProductProduct) {

	if page < 1 {
		page = 1
	}

	if offset < 1 {
		offset, _ = beego.AppConfig.Int64("pageoffset")
	}

	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(new(ProductProduct))
	// qs = qs.RelatedSel()
	cond := orm.NewCondition()

	var (
		productProducts []ProductProduct
		num             int64
		err             error
	)
	var paginator utils.Paginator

	//后面再考虑查看权限的问题
	qs = qs.SetCond(cond)
	qs = qs.RelatedSel()
	if cnt, err := qs.Count(); err == nil {
		paginator = utils.GenPaginator(page, offset, cnt)
	}

	start := (page - 1) * offset
	if num, err = qs.OrderBy("-id").Limit(offset, start).All(&productProducts); err == nil {
		paginator.CurrentPageSize = num
	}

	return paginator, err, productProducts
}

//添加属性
func AddProductProduct(obj ProductProduct, user base.User) (int64, error) {
	o := orm.NewOrm()
	o.Using("default")
	productProduct := new(ProductProduct)
	productProduct.Name = obj.Name
	productProduct.CreateUser = &user
	productProduct.UpdateUser = &user
	id, err := o.Insert(productProduct)
	return id, err
}

//获得某一个属性信息
func GetProductProduct(id int64) (ProductProduct, error) {
	o := orm.NewOrm()
	o.Using("default")
	productProduct := ProductProduct{Base: base.Base{Id: id}}
	err := o.Read(&productProduct)
	return productProduct, err
}
