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
