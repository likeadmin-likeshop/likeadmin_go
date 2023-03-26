package generator

import (
	"gorm.io/gorm"
)

var GenUtil = genUtil{}

//genUtil 代码生成工具
type genUtil struct{}

//GetDbTablesQuery 查询库中的数据表
func (gu genUtil) GetDbTablesQuery(db *gorm.DB, tableName string, tableComment string) *gorm.DB {
	whereStr := ""
	if tableName != "" {
		whereStr += `lower(table_name) like lower("%` + tableName + `%")`
	}
	if tableComment != "" {
		whereStr += `lower(table_comment) like lower("%` + tableComment + `%")`
	}
	query := db.Table("information_schema.tables").Where(
		`table_schema = (SELECT database()) 
			AND table_name NOT LIKE "qrtz_%" 
			AND table_name NOT LIKE "gen_%" 
			AND table_name NOT IN (select table_name from la_gen_table) ` + whereStr).Select(
		"table_name, table_comment, create_time, update_time")
	return query
}
