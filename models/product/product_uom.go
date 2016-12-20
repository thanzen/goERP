package product

import (
	"fmt"
	"pms/models/base"
	"pms/utils"

	"github.com/astaxie/beego/orm"
)

type ProductUom struct {
	base.Base
	Name      string           `orm:"unique" form:"name"`          //计量单位名称
	Active    bool             `orm:"default(true)" form:"active"` //有效
	Category  *ProductUomCateg `orm:"rel(fk)"`                     //计量单位类别
	Factor    float64          `form:"factor"`                     //比率
	FactorInv float64          `form:"factorInv"`                  //更大比率
	Rounding  float64          `form:"rounding"`                   //舍入精度
	Type      int64            `form:"type"`                       //类型
}

//列出记录
func ListProductUom(condArr map[string]interface{}, start, length int64) (utils.Paginator, []ProductUom, error) {

	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(new(ProductUom))
	// qs = qs.RelatedSel()
	cond := orm.NewCondition()
	if name, ok := condArr["name"]; ok {
		cond = cond.And("name__icontains", name)
	}
	var (
		arrs []ProductUom
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

//添加产品分类
func CreateProductUom(obj *ProductUom, user base.User) (int64, error) {
	o := orm.NewOrm()
	o.Using("default")
	obj.CreateUser = &user
	obj.UpdateUser = &user
	id, err := o.Insert(obj)
	return id, err
}

//修改产品分类
func UpdateProductUom(obj *ProductUom, user base.User) (int64, error) {
	o := orm.NewOrm()
	o.Using("default")
	updateObj := ProductUom{Base: base.Base{Id: obj.Id}}
	updateObj.UpdateUser = &user
	updateObj.Name = obj.Name
	if num, err := o.Update(&updateObj, "Name", "UpdateUser", "UpdateDate"); err == nil {
		return num, err
	} else {
		return 0, err
	}
}

//根据ID查询类别
func GetProductUomByID(id int64) (ProductUom, error) {
	o := orm.NewOrm()
	o.Using("default")
	obj := ProductUom{Base: base.Base{Id: id}}
	err := o.Read(&obj)
	if obj.Category != nil {
		o.Read(obj.Category)
	}
	return obj, err
}

//根据名称查询类别
func GetProductUomByName(name string) (ProductUom, error) {
	o := orm.NewOrm()
	o.Using("default")
	var (
		obj ProductUom
		err error
	)
	cond := orm.NewCondition()
	qs := o.QueryTable(new(ProductUom))

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
