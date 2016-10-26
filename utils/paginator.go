package utils

type Paginator struct {
	CurrentPage     int64   //当前页
	PageSize        int64   //每页数量
	CurrentPageSize int64   //当前页数量
	TotalPage       int64   //总页数
	TotalCount      int64   //总数量
	FirstPage       bool    //为第一页
	LastPage        bool    //为最后一页
	PageList        []int64 //显示的页
}

func GenPaginator(page, offset, count int64) Paginator {
	var paginator Paginator
	paginator.TotalCount = count
	paginator.TotalPage = (count + offset - 1) / offset
	paginator.CurrentPage = page
	paginator.PageSize = offset
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
	if page >= 3 {
		paginator.PageList = []int64{page - 2, page - 1, page, page + 1, page + 2}
	} else {
		list := make([]int64, 0, 1)
		for index := 1; index < int(paginator.TotalPage+1); index++ {
			list = append(list, int64(index))
		}
		paginator.PageList = list
	}
	return paginator

}
