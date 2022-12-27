package config

import (
	"flag"
	"github.com/spf13/viper"
	"log"
)

var Config = loadConfig(".")

//envConfig 环境配置
type envConfig struct {
	GinMode                string `mapstructure:"GIN_MODE"`         // gin运行模式
	ServerPort             int    `mapstructure:"SERVER_PORT"`      // 服务运行端口
	UploadDirectory        string `mapstructure:"UPLOAD_DIRECTORY"` // 上传文件路径
	RedisUrl               string `mapstructure:"REDIS_URL"`        // Redis源配置
	RedisPoolSize          int    // Redis连接池大小
	DatabaseUrl            string `mapstructure:"DATABASE_URL"` // 数据源配置
	DbTablePrefix          string // Mysql表前缀
	DbDefaultStringSize    uint   // 数据库string类型字段的默认长度
	DbMaxIdleConns         int    // 数据库空闲连接池最大值
	DbMaxOpenConns         int    // 数据库连接池最大值
	DbConnMaxLifetimeHours int16  // 连接可复用的最大时间(小时)
	Secret                 string // 系统加密字符
	RedisPrefix            string // Redis键前缀
}

//loadConfig 加载配置
func loadConfig(path string) envConfig {
	var cfgPath string
	flag.StringVar(&cfgPath, "c", "", "config file path.")
	flag.Parse()
	if cfgPath == "" {
		viper.AddConfigPath(path)
		viper.SetConfigFile(".env")
	} else {
		viper.SetConfigFile(cfgPath)
	}
	viper.AutomaticEnv()
	config := envConfig{
		GinMode:    "debug",
		ServerPort: 8000,
		// 上传文件路径
		UploadDirectory: "/tmp/uploads/likeadmin-python/",
		// Redis源配置
		RedisUrl:      "redis://localhost:6379",
		RedisPoolSize: 100,
		// 数据源配置
		DatabaseUrl:            "mysql+pymysql://root:root@localhost:3306/likeadmin?charset=utf8mb4",
		DbTablePrefix:          "la_",
		DbDefaultStringSize:    256,
		DbMaxIdleConns:         10,
		DbMaxOpenConns:         100,
		DbConnMaxLifetimeHours: 2,
		// 全局配置
		// 系统加密字符
		Secret: "UVTIyzCy",
		// Redis键前缀
		RedisPrefix: "Like:",
	}
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("loadConfig ReadInConfig err:", err)
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatal("loadConfig Unmarshal err:", err)
	}
	return config
}
