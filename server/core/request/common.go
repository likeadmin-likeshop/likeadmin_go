package request

//PageReq 分页请求参数
type PageReq struct {
	PageNo   int `form:"pageNo"`   // 页码
	PageSize int `form:"pageSize"` // 每页大小
}
