package routers

import (
	"github.com/gin-gonic/gin"
	"likeadmin/admin/schemas/req"
	"likeadmin/admin/service/system"
	"likeadmin/config"
	"likeadmin/core"
	"likeadmin/core/request"
	"likeadmin/core/response"
	"likeadmin/middleware"
	"likeadmin/utils"
)

var Group = core.Group("/system")

func init() {
	Group.AddPOST("/login", login)
	Group.AddPOST("/logout", logout)
	Group.AddGET("/admin/self", adminSelf)
	Group.AddGET("/admin/list", adminList)
	Group.AddGET("/admin/detail", adminDetail)
	Group.AddPOST("/admin/add", adminAdd, middleware.RecordLog("管理员新增"))
	Group.AddPOST("/admin/edit", adminEdit, middleware.RecordLog("管理员编辑"))
	Group.AddPOST("/admin/upInfo", adminUpInfo, middleware.RecordLog("管理员更新"))
	Group.AddPOST("/admin/del", adminDel, middleware.RecordLog("管理员删除"))
	Group.AddPOST("/admin/disable", adminDisable, middleware.RecordLog("管理员状态切换"))
	Group.AddGET("/role/all", roleAll)
	Group.AddGET("/role/list", roleList, middleware.RecordLog("角色列表"))
	Group.AddGET("/role/detail", roleDetail, middleware.RecordLog("角色详情"))
	Group.AddPOST("/role/add", roleAdd, middleware.RecordLog("角色新增"))
	Group.AddPOST("/role/edit", roleEdit, middleware.RecordLog("角色编辑"))
	Group.AddPOST("/role/del", roleDel, middleware.RecordLog("角色删除"))
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

//roleAll 角色所有
func roleAll(c *gin.Context) {
	response.OkWithData(c, system.SystemAuthRoleService.All())
}

//roleList 角色列表
func roleList(c *gin.Context) {
	var page request.PageReq
	utils.VerifyUtil.VerifyQuery(c, &page)
	response.OkWithData(c, system.SystemAuthRoleService.List(page))
}

//roleDetail 角色详情
func roleDetail(c *gin.Context) {
	var detailReq req.SystemAuthRoleDetailReq
	utils.VerifyUtil.VerifyQuery(c, &detailReq)
	response.OkWithData(c, system.SystemAuthRoleService.Detail(detailReq.ID))
}

//roleAdd 新增角色
func roleAdd(c *gin.Context) {
	var addReq req.SystemAuthRoleAddReq
	utils.VerifyUtil.VerifyJSON(c, &addReq)
	system.SystemAuthRoleService.Add(addReq)
	response.Ok(c)
}

//roleEdit 编辑角色
func roleEdit(c *gin.Context) {
	var editReq req.SystemAuthRoleEditReq
	utils.VerifyUtil.VerifyJSON(c, &editReq)
	system.SystemAuthRoleService.Edit(editReq)
	response.Ok(c)
}

//roleDel 删除角色
func roleDel(c *gin.Context) {
	var delReq req.SystemAuthRoleDelReq
	utils.VerifyUtil.VerifyJSON(c, &delReq)
	system.SystemAuthRoleService.Del(delReq.ID)
	response.Ok(c)
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
