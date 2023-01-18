package request

//PageReq 分页请求参数
type PageReq struct {
	PageNo   int `form:"pageNo,default=1" validate:"omitempty,gte=1"`          // 页码
	PageSize int `form:"pageSize,default=20" validate:"omitempty,gt=0,lte=60"` // 每页大小
}
