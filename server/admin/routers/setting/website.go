package setting

import (
	"github.com/gin-gonic/gin"
	"likeadmin/admin/schemas/req"
	"likeadmin/admin/service/setting"
	"likeadmin/core"
	"likeadmin/core/response"
	"likeadmin/middleware"
	"likeadmin/util"
)

var WebsiteGroup = core.Group("/setting", newWebsiteHandler, regWebsite, middleware.TokenAuth())

func newWebsiteHandler(srv setting.ISettingWebsiteService) *websiteHandler {
	return &websiteHandler{srv: srv}
}

func regWebsite(rg *gin.RouterGroup, group *core.GroupBase) error {
	return group.Reg(func(handle *websiteHandler) {
		rg.GET("/website/detail", handle.detail)
		rg.POST("/website/save", handle.save)
	})
}

type websiteHandler struct {
	srv setting.ISettingWebsiteService
}

//detail 获取网站信息
func (wh websiteHandler) detail(c *gin.Context) {
	res, err := wh.srv.Detail()
	response.CheckAndRespWithData(c, res, err)
}

//save 保存网站信息
func (wh websiteHandler) save(c *gin.Context) {
	var wsReq req.SettingWebsiteReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &wsReq)) {
		return
	}
	response.CheckAndResp(c, wh.srv.Save(wsReq))
}
