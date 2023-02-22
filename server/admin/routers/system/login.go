package system

import (
	"github.com/gin-gonic/gin"
	"likeadmin/admin/schemas/req"
	"likeadmin/admin/service/system"
	"likeadmin/core"
	"likeadmin/core/response"
	"likeadmin/middleware"
	"likeadmin/util"
)

var LoginGroup = core.Group("/system", newLoginHandler, regLogin, middleware.TokenAuth())

func newLoginHandler(srv system.ISystemLoginService) *loginHandler {
	return &loginHandler{srv: srv}
}

func regLogin(rg *gin.RouterGroup, group *core.GroupBase) error {
	return group.Reg(func(handle *loginHandler) {
		rg.POST("/login", handle.login)
		rg.POST("/logout", handle.logout)
	})
}

type loginHandler struct {
	srv system.ISystemLoginService
}

//login 登录系统
func (lh loginHandler) login(c *gin.Context) {
	var loginReq req.SystemLoginReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &loginReq)) {
		return
	}
	res, err := lh.srv.Login(c, &loginReq)
	response.CheckAndRespWithData(c, res, err)
}

//logout 登录退出
func (lh loginHandler) logout(c *gin.Context) {
	var logoutReq req.SystemLogoutReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyHeader(c, &logoutReq)) {
		return
	}
	response.CheckAndResp(c, lh.srv.Logout(&logoutReq))
}
