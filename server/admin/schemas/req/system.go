package req

//SystemLoginReq 系统登录参数
type SystemLoginReq struct {
	Username string `json:"username" binding:"required,min=2,max=20"` // 账号
	Password string `json:"password" binding:"required,min=6,max=32"` // 密码
}

//SystemLogoutReq 登录退出参数
type SystemLogoutReq struct {
	Token string `header:"token" binding:"required"` // 令牌
}

//SystemAuthAdminListReq 管理员列表参数
type SystemAuthAdminListReq struct {
	Username string `form:"username"`        // 账号
	Nickname string `form:"nickname"`        // 昵称
	Role     int    `form:"role,default=-1"` // 角色ID
}

//SystemAuthAdminDetailReq 管理员详情参数
type SystemAuthAdminDetailReq struct {
	ID uint `form:"id" binding:"required,gt=0"` // 主键
}
