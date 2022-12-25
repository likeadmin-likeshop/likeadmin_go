package req

//SystemLoginReq 系统登录参数
type SystemLoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

//SystemLogoutReq 登录退出参数
type SystemLogoutReq struct {
	Token string `header:"token" binding:"required"`
}
