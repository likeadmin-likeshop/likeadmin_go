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

var RoleGroup = core.Group("/system", newRoleHandler, regRole, middleware.TokenAuth())

func newRoleHandler(srv system.ISystemAuthRoleService) *roleHandler {
	return &roleHandler{srv: srv}
}

func regRole(rg *gin.RouterGroup, group *core.GroupBase) error {
	return group.Reg(func(handle *roleHandler) {
		rg.GET("/role/all", handle.all)
		rg.GET("/role/list", middleware.RecordLog("角色列表"), handle.list)
		rg.GET("/role/detail", middleware.RecordLog("角色详情"), handle.detail)
		rg.POST("/role/add", middleware.RecordLog("角色新增"), handle.add)
		rg.POST("/role/edit", middleware.RecordLog("角色编辑"), handle.edit)
		rg.POST("/role/del", middleware.RecordLog("角色删除"), handle.del)
	})
}

type roleHandler struct {
	srv system.ISystemAuthRoleService
}

//all 角色所有
func (rh roleHandler) all(c *gin.Context) {
	res, err := rh.srv.All()
	response.CheckAndRespWithData(c, res, err)
}

//list 角色列表
func (rh roleHandler) list(c *gin.Context) {
	var page request.PageReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	res, err := rh.srv.List(page)
	response.CheckAndRespWithData(c, res, err)
}

//detail 角色详情
func (rh roleHandler) detail(c *gin.Context) {
	var detailReq req.SystemAuthRoleDetailReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := rh.srv.Detail(detailReq.ID)
	response.CheckAndRespWithData(c, res, err)
}

//add 新增角色
func (rh roleHandler) add(c *gin.Context) {
	var addReq req.SystemAuthRoleAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &addReq)) {
		return
	}
	response.CheckAndResp(c, rh.srv.Add(addReq))
}

//edit 编辑角色
func (rh roleHandler) edit(c *gin.Context) {
	var editReq req.SystemAuthRoleEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &editReq)) {
		return
	}
	response.CheckAndResp(c, rh.srv.Edit(editReq))
}

//del 删除角色
func (rh roleHandler) del(c *gin.Context) {
	var delReq req.SystemAuthRoleDelReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &delReq)) {
		return
	}
	response.CheckAndResp(c, rh.srv.Del(delReq.ID))
}
