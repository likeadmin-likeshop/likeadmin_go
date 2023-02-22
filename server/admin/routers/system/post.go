package system

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

var PostGroup = core.Group("/system", newPostHandler, regPost, middleware.TokenAuth())

func newPostHandler(srv system.ISystemAuthPostService) *postHandler {
	return &postHandler{srv: srv}
}

func regPost(rg *gin.RouterGroup, group *core.GroupBase) error {
	return group.Reg(func(handle *postHandler) {
		rg.GET("/post/all", handle.all)
		rg.GET("/post/list", handle.list)
		rg.GET("/post/detail", handle.detail)
		rg.POST("/post/add", handle.add)
		rg.POST("/post/edit", handle.edit)
		rg.POST("/post/del", handle.del)
	})
}

type postHandler struct {
	srv system.ISystemAuthPostService
}

//all 岗位所有
func (ph postHandler) all(c *gin.Context) {
	res, err := ph.srv.All()
	response.CheckAndRespWithData(c, res, err)
}

//list 岗位列表
func (ph postHandler) list(c *gin.Context) {
	var page request.PageReq
	var listReq req.SystemAuthPostListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := ph.srv.List(page, listReq)
	response.CheckAndRespWithData(c, res, err)
}

//detail 岗位详情
func (ph postHandler) detail(c *gin.Context) {
	var detailReq req.SystemAuthPostDetailReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := ph.srv.Detail(detailReq.ID)
	response.CheckAndRespWithData(c, res, err)
}

//add 岗位新增
func (ph postHandler) add(c *gin.Context) {
	var addReq req.SystemAuthPostAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &addReq)) {
		return
	}
	response.CheckAndResp(c, ph.srv.Add(addReq))
}

//edit 岗位编辑
func (ph postHandler) edit(c *gin.Context) {
	var editReq req.SystemAuthPostEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &editReq)) {
		return
	}
	response.CheckAndResp(c, ph.srv.Edit(editReq))
}

//del 岗位删除
func (ph postHandler) del(c *gin.Context) {
	var delReq req.SystemAuthPostDelReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &delReq)) {
		return
	}
	response.CheckAndResp(c, ph.srv.Del(delReq.ID))
}
