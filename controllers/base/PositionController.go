package base

import (
	"encoding/json"
	"fmt"
	mb "pms/models/base"
	"strconv"
	"strings"
)

type PositionController struct {
	BaseController
}

func (this *PositionController) Post() {

	action := this.Input().Get("action")
	switch action {
	case "validator":
		this.Validator()
	case "table": //bootstrap table的post请求
		this.PostList()
	case "selectSearch":
		this.PostList()
	default:
		this.PostList()
	}
}

func (this *PositionController) Get() {
	this.GetList()

	this.URL = "/position"
	this.Data["URL"] = this.URL
	this.Layout = "base/base.html"
	this.Data["MenuPositionActive"] = "active"
}

func (this *PositionController) Validator() {
	name := this.GetString("name")
	name = strings.TrimSpace(name)
	result := make(map[string]bool)
	if _, err := mb.GetPositionByName(name); err != nil {
		result["valid"] = true
	} else {
		result["valid"] = false
	}
	this.Data["json"] = result
	this.ServeJSON()
}

// 获得符合要求的城市数据
func (this *PositionController) positionList(start, length int64, condArr map[string]interface{}) (map[string]interface{}, error) {

	var positions []mb.Position
	paginator, positions, err := mb.ListPosition(condArr, start, length)
	fmt.Println(positions)
	result := make(map[string]interface{})
	if err == nil {

		// result["recordsFiltered"] = paginator.TotalCount
		tableLines := make([]interface{}, 0, 4)
		for _, position := range positions {
			oneLine := make(map[string]interface{})

			oneLine["Id"] = position.Id
			oneLine["name"] = position.Name

			tableLines = append(tableLines, oneLine)
		}
		fmt.Println(tableLines)
		result["data"] = tableLines
		if jsonResult, er := json.Marshal(&paginator); er == nil {
			result["paginator"] = string(jsonResult)
			result["total"] = paginator.TotalCount
		}
	}
	return result, err
}
func (this *PositionController) PostList() {
	condArr := make(map[string]interface{})
	start := this.Input().Get("offset")
	length := this.Input().Get("limit")
	var (
		startInt64  int64
		lengthInt64 int64
	)
	if startInt, ok := strconv.Atoi(start); ok == nil {
		startInt64 = int64(startInt)
	}
	if lengthInt, ok := strconv.Atoi(length); ok == nil {
		lengthInt64 = int64(lengthInt)
	}
	if result, err := this.positionList(startInt64, lengthInt64, condArr); err == nil {
		this.Data["json"] = result
	}
	this.ServeJSON()

}

func (this *PositionController) GetList() {
	this.Data["tableId"] = "table-position"
	this.TplName = "base/table_base.html"
}
