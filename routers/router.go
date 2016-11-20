package routers

import (
	"pms/controllers/address"
	"pms/controllers/base"
	"pms/controllers/product"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &base.IndexController{})
	//=======================================基本操作===========================================
	//登录
	beego.Router("/login/:action([A-Za-z]+)/", &base.LoginController{})
	//用户
	beego.Router("/user/:action([A-Za-z]+)/?:id", &base.UserController{})
	//部门
	beego.Router("/department/:action([A-Za-z]+)/?:id", &base.DepartmentController{})
	//职位
	beego.Router("/position/:action([A-Za-z]+)/?:id", &base.PositionController{})
	//登录日志
	beego.Router("/record/:action([A-Za-z]+)/?:id", &base.RecordController{})
	//国家
	beego.Router("/country/:action([A-Za-z]+)/?:id", &address.CountryController{})
	//省份
	beego.Router("/province/:action([A-Za-z]+)/?:id", &address.ProvinceController{})
	//城市
	beego.Router("/city/:action([A-Za-z]+)/?:id", &address.CityController{})
	//区县
	beego.Router("/district/:action([A-Za-z]+)/?:id", &address.DistrictController{})
	//=======================================产品管理===========================================
	//属性
	beego.Router("/product/attribute/:action([A-Za-z]+)/?:id", &product.ProductAttributeController{})
	//属性值
	beego.Router("/product/attributevalue/:action([A-Za-z]+)/?:id", &product.ProductAttributeValueController{})
	//属性值明细
	beego.Router("/product/attributeline/:action([A-Za-z]+)/?:id", &product.ProductAttributeLineController{})
	//产品款式
	beego.Router("/product/template/:action([A-Za-z]+)/?:id", &product.ProductTemplateController{})
	//产品规格
	beego.Router("/product/product/:action([A-Za-z]+)/?:id", &product.ProductProductController{})
	//产品类别
	beego.Router("/product/category/:action([A-Za-z]+)/?:id", &product.ProductCategoryController{})
	//产品标签
	beego.Router("/product/tag/:action([A-Za-z]+)/?:id", &product.ProductTagController{})
	//产品包装
	beego.Router("/product/packaging/:action([A-Za-z]+)/?:id", &product.ProductPackagingController{})
	//产品属性价格
	beego.Router("/product/attributeprice/:action([A-Za-z]+)/?:id", &product.ProductAttributePriceController{})
	//产品计量单位
	beego.Router("/product/uom/:action([A-Za-z]+)/?:id", &product.ProductUomController{})
	//产品计量单位类别
	beego.Router("/product/uomcateg/:action([A-Za-z]+)/?:id", &product.ProductUomCategController{})

}
