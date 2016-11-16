package base

type Company struct {
	Base
	Name       string        `orm:"unique"`        //公司名称
	Children   []*Company    `orm:"reverse(many)"` //子公司
	Parent     *Company      `orm:"rel(fk);null"`  //上级公司
	Department []*Department `orm:"reverse(many)"` //部门
	Country    *Country      `orm:"rel(fk);null"`  //国家
	Province   *Province     `orm:"rel(fk);null"`  //身份
	City       *City         `orm:"rel(fk);null"`  //城市
	District   *District     `orm:"rel(fk);null"`  //区县
	Street     string        `orm:"default(\"\")"` //街道
}
