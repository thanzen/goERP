package product

import (
	"fmt"
	"pms/models/base"
	"pms/utils"

	"github.com/astaxie/beego/orm"
)

type ProductUomCateg struct {
	base.Base
	Name string        `orm:"unique" form:"name"` //计量单位分类
	Uoms []*ProductUom `orm:"reverse(many)"`      //计量单位
}

//列出记录
func ListProductUomCateg(condArr map[string]interface{}, start, length int64) (utils.Paginator, []ProductUomCateg, error) {

	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(new(ProductUomCateg))
	// qs = qs.RelatedSel()
	cond := orm.NewCondition()
	if name, ok := condArr["name"]; ok {
		cond = cond.And("name__icontains", name)
	}
	var (
		arrs []ProductUomCateg
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
	for i, _ := range arrs {
		o.LoadRelated(&arrs[i], "Uoms")
	}
	return paginator, arrs, err
}

//添加产品分类
func CreateProductUomCateg(obj *ProductUomCateg, user base.User) (int64, error) {
	o := orm.NewOrm()
	o.Using("default")
	obj.CreateUser = &user
	obj.UpdateUser = &user
	id, err := o.Insert(obj)
	return id, err
}

//修改产品分类
func UpdateProductUomCateg(obj *ProductUomCateg, user base.User) (int64, error) {
	o := orm.NewOrm()
	o.Using("default")
	updateObj := ProductUomCateg{Base: base.Base{Id: obj.Id}}
	updateObj.UpdateUser = &user
	updateObj.Name = obj.Name
	if num, err := o.Update(&updateObj, "Name", "UpdateUser", "UpdateDate"); err == nil {
		return num, err
	} else {
		return 0, err
	}
}

//根据ID查询类别
func GetProductUomCategByID(id int64) (ProductUomCateg, error) {
	o := orm.NewOrm()
	o.Using("default")
	obj := ProductUomCateg{Base: base.Base{Id: id}}
	err := o.Read(&obj)
	o.LoadRelated(&obj, "Uoms")
	return obj, err
}

//根据名称查询类别
func GetProductUomCategByName(name string) (ProductUomCateg, error) {
	o := orm.NewOrm()
	o.Using("default")
	var (
		obj ProductUomCateg
		err error
	)
	cond := orm.NewCondition()
	qs := o.QueryTable(new(ProductUomCateg))

	if name != "" {
		cond = cond.And("name", name)
		qs = qs.SetCond(cond)
		qs = qs.RelatedSel()
		err = qs.One(&obj)
		o.LoadRelated(&obj, "Uoms")
	} else {
		err = fmt.Errorf("%s", "查询条件不成立")
	}

	return obj, err
}
