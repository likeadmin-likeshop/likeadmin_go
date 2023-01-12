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

//storageList 存储列表
func storageList(c *gin.Context) {
	response.OkWithData(c, setting.SettingStorageService.List())
}

//storageDetail 存储详情
func storageDetail(c *gin.Context) {
	var detailReq req.SettingStorageDetailReq
	util.VerifyUtil.VerifyQuery(c, &detailReq)
	response.OkWithData(c, setting.SettingStorageService.Detail(detailReq.Alias))
}

//storageEdit 存储编辑
func storageEdit(c *gin.Context) {
	var editReq req.SettingStorageEditReq
	util.VerifyUtil.VerifyBody(c, &editReq)
	setting.SettingStorageService.Edit(editReq)
	response.Ok(c)
}

//storageChange 存储切换
func storageChange(c *gin.Context) {
	var changeReq req.SettingStorageChangeReq
	util.VerifyUtil.VerifyBody(c, &changeReq)
	setting.SettingStorageService.Change(changeReq.Alias, changeReq.Status)
	response.Ok(c)
}
