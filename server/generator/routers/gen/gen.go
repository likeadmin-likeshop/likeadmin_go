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
		rg.GET("/list", handle.List)
		rg.GET("/detail", handle.Detail)
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

//List 生成列表
func (gh genHandler) List(c *gin.Context) {
	var page request.PageReq
	var listReq req.ListTableReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := gh.srv.List(page, listReq)
	response.CheckAndRespWithData(c, res, err)
}

//Detail 生成详情
func (gh genHandler) Detail(c *gin.Context) {
	var detailReq req.DetailTableReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := gh.srv.Detail(detailReq.ID)
	response.CheckAndRespWithData(c, res, err)
}
