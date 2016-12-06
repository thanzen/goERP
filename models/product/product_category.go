package product

import (
	"fmt"
	"pms/models/base"
	"pms/utils"

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

func ListProductCategory(condArr map[string]interface{}, start, length int64) (utils.Paginator, []ProductCategory, error) {

	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(new(ProductCategory))
	// qs = qs.RelatedSel()
	cond := orm.NewCondition()
	if name, ok := condArr["name"]; ok {
		cond = cond.And("name_icontains", name)
	}
	var (
		arrs []ProductCategory
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

	if num, err = qs.OrderBy("id").Limit(length, start).All(&arrs); err == nil {
		paginator.CurrentPageSize = num
	}

	return paginator, arrs, err
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

//根据ID查询类别
func GetProductCategoryByID(id int64) (ProductCategory, error) {
	o := orm.NewOrm()
	o.Using("default")
	var (
		obj ProductCategory
		err error
	)
	cond := orm.NewCondition()
	cond = cond.And("id", id)
	qs := o.QueryTable(new(ProductCategory))
	qs = qs.RelatedSel()
	err = qs.One(&obj)
	return obj, err
}

//根据名称查询类别
func GetProductCategoryByName(name string) (ProductCategory, error) {
	o := orm.NewOrm()
	o.Using("default")
	var (
		obj ProductCategory
		err error
	)
	cond := orm.NewCondition()
	qs := o.QueryTable(new(ProductCategory))

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
