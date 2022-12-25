package response

import (
	"github.com/gin-gonic/gin"
)

//NoRoute 无路由的响应
func NoRoute(c *gin.Context) {
	Fail(c, Request404Error)
}

//NoMethod 无方法的响应
func NoMethod(c *gin.Context) {
	Fail(c, Request405Error)
}
