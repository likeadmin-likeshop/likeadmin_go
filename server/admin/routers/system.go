package routers

import (
	"github.com/gin-gonic/gin"
	"likeadmin/admin/schemas/req"
	"likeadmin/admin/schemas/resp"
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
	var menuResps []resp.SystemAuthMenuResp
	response.Copy(&menuResps, system.SystemAuthMenuService.List())
	menuTree := utils.ArrayUtil.ListToTree(
		utils.ConvertUtil.StructsToMaps(menuResps), "id", "pid", "children")
	response.OkWithData(c, menuTree)
}
