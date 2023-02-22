package system

import (
	"github.com/gin-gonic/gin"
	"likeadmin/admin/schemas/req"
	"likeadmin/admin/service/system"
	"likeadmin/config"
	"likeadmin/core"
	"likeadmin/core/response"
	"likeadmin/middleware"
	"likeadmin/util"
)

var MenuGroup = core.Group("/system", newMenuHandler, regMenu, middleware.TokenAuth())

func newMenuHandler(srv system.ISystemAuthMenuService) *menuHandler {
	return &menuHandler{srv: srv}
}

func regMenu(rg *gin.RouterGroup, group *core.GroupBase) error {
	return group.Reg(func(handle *menuHandler) {
		rg.GET("/menu/route", handle.route)
		rg.GET("/menu/list", handle.list)
		rg.GET("/menu/detail", handle.detail)
		rg.POST("/menu/add", handle.add)
		rg.POST("/menu/edit", handle.edit)
		rg.POST("/menu/del", handle.del)
	})
}

type menuHandler struct {
	srv system.ISystemAuthMenuService
}

//route 菜单路由
func (mh menuHandler) route(c *gin.Context) {
	adminId := config.AdminConfig.GetAdminId(c)
	res, err := mh.srv.SelectMenuByRoleId(c, adminId)
	response.CheckAndRespWithData(c, res, err)
}

//list 菜单列表
func (mh menuHandler) list(c *gin.Context) {
	res, err := mh.srv.List()
	response.CheckAndRespWithData(c, res, err)
}

//detail 菜单详情
func (mh menuHandler) detail(c *gin.Context) {
	var detailReq req.SystemAuthMenuDetailReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := mh.srv.Detail(detailReq.ID)
	response.CheckAndRespWithData(c, res, err)
}

//add 新增菜单
func (mh menuHandler) add(c *gin.Context) {
	var addReq req.SystemAuthMenuAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &addReq)) {
		return
	}
	response.CheckAndResp(c, mh.srv.Add(addReq))
}

//edit 编辑菜单
func (mh menuHandler) edit(c *gin.Context) {
	var editReq req.SystemAuthMenuEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &editReq)) {
		return
	}
	response.CheckAndResp(c, mh.srv.Edit(editReq))
}

//del 删除菜单
func (mh menuHandler) del(c *gin.Context) {
	var delReq req.SystemAuthMenuDelReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &delReq)) {
		return
	}
	response.CheckAndResp(c, mh.srv.Del(delReq.ID))
}
