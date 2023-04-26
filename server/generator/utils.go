package generator

import (
	"gorm.io/gorm"
	"likeadmin/config"
	"likeadmin/model/gen"
	"likeadmin/util"
	"strconv"
	"strings"
	"time"
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

//GetDbTablesQueryByNames 根据表名集查询表
func (gu genUtil) GetDbTablesQueryByNames(db *gorm.DB, tableNames []string) *gorm.DB {
	query := db.Table("information_schema.tables").Where(
		`table_schema = (SELECT database()) 
			AND table_name NOT LIKE "qrtz_%" 
			AND table_name NOT LIKE "gen_%" 
			AND table_name IN ?`, tableNames).Select(
		"table_name, table_comment, create_time, update_time")
	return query
}

//GetDbTableColumnsQueryByName 根据表名查询列信息
func (gu genUtil) GetDbTableColumnsQueryByName(db *gorm.DB, tableName string) *gorm.DB {
	query := db.Table("information_schema.columns").Where(
		`table_schema = (SELECT database()) 
			AND table_name = ?`, tableName).Order("ordinal_position").Select(
		`column_name, 
			(CASE WHEN (is_nullable = "no" && column_key != "PRI") THEN "1" ELSE NULL END) AS is_required,
			(CASE WHEN column_key = "PRI" THEN "1" ELSE "0" END) AS is_pk,
			ordinal_position AS sort, column_comment,
			(CASE WHEN extra = "auto_increment" THEN "1" ELSE "0" END) AS is_increment, column_type`)
	return query
}

//InitTable 初始化表
func (gu genUtil) InitTable(table gen.GenTable) gen.GenTable {
	return gen.GenTable{
		TableName:    table.TableName,
		TableComment: table.TableComment,
		AuthorName:   "",
		EntityName:   gu.ToClassName(table.TableName),
		ModuleName:   gu.ToModuleName(table.TableName),
		FunctionName: strings.Replace(table.TableComment, "表", "", -1),
		CreateTime:   time.Now().Unix(),
		UpdateTime:   time.Now().Unix(),
	}
}

//InitColumn 初始化字段列
func (gu genUtil) InitColumn(tableId uint, column gen.GenTableColumn) gen.GenTableColumn {
	columnType := gu.GetDbType(column.ColumnType)
	columnLen := gu.GetColumnLength(column.ColumnType)
	col := gen.GenTableColumn{
		TableID:       tableId,
		ColumnName:    column.ColumnName,
		ColumnComment: column.ColumnComment,
		ColumnType:    columnType,
		ColumnLength:  columnLen,
		JavaField:     column.ColumnName,
		JavaType:      GoConstants.TypeString,
		QueryType:     GenConstants.QueryEq,
		Sort:          column.Sort,
		IsPk:          column.IsPk,
		IsIncrement:   column.IsIncrement,
		IsRequired:    column.IsRequired,
		CreateTime:    time.Now().Unix(),
		UpdateTime:    time.Now().Unix(),
	}
	if util.ToolsUtil.Contains(append(SqlConstants.ColumnTypeStr, SqlConstants.ColumnTypeText...), columnType) {
		//文本域组
		if columnLen >= 500 || util.ToolsUtil.Contains(SqlConstants.ColumnTypeText, columnType) {
			col.HtmlType = HtmlConstants.HtmlTextarea
		} else {
			col.HtmlType = HtmlConstants.HtmlInput
		}
	} else if util.ToolsUtil.Contains(SqlConstants.ColumnTypeTime, columnType) {
		//日期字段
		col.JavaType = GoConstants.TypeDate
		col.HtmlType = HtmlConstants.HtmlDatetime
	} else if util.ToolsUtil.Contains(SqlConstants.ColumnTimeName, columnType) {
		//时间字段
		col.JavaType = GoConstants.TypeInt
		col.HtmlType = HtmlConstants.HtmlDatetime
	} else if util.ToolsUtil.Contains(SqlConstants.ColumnTypeNumber, columnType) {
		//数字字段
		col.HtmlType = HtmlConstants.HtmlInput
		if strings.Contains(columnType, ",") {
			col.JavaType = GoConstants.TypeFloat
		} else {
			col.JavaType = GoConstants.TypeInt
		}
	}
	//非必填字段
	if util.ToolsUtil.Contains(SqlConstants.ColumnNameNotEdit, col.ColumnName) {
		col.IsRequired = 0
	}
	//需插入字段
	if !util.ToolsUtil.Contains(SqlConstants.ColumnNameNotAdd, col.ColumnName) {
		col.IsInsert = GenConstants.Require
	}
	//需编辑字段
	if !util.ToolsUtil.Contains(SqlConstants.ColumnNameNotEdit, col.ColumnName) {
		col.IsEdit = GenConstants.Require
		col.IsRequired = GenConstants.Require
	}
	//需列表字段
	if !util.ToolsUtil.Contains(SqlConstants.ColumnNameNotList, col.ColumnName) && col.IsPk == 0 {
		col.IsList = GenConstants.Require
	}
	//需查询字段
	if !util.ToolsUtil.Contains(SqlConstants.ColumnNameNotQuery, col.ColumnName) && col.IsPk == 0 {
		col.IsQuery = GenConstants.Require
	}
	lowerColName := strings.ToLower(col.ColumnName)
	//模糊查字段
	if strings.HasSuffix(lowerColName, "name") || util.ToolsUtil.Contains([]string{"title", "mobile"}, lowerColName) {
		col.QueryType = GenConstants.QueryLike
	}
	//根据字段设置
	if strings.HasSuffix(lowerColName, "status") || util.ToolsUtil.Contains([]string{"is_show", "is_disable"}, lowerColName) {
		//状态字段设置单选框
		col.HtmlType = HtmlConstants.HtmlRadio
	} else if strings.HasSuffix(lowerColName, "type") || strings.HasSuffix(lowerColName, "sex") {
		//类型&性别字段设置下拉框
		col.HtmlType = HtmlConstants.HtmlSelect
	} else if strings.HasSuffix(lowerColName, "image") {
		//图片字段设置图片上传
		col.HtmlType = HtmlConstants.HtmlImageUpload
	} else if strings.HasSuffix(lowerColName, "file") {
		//文件字段设置文件上传
		col.HtmlType = HtmlConstants.HtmlFileUpload
	} else if strings.HasSuffix(lowerColName, "content") {
		//富文本字段设置富文本编辑器
		col.HtmlType = HtmlConstants.HtmlEditor
	}
	return col
}

//ToModuleName 表名转业务名
func (gu genUtil) ToModuleName(name string) string {
	names := strings.Split(name, "_")
	return names[len(names)-1]
}

//ToClassName 表名转类名
func (gu genUtil) ToClassName(name string) string {
	tablePrefix := config.Config.DbTablePrefix
	if config.GenConfig.IsRemoveTablePrefix && tablePrefix != "" {
		if strings.HasPrefix(name, tablePrefix) {
			name = name[len(tablePrefix):]
		}
	}
	return util.StringUtil.ToCamelCase(name)
}

//GetDbType 获取数据库类型字段
func (gu genUtil) GetDbType(columnType string) string {
	index := strings.IndexRune(columnType, '(')
	if index < 0 {
		return columnType
	}
	return columnType[:index]
}

//GetColumnLength 获取字段长度
func (gu genUtil) GetColumnLength(columnType string) int {
	index := strings.IndexRune(columnType, '(')
	if index < 0 {
		return 0
	}
	length, err := strconv.Atoi(columnType[index+1 : strings.IndexRune(columnType, ')')])
	if err != nil {
		return 0
	}
	return length
}

//GetTablePriCol 获取主键列名称
func (gu genUtil) GetTablePriCol(columns []gen.GenTableColumn) (res gen.GenTableColumn) {
	for _, col := range columns {
		if col.IsPk == 1 {
			res = col
			return
		}
	}
	return
}
