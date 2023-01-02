package system

//SystemAuthAdmin 系统管理员实体
type SystemAuthAdmin struct {
	ID            uint   `gorm:"primarykey;comment:'主键'"`
	DeptId        uint   `gorm:"not null;default:0;comment:'部门ID'"`
	PostId        uint   `gorm:"not null;default:0;comment:'岗位ID'"`
	Username      string `gorm:"not null;default:'';comment:'用户账号''"`
	Nickname      string `gorm:"not null;default:'';comment:'用户昵称'"`
	Password      string `gorm:"not null;default:'';comment:'用户密码'"`
	Avatar        string `gorm:"not null;default:'';comment:'用户头像'"`
	Role          string `gorm:"not null;default:'';comment:'角色主键'"`
	Salt          string `gorm:"not null;default:'';comment:'加密盐巴'"`
	Sort          uint16 `gorm:"not null;default:0;comment:'排序编号'"`
	IsMultipoint  uint8  `gorm:"not null;default:0;comment:'多端登录: 0=否, 1=是''"`
	IsDisable     uint8  `gorm:"not null;default:0;comment:'是否禁用: 0=否, 1=是'"`
	IsDelete      uint8  `gorm:"not null;default:0;comment:'是否删除: [0=否, 1=是]'"`
	LastLoginIp   string `gorm:"not null;default:'';comment:'最后登录IP'"`
	LastLoginTime int64  `gorm:"not null;default:0;comment:最后登录时间"`
	CreateTime    int64  `gorm:"autoCreateTime;not null;comment:创建时间"`
	UpdateTime    int64  `gorm:"autoUpdateTime;not null;comment:更新时间"`
	DeleteTime    int64  `gorm:"not null;default:0;comment:删除时间"`
}

//SystemAuthMenu 系统菜单实体
type SystemAuthMenu struct {
	ID         uint   `gorm:"primarykey;comment:'主键'"`
	Pid        uint   `gorm:"not null;default:0;comment:'上级菜单'"`
	MenuType   string `gorm:"not null;default:'';comment:'权限类型: M=目录，C=菜单，A=按钮''"`
	MenuName   string `gorm:"not null;default:'';comment:'菜单名称'"`
	MenuIcon   string `gorm:"not null;default:'';comment:'菜单图标'"`
	MenuSort   uint16 `gorm:"not null;default:0;comment:'菜单排序'"`
	Perms      string `gorm:"not null;default:'';comment:'权限标识'"`
	Paths      string `gorm:"not null;default:'';comment:'路由地址'"`
	Component  string `gorm:"not null;default:'';comment:'前端组件'"`
	Selected   string `gorm:"not null;default:'';comment:'选中路径'"`
	Params     string `gorm:"not null;default:'';comment:'路由参数'"`
	IsCache    uint8  `gorm:"not null;default:0;comment:'是否缓存: 0=否, 1=是''"`
	IsShow     uint8  `gorm:"not null;default:1;comment:'是否显示: 0=否, 1=是'"`
	IsDisable  uint8  `gorm:"not null;default:0;comment:'是否禁用: 0=否, 1=是'"`
	CreateTime int64  `gorm:"autoCreateTime;not null;comment:创建时间"`
	UpdateTime int64  `gorm:"autoUpdateTime;not null;comment:更新时间"`
}

//SystemAuthPerm 系统角色菜单实体
type SystemAuthPerm struct {
	ID     string `gorm:"primarykey;comment:'主键'"`
	RoleId uint   `gorm:"not null;default:0;comment:'角色ID'"`
	MenuId uint   `gorm:"not null;default:0;comment:'菜单ID'"`
}

//SystemAuthRole 系统角色实体
type SystemAuthRole struct {
	ID         uint   `gorm:"primarykey;comment:'主键'"`
	Name       string `gorm:"not null;default:'';comment:'角色名称''"`
	Remark     string `gorm:"not null;default:'';comment:'备注信息'"`
	IsDisable  uint8  `gorm:"not null;default:0;comment:'是否禁用: 0=否, 1=是'"`
	Sort       uint16 `gorm:"not null;default:0;comment:'角色排序'"`
	CreateTime int64  `gorm:"autoCreateTime;not null;comment:创建时间"`
	UpdateTime int64  `gorm:"autoUpdateTime;not null;comment:更新时间"`
}

//SystemAuthDept 系统部门实体
type SystemAuthDept struct {
	ID         uint   `gorm:"primarykey;comment:'主键'"`
	Pid        uint   `gorm:"not null;default:0;comment:'上级主键'"`
	Name       string `gorm:"not null;default:'';comment:'部门名称''"`
	Duty       string `gorm:"not null;default:'';comment:'负责人名'"`
	Mobile     string `gorm:"not null;default:'';comment:'联系电话'"`
	Sort       uint16 `gorm:"not null;default:0;comment:'排序编号'"`
	IsStop     uint8  `gorm:"not null;default:0;comment:'是否停用: 0=否, 1=是'"`
	IsDelete   uint8  `gorm:"not null;default:0;comment:'是否删除: 0=否, 1=是'"`
	CreateTime int64  `gorm:"autoCreateTime;not null;comment:创建时间"`
	UpdateTime int64  `gorm:"autoUpdateTime;not null;comment:更新时间"`
	DeleteTime int64  `gorm:"not null;default:0;comment:删除时间"`
}

//SystemLogLogin 系统登录日志实体
type SystemLogLogin struct {
	ID         uint   `gorm:"primarykey;comment:'主键'"`
	AdminId    uint   `gorm:"not null;default:0;comment:'管理员ID'"`
	Username   string `gorm:"not null;default:'';comment:'登录账号'"`
	Ip         string `gorm:"not null;default:'';comment:'登录地址'"`
	Os         string `gorm:"not null;default:'';comment:'操作系统'"`
	Browser    string `gorm:"not null;default:'';comment:'浏览器'"`
	Status     uint8  `gorm:"not null;default:0;comment:'操作状态: 1=成功, 0=失败'"`
	CreateTime int64  `gorm:"autoCreateTime;not null;comment:创建时间"`
}

//SystemLogOperate 系统操作日志实体
type SystemLogOperate struct {
	ID         uint   `gorm:"primarykey;comment:'主键'"`
	AdminId    uint   `gorm:"not null;default:0;comment:'操作人ID'"`
	Type       string `gorm:"not null;default:'';comment:'请求类型: GET/POST/PUT'"`
	Title      string `gorm:"default:'';comment:'操作标题'"`
	Ip         string `gorm:"not null;default:'';comment:'请求IP'"`
	Url        string `gorm:"not null;default:'';comment:'请求接口'"`
	Method     string `gorm:"not null;default:'';comment:'请求方法'"`
	Args       string `gorm:"comment:'请求参数'"`
	Error      string `gorm:"comment:'错误信息'"`
	Status     uint8  `gorm:"not null;default:0;comment:'执行状态: 1=成功, 2=失败'"`
	StartTime  int64  `gorm:"not null;default:0;comment:开始时间"`
	EndTime    int64  `gorm:"not null;default:0;comment:结束时间"`
	TaskTime   int64  `gorm:"not null;default:0;comment:执行耗时"`
	CreateTime int64  `gorm:"autoCreateTime;not null;comment:创建时间"`
}
