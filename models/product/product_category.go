package product

import (
	"pms/models/base"
	"pms/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type ProductCategory struct {
	base.Base
	Name           string             `orm:"unique"`        //产品属性名称
	Parent         *ProductCategory   `orm:"rel(fk);null"`  //上级分类
	Childs         []*ProductCategory `orm:"reverse(many)"` //下级分类
	Sequence       int64              //序列
	ParentFullPath string             //上级全路径
}

//列出记录
func ListProductCategory(condArr map[string]interface{}, page, offset int64) (utils.Paginator, error, []ProductCategory) {

	if page < 1 {
		page = 1
	}

	if offset < 1 {
		offset, _ = beego.AppConfig.Int64("pageoffset")
	}

	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(new(ProductCategory))
	// qs = qs.RelatedSel()
	cond := orm.NewCondition()

	var (
		productCategories []ProductCategory
		num               int64
		err               error
	)
	var paginator utils.Paginator

	//后面再考虑查看权限的问题
	qs = qs.SetCond(cond)
	qs = qs.RelatedSel()
	if cnt, err := qs.Count(); err == nil {
		paginator = utils.GenPaginator(page, offset, cnt)
	}

	start := (page - 1) * offset
	if num, err = qs.OrderBy("-id").Limit(offset, start).All(&productCategories); err == nil {
		paginator.CurrentPageSize = num
	}

	return paginator, err, productCategories
}

//添加属性
func AddProductCategory(obj ProductCategory, user base.User) (int64, error) {
	o := orm.NewOrm()
	o.Using("default")

	productCategory := new(ProductCategory)
	productCategory.Name = obj.Name
	productCategory.CreateUser = &user
	productCategory.UpdateUser = &user
	id, err := o.Insert(productCategory)
	return id, err
}

//获得某一个属性信息
func GetProductCategory(id int64) (ProductCategory, error) {
	var productCategory ProductCategory
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(new(ProductCategory))
	qs = qs.RelatedSel()
	err := qs.Filter("id", id).Limit(1).One(&productCategory)
	return productCategory, err
}
