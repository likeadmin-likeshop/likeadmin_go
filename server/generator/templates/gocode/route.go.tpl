package {{{ .PackageName }}}

import (
	"github.com/gin-gonic/gin"
	"likeadmin/admin/schemas/req"
	"likeadmin/admin/service/system"
	"likeadmin/core"
	"likeadmin/core/request"
	"likeadmin/core/response"
	"likeadmin/middleware"
	"likeadmin/util"
)

var {{{ title (toCamelCase .ModuleName) }}}Group = core.Group("/", new{{{ title (toCamelCase .ModuleName) }}}Handler, reg{{{ title (toCamelCase .ModuleName) }}}, middleware.TokenAuth())

func new{{{ title (toCamelCase .ModuleName) }}}Handler(srv I{{{ title (toCamelCase .EntityName) }}}Service) *{{{ toCamelCase .ModuleName }}}Handler {
	return &{{{ toCamelCase .ModuleName }}}Handler{srv: srv}
}

func reg{{{ title (toCamelCase .ModuleName) }}}(rg *gin.RouterGroup, group *core.GroupBase) error {
	return group.Reg(func(handle *{{{ toCamelCase .ModuleName }}}Handler) {
		rg.GET("/{{{ .ModuleName }}}/list", handle.list)
		rg.GET("/{{{ .ModuleName }}}/detail", handle.detail)
		rg.POST("/{{{ .ModuleName }}}/add", handle.add)
		rg.POST("/{{{ .ModuleName }}}/edit", handle.edit)
		rg.POST("/{{{ .ModuleName }}}/del", handle.del)
	})
}

type {{{ toCamelCase .ModuleName }}}Handler struct {
	srv I{{{ title (toCamelCase .EntityName) }}}Service
}

//list {{{ .ModuleName }}}列表
func (hd {{{ toCamelCase .ModuleName }}}Handler) list(c *gin.Context) {
	var page request.PageReq
	var listReq {{{ title (toCamelCase .EntityName) }}}ListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := ph.srv.List(page, listReq)
	response.CheckAndRespWithData(c, res, err)
}

//detail {{{ .ModuleName }}}详情
func (hd {{{ toCamelCase .ModuleName }}}Handler) detail(c *gin.Context) {
	var detailReq {{{ title (toCamelCase .EntityName) }}}DetailReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := hd.srv.Detail(detailReq.{{{ title (toCamelCase .PrimaryKey) }}})
	response.CheckAndRespWithData(c, res, err)
}

//add {{{ .ModuleName }}}新增
func (hd {{{ toCamelCase .ModuleName }}}Handler) add(c *gin.Context) {
	var addReq {{{ title (toCamelCase .EntityName) }}}AddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &addReq)) {
		return
	}
	response.CheckAndResp(c, hd.srv.Add(addReq))
}

//edit {{{ .ModuleName }}}编辑
func (hd {{{ toCamelCase .ModuleName }}}Handler) edit(c *gin.Context) {
	var editReq {{{ title (toCamelCase .EntityName) }}}EditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &editReq)) {
		return
	}
	response.CheckAndResp(c, hd.srv.Edit(editReq))
}

//del {{{ .ModuleName }}}删除
func (hd {{{ toCamelCase .ModuleName }}}Handler) del(c *gin.Context) {
	var delReq {{{ title (toCamelCase .EntityName) }}}DelReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &delReq)) {
		return
	}
	response.CheckAndResp(c, hd.srv.Del(delReq.{{{ title (toCamelCase .PrimaryKey) }}}))
}
