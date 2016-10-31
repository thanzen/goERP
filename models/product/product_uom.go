package product

import (
	"pms/models/base"
	"pms/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type ProductUom struct {
	base.Base
	Name      string           `orm:"unique"`        //计量单位名称
	Active    bool             `orm:"default(true)"` //有效
	Category  *ProductUomCateg `orm:"rel(fk)"`       //计量单位类别
	Factor    float64          //比率
	FactorInv float64          //更大比率
	Rounding  float64          //舍入精度
}

//列出记录
func ListProductUom(condArr map[string]interface{}, page, offset int64) (utils.Paginator, error, []ProductUom) {

	if page < 1 {
		page = 1
	}

	if offset < 1 {
		offset, _ = beego.AppConfig.Int64("pageoffset")
	}

	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(new(ProductUom))
	// qs = qs.RelatedSel()
	cond := orm.NewCondition()

	var (
		productUoms []ProductUom
		num         int64
		err         error
	)
	var paginator utils.Paginator

	//后面再考虑查看权限的问题
	qs = qs.SetCond(cond)
	qs = qs.RelatedSel()
	if cnt, err := qs.Count(); err == nil {
		paginator = utils.GenPaginator(page, offset, cnt)
	}

	start := (page - 1) * offset
	if num, err = qs.OrderBy("-id").Limit(offset, start).All(&productUoms); err == nil {
		paginator.CurrentPageSize = num
	}

	return paginator, err, productUoms
}

//添加属性
func AddProductUom(obj ProductUom, user base.User) (int64, error) {
	o := orm.NewOrm()
	o.Using("default")
	productUom := new(ProductUom)
	productUom.CreateUser = &user
	productUom.UpdateUser = &user

	id, err := o.Insert(productUom)
	return id, err
}

//获得某一个属性信息
func GetProductUom(id int64) (ProductUom, error) {
	o := orm.NewOrm()
	o.Using("default")
	productUom := ProductUom{Base: base.Base{Id: id}}
	err := o.Read(&productUom)
	return productUom, err
}
