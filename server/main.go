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

//initRouter 初始化router
func initRouter() *gin.Engine {
	// 初始化gin
	gin.SetMode(config.Config.GinMode)
	router := gin.New()
	// 设置静态路径
	router.Static(config.Config.PublicPrefix, config.Config.UploadDirectory)
	router.Static(config.Config.StaticPath, config.Config.StaticDirectory)
	// 设置中间件
	router.Use(gin.Logger(), middleware.Cors(), middleware.ErrorRecover())
	// 演示模式
	if config.Config.DisallowModify {
		router.Use(middleware.ShowMode())
	}
	// 特殊异常处理
	router.NoMethod(response.NoMethod)
	router.NoRoute(response.NoRoute)
	// 注册路由
	group := router.Group("/api")
	core.RegisterGroup(group, routers.CommonGroup, middleware.TokenAuth())
	core.RegisterGroup(group, routers.MonitorGroup, middleware.TokenAuth())
	core.RegisterGroup(group, routers.SettingGroup, middleware.TokenAuth())
	core.RegisterGroup(group, routers.SystemGroup, middleware.TokenAuth())
	return router
}

//initServer 初始化server
func initServer(router *gin.Engine) *http.Server {
	return &http.Server{
		Addr:           ":" + strconv.Itoa(config.Config.ServerPort),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}

func main() {
	// 刷新日志缓冲
	defer core.Logger.Sync()
	// 程序结束前关闭数据库连接
	if core.DB != nil {
		db, _ := core.DB.DB()
		defer db.Close()
	}
	// 初始化router
	router := initRouter()
	// 初始化server
	s := initServer(router)
	// 运行服务
	s.ListenAndServe().Error()
}
