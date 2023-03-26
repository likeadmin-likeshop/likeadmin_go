package req

//DbTablesReq 库表列表参数
type DbTablesReq struct {
	TableName    string `form:"tableName"`    // 表名称
	TableComment string `form:"tableComment"` // 表描述
}
