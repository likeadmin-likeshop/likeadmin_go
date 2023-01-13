package middleware

import (
	"github.com/gin-gonic/gin"
	"likeadmin/config"
	"likeadmin/core/response"
	"likeadmin/util"
	"strings"
)

//ShowMode 演示模式中间件，演示模式禁止POST
func ShowMode() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 路由转权限
		auths := strings.ReplaceAll(strings.Replace(c.Request.URL.Path, "/api/", "", 1), "/", ":")
		// 禁止修改操作 (演示功能,限制POST请求)
		if c.Request.Method == "POST" && !util.ToolsUtil.Contains(config.AdminConfig.ShowWhitelistUri, auths) {
			response.FailWithMsg(c, response.NoPermission, "演示环境不支持修改数据，请下载源码本地部署体验!")
			c.Abort()
			return
		}
	}
}
