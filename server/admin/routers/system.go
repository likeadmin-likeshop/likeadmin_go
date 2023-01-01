package routers

import (
	"github.com/gin-gonic/gin"
	"likeadmin/admin/schemas/req"
	"likeadmin/admin/service/system"
	"likeadmin/config"
	"likeadmin/core"
	"likeadmin/core/request"
	"likeadmin/core/response"
	"likeadmin/utils"
)

var Group = core.Group("/system")

func init() {
	Group.AddPOST("/login", login)
	Group.AddPOST("/logout", logout)
	Group.AddGET("/admin/self", adminSelf)
	Group.AddGET("/admin/list", adminList)
	Group.AddGET("/admin/detail", adminDetail)
	Group.AddPOST("/admin/add", adminAdd)
	Group.AddPOST("/admin/edit", adminEdit)
	Group.AddPOST("/admin/upInfo", adminUpInfo)
	Group.AddPOST("/admin/del", adminDel)
	Group.AddPOST("/admin/disable", adminDisable)
	Group.AddGET("/role/list", roleList)
	Group.AddGET("/menu/route", menuRoute)
	Group.AddGET("/menu/list", menuList)
}

//login 登录系统
func login(c *gin.Context) {
	var loginReq req.SystemLoginReq
	utils.VerifyUtil.VerifyJSON(c, &loginReq)
	resp := system.SystemLoginService.Login(c, &loginReq)
	response.OkWithData(c, resp)
}

//logout 登录退出
func logout(c *gin.Context) {
	var logoutReq req.SystemLogoutReq
	utils.VerifyUtil.VerifyHeader(c, &logoutReq)
	system.SystemLoginService.Logout(&logoutReq)
	response.Ok(c)
}

//adminSelf 管理员信息
func adminSelf(c *gin.Context) {
	adminId := config.AdminConfig.GetAdminId(c)
	response.OkWithData(c, system.SystemAuthAdminService.Self(adminId))
}

//adminList 管理员列表
func adminList(c *gin.Context) {
	var page request.PageReq
	var listReq req.SystemAuthAdminListReq
	utils.VerifyUtil.VerifyQuery(c, &page)
	utils.VerifyUtil.VerifyQuery(c, &listReq)
	response.OkWithData(c, system.SystemAuthAdminService.List(page, listReq))
}

//adminDetail 管理员详细
func adminDetail(c *gin.Context) {
	var detailReq req.SystemAuthAdminDetailReq
	utils.VerifyUtil.VerifyQuery(c, &detailReq)
	response.OkWithData(c, system.SystemAuthAdminService.Detail(detailReq.ID))
}

//adminAdd 管理员新增
func adminAdd(c *gin.Context) {
	var addReq req.SystemAuthAdminAddReq
	utils.VerifyUtil.VerifyJSON(c, &addReq)
	system.SystemAuthAdminService.Add(addReq)
	response.Ok(c)
}

//adminEdit 管理员编辑
func adminEdit(c *gin.Context) {
	var editReq req.SystemAuthAdminEditReq
	utils.VerifyUtil.VerifyJSON(c, &editReq)
	system.SystemAuthAdminService.Edit(c, editReq)
	response.Ok(c)
}

//adminUpInfo 管理员更新
func adminUpInfo(c *gin.Context) {
	// TODO: 管理员更新
	var updateReq req.SystemAuthAdminUpdateReq
	utils.VerifyUtil.VerifyJSON(c, &updateReq)
	system.SystemAuthAdminService.Update(c, updateReq, config.AdminConfig.GetAdminId(c))
	response.Ok(c)
}

//adminDel 管理员删除
func adminDel(c *gin.Context) {
	var delReq req.SystemAuthAdminDelReq
	utils.VerifyUtil.VerifyJSON(c, &delReq)
	system.SystemAuthAdminService.Del(c, delReq.ID)
	response.Ok(c)
}

//adminDisable 管理员状态切换
func adminDisable(c *gin.Context) {
	var disableReq req.SystemAuthAdminDisableReq
	utils.VerifyUtil.VerifyJSON(c, &disableReq)
	system.SystemAuthAdminService.Disable(c, disableReq.ID)
	response.Ok(c)
}

//roleList 角色列表
func roleList(c *gin.Context) {
	var page request.PageReq
	utils.VerifyUtil.VerifyQuery(c, &page)
	response.OkWithData(c, system.SystemAuthRoleService.List(page))
}

//menuRoute 菜单路由
func menuRoute(c *gin.Context) {
	adminId := config.AdminConfig.GetAdminId(c)
	response.OkWithData(c, system.SystemAuthMenuService.SelectMenuByRoleId(c, adminId))
}

//menuList 菜单列表
func menuList(c *gin.Context) {
	response.OkWithData(c, system.SystemAuthMenuService.List())
}
