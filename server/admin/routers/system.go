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
	"likeadmin/util"
)

var SystemGroup = core.Group("/system")

func init() {
	group := SystemGroup
	group.AddPOST("/login", login)
	group.AddPOST("/logout", logout)
	group.AddGET("/admin/self", adminSelf)
	group.AddGET("/admin/list", adminList)
	group.AddGET("/admin/detail", adminDetail)
	group.AddPOST("/admin/add", adminAdd, middleware.RecordLog("管理员新增"))
	group.AddPOST("/admin/edit", adminEdit, middleware.RecordLog("管理员编辑"))
	group.AddPOST("/admin/upInfo", adminUpInfo, middleware.RecordLog("管理员更新"))
	group.AddPOST("/admin/del", adminDel, middleware.RecordLog("管理员删除"))
	group.AddPOST("/admin/disable", adminDisable, middleware.RecordLog("管理员状态切换"))
	group.AddGET("/role/all", roleAll)
	group.AddGET("/role/list", roleList, middleware.RecordLog("角色列表"))
	group.AddGET("/role/detail", roleDetail, middleware.RecordLog("角色详情"))
	group.AddPOST("/role/add", roleAdd, middleware.RecordLog("角色新增"))
	group.AddPOST("/role/edit", roleEdit, middleware.RecordLog("角色编辑"))
	group.AddPOST("/role/del", roleDel, middleware.RecordLog("角色删除"))
	group.AddGET("/menu/route", menuRoute)
	group.AddGET("/menu/list", menuList)
	group.AddGET("/menu/detail", menuDetail)
	group.AddPOST("/menu/add", menuAdd)
	group.AddPOST("/menu/edit", menuEdit)
	group.AddPOST("/menu/del", menuDel)
}

//login 登录系统
func login(c *gin.Context) {
	var loginReq req.SystemLoginReq
	util.VerifyUtil.VerifyJSON(c, &loginReq)
	resp := system.SystemLoginService.Login(c, &loginReq)
	response.OkWithData(c, resp)
}

//logout 登录退出
func logout(c *gin.Context) {
	var logoutReq req.SystemLogoutReq
	util.VerifyUtil.VerifyHeader(c, &logoutReq)
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
	util.VerifyUtil.VerifyQuery(c, &page)
	util.VerifyUtil.VerifyQuery(c, &listReq)
	response.OkWithData(c, system.SystemAuthAdminService.List(page, listReq))
}

//adminDetail 管理员详细
func adminDetail(c *gin.Context) {
	var detailReq req.SystemAuthAdminDetailReq
	util.VerifyUtil.VerifyQuery(c, &detailReq)
	response.OkWithData(c, system.SystemAuthAdminService.Detail(detailReq.ID))
}

//adminAdd 管理员新增
func adminAdd(c *gin.Context) {
	var addReq req.SystemAuthAdminAddReq
	util.VerifyUtil.VerifyJSON(c, &addReq)
	system.SystemAuthAdminService.Add(addReq)
	response.Ok(c)
}

//adminEdit 管理员编辑
func adminEdit(c *gin.Context) {
	var editReq req.SystemAuthAdminEditReq
	util.VerifyUtil.VerifyJSON(c, &editReq)
	system.SystemAuthAdminService.Edit(c, editReq)
	response.Ok(c)
}

//adminUpInfo 管理员更新
func adminUpInfo(c *gin.Context) {
	var updateReq req.SystemAuthAdminUpdateReq
	util.VerifyUtil.VerifyJSON(c, &updateReq)
	system.SystemAuthAdminService.Update(c, updateReq, config.AdminConfig.GetAdminId(c))
	response.Ok(c)
}

//adminDel 管理员删除
func adminDel(c *gin.Context) {
	var delReq req.SystemAuthAdminDelReq
	util.VerifyUtil.VerifyJSON(c, &delReq)
	system.SystemAuthAdminService.Del(c, delReq.ID)
	response.Ok(c)
}

//adminDisable 管理员状态切换
func adminDisable(c *gin.Context) {
	var disableReq req.SystemAuthAdminDisableReq
	util.VerifyUtil.VerifyJSON(c, &disableReq)
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
	util.VerifyUtil.VerifyQuery(c, &page)
	response.OkWithData(c, system.SystemAuthRoleService.List(page))
}

//roleDetail 角色详情
func roleDetail(c *gin.Context) {
	var detailReq req.SystemAuthRoleDetailReq
	util.VerifyUtil.VerifyQuery(c, &detailReq)
	response.OkWithData(c, system.SystemAuthRoleService.Detail(detailReq.ID))
}

//roleAdd 新增角色
func roleAdd(c *gin.Context) {
	var addReq req.SystemAuthRoleAddReq
	util.VerifyUtil.VerifyJSON(c, &addReq)
	system.SystemAuthRoleService.Add(addReq)
	response.Ok(c)
}

//roleEdit 编辑角色
func roleEdit(c *gin.Context) {
	var editReq req.SystemAuthRoleEditReq
	util.VerifyUtil.VerifyJSON(c, &editReq)
	system.SystemAuthRoleService.Edit(editReq)
	response.Ok(c)
}

//roleDel 删除角色
func roleDel(c *gin.Context) {
	var delReq req.SystemAuthRoleDelReq
	util.VerifyUtil.VerifyJSON(c, &delReq)
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

//menuDetail 菜单详情
func menuDetail(c *gin.Context) {
	var detailReq req.SystemAuthMenuDetailReq
	util.VerifyUtil.VerifyQuery(c, &detailReq)
	response.OkWithData(c, system.SystemAuthMenuService.Detail(detailReq.ID))
}

//menuAdd 新增菜单
func menuAdd(c *gin.Context) {
	var addReq req.SystemAuthMenuAddReq
	util.VerifyUtil.VerifyJSON(c, &addReq)
	system.SystemAuthMenuService.Add(addReq)
	response.Ok(c)
}

//menuEdit 编辑菜单
func menuEdit(c *gin.Context) {
	var editReq req.SystemAuthMenuEditReq
	util.VerifyUtil.VerifyJSON(c, &editReq)
	system.SystemAuthMenuService.Edit(editReq)
	response.Ok(c)
}

//menuDel 删除菜单
func menuDel(c *gin.Context) {
	var delReq req.SystemAuthMenuDelReq
	util.VerifyUtil.VerifyJSON(c, &delReq)
	system.SystemAuthMenuService.Del(delReq.ID)
	response.Ok(c)
}
