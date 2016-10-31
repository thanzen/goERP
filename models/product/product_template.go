package product

import (
	"pms/models/base"
	"pms/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type ProductTemplate struct {
	base.Base
	Name               string                  `orm:"unique"` //产品属性名称
	Sequence           int32                   //序列号
	Description        string                  `orm:"type(text)"`     //描述
	DescriptioSale     string                  `orm:"type(text)"`     //销售描述
	DescriptioPurchase string                  `orm:"type(text)"`     //采购描述
	Rental             bool                    `orm:"default(false)"` //代售品
	Categ              *ProductCategory        `orm:"rel(fk)"`        //产品类别
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

//列出记录
func ListProductTemplate(condArr map[string]interface{}, page, offset int64) (utils.Paginator, error, []ProductTemplate) {

	if page < 1 {
		page = 1
	}

	if offset < 1 {
		offset, _ = beego.AppConfig.Int64("pageoffset")
	}

	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(new(ProductTemplate))
	// qs = qs.RelatedSel()
	cond := orm.NewCondition()

	var (
		productTemplates []ProductTemplate
		num              int64
		err              error
	)
	var paginator utils.Paginator

	//后面再考虑查看权限的问题
	qs = qs.SetCond(cond)
	qs = qs.RelatedSel()
	if cnt, err := qs.Count(); err == nil {
		paginator = utils.GenPaginator(page, offset, cnt)
	}

	start := (page - 1) * offset
	if num, err = qs.OrderBy("-id").Limit(offset, start).All(&productTemplates); err == nil {
		paginator.CurrentPageSize = num
	}

	return paginator, err, productTemplates
}

//添加属性
func AddProductTemplate(obj ProductTemplate, user base.User) (int64, error) {
	o := orm.NewOrm()
	o.Using("default")
	productTemplate := new(ProductTemplate)
	productTemplate.Name = obj.Name
	productTemplate.CreateUser = &user
	productTemplate.UpdateUser = &user
	id, err := o.Insert(productTemplate)
	return id, err
}

//获得某一个属性信息
func GetProductTemplate(id int64) (ProductTemplate, error) {
	o := orm.NewOrm()
	o.Using("default")
	productTemplate := ProductTemplate{Base: base.Base{Id: id}}
	err := o.Read(&productTemplate)
	return productTemplate, err
}
