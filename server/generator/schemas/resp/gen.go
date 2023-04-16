package resp

import (
	"likeadmin/core"
)

//DbTableResp 数据表返回信息
type DbTableResp struct {
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

//GenTableBaseResp 生成表基本返回信息
type GenTableBaseResp struct {
	ID           uint        `json:"id" structs:"id"`                     // 主键
	TableName    string      `json:"tableName" structs:"tableName"`       // 表的名称
	TableComment string      `json:"tableComment" structs:"tableComment"` // 表的描述
	EntityName   string      `json:"entityName" structs:"entityName"`     // 实体名称
	AuthorName   string      `json:"authorName" structs:"authorName"`     // 作者名称
	Remarks      string      `json:"remarks" structs:"remarks"`           // 备注信息
	CreateTime   core.TsTime `json:"createTime" structs:"createTime"`     // 创建时间
	UpdateTime   core.TsTime `json:"updateTime" structs:"updateTime"`     // 更新时间
}

//GenTableGenResp 生成表生成返回信息
type GenTableGenResp struct {
	GenTpl       string `json:"genTpl" structs:"genTpl"`             // 生成模板方式: [crud=单表, tree=树表]
	GenType      int    `json:"genType" structs:"genType"`           // 生成代码方式: [0=zip压缩包, 1=自定义路径]
	GenPath      string `json:"genPath" structs:"genPath"`           // 生成代码路径: [不填默认项目路径]
	ModuleName   string `json:"moduleName" structs:"moduleName"`     // 生成模块名
	FunctionName string `json:"functionName" structs:"functionName"` // 生成功能名
	TreePrimary  string `json:"treePrimary" structs:"treePrimary"`   // 树主键字段
	TreeParent   string `json:"treeParent" structs:"treeParent"`     // 树父级字段
	TreeName     string `json:"treeName" structs:"treeName"`         // 树显示字段
	SubTableName string `json:"subTableName" structs:"subTableName"` // 关联表名称
	SubTableFk   string `json:"subTableFk" structs:"subTableFk"`     // 关联表外键
}

//GenColumnResp 生成列返回信息
type GenColumnResp struct {
	ID            uint        `json:"id" structs:"id"`                       // 字段主键
	ColumnName    string      `json:"columnName" structs:"columnName"`       // 字段名称
	ColumnComment string      `json:"columnComment" structs:"columnComment"` // 字段描述
	ColumnLength  int         `json:"columnLength" structs:"columnLength"`   // 字段长度
	ColumnType    string      `json:"columnType" structs:"columnType"`       // 字段类型
	JavaType      string      `json:"goType" structs:"goType"`               // Go类型
	JavaField     string      `json:"goField" structs:"goField"`             // Go字段
	IsRequired    uint8       `json:"isRequired" structs:"isRequired"`       // 是否必填
	IsInsert      uint8       `json:"isInsert" structs:"isInsert"`           // 是否为插入字段
	IsEdit        uint8       `json:"isEdit" structs:"isEdit"`               // 是否编辑字段
	IsList        uint8       `json:"isList" structs:"isList"`               // 是否列表字段
	IsQuery       uint8       `json:"isQuery" structs:"isQuery"`             // 是否查询字段
	QueryType     string      `json:"queryType" structs:"queryType"`         // 查询方式: [等于、不等于、大于、小于、范围]
	HtmlType      string      `json:"htmlType" structs:"htmlType"`           // 显示类型: [文本框、文本域、下拉框、复选框、单选框、日期控件]
	DictType      string      `json:"dictType" structs:"dictType"`           // 字典类型
	CreateTime    core.TsTime `json:"createTime" structs:"createTime"`       // 创建时间
	UpdateTime    core.TsTime `json:"updateTime" structs:"updateTime"`       // 更新时间
}

//GenTableDetailResp 生成表详情返回信息
type GenTableDetailResp struct {
	Base   GenTableBaseResp `json:"base" structs:"base"`     // 基本信息
	Gen    GenTableGenResp  `json:"gen" structs:"gen"`       // 生成信息
	Column []GenColumnResp  `json:"column" structs:"column"` // 字段列表
}
