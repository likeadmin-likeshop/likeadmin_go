package main

import (
	"github.com/gin-gonic/gin"
	"likeadmin/admin/routers"
	"likeadmin/config"
	"likeadmin/core"
	"likeadmin/core/response"
	"likeadmin/middleware"
	"net/http"
	"strconv"
	"time"
)

func main() {
	// 刷新日志缓冲
	defer core.Logger.Sync()
	// 程序结束前关闭数据库连接
	if core.DB != nil {
		db, _ := core.DB.DB()
		defer db.Close()
	}
	// 初始化gin
	gin.SetMode(config.Config.GinMode)
	router := gin.New()
	// 设置静态路径
	router.Static(config.Config.PublicPrefix, config.Config.UploadDirectory)
	// 设置中间件
	router.Use(gin.Logger(), middleware.ErrorRecover(), middleware.TokenAuth())
	// 特殊异常处理
	router.NoMethod(response.NoMethod)
	router.NoRoute(response.NoRoute)
	// 配置路由
	group := router.Group("/api")
	core.RegisterGroup(group, routers.Group, nil)

	// 运行服务
	s := &http.Server{
		Addr:           ":" + strconv.Itoa(config.Config.ServerPort),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe().Error()
}
