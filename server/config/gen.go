package config

//GenConfig 代码生成器公共配置
var GenConfig = genConfig{
	// 基础包名
	PackageName: "like",
	// 后台包名
	AdminPackage: "like.admin",
	// 实体包名
	ModelsPackage: "like.models",
	// 是否去除表前缀
	IsRemoveTablePrefix: true,
	// 生成代码根路径
	GenRootPath: "target",
}

type genConfig struct {
	PackageName         string
	AdminPackage        string
	ModelsPackage       string
	IsRemoveTablePrefix bool
	GenRootPath         string
}
