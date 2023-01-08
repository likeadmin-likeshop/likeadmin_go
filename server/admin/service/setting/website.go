package setting

import (
	"likeadmin/admin/schemas/req"
	"likeadmin/util"
)

var SettingWebsiteService = settingWebsiteService{}

//settingWebsiteService 网站信息配置服务实现类
type settingWebsiteService struct{}

//Detail 获取网站信息
func (wSrv settingWebsiteService) Detail() map[string]string {
	data, err := util.ConfigUtil.Get("website")
	util.CheckUtil.CheckErr(err, "Detail Get err")
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
	err := util.ConfigUtil.Set("website", "name", wsReq.Name)
	util.CheckUtil.CheckErr(err, "Save Set name err")
	err = util.ConfigUtil.Set("website", "logo", util.UrlUtil.ToRelativeUrl(wsReq.Logo))
	util.CheckUtil.CheckErr(err, "Save Set logo err")
	err = util.ConfigUtil.Set("website", "favicon", util.UrlUtil.ToRelativeUrl(wsReq.Favicon))
	util.CheckUtil.CheckErr(err, "Save Set favicon err")
	err = util.ConfigUtil.Set("website", "backdrop", util.UrlUtil.ToRelativeUrl(wsReq.Backdrop))
	util.CheckUtil.CheckErr(err, "Save Set backdrop err")
	err = util.ConfigUtil.Set("website", "shopName", wsReq.ShopName)
	util.CheckUtil.CheckErr(err, "Save Set shopName err")
	err = util.ConfigUtil.Set("website", "shopLogo", util.UrlUtil.ToRelativeUrl(wsReq.ShopLogo))
	util.CheckUtil.CheckErr(err, "Save Set shopLogo err")
}
