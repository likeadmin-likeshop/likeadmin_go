package routers

import (
	"github.com/gin-gonic/gin"
	"likeadmin/admin/schemas/req"
	"likeadmin/admin/service/setting"
	"likeadmin/core"
	"likeadmin/core/response"
	"likeadmin/util"
)

var SettingGroup = core.Group("/setting")

func init() {
	group := SettingGroup
	group.AddGET("/website/detail", websiteDetail)
	group.AddPOST("/website/save", websiteSave)
	group.AddGET("/copyright/detail", copyrightDetail)
	group.AddPOST("/copyright/save", copyrightSave)
	group.AddGET("/protocol/detail", protocolDetail)
	group.AddPOST("/protocol/save", protocolSave)
}

//websiteDetail 获取网站信息
func websiteDetail(c *gin.Context) {
	response.OkWithData(c, setting.SettingWebsiteService.Detail())
}

//websiteSave 保存网站信息
func websiteSave(c *gin.Context) {
	var wsReq req.SettingWebsiteReq
	util.VerifyUtil.VerifyJSON(c, &wsReq)
	setting.SettingWebsiteService.Save(wsReq)
	response.Ok(c)
}

//copyrightDetail 获取备案信息
func copyrightDetail(c *gin.Context) {
	response.OkWithData(c, setting.SettingCopyrightService.Detail())
}

//copyrightSave 保存备案信息
func copyrightSave(c *gin.Context) {
	var cReqs []req.SettingCopyrightItemReq
	util.VerifyUtil.VerifyJSONArray(c, &cReqs)
	setting.SettingCopyrightService.Save(cReqs)
	response.Ok(c)
}

//protocolDetail 获取政策信息
func protocolDetail(c *gin.Context) {
	response.OkWithData(c, setting.SettingProtocolService.Detail())
}

//protocolSave 保存政策信息
func protocolSave(c *gin.Context) {
	var pReq req.SettingProtocolReq
	util.VerifyUtil.VerifyJSON(c, &pReq)
	setting.SettingProtocolService.Save(pReq)
	response.Ok(c)
}
