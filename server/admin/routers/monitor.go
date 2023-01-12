package routers

import (
	"github.com/gin-gonic/gin"
	"likeadmin/admin/service/common"
	"likeadmin/core"
	"likeadmin/core/response"
	"likeadmin/util"
)

var MonitorGroup = core.Group("/monitor")

func init() {
	group := MonitorGroup
	group.AddGET("/cache", cache)
	group.AddGET("/server", server)
}

//cache 缓存监控
func cache(c *gin.Context) {
	// TODO: cache
	response.OkWithData(c, common.IndexService.Console())
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
