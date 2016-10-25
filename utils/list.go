package utils

type CellInfo struct {
	Id    interface{}
	Name  interface{}
	Value interface{}
	Url   string
}
type TableLineInfo struct {
	Url      string
	PageName string
	Title    map[string]interface{}
	Body     map[string]interface{}
}
type TableTitle struct {
	TitleCell map[int]interface{} //表头内容
	Len       int64               //表头长度
}
