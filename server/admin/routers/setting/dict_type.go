package setting

import (
	"github.com/gin-gonic/gin"
	"likeadmin/admin/schemas/req"
	"likeadmin/admin/service/setting"
	"likeadmin/core"
	"likeadmin/core/request"
	"likeadmin/core/response"
	"likeadmin/middleware"
	"likeadmin/util"
)

var DictTypeGroup = core.Group("/setting", newDictTypeHandler, regDictType, middleware.TokenAuth())

func newDictTypeHandler(srv setting.ISettingDictTypeService) *dictTypeHandler {
	return &dictTypeHandler{srv: srv}
}

func regDictType(rg *gin.RouterGroup, group *core.GroupBase) error {
	return group.Reg(func(handle *dictTypeHandler) {
		rg.GET("/dict/type/all", handle.all)
		rg.GET("/dict/type/list", handle.list)
		rg.GET("/dict/type/detail", handle.detail)
		rg.POST("/dict/type/add", handle.add)
	})
}

type dictTypeHandler struct {
	srv setting.ISettingDictTypeService
}

//all 字典类型所有
func (dth dictTypeHandler) all(c *gin.Context) {
	res, err := dth.srv.All()
	response.CheckAndRespWithData(c, res, err)
}

//list 字典类型列表
func (dth dictTypeHandler) list(c *gin.Context) {
	var page request.PageReq
	var listReq req.SettingDictTypeListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := dth.srv.List(page, listReq)
	response.CheckAndRespWithData(c, res, err)
}

//detail 字典类型详情
func (dth dictTypeHandler) detail(c *gin.Context) {
	var detailReq req.SettingDictTypeDetailReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := dth.srv.Detail(detailReq.ID)
	response.CheckAndRespWithData(c, res, err)
}

//detail 字典类型新增
func (dth dictTypeHandler) add(c *gin.Context) {
	var addReq req.SettingDictTypeAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &addReq)) {
		return
	}
	response.CheckAndResp(c, dth.srv.Add(addReq))
}
