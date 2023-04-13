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
	"strings"
)

var GenGroup = core.Group("/gen", newGenHandler, regGen, middleware.TokenAuth())

func newGenHandler(srv gen.IGenerateService) *genHandler {
	return &genHandler{srv: srv}
}

func regGen(rg *gin.RouterGroup, group *core.GroupBase) error {
	return group.Reg(func(handle *genHandler) {
		rg.GET("/db", handle.dbTables)
		rg.GET("/list", handle.list)
		rg.GET("/detail", handle.detail)
		rg.POST("/importTable", handle.importTable)
		rg.POST("/delTable", handle.delTable)
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
func (gh genHandler) list(c *gin.Context) {
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
func (gh genHandler) detail(c *gin.Context) {
	var detailReq req.DetailTableReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := gh.srv.Detail(detailReq.ID)
	response.CheckAndRespWithData(c, res, err)
}

//ImportTable 导入表结构
func (gh genHandler) importTable(c *gin.Context) {
	var importReq req.ImportTableReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &importReq)) {
		return
	}
	err := gh.srv.ImportTable(strings.Split(importReq.Tables, ","))
	response.CheckAndResp(c, err)
}

//DelTable 删除表结构
func (gh genHandler) delTable(c *gin.Context) {
	var delReq req.DelTableReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &delReq)) {
		return
	}
	err := gh.srv.DelTable(delReq.Ids)
	response.CheckAndResp(c, err)
}
