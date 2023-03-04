package resp

import "likeadmin/core"

//SettingDictTypeResp 字典类型返回信息
type SettingDictTypeResp struct {
	ID         uint        `json:"id" structs:"id"`                 // 主键
	DictName   string      `json:"dictName" structs:"dictName"`     // 字典名称
	DictType   string      `json:"dictType" structs:"dictType"`     // 字典类型
	DictRemark string      `json:"dictRemark" structs:"dictRemark"` // 字典备注
	DictStatus uint8       `json:"dictStatus" structs:"dictStatus"` // 字典状态
	CreateTime core.TsTime `json:"createTime" structs:"createTime"` // 创建时间
	UpdateTime core.TsTime `json:"updateTime" structs:"updateTime"` // 更新时间
}

//SettingDictDataResp 字典数据返回信息
type SettingDictDataResp struct {
	ID         uint        `json:"id" structs:"id"`                 // 主键
	TypeId     uint        `json:"typeId" structs:"typeId"`         // 类型
	Name       string      `json:"name" structs:"name"`             // 键
	Value      string      `json:"value" structs:"value"`           // 值
	Remark     string      `json:"remark" structs:"remark"`         // 备注
	Sort       uint16      `json:"sort" structs:"sort"`             // 排序
	Status     uint8       `json:"status" structs:"status"`         // 状态: [0=停用, 1=禁用]
	CreateTime core.TsTime `json:"createTime" structs:"createTime"` // 创建时间
	UpdateTime core.TsTime `json:"updateTime" structs:"updateTime"` // 更新时间
}
