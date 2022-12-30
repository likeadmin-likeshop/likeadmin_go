package core

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"likeadmin/config"
	"log"
	"os"
	"time"
)

var DB = initMysql()

//initMysql 初始化mysql会话
func initMysql() *gorm.DB {
	// 日志配置
	slowThreshold := time.Second
	ignoreRecordNotFoundError := true
	if config.Config.GinMode == "debug" {
		slowThreshold = 200 * time.Millisecond
		ignoreRecordNotFoundError = false
	}
	logger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             slowThreshold,             // 慢 SQL 阈值
			LogLevel:                  logger.Warn,               // 日志级别
			IgnoreRecordNotFoundError: ignoreRecordNotFoundError, // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,                      // 彩色打印
		},
	)
	// 初始化会话
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       config.Config.DatabaseUrl,         // DSN data source name
		DefaultStringSize:         config.Config.DbDefaultStringSize, // string 类型字段的默认长度
		SkipInitializeWithVersion: false,                             // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		SkipDefaultTransaction: true, // 禁用默认事务
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   config.Config.DbTablePrefix, // 表名前缀
			SingularTable: true,                        // 使用单一表名, eg. `User` => `user`
		},
		DisableForeignKeyConstraintWhenMigrating: true,   // 禁用自动创建外键约束
		Logger:                                   logger, // 自定义Logger
	})
	if err != nil {
		log.Fatal("initMysql gorm.Open err:", err)
	}
	db.InstanceSet("gorm:table_options", "ENGINE=InnoDB")
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("initMysql db.DB err:", err)
	}
	// 数据库空闲连接池最大值
	sqlDB.SetMaxIdleConns(config.Config.DbMaxIdleConns)
	// 数据库连接池最大值
	sqlDB.SetMaxOpenConns(config.Config.DbMaxOpenConns)
	// 连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Duration(config.Config.DbConnMaxLifetimeHours) * time.Hour)
	return db
}

func DBTableName(model interface{}) string {
	stmt := &gorm.Statement{DB: DB}
	stmt.Parse(model)
	return stmt.Schema.Table
}
