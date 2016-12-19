package product

import (
	"pms/models/base"
	"pms/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type ProductTag struct {
	base.Base
	Name string `orm:"unique"` //产品标记名称

}

//列出记录
func ListProductTag(condArr map[string]interface{}, page, offset int64) (utils.Paginator, error, []ProductTag) {

	if page < 1 {
		page = 1
	}

	if offset < 1 {
		offset, _ = beego.AppConfig.Int64("pageoffset")
	}

	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(new(ProductTag))
	// qs = qs.RelatedSel()
	cond := orm.NewCondition()

	var (
		productTags []ProductTag
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
	if num, err = qs.OrderBy("-id").Limit(offset, start).All(&productTags); err == nil {
		paginator.CurrentPageSize = num
	}

	return paginator, err, productTags
}

//添加属性
func CreateProductTag(obj ProductTag, user base.User) (int64, error) {
	o := orm.NewOrm()
	o.Using("default")
	productTag := new(ProductTag)
	productTag.CreateUser = &user
	productTag.UpdateUser = &user

	id, err := o.Insert(productTag)
	return id, err
}

//获得某一个属性信息
func GetProductTag(id int64) (ProductTag, error) {
	o := orm.NewOrm()
	o.Using("default")
	productTag := ProductTag{Base: base.Base{Id: id}}
	err := o.Read(&productTag)
	return productTag, err
}
