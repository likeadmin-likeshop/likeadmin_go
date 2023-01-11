package req

import (
	"time"
)

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

//SystemAuthRoleAddReq 新增角色参数
type SystemAuthRoleAddReq struct {
	Name      string `form:"name" binding:"required,min=1,max=30"` // 角色名称
	Sort      int    `form:"sort" binding:"gte=0"`                 // 角色排序
	IsDisable uint8  `form:"isDisable" binding:"oneof=0 1"`        // 是否禁用: [0=否, 1=是]
	Remark    string `form:"remark" binding:"max=200"`             // 角色备注
	MenuIds   string `form:"menuIds"`                              // 关联菜单
}

//SystemAuthRoleEditReq 编辑角色参数
type SystemAuthRoleEditReq struct {
	ID        uint   `form:"id" binding:"required,gt=0"`           // 主键
	Name      string `form:"name" binding:"required,min=1,max=30"` // 角色名称
	Sort      int    `form:"sort" binding:"gte=0"`                 // 角色排序
	IsDisable uint8  `form:"isDisable" binding:"oneof=0 1"`        // 是否禁用: [0=否, 1=是]
	Remark    string `form:"remark" binding:"max=200"`             // 角色备注
	MenuIds   string `form:"menuIds"`                              // 关联菜单
}

//SystemAuthRoleDelReq 删除角色参数
type SystemAuthRoleDelReq struct {
	ID uint `form:"id" binding:"required,gt=0"` // 主键
}

//SystemAuthMenuDetailReq 菜单详情参数
type SystemAuthMenuDetailReq struct {
	ID uint `form:"id" binding:"required,gt=0"` // 主键
}

//SystemAuthMenuAddReq 新增菜单参数
type SystemAuthMenuAddReq struct {
	Pid       uint   `form:"pid" binding:"required,gt=0"`              // 上级菜单
	MenuType  string `form:"menuType" binding:"oneof=M C A"`           // 权限类型: [M=目录, C=菜单, A=按钮]
	MenuName  string `form:"menuName" binding:"required,min=1,max=30"` // 菜单名称
	MenuIcon  string `form:"menuIcon" binding:"max=100"`               // 菜单图标
	MenuSort  int    `form:"menuSort" binding:"required,gte=0"`        // 菜单排序
	Perms     string `form:"perms" binding:"max=100"`                  // 权限标识
	Paths     string `form:"paths" binding:"max=200"`                  // 路由地址
	Component string `form:"component" binding:"max=200"`              // 前端组件
	Selected  string `form:"selected" binding:"max=200"`               // 选中路径
	Params    string `form:"params" binding:"max=200"`                 // 路由参数
	IsCache   uint8  `form:"isCache" binding:"oneof=0 1"`              // 是否缓存: [0=否, 1=是]
	IsShow    uint8  `form:"isShow" binding:"oneof=0 1"`               // 是否显示: [0=否, 1=是]
	IsDisable uint8  `form:"isDisable" binding:"oneof=0 1"`            // 是否禁用: [0=否, 1=是]
}

//SystemAuthMenuEditReq 编辑菜单参数
type SystemAuthMenuEditReq struct {
	ID        uint   `form:"id" binding:"required,gt=0"`               // 主键
	Pid       uint   `form:"pid" binding:"required,gt=0"`              // 上级菜单
	MenuType  string `form:"menuType" binding:"oneof=M C A"`           // 权限类型: [M=目录, C=菜单, A=按钮]
	MenuName  string `form:"menuName" binding:"required,min=1,max=30"` // 菜单名称
	MenuIcon  string `form:"menuIcon" binding:"max=100"`               // 菜单图标
	MenuSort  int    `form:"menuSort" binding:"required,gte=0"`        // 菜单排序
	Perms     string `form:"perms" binding:"max=100"`                  // 权限标识
	Paths     string `form:"paths" binding:"max=200"`                  // 路由地址
	Component string `form:"component" binding:"max=200"`              // 前端组件
	Selected  string `form:"selected" binding:"max=200"`               // 选中路径
	Params    string `form:"params" binding:"max=200"`                 // 路由参数
	IsCache   uint8  `form:"isCache" binding:"oneof=0 1"`              // 是否缓存: [0=否, 1=是]
	IsShow    uint8  `form:"isShow" binding:"oneof=0 1"`               // 是否显示: [0=否, 1=是]
	IsDisable uint8  `form:"isDisable" binding:"oneof=0 1"`            // 是否禁用: [0=否, 1=是]
}

//SystemAuthMenuDelReq 删除菜单参数
type SystemAuthMenuDelReq struct {
	ID uint `form:"id" binding:"required,gt=0"` // 主键
}

//SystemAuthDeptListReq 部门列表参数
type SystemAuthDeptListReq struct {
	Name   string `form:"name"`                                     // 部门名称
	IsStop int8   `form:"isStop,default=-1" binding:"oneof=-1 0 1"` // 是否停用: [0=否, 1=是]
}

//SystemAuthDeptDetailReq 部门详情参数
type SystemAuthDeptDetailReq struct {
	ID uint `form:"id" binding:"required,gt=0"` // 主键
}

//SystemAuthDeptAddReq 部门新增参数
type SystemAuthDeptAddReq struct {
	Pid    uint   `form:"pid" binding:"gte=0"`                   // 部门父级
	Name   string `form:"name" binding:"required,min=1,max=100"` // 部门名称
	Duty   string `form:"duty" binding:"omitempty,min=1,max=30"` // 负责人
	Mobile string `form:"mobile" binding:"omitempty,len=11"`     // 联系电话
	IsStop uint8  `form:"isStop" binding:"oneof=0 1"`            // 是否停用: [0=否, 1=是]
	Sort   int    `form:"sort" binding:"gte=0,lte=9999"`         // 排序编号
}

//SystemAuthDeptEditReq 部门编辑参数
type SystemAuthDeptEditReq struct {
	ID     uint   `form:"id" binding:"required,gt=0"`            // 主键
	Pid    uint   `form:"pid" binding:"gte=0"`                   // 部门父级
	Name   string `form:"name" binding:"required,min=1,max=100"` // 部门名称
	Duty   string `form:"duty" binding:"omitempty,min=1,max=30"` // 负责人
	Mobile string `form:"mobile" binding:"omitempty,len=11"`     // 联系电话
	IsStop uint8  `form:"isStop" binding:"oneof=0 1"`            // 是否停用: [0=否, 1=是]
	Sort   int    `form:"sort" binding:"gte=0,lte=9999"`         // 排序编号
}

//SystemAuthDeptDelReq 部门删除参数
type SystemAuthDeptDelReq struct {
	ID uint `form:"id" binding:"required,gt=0"` // 主键
}

//SystemAuthPostListReq 岗位列表参数
type SystemAuthPostListReq struct {
	Code   string `form:"code"`                                     // 岗位编码
	Name   string `form:"name"`                                     // 岗位名称
	IsStop int8   `form:"isStop,default=-1" binding:"oneof=-1 0 1"` // 是否停用: [0=否, 1=是]
}

//SystemAuthPostDetailReq 岗位详情参数
type SystemAuthPostDetailReq struct {
	ID uint `form:"id" binding:"required,gt=0"` // 主键
}

//SystemAuthPostAddReq 岗位新增参数
type SystemAuthPostAddReq struct {
	Code    string `form:"code" binding:"omitempty,min=1,max=30"` // 岗位编码
	Name    string `form:"name" binding:"required,min=1,max=30"`  // 岗位名称
	Remarks string `form:"remarks" binding:"max=250"`             // 岗位备注
	IsStop  uint8  `form:"isStop" binding:"oneof=0 1"`            // 是否停用: [0=否, 1=是]
	Sort    int    `form:"sort" binding:"gte=0"`                  // 排序编号
}

//SystemAuthPostEditReq 岗位编辑参数
type SystemAuthPostEditReq struct {
	ID      uint   `form:"id" binding:"required,gt=0"`            // 主键
	Code    string `form:"code" binding:"omitempty,min=1,max=30"` // 岗位编码
	Name    string `form:"name" binding:"required,min=1,max=30"`  // 岗位名称
	Remarks string `form:"remarks" binding:"max=250"`             // 岗位备注
	IsStop  uint8  `form:"isStop" binding:"oneof=0 1"`            // 是否停用: [0=否, 1=是]
	Sort    int    `form:"sort" binding:"gte=0"`                  // 排序编号
}

//SystemAuthPostDelReq 岗位删除参数
type SystemAuthPostDelReq struct {
	ID uint `form:"id" binding:"required,gt=0"` // 主键
}

//SystemLogOperateReq 操作日志列表参数
type SystemLogOperateReq struct {
	Title     string    `form:"title"`                                       // 操作标题
	Username  string    `form:"username"`                                    // 用户账号
	Ip        string    `form:"ip"`                                          // 请求IP
	Type      string    `form:"type" binding:"omitempty,oneof=GET POST PUT"` // 请求类型: GET/POST/PUT
	Status    int       `form:"status" binding:"omitempty,oneof=1 2"`        // 执行状态: [1=成功, 2=失败]
	Url       string    `form:"url"`                                         // 请求地址
	StartTime time.Time `form:"startTime" time_format:"2006-01-02"`          // 开始时间
	EndTime   time.Time `form:"endTime" time_format:"2006-01-02"`            // 结束时间
}

//SystemLogLoginReq 登录日志列表参数
type SystemLogLoginReq struct {
	Username  string    `form:"username"`                             // 登录账号
	Status    int       `form:"status" binding:"omitempty,oneof=1 2"` // 执行状态: [1=成功, 2=失败]
	StartTime time.Time `form:"startTime" time_format:"2006-01-02"`   // 开始时间
	EndTime   time.Time `form:"endTime" time_format:"2006-01-02"`     // 结束时间
}
