package resp

import "likeadmin/core"

//SystemLoginResp 系统登录返回信息
type SystemLoginResp struct {
	Token string `json:"token"`
}

//SystemAuthRoleResp 系统角色返回信息
type SystemAuthRoleResp struct {
	ID         uint        `json:"id" structs:"id"`                 // 主键
	Name       string      `json:"name" structs:"name"`             // 角色名称
	Remark     string      `json:"remark" structs:"remark"`         // 角色备注
	Menus      []uint      `json:"menus" structs:"menus"`           // 关联菜单
	Member     int64       `json:"member" structs:"member"`         // 成员数量
	Sort       uint16      `json:"sort" structs:"sort"`             // 角色排序
	IsDisable  uint8       `json:"isDisable" structs:"isDisable"`   // 是否禁用: [0=否, 1=是]
	CreateTime core.TsTime `json:"createTime" structs:"createTime"` // 创建时间
	UpdateTime core.TsTime `json:"updateTime" structs:"updateTime"` // 更新时间
}

//SystemAuthMenuResp 系统菜单返回信息
type SystemAuthMenuResp struct {
	ID         uint                 `json:"id" structs:"id"`                       // 主键
	Pid        uint                 `json:"pid" structs:"pid"`                     // 上级菜单
	MenuType   string               `json:"menuType" structs:"menuType"`           // 权限类型: [M=目录, C=菜单, A=按钮]
	MenuName   string               `json:"menuName" structs:"menuName"`           // 菜单名称
	MenuIcon   string               `json:"menuIcon" structs:"menuIcon"`           // 菜单图标
	MenuSort   uint16               `json:"menuSort" structs:"menuSort"`           // 菜单排序
	Perms      string               `json:"perms" structs:"perms"`                 // 权限标识
	Paths      string               `json:"paths" structs:"paths"`                 // 路由地址
	Component  string               `json:"component" structs:"component"`         // 前端组件
	Selected   string               `json:"selected" structs:"selected"`           // 选中路径
	Params     string               `json:"params" structs:"params"`               // 路由参数
	IsCache    uint8                `json:"isCache" structs:"isCache"`             // 是否缓存: [0=否, 1=是]
	IsShow     uint8                `json:"isShow" structs:"isShow"`               // 是否显示: [0=否, 1=是]
	IsDisable  uint8                `json:"isDisable" structs:"isDisable"`         // 是否禁用: [0=否, 1=是]
	CreateTime core.TsTime          `json:"createTime" structs:"createTime"`       // 创建时间
	UpdateTime core.TsTime          `json:"updateTime" structs:"updateTime"`       // 更新时间
	Children   []SystemAuthMenuResp `json:"children,omitempty" structs:"children"` // 子集
}
