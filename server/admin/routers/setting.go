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
	group.AddGET("/storage/list", storageList)
	group.AddGET("/storage/detail", storageDetail)
	group.AddPOST("/storage/edit", storageEdit)
	group.AddPOST("/storage/change", storageChange)
}

//websiteDetail 获取网站信息
func websiteDetail(c *gin.Context) {
	res, err := setting.SettingWebsiteService.Detail()
	response.CheckAndRespWithData(c, res, err)
}

//websiteSave 保存网站信息
func websiteSave(c *gin.Context) {
	var wsReq req.SettingWebsiteReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &wsReq)) {
		return
	}
	response.CheckAndResp(c, setting.SettingWebsiteService.Save(wsReq))
}

//copyrightDetail 获取备案信息
func copyrightDetail(c *gin.Context) {
	res, err := setting.SettingCopyrightService.Detail()
	response.CheckAndRespWithData(c, res, err)
}

//copyrightSave 保存备案信息
func copyrightSave(c *gin.Context) {
	var cReqs []req.SettingCopyrightItemReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSONArray(c, &cReqs)) {
		return
	}
	response.CheckAndResp(c, setting.SettingCopyrightService.Save(cReqs))
}

//protocolDetail 获取政策信息
func protocolDetail(c *gin.Context) {
	res, err := setting.SettingProtocolService.Detail()
	response.CheckAndRespWithData(c, res, err)
}

//protocolSave 保存政策信息
func protocolSave(c *gin.Context) {
	var pReq req.SettingProtocolReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &pReq)) {
		return
	}
	response.CheckAndResp(c, setting.SettingProtocolService.Save(pReq))
}

//storageList 存储列表
func storageList(c *gin.Context) {
	res, err := setting.SettingStorageService.List()
	response.CheckAndRespWithData(c, res, err)
}

//storageDetail 存储详情
func storageDetail(c *gin.Context) {
	var detailReq req.SettingStorageDetailReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := setting.SettingStorageService.Detail(detailReq.Alias)
	response.CheckAndRespWithData(c, res, err)
}

//storageEdit 存储编辑
func storageEdit(c *gin.Context) {
	var editReq req.SettingStorageEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &editReq)) {
		return
	}
	response.CheckAndResp(c, setting.SettingStorageService.Edit(editReq))
}

//storageChange 存储切换
func storageChange(c *gin.Context) {
	var changeReq req.SettingStorageChangeReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &changeReq)) {
		return
	}
	response.CheckAndResp(c, setting.SettingStorageService.Change(changeReq.Alias, changeReq.Status))
}
