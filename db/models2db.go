package db

import (
	"pms/models/base"
	"pms/models/partner"
	"pms/models/product"

	"github.com/astaxie/beego/orm"
)

func init() {
	//============================基本表============================
	orm.RegisterModel(new(base.User), new(base.Record))
	orm.RegisterModel(new(base.Company), new(base.Group), new(base.Department))
	orm.RegisterModel(new(base.Country), new(base.Province), new(base.City), new(base.District))
	//============================客户表============================
	orm.RegisterModel(new(partner.Partner))
	//============================产品表============================
	//属性
	orm.RegisterModel(new(product.ProductAttribute))
	//属性值
	orm.RegisterModel(new(product.ProductAttributeValue))
	//属性值明细
	orm.RegisterModel(new(product.ProductAttributeLine))
	//产品款式
	orm.RegisterModel(new(product.ProductTemplate))
	//产品规格
	orm.RegisterModel(new(product.ProductProduct))
	//产品类别
	orm.RegisterModel(new(product.ProductCategory))
	//产品标签
	orm.RegisterModel(new(product.ProductTag))
	//产品供应商
	orm.RegisterModel(new(product.ProductSupplier))
	//产品属性价格
	orm.RegisterModel(new(product.ProductAttributePrice))
	//产品包装
	orm.RegisterModel(new(product.ProductPackaging))
	//产品计量单位
	orm.RegisterModel(new(product.ProductUom))
	//产品计量单位类别
	orm.RegisterModel(new(product.ProductUomCateg))

}
