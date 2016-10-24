package product

import (
	"pms/models/base"
)

type ProductTemplate struct {
	base.Base
	Name               string                  `orm:"unique"` //产品属性名称
	Sequence           int32                   //序列号
	Description        string                  `orm:"type(text)"`     //描述
	DescriptioSale     string                  `orm:"type(text)"`     //销售描述
	DescriptioPurchase string                  `orm:"type(text)"`     //采购描述
	Rental             bool                    `orm:"default(false)"` //代售品
	Categ              *ProductCatgory         `orm:"rel(fk)"`        //产品类别
	Price              float64                 //模版产品价格
	SaleOk             bool                    `orm:"default(true)"` //可销售
	Active             bool                    `orm:"default(true)"` //有效
	IsProductVariant   bool                    `orm:"default(true)"` //是变形产品
	FirstSaleUom       *ProductUom             `orm:"rel(fk)"`       //第一销售单位
	SecondSaleUom      *ProductUom             `orm:"rel(fk)"`       //第二销售单位
	FirstPurchaseUom   *ProductUom             `orm:"rel(fk)"`       //第一采购单位
	SecondPurchaseUom  *ProductUom             `orm:"rel(fk)"`       //第二采购单位
	AttributeLines     []*ProductAttributeLine `orm:"reverse(many)"` //属性明细
	ProductVariants    []*ProductProduct       `orm:"reverse(many)"` //产品规格明细
	TemplatePackagings []*ProductPackaging     `orm:"reverse(many)"` //打包方式
	VariantCount       int32                   //产品规格数量
	Barcode            string                  //条码,如ean13
	DefaultCode        string                  //产品编码
	// ProductPricelistItems []*ProductPricelistItem `orm:"reverse(many)"`
	PackagingDependTemp bool `orm:"default(true)"` //根据款式打包
	PurchaseDependTemp  bool `orm:"default(true)"` //根据款式采购，ture一个供应商可以供应所有的款式
}
