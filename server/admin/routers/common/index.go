package common

import (
	"github.com/gin-gonic/gin"
	"likeadmin/admin/service/common"
	"likeadmin/core"
	"likeadmin/core/response"
	"likeadmin/middleware"
)

var IndexGroup = core.Group("/common", newIndexHandler, regIndex, middleware.TokenAuth())

func newIndexHandler(srv common.IIndexService) *indexHandler {
	return &indexHandler{srv: srv}
}

func regIndex(rg *gin.RouterGroup, group *core.GroupBase) error {
	return group.Reg(func(handle *indexHandler) {
		rg.GET("/index/console", handle.console)
		rg.GET("/index/config", handle.config)
	})
}

type indexHandler struct {
	srv common.IIndexService
}

//console 控制台
func (ih indexHandler) console(c *gin.Context) {
	res, err := ih.srv.Console()
	response.CheckAndRespWithData(c, res, err)
}

//config 公共配置
func (ih indexHandler) config(c *gin.Context) {
	res, err := ih.srv.Config()
	response.CheckAndRespWithData(c, res, err)
}
