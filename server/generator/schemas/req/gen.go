package req

import "time"

//DbTablesReq 库表列表参数
type DbTablesReq struct {
	TableName    string `form:"tableName"`    // 表名称
	TableComment string `form:"tableComment"` // 表描述
}

//ListTableReq 生成列表参数
type ListTableReq struct {
	TableName    string    `form:"tableName"`                          // 表名称
	TableComment string    `form:"tableComment"`                       // 表描述
	StartTime    time.Time `form:"startTime" time_format:"2006-01-02"` // 开始时间
	EndTime      time.Time `form:"endTime" time_format:"2006-01-02"`   // 结束时间
}
