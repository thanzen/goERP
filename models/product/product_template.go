package product

import (
	"fmt"
	"pms/models/base"
	"pms/utils"

	"github.com/astaxie/beego/orm"
)

type ProductTemplate struct {
	base.Base
	Name               string                  `orm:"unique"` //产品属性名称
	Sequence           int32                   //序列号
	Description        string                  `orm:"type(text);null"` //描述
	DescriptioSale     string                  `orm:"type(text);null"` //销售描述
	DescriptioPurchase string                  `orm:"type(text);null"` //采购描述
	Rental             bool                    `orm:"default(false)"`  //代售品
	Categ              *ProductCategory        `orm:"rel(fk)"`         //产品类别
	Price              float64                 //模版产品价格
	StandardPrice      float64                 //成本价格
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
	ProductType        string                  `orm:"default(\"stock\")"` //产品类型
	ProductMethod      string                  `orm:"default(\"hand\")"`  //产品规格创建方式
	// ProductPricelistItems []*ProductPricelistItem `orm:"reverse(many)"`
	PackagingDependTemp bool `orm:"default(true)"` //根据款式打包
	PurchaseDependTemp  bool `orm:"default(true)"` //根据款式采购，ture一个供应商可以供应所有的款式
}

//列出记录
func ListProductTemplate(condArr map[string]interface{}, start, length int64) (utils.Paginator, []ProductTemplate, error) {

	o := orm.NewOrm()

	o.Using("default")
	qs := o.QueryTable(new(ProductTemplate))
	// qs = qs.RelatedSel()
	cond := orm.NewCondition()

	var (
		arrs []ProductTemplate
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

//添加产品模版
func CreateProductTemplate(obj *ProductTemplate, user base.User) (int64, error) {
	o := orm.NewOrm()
	o.Using("default")
	productTemplate := new(ProductTemplate)
	productTemplate.Name = obj.Name
	productTemplate.CreateUser = &user
	productTemplate.UpdateUser = &user
	id, err := o.Insert(productTemplate)
	return id, err
}

func UpdateProductTemplate(obj *ProductTemplate, user base.User) (int64, error) {
	o := orm.NewOrm()
	o.Using("default")
	updateObj := ProductTemplate{Base: base.Base{Id: obj.Id}}
	updateObj.UpdateUser = &user
	updateObj.Name = obj.Name
	return o.Update(&updateObj, "Name", "UpdateUser", "UpdateDate")
}

//获得某一个产品模版信息
func GetProductTemplateByID(id int64) (ProductTemplate, error) {
	o := orm.NewOrm()
	o.Using("default")
	obj := ProductTemplate{Base: base.Base{Id: id}}
	err := o.Read(&obj)
	if obj.Categ != nil {
		o.Read(obj.Categ)
	}
	if obj.FirstSaleUom != nil {
		o.Read(obj.FirstSaleUom)
	}
	if obj.SecondSaleUom != nil {
		o.Read(obj.SecondSaleUom)
	}
	if obj.FirstPurchaseUom != nil {
		o.Read(obj.FirstPurchaseUom)
	}
	if obj.SecondPurchaseUom != nil {
		o.Read(obj.SecondPurchaseUom)
	}
	return obj, err
}

func GetProductTemplateByName(name string) (ProductTemplate, error) {
	o := orm.NewOrm()
	o.Using("default")
	var (
		obj ProductTemplate
		err error
	)
	cond := orm.NewCondition()
	qs := o.QueryTable(new(ProductTemplate))

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
