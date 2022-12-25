package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"likeadmin/core"
	"likeadmin/core/response"
	"runtime/debug"
)

//ErrorRecover 异常恢复中间件
func ErrorRecover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				switch v := r.(type) {
				// 自定义类型
				case response.RespType:
					core.Logger.WithOptions(zap.AddCallerSkip(2)).Warnf(
						"Request Fail by recover: url=[%s], resp=[%+v]", c.Request.URL.Path, v)
					response.Result(c, v, []string{})
				// 其他类型
				default:
					core.Logger.Errorf("stacktrace from panic: %+v\n%s", r, string(debug.Stack()))
					response.Fail(c, response.SystemError)
				}
				c.Abort()
			}
		}()
		c.Next()
	}
}
