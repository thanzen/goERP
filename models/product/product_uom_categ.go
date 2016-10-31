package product

import (
	"pms/models/base"
	"pms/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type ProductUomCateg struct {
	base.Base
	Name string `orm:"unique"` //计量单位分类
}

//列出记录
func ListProductUomCateg(condArr map[string]interface{}, page, offset int64) (utils.Paginator, error, []ProductUomCateg) {

	if page < 1 {
		page = 1
	}

	if offset < 1 {
		offset, _ = beego.AppConfig.Int64("pageoffset")
	}

	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(new(ProductUomCateg))
	// qs = qs.RelatedSel()
	cond := orm.NewCondition()

	var (
		productUomCategs []ProductUomCateg
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
	if num, err = qs.OrderBy("-id").Limit(offset, start).All(&productUomCategs); err == nil {
		paginator.CurrentPageSize = num
	}

	return paginator, err, productUomCategs
}

//添加属性
func AddProductUomCateg(obj ProductUomCateg, user base.User) (int64, error) {
	o := orm.NewOrm()
	o.Using("default")
	productUomCateg := new(ProductUomCateg)
	productUomCateg.CreateUser = &user
	productUomCateg.UpdateUser = &user

	id, err := o.Insert(productUomCateg)
	return id, err
}

//获得某一个属性信息
func GetProductUomCateg(id int64) (ProductUomCateg, error) {
	o := orm.NewOrm()
	o.Using("default")
	productUomCateg := ProductUomCateg{Base: base.Base{Id: id}}
	err := o.Read(&productUomCateg)
	return productUomCateg, err
}
