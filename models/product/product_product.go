package product

import (
	"pms/models/base"
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
