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
		rg.POST("/syncTable", handle.syncTable)
		rg.POST("/editTable", handle.editTable)
		rg.POST("/delTable", handle.delTable)
		rg.GET("/previewCode", handle.previewCode)
		rg.GET("/genCode", handle.genCode)
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

//list 生成列表
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

//detail 生成详情
func (gh genHandler) detail(c *gin.Context) {
	var detailReq req.DetailTableReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := gh.srv.Detail(detailReq.ID)
	response.CheckAndRespWithData(c, res, err)
}

//importTable 导入表结构
func (gh genHandler) importTable(c *gin.Context) {
	var importReq req.ImportTableReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &importReq)) {
		return
	}
	err := gh.srv.ImportTable(strings.Split(importReq.Tables, ","))
	response.CheckAndResp(c, err)
}

//syncTable 同步表结构
func (gh genHandler) syncTable(c *gin.Context) {
	var syncReq req.SyncTableReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &syncReq)) {
		return
	}
	err := gh.srv.SyncTable(syncReq.ID)
	response.CheckAndResp(c, err)
}

//editTable 编辑表结构
func (gh genHandler) editTable(c *gin.Context) {
	var editReq req.EditTableReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &editReq)) {
		return
	}
	err := gh.srv.EditTable(editReq)
	response.CheckAndResp(c, err)
}

//delTable 删除表结构
func (gh genHandler) delTable(c *gin.Context) {
	var delReq req.DelTableReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &delReq)) {
		return
	}
	err := gh.srv.DelTable(delReq.Ids)
	response.CheckAndResp(c, err)
}

//previewCode 预览代码
func (gh genHandler) previewCode(c *gin.Context) {
	var previewReq req.PreviewCodeReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &previewReq)) {
		return
	}
	res, err := gh.srv.PreviewCode(previewReq.ID)
	response.CheckAndRespWithData(c, res, err)
}

//genCode 生成代码
func (gh genHandler) genCode(c *gin.Context) {
	var genReq req.GenCodeReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &genReq)) {
		return
	}
	for _, table := range strings.Split(genReq.Tables, ",") {
		err := gh.srv.GenCode(table)
		if response.IsFailWithResp(c, err) {
			return
		}
	}
	response.Ok(c)
}
