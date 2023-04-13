package gen

//GenTable 代码生成业务实体
type GenTable struct {
	ID           uint   `gorm:"primarykey;comment:'主键'"`
	TableName    string `gorm:"not null;default:'';comment:'表名称''"`
	TableComment string `gorm:"not null;default:'';comment:'表描述'"`
	SubTableName string `gorm:"not null;default:'';comment:'关联表名称'"`
	SubTableFk   string `gorm:"not null;default:'';comment:'关联表外键'"`
	AuthorName   string `gorm:"not null;default:'';comment:'作者的名称'"`
	EntityName   string `gorm:"not null;default:'';comment:'实体的名称'"`
	ModuleName   string `gorm:"not null;default:'';comment:'生成模块名'"`
	FunctionName string `gorm:"not null;default:'';comment:'生成功能名'"`
	TreePrimary  string `gorm:"not null;default:'';comment:'树主键字段'"`
	TreeParent   string `gorm:"not null;default:'';comment:'树父级字段'"`
	TreeName     string `gorm:"not null;default:'';comment:'树显示字段'"`
	GenTpl       string `gorm:"not null;default:'crud';comment:'生成模板方式: [crud=单表, tree=树表]'"`
	GenType      int    `gorm:"not null;default:0;comment:'生成代码方式: [0=zip压缩包, 1=自定义路径]'"`
	GenPath      string `gorm:"not null;default:'/';comment:'生成代码路径: [不填默认项目路径]'"`
	Remarks      string `gorm:"not null;default:'';comment:'备注信息'"`
	CreateTime   int64  `gorm:"autoCreateTime;not null;comment:'创建时间'"`
	UpdateTime   int64  `gorm:"autoUpdateTime;not null;comment:'更新时间'"`
}

//GenTableColumn 代码生成表列实体
type GenTableColumn struct {
	ID            uint   `gorm:"primarykey;comment:'列主键'"`
	TableID       uint   `gorm:"not null;default:0;comment:'表外键'"`
	ColumnName    string `gorm:"not null;default:'';comment:'列名称'"`
	ColumnComment string `gorm:"not null;default:'';comment:'列描述'"`
	ColumnLength  int    `gorm:"not null;default:0;comment:'列长度'"`
	ColumnType    string `gorm:"not null;default:'';comment:'列类型'"`
	JavaType      string `gorm:"not null;default:'';comment:'类型'"`
	JavaField     string `gorm:"not null;default:'';comment:'字段名'"`
	IsPk          uint8  `gorm:"not null;default:0;comment:'是否主键: [1=是, 0=否]'"`
	IsIncrement   uint8  `gorm:"not null;default:0;comment:'是否自增: [1=是, 0=否]'"`
	IsRequired    uint8  `gorm:"not null;default:0;comment:'是否必填: [1=是, 0=否]'"`
	IsInsert      uint8  `gorm:"not null;default:0;comment:'是否为插入字段: [1=是, 0=否]'"`
	IsEdit        uint8  `gorm:"not null;default:0;comment:'是否编辑字段: [1=是, 0=否]'"`
	IsList        uint8  `gorm:"not null;default:0;comment:'是否列表字段: [1=是, 0=否]'"`
	IsQuery       uint8  `gorm:"not null;default:0;comment:'是否查询字段: [1=是, 0=否]'"`
	QueryType     string `gorm:"not null;default:'=';comment:'查询方式: [等于、不等于、大于、小于、范围]'"`
	HtmlType      string `gorm:"not null;default:'';comment:'显示类型: [文本框、文本域、下拉框、复选框、单选框、日期控件]'"`
	DictType      string `gorm:"not null;default:'';comment:'字典类型'"`
	Sort          int    `gorm:"not null;default:0;comment:'排序编号'"`
	CreateTime    int64  `gorm:"autoCreateTime;not null;comment:'创建时间'"`
	UpdateTime    int64  `gorm:"autoUpdateTime;not null;comment:'更新时间'"`
}
