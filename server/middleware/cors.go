package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

//Cors CORS（跨域资源共享）中间件
func Cors() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
		AllowMethods: []string{"OPTIONS", "GET", "POST", "POST", "DELETE", "PUT"},
		MaxAge:       1 * time.Hour,
	})
}
