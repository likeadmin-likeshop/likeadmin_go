package gen

import (
	"github.com/gin-gonic/gin"
	"likeadmin/core"
	"likeadmin/core/request"
	"likeadmin/core/response"
	"likeadmin/generator/schemas/req"
	"likeadmin/generator/service/gen"
	"likeadmin/middleware"
	"likeadmin/util"
)

var GenGroup = core.Group("/gen", newGenHandler, regGen, middleware.TokenAuth())

func newGenHandler(srv gen.IGenerateService) *genHandler {
	return &genHandler{srv: srv}
}

func regGen(rg *gin.RouterGroup, group *core.GroupBase) error {
	return group.Reg(func(handle *genHandler) {
		rg.GET("/db", handle.dbTables)
	})
}

type genHandler struct {
	srv gen.IGenerateService
}

//dbTables 数据表列表
func (gh genHandler) dbTables(c *gin.Context) {
	var page request.PageReq
	var tbReq req.DbTablesReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &tbReq)) {
		return
	}
	res, err := gh.srv.DbTables(page, tbReq)
	response.CheckAndRespWithData(c, res, err)
}
