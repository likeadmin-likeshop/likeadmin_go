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
