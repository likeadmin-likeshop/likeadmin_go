package routers

import (
	"github.com/gin-gonic/gin"
	"likeadmin/core"
	"likeadmin/core/response"
	"likeadmin/middleware"
	"likeadmin/util"
	"strings"
)

var MonitorGroup = core.Group("/monitor")

func init() {
	group := MonitorGroup
	group.AddGET("/cache", cache, middleware.RecordLog("缓存监控"))
	group.AddGET("/server", server, middleware.RecordLog("服务监控"))
}

//cache 缓存监控
func cache(c *gin.Context) {
	cmdStatsMap := util.RedisUtil.Info("commandstats")
	var stats []map[string]string
	for k, v := range cmdStatsMap {
		stats = append(stats, map[string]string{
			"name":  strings.Split(k, "_")[1],
			"value": v[strings.Index(v, "=")+1 : strings.Index(v, ",")],
		})
	}
	response.OkWithData(c, map[string]interface{}{
		"info":         util.RedisUtil.Info(),
		"commandStats": stats,
		"dbSize":       util.RedisUtil.DBSize(),
	})
}

//server 服务监控
func server(c *gin.Context) {
	response.OkWithData(c, map[string]interface{}{
		"cpu":  util.ServerUtil.GetCpuInfo(),
		"mem":  util.ServerUtil.GetMemInfo(),
		"sys":  util.ServerUtil.GetSysInfo(),
		"disk": util.ServerUtil.GetDiskInfo(),
		"go":   util.ServerUtil.GetGoInfo(),
	})
}
