package resp

import (
	"likeadmin/core"
)

//DbTablesResp 数据表返回信息
type DbTablesResp struct {
	TableName    string              `json:"tableName" structs:"tableName"`       // 表的名称
	TableComment string              `json:"tableComment" structs:"tableComment"` // 表的描述
	CreateTime   core.OnlyRespTsTime `json:"createTime" structs:"createTime"`     // 创建时间
	UpdateTime   core.OnlyRespTsTime `json:"updateTime" structs:"updateTime"`     // 更新时间
}

//GenTableResp 生成表返回信息
type GenTableResp struct {
	ID           uint        `json:"id" structs:"id"`                     // 主键
	GenType      int         `json:"genType" structs:"genType"`           // 生成类型
	TableName    string      `json:"tableName" structs:"tableName"`       // 表名称
	TableComment string      `json:"tableComment" structs:"tableComment"` // 表描述
	CreateTime   core.TsTime `json:"createTime" structs:"createTime"`     // 创建时间
	UpdateTime   core.TsTime `json:"updateTime" structs:"updateTime"`     // 更新时间
}
