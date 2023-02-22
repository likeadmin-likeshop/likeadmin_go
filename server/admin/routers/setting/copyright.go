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

var CopyrightGroup = core.Group("/setting", newCopyrightHandler, regCopyright, middleware.TokenAuth())

func newCopyrightHandler(srv setting.ISettingCopyrightService) *copyrightHandler {
	return &copyrightHandler{srv: srv}
}

func regCopyright(rg *gin.RouterGroup, group *core.GroupBase) error {
	return group.Reg(func(handle *copyrightHandler) {
		rg.GET("/copyright/detail", handle.detail)
		rg.POST("/copyright/save", handle.save)
	})
}

type copyrightHandler struct {
	srv setting.ISettingCopyrightService
}

//detail 获取备案信息
func (ch copyrightHandler) detail(c *gin.Context) {
	res, err := ch.srv.Detail()
	response.CheckAndRespWithData(c, res, err)
}

//save 保存备案信息
func (ch copyrightHandler) save(c *gin.Context) {
	var cReqs []req.SettingCopyrightItemReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSONArray(c, &cReqs)) {
		return
	}
	response.CheckAndResp(c, ch.srv.Save(cReqs))
}
