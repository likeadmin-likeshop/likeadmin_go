package req

//SystemLoginReq 系统登录参数
type SystemLoginReq struct {
	Username string `json:"username" binding:"required,min=2,max=20"`
	Password string `json:"password" binding:"required,min=6,max=32"`
}

//SystemLogoutReq 登录退出参数
type SystemLogoutReq struct {
	Token string `header:"token" binding:"required"`
}
