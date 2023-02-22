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

var StorageGroup = core.Group("/setting", newStorageHandler, regStorage, middleware.TokenAuth())

func newStorageHandler(srv setting.ISettingStorageService) *storageHandler {
	return &storageHandler{srv: srv}
}

func regStorage(rg *gin.RouterGroup, group *core.GroupBase) error {
	return group.Reg(func(handle *storageHandler) {
		rg.GET("/storage/list", handle.list)
		rg.GET("/storage/detail", handle.detail)
		rg.POST("/storage/edit", handle.edit)
		rg.POST("/storage/change", handle.change)
	})
}

type storageHandler struct {
	srv setting.ISettingStorageService
}

//list 存储列表
func (sh storageHandler) list(c *gin.Context) {
	res, err := sh.srv.List()
	response.CheckAndRespWithData(c, res, err)
}

//detail 存储详情
func (sh storageHandler) detail(c *gin.Context) {
	var detailReq req.SettingStorageDetailReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := sh.srv.Detail(detailReq.Alias)
	response.CheckAndRespWithData(c, res, err)
}

//edit 存储编辑
func (sh storageHandler) edit(c *gin.Context) {
	var editReq req.SettingStorageEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &editReq)) {
		return
	}
	response.CheckAndResp(c, sh.srv.Edit(editReq))
}

//change 存储切换
func (sh storageHandler) change(c *gin.Context) {
	var changeReq req.SettingStorageChangeReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &changeReq)) {
		return
	}
	response.CheckAndResp(c, sh.srv.Change(changeReq.Alias, changeReq.Status))
}
