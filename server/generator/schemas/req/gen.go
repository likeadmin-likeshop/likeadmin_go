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

//DetailTableReq 生成详情参数
type DetailTableReq struct {
	ID uint `form:"id" binding:"required,gt=0"` // 主键
}

//ImportTableReq 导入表结构参数
type ImportTableReq struct {
	Tables string `form:"tables" binding:"required"` // 导入的表, 用","分隔
}

//SyncTableReq 同步表结构参数
type SyncTableReq struct {
	ID uint `form:"id" binding:"required,gt=0"` // 主键
}

//EditColumn 表编辑列
type EditColumn struct {
	ID            uint   `form:"id" binding:"required,gt=0"`               // 主键
	ColumnComment string `form:"columnComment" binding:"required,max=200"` // 列描述
	JavaField     string `form:"goField" binding:"required,max=100"`       // 字段
	IsRequired    uint8  `form:"isStop" binding:"oneof=0 1"`               // 是否必填: [0=否, 1=是]
	IsInsert      uint8  `form:"isInsert" binding:"oneof=0 1"`             // 是否新增字段: [0=否, 1=是]
	IsEdit        uint8  `form:"isEdit" binding:"oneof=0 1"`               // 是否编辑字段: [0=否, 1=是]
	IsList        uint8  `form:"isList" binding:"oneof=0 1"`               // 是否列表字段: [0=否, 1=是]
	IsQuery       uint8  `form:"isQuery" binding:"oneof=0 1"`              // 是否查询字段: [0=否, 1=是]
	QueryType     string `form:"queryType" binding:"required,max=30"`      // 查询方式
	HtmlType      string `form:"htmlType" binding:"required,max=30"`       // 表单类型
	DictType      string `form:"dictType" binding:"required,max=200"`      // 字典类型
}

//EditTableReq 编辑表结构参数
type EditTableReq struct {
	ID           uint         `form:"id" binding:"required,gt=0"`                    // 主键
	TableName    string       `form:"tableName" binding:"required,min=1,max=200"`    // 表名称
	EntityName   string       `form:"entityName" binding:"required,min=1,max=200"`   // 实体名称
	TableComment string       `form:"tableComment" binding:"required,min=1,max=200"` // 表描述
	AuthorName   string       `form:"authorName" binding:"required,min=1,max=100"`   // 作者名称
	Remarks      string       `form:"remarks" binding:"max=60"`                      // 备注信息
	GenTpl       string       `form:"genTpl" binding:"oneof=crud tree"`              // 生成模板方式: [crud=单表, tree=树表]
	ModuleName   string       `form:"moduleName" binding:"required,min=1,max=60"`    // 生成模块名
	FunctionName string       `form:"functionName" binding:"required,min=1,max=60"`  // 生成功能名
	GenType      int          `form:"genType" binding:"oneof=0 1"`                   // 生成代码方式: [0=zip压缩包, 1=自定义路径]
	GenPath      string       `form:"genPath,default=/" binding:"required,max=60"`   // 生成路径
	TreePrimary  string       `form:"treePrimary"`                                   // 树表主键
	TreeParent   string       `form:"treeParent"`                                    // 树表父键
	TreeName     string       `form:"treeName"`                                      // 树表名称
	SubTableName string       `form:"subTableName"`                                  // 子表名称
	SubTableFk   string       `form:"subTableFk"`                                    // 子表外键
	Columns      []EditColumn `form:"columns" binding:"required"`                    // 字段列表
}

//DelTableReq 删除表结构参数
type DelTableReq struct {
	Ids []uint `form:"ids" binding:"required"` // 主键
}
