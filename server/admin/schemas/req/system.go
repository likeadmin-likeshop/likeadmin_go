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

//SystemAuthAdminAddReq 管理员新增参数
type SystemAuthAdminAddReq struct {
	DeptId       uint   `form:"deptId" binding:"required,gt=0"`           // 部门ID
	PostId       uint   `form:"postId" binding:"required,gt=0"`           // 岗位ID
	Username     string `form:"username" binding:"required,min=2,max=20"` // 账号
	Nickname     string `form:"nickname" binding:"required,min=2,max=30"` // 昵称
	Password     string `form:"password" binding:"required"`              // 密码
	Avatar       string `form:"avatar" binding:"required"`                // 头像
	Role         uint   `form:"role" binding:"required,gt=0"`             // 角色
	Sort         int    `form:"sort" binding:"required,gte=0"`            // 排序
	IsDisable    uint8  `form:"isDisable" binding:"oneof=0 1"`            // 是否禁用: [0=否, 1=是]
	IsMultipoint uint8  `form:"isMultipoint" binding:"oneof=0 1"`         // 多端登录: [0=否, 1=是]
}

//SystemAuthAdminEditReq 管理员编辑参数
type SystemAuthAdminEditReq struct {
	ID           uint   `form:"id" binding:"required,gt=0"`               // 主键
	DeptId       uint   `form:"deptId" binding:"required,gt=0"`           // 部门ID
	PostId       uint   `form:"postId" binding:"required,gt=0"`           // 岗位ID
	Username     string `form:"username" binding:"required,min=2,max=20"` // 账号
	Nickname     string `form:"nickname" binding:"required,min=2,max=30"` // 昵称
	Password     string `form:"password" binding:"required"`              // 密码
	Avatar       string `form:"avatar"`                                   // 头像
	Role         uint   `form:"role" binding:"required,gt=0"`             // 角色
	Sort         int    `form:"sort" binding:"required,gte=0"`            // 排序
	IsDisable    uint8  `form:"isDisable" binding:"oneof=0 1"`            // 是否禁用: [0=否, 1=是]
	IsMultipoint uint8  `form:"isMultipoint" binding:"oneof=0 1"`         // 多端登录: [0=否, 1=是]
}

//SystemAuthAdminUpdateReq 管理员更新参数
type SystemAuthAdminUpdateReq struct {
	Nickname     string `form:"nickname" binding:"required,min=2,max=30"`     // 昵称
	Avatar       string `form:"avatar"`                                       // 头像
	Password     string `form:"password" binding:"required"`                  // 密码
	CurrPassword string `form:"currPassword" binding:"required,min=6,max=32"` // 密码
}

//SystemAuthAdminDelReq 管理员删除参数
type SystemAuthAdminDelReq struct {
	ID uint `form:"id" binding:"required,gt=0"` // 主键
}

//SystemAuthAdminDisableReq 管理员状态切换参数
type SystemAuthAdminDisableReq struct {
	ID uint `form:"id" binding:"required,gt=0"` // 主键
}

//SystemAuthRoleDetailReq 角色详情参数
type SystemAuthRoleDetailReq struct {
	ID uint `form:"id" binding:"required,gt=0"` // 主键
}
