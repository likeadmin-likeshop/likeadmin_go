package routers

import (
	"github.com/gin-gonic/gin"
	"likeadmin/admin/schemas/req"
	"likeadmin/admin/schemas/resp"
	"likeadmin/admin/service/system"
	"likeadmin/core"
	"likeadmin/core/response"
	"likeadmin/utils"
)

var Group = core.Group("/system")

func init() {
	Group.AddPOST("/login", login)
	Group.AddPOST("/logout", logout)
	Group.AddGET("/menu/list", menuList)
}

//login 登录系统
func login(c *gin.Context) {
	var loginReq req.SystemLoginReq
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		response.FailWithData(c, response.ParamsValidError, err.Error())
		return
	}
	resp := system.SystemLoginService.Login(c, &loginReq)
	response.OkWithData(c, resp)
}

//logout 登录退出
func logout(c *gin.Context) {
	var logoutReq req.SystemLogoutReq
	if err := c.ShouldBindHeader(&logoutReq); err != nil {
		response.FailWithData(c, response.ParamsValidError, err.Error())
		return
	}
	system.SystemLoginService.Logout(&logoutReq)
	response.Ok(c)
}

//func menuList(c *gin.Context) {
//	var menus []system.SystemAuthMenu
//	result := core.DB.Find(&menus)
//	var menuResps []resp.SystemAuthMenuResp
//	response.Copy(c, &menuResps, &menus)
//	response.OkWithData(c, response.PageResp{
//		Count:    result.RowsAffected,
//		PageNo:   1,
//		PageSize: 20,
//		Lists:    menuResps,
//	})
//}

//menuList 菜单列表
func menuList(c *gin.Context) {
	var menuResps []resp.SystemAuthMenuResp
	response.Copy(c, &menuResps, system.SystemAuthMenuService.List())
	menuTree := utils.ArrayUtil.ListToTree(
		utils.ConvertUtil.StructsToMaps(menuResps), "id", "pid", "children")
	response.OkWithData(c, menuTree)
}
