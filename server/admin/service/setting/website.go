package setting

import (
	"likeadmin/admin/schemas/req"
	"likeadmin/core"
	"likeadmin/core/response"
	"likeadmin/util"
)

var SettingWebsiteService = settingWebsiteService{}

//settingWebsiteService 网站信息配置服务实现类
type settingWebsiteService struct{}

//Detail 获取网站信息
func (wSrv settingWebsiteService) Detail() map[string]string {
	data, err := util.ConfigUtil.Get("website")
	if err != nil {
		core.Logger.Errorf("Detail Get err: err=[%+v]", err)
		panic(response.SystemError)
	}
	return map[string]string{
		"name":     data["name"],
		"logo":     util.UrlUtil.ToAbsoluteUrl(data["logo"]),
		"favicon":  util.UrlUtil.ToAbsoluteUrl(data["favicon"]),
		"backdrop": util.UrlUtil.ToAbsoluteUrl(data["backdrop"]),
		"shopName": data["shopName"],
		"shopLogo": util.UrlUtil.ToAbsoluteUrl(data["shopLogo"]),
	}
}

//Save 保存网站信息
func (wSrv settingWebsiteService) Save(wsReq req.SettingWebsiteReq) {
	if err := util.ConfigUtil.Set("website", "name", wsReq.Name); err != nil {
		core.Logger.Errorf("Save Set name err: err=[%+v]", err)
		panic(response.SystemError)
	}
	if err := util.ConfigUtil.Set("website", "logo", util.UrlUtil.ToRelativeUrl(wsReq.Logo)); err != nil {
		core.Logger.Errorf("Save Set logo err: err=[%+v]", err)
		panic(response.SystemError)
	}
	if err := util.ConfigUtil.Set("website", "favicon", util.UrlUtil.ToRelativeUrl(wsReq.Favicon)); err != nil {
		core.Logger.Errorf("Save Set favicon err: err=[%+v]", err)
		panic(response.SystemError)
	}
	if err := util.ConfigUtil.Set("website", "backdrop", util.UrlUtil.ToRelativeUrl(wsReq.Backdrop)); err != nil {
		core.Logger.Errorf("Save Set backdrop err: err=[%+v]", err)
		panic(response.SystemError)
	}
	if err := util.ConfigUtil.Set("website", "shopName", wsReq.ShopName); err != nil {
		core.Logger.Errorf("Save Set shopName err: err=[%+v]", err)
		panic(response.SystemError)
	}
	if err := util.ConfigUtil.Set("website", "shopLogo", util.UrlUtil.ToRelativeUrl(wsReq.ShopLogo)); err != nil {
		core.Logger.Errorf("Save Set shopLogo err: err=[%+v]", err)
		panic(response.SystemError)
	}

}
