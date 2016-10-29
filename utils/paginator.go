package utils

import "fmt"

type Paginator struct {
	CurrentPage     int64   //当前页
	NextPage        int64   //下一页
	PrePage         int64   //上一页
	PageSize        int64   //每页数量
	CurrentPageSize int64   //当前页数量
	TotalPage       int64   //总页数
	TotalCount      int64   //总数量
	FirstPage       bool    //为第一页
	LastPage        bool    //为最后一页
	PageList        []int64 //显示的页
	Max             int64
	Url             string //Url地址头
}

func GenPaginator(page, offset, count int64) Paginator {
	var paginator Paginator
	paginator.TotalCount = count
	fmt.Println("====================")
	fmt.Println(offset)
	paginator.TotalPage = (count + offset - 1) / offset
	paginator.PageSize = offset
	if page < 1 {
		page = 1
	}
	if page == 1 {
		paginator.FirstPage = true
	} else {
		paginator.FirstPage = false
	}
	if offset == paginator.TotalPage {
		paginator.LastPage = true
	} else {
		paginator.LastPage = false
	}
	if page > paginator.TotalPage {
		page = paginator.TotalPage
	}
	paginator.CurrentPage = page
	fmt.Println("paginator.CurrentPage")
	fmt.Println(paginator.CurrentPage)
	list := make([]int64, 0, 1)
	if paginator.TotalPage <= 5 {
		for index := 1; index <= int(paginator.TotalPage); index++ {
			list = append(list, int64(index))
		}
		paginator.PageList = list
	} else {
		if page+2 >= paginator.TotalPage {
			paginator.PageList = []int64{paginator.TotalPage - 4, paginator.TotalPage - 3, paginator.TotalPage - 2, paginator.TotalPage - 1, paginator.TotalPage}
		} else if page <= 2 {
			paginator.PageList = []int64{1, 2, 3, 4, 5}
		} else {
			paginator.PageList = []int64{page - 2, page - 1, page, page + 1, page + 2}
		}
	}
	paginator.NextPage = page + 1
	paginator.PrePage = page - 1
	fmt.Println("paginator.CurrentPage")
	fmt.Println(paginator.CurrentPage)
	if paginator.TotalCount > 0 && paginator.CurrentPage > 0 {
		paginator.Max = paginator.TotalCount / paginator.CurrentPage
	} else {
		paginator.Max = 0
	}
	return paginator

}
