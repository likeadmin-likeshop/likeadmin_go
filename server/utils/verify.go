package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"likeadmin/core/response"
)

var VerifyUtil = verifyUtil{}

//verifyUtil 参数验证工具类
type verifyUtil struct{}

func (vu verifyUtil) VerifyJSON(c *gin.Context, obj any) {
	if err := c.ShouldBindBodyWith(obj, binding.JSON); err != nil {
		panic(response.ParamsValidError.MakeData(err.Error()))
		return
	}
}

func (vu verifyUtil) VerifyHeader(c *gin.Context, obj any) {
	if err := c.ShouldBindHeader(obj); err != nil {
		panic(response.ParamsValidError.MakeData(err.Error()))
		return
	}
}

func (vu verifyUtil) VerifyQuery(c *gin.Context, obj any) {
	if err := c.ShouldBindQuery(obj); err != nil {
		panic(response.ParamsValidError.MakeData(err.Error()))
		return
	}
}
