package config

//GenConfig 代码生成器公共配置
var GenConfig = genConfig{
	// 基础包名
	PackageName: "gencode",
	// 是否去除表前缀
	IsRemoveTablePrefix: true,
	// 生成代码根路径
	GenRootPath: "/tmp/target",
}

type genConfig struct {
	PackageName         string
	IsRemoveTablePrefix bool
	GenRootPath         string
}
