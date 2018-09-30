package Models

type PageRows struct {
	// current page
	Cur       int
	TotalPage int
	Data      interface{}
}

func PageRowsInit(pageRows *PageRows, cur, totalPage int, data interface{}) {
	pageRows.Cur = cur
	pageRows.TotalPage = totalPage
	pageRows.Data = data
}
