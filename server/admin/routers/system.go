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
	group.AddGET("/dept/all", deptAll)
	group.AddGET("/dept/list", deptList)
	group.AddGET("/dept/detail", deptDetail)
	group.AddPOST("/dept/add", deptAdd)
	group.AddPOST("/dept/edit", deptEdit)
	group.AddPOST("/dept/del", deptDel)
	group.AddGET("/post/all", postAll)
	group.AddGET("/post/list", postList)
	group.AddGET("/post/detail", postDetail)
	group.AddPOST("/post/add", postAdd)
	group.AddPOST("/post/edit", postEdit)
	group.AddPOST("/post/del", postDel)
	group.AddGET("/log/operate", logOperate)
	group.AddGET("/log/login", logLogin)
}

//login 登录系统
func login(c *gin.Context) {
	var loginReq req.SystemLoginReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &loginReq)) {
		return
	}
	res, err := system.SystemLoginService.Login(c, &loginReq)
	response.CheckAndRespWithData(c, res, err)
}

//logout 登录退出
func logout(c *gin.Context) {
	var logoutReq req.SystemLogoutReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyHeader(c, &logoutReq)) {
		return
	}
	response.CheckAndResp(c, system.SystemLoginService.Logout(&logoutReq))
}

//adminSelf 管理员信息
func adminSelf(c *gin.Context) {
	adminId := config.AdminConfig.GetAdminId(c)
	res, err := system.SystemAuthAdminService.Self(adminId)
	response.CheckAndRespWithData(c, res, err)
}

//adminList 管理员列表
func adminList(c *gin.Context) {
	var page request.PageReq
	var listReq req.SystemAuthAdminListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := system.SystemAuthAdminService.List(page, listReq)
	response.CheckAndRespWithData(c, res, err)
}

//adminDetail 管理员详细
func adminDetail(c *gin.Context) {
	var detailReq req.SystemAuthAdminDetailReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := system.SystemAuthAdminService.Detail(detailReq.ID)
	response.CheckAndRespWithData(c, res, err)
}

//adminAdd 管理员新增
func adminAdd(c *gin.Context) {
	var addReq req.SystemAuthAdminAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &addReq)) {
		return
	}
	response.CheckAndResp(c, system.SystemAuthAdminService.Add(addReq))
}

//adminEdit 管理员编辑
func adminEdit(c *gin.Context) {
	var editReq req.SystemAuthAdminEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &editReq)) {
		return
	}
	response.CheckAndResp(c, system.SystemAuthAdminService.Edit(c, editReq))
}

//adminUpInfo 管理员更新
func adminUpInfo(c *gin.Context) {
	var updateReq req.SystemAuthAdminUpdateReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &updateReq)) {
		return
	}
	response.CheckAndResp(c, system.SystemAuthAdminService.Update(c, updateReq, config.AdminConfig.GetAdminId(c)))
}

//adminDel 管理员删除
func adminDel(c *gin.Context) {
	var delReq req.SystemAuthAdminDelReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &delReq)) {
		return
	}
	response.CheckAndResp(c, system.SystemAuthAdminService.Del(c, delReq.ID))
}

//adminDisable 管理员状态切换
func adminDisable(c *gin.Context) {
	var disableReq req.SystemAuthAdminDisableReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &disableReq)) {
		return
	}
	response.CheckAndResp(c, system.SystemAuthAdminService.Disable(c, disableReq.ID))
}

//roleAll 角色所有
func roleAll(c *gin.Context) {
	res, err := system.SystemAuthRoleService.All()
	response.CheckAndRespWithData(c, res, err)
}

//roleList 角色列表
func roleList(c *gin.Context) {
	var page request.PageReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	res, err := system.SystemAuthRoleService.List(page)
	response.CheckAndRespWithData(c, res, err)
}

//roleDetail 角色详情
func roleDetail(c *gin.Context) {
	var detailReq req.SystemAuthRoleDetailReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := system.SystemAuthRoleService.Detail(detailReq.ID)
	response.CheckAndRespWithData(c, res, err)
}

//roleAdd 新增角色
func roleAdd(c *gin.Context) {
	var addReq req.SystemAuthRoleAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &addReq)) {
		return
	}
	response.CheckAndResp(c, system.SystemAuthRoleService.Add(addReq))
}

//roleEdit 编辑角色
func roleEdit(c *gin.Context) {
	var editReq req.SystemAuthRoleEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &editReq)) {
		return
	}
	response.CheckAndResp(c, system.SystemAuthRoleService.Edit(editReq))
}

//roleDel 删除角色
func roleDel(c *gin.Context) {
	var delReq req.SystemAuthRoleDelReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &delReq)) {
		return
	}
	response.CheckAndResp(c, system.SystemAuthRoleService.Del(delReq.ID))
}

//menuRoute 菜单路由
func menuRoute(c *gin.Context) {
	adminId := config.AdminConfig.GetAdminId(c)
	res, err := system.SystemAuthMenuService.SelectMenuByRoleId(c, adminId)
	response.CheckAndRespWithData(c, res, err)
}

//menuList 菜单列表
func menuList(c *gin.Context) {
	res, err := system.SystemAuthMenuService.List()
	response.CheckAndRespWithData(c, res, err)
}

//menuDetail 菜单详情
func menuDetail(c *gin.Context) {
	var detailReq req.SystemAuthMenuDetailReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := system.SystemAuthMenuService.Detail(detailReq.ID)
	response.CheckAndRespWithData(c, res, err)
}

//menuAdd 新增菜单
func menuAdd(c *gin.Context) {
	var addReq req.SystemAuthMenuAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &addReq)) {
		return
	}
	response.CheckAndResp(c, system.SystemAuthMenuService.Add(addReq))
}

//menuEdit 编辑菜单
func menuEdit(c *gin.Context) {
	var editReq req.SystemAuthMenuEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &editReq)) {
		return
	}
	response.CheckAndResp(c, system.SystemAuthMenuService.Edit(editReq))
}

//menuDel 删除菜单
func menuDel(c *gin.Context) {
	var delReq req.SystemAuthMenuDelReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &delReq)) {
		return
	}
	response.CheckAndResp(c, system.SystemAuthMenuService.Del(delReq.ID))
}

//deptAll 部门所有
func deptAll(c *gin.Context) {
	res, err := system.SystemAuthDeptService.All()
	response.CheckAndRespWithData(c, res, err)
}

//deptList 部门列表
func deptList(c *gin.Context) {
	var listReq req.SystemAuthDeptListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := system.SystemAuthDeptService.List(listReq)
	response.CheckAndRespWithData(c, res, err)
}

//deptDetail 部门详情
func deptDetail(c *gin.Context) {
	var detailReq req.SystemAuthDeptDetailReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := system.SystemAuthDeptService.Detail(detailReq.ID)
	response.CheckAndRespWithData(c, res, err)
}

//deptAdd 部门新增
func deptAdd(c *gin.Context) {
	var addReq req.SystemAuthDeptAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &addReq)) {
		return
	}
	response.CheckAndResp(c, system.SystemAuthDeptService.Add(addReq))
}

//deptEdit 部门编辑
func deptEdit(c *gin.Context) {
	var editReq req.SystemAuthDeptEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &editReq)) {
		return
	}
	response.CheckAndResp(c, system.SystemAuthDeptService.Edit(editReq))
}

//deptDel 部门删除
func deptDel(c *gin.Context) {
	var delReq req.SystemAuthDeptDelReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &delReq)) {
		return
	}
	response.CheckAndResp(c, system.SystemAuthDeptService.Del(delReq.ID))
}

//postAll 岗位所有
func postAll(c *gin.Context) {
	res, err := system.SystemAuthPostService.All()
	response.CheckAndRespWithData(c, res, err)
}

//postList 岗位列表
func postList(c *gin.Context) {
	var page request.PageReq
	var listReq req.SystemAuthPostListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := system.SystemAuthPostService.List(page, listReq)
	response.CheckAndRespWithData(c, res, err)
}

//postDetail 岗位详情
func postDetail(c *gin.Context) {
	var detailReq req.SystemAuthPostDetailReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := system.SystemAuthPostService.Detail(detailReq.ID)
	response.CheckAndRespWithData(c, res, err)
}

//postAdd 岗位新增
func postAdd(c *gin.Context) {
	var addReq req.SystemAuthPostAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &addReq)) {
		return
	}
	response.CheckAndResp(c, system.SystemAuthPostService.Add(addReq))
}

//postEdit 岗位编辑
func postEdit(c *gin.Context) {
	var editReq req.SystemAuthPostEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &editReq)) {
		return
	}
	response.CheckAndResp(c, system.SystemAuthPostService.Edit(editReq))
}

//postDel 岗位删除
func postDel(c *gin.Context) {
	var delReq req.SystemAuthPostDelReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &delReq)) {
		return
	}
	response.CheckAndResp(c, system.SystemAuthPostService.Del(delReq.ID))
}

//logOperate 操作日志
func logOperate(c *gin.Context) {
	var page request.PageReq
	var logReq req.SystemLogOperateReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &logReq)) {
		return
	}
	res, err := system.SystemLogsServer.Operate(page, logReq)
	response.CheckAndRespWithData(c, res, err)
}

//logLogin 登录日志
func logLogin(c *gin.Context) {
	var page request.PageReq
	var logReq req.SystemLogLoginReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &logReq)) {
		return
	}
	res, err := system.SystemLogsServer.Login(page, logReq)
	response.CheckAndRespWithData(c, res, err)
}
