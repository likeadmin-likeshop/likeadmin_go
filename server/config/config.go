package config

import (
	"flag"
	"github.com/spf13/viper"
	"log"
	"path"
	"runtime"
	"strconv"
)

var Config = loadConfig(".")

//envConfig 环境配置
type envConfig struct {
	RootPath               string   // 项目根目录
	GinMode                string   `mapstructure:"GIN_MODE"`        // gin运行模式
	PublicUrl              string   `mapstructure:"PUBLIC_URL"`      // 对外发布的Url
	ServerPort             int      `mapstructure:"SERVER_PORT"`     // 服务运行端口
	DisallowModify         bool     `mapstructure:"DISALLOW_MODIFY"` // 禁止修改操作 (演示功能,限制POST请求)
	PublicPrefix           string   // 资源访问前缀
	UploadDirectory        string   `mapstructure:"UPLOAD_DIRECTORY"` // 上传文件路径
	RedisUrl               string   `mapstructure:"REDIS_URL"`        // Redis源配置
	RedisPoolSize          int      // Redis连接池大小
	DatabaseUrl            string   `mapstructure:"DATABASE_URL"` // 数据源配置
	DbTablePrefix          string   // Mysql表前缀
	DbDefaultStringSize    uint     // 数据库string类型字段的默认长度
	DbMaxIdleConns         int      // 数据库空闲连接池最大值
	DbMaxOpenConns         int      // 数据库连接池最大值
	DbConnMaxLifetimeHours int16    // 连接可复用的最大时间(小时)
	Version                string   // 版本
	Secret                 string   // 系统加密字符
	StaticPath             string   // 静态资源URL路径
	StaticDirectory        string   // 静态资源本地路径
	RedisPrefix            string   // Redis键前缀
	UploadImageSize        int64    // 上传图片限制
	UploadVideoSize        int64    // 上传视频限制
	UploadImageExt         []string // 上传图片扩展
	UploadVideoExt         []string // 上传视频扩展
}

//loadConfig 加载配置
func loadConfig(envPath string) envConfig {
	var cfgPath string
	flag.StringVar(&cfgPath, "c", "", "config file envPath.")
	flag.Parse()
	if cfgPath == "" {
		viper.AddConfigPath(envPath)
		viper.SetConfigFile(".env")
	} else {
		viper.SetConfigFile(cfgPath)
	}
	viper.AutomaticEnv()
	var runPath string
	if _, filename, _, ok := runtime.Caller(0); ok {
		runPath = path.Dir(path.Dir(filename))
	}
	config := envConfig{
		RootPath: runPath,
		GinMode:  "debug",
		// 服务运行端口
		ServerPort: 8000,
		// 禁止修改操作 (演示功能,限制POST请求)
		DisallowModify: false,
		// 资源访问前缀
		PublicPrefix: "/api/uploads",
		// 上传文件路径
		UploadDirectory: "/tmp/uploads/likeadmin-go/",
		// Redis源配置
		RedisUrl:      "redis://localhost:6379",
		RedisPoolSize: 100,
		// 数据源配置
		DatabaseUrl:            "root:root@tcp(localhost:3306)/likeadmin?charset=utf8mb4&parseTime=True&loc=Local",
		DbTablePrefix:          "la_",
		DbDefaultStringSize:    256,
		DbMaxIdleConns:         10,
		DbMaxOpenConns:         100,
		DbConnMaxLifetimeHours: 2,
		// 全局配置
		// 版本
		Version: "v1.0.0",
		// 系统加密字符
		Secret: "UVTIyzCy",
		// 静态资源URL路径
		StaticPath: "/api/static",
		// 静态资源本地路径
		StaticDirectory: "static",
		// Redis键前缀
		RedisPrefix: "Like:",
		// 上传图片限制
		UploadImageSize: 1024 * 1024 * 10,
		// 上传视频限制
		UploadVideoSize: 1024 * 1024 * 30,
		// 上传图片扩展
		UploadImageExt: []string{"png", "jpg", "jpeg", "gif", "ico", "bmp"},
		// 上传视频扩展
		UploadVideoExt: []string{"mp4", "mp3", "avi", "flv", "rmvb", "mov"},
	}
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("loadConfig ReadInConfig err:", err)
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatal("loadConfig Unmarshal err:", err)
	}
	// PublicUrl未设置设置默认值
	if config.PublicUrl == "" {
		config.PublicUrl = "http://127.0.0.1:" + strconv.Itoa(config.ServerPort)
	}
	return config
}
