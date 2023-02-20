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
	response.DI(c, func(srv *setting.SettingWebsiteService) {
		res, err := srv.Detail()
		response.CheckAndRespWithData(c, res, err)
	})
}

//websiteSave 保存网站信息
func websiteSave(c *gin.Context) {
	response.DI(c, func(srv *setting.SettingWebsiteService) {
		var wsReq req.SettingWebsiteReq
		if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &wsReq)) {
			return
		}
		response.CheckAndResp(c, srv.Save(wsReq))
	})
}

//copyrightDetail 获取备案信息
func copyrightDetail(c *gin.Context) {
	response.DI(c, func(srv *setting.SettingCopyrightService) {
		res, err := srv.Detail()
		response.CheckAndRespWithData(c, res, err)
	})
}

//copyrightSave 保存备案信息
func copyrightSave(c *gin.Context) {
	response.DI(c, func(srv *setting.SettingCopyrightService) {
		var cReqs []req.SettingCopyrightItemReq
		if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSONArray(c, &cReqs)) {
			return
		}
		response.CheckAndResp(c, srv.Save(cReqs))
	})
}

//protocolDetail 获取政策信息
func protocolDetail(c *gin.Context) {
	response.DI(c, func(srv *setting.SettingProtocolService) {
		res, err := srv.Detail()
		response.CheckAndRespWithData(c, res, err)
	})
}

//protocolSave 保存政策信息
func protocolSave(c *gin.Context) {
	response.DI(c, func(srv *setting.SettingProtocolService) {
		var pReq req.SettingProtocolReq
		if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &pReq)) {
			return
		}
		response.CheckAndResp(c, srv.Save(pReq))
	})
}

//storageList 存储列表
func storageList(c *gin.Context) {
	response.DI(c, func(srv *setting.SettingStorageService) {
		res, err := srv.List()
		response.CheckAndRespWithData(c, res, err)
	})
}

//storageDetail 存储详情
func storageDetail(c *gin.Context) {
	response.DI(c, func(srv *setting.SettingStorageService) {
		var detailReq req.SettingStorageDetailReq
		if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
			return
		}
		res, err := srv.Detail(detailReq.Alias)
		response.CheckAndRespWithData(c, res, err)
	})
}

//storageEdit 存储编辑
func storageEdit(c *gin.Context) {
	response.DI(c, func(srv *setting.SettingStorageService) {
		var editReq req.SettingStorageEditReq
		if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &editReq)) {
			return
		}
		response.CheckAndResp(c, srv.Edit(editReq))
	})
}

//storageChange 存储切换
func storageChange(c *gin.Context) {
	response.DI(c, func(srv *setting.SettingStorageService) {
		var changeReq req.SettingStorageChangeReq
		if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &changeReq)) {
			return
		}
		response.CheckAndResp(c, srv.Change(changeReq.Alias, changeReq.Status))
	})
}
