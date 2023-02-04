package setting

import (
	"likeadmin/admin/schemas/req"
	"likeadmin/core/response"
	"likeadmin/util"
)

var SettingWebsiteService = settingWebsiteService{}

//settingWebsiteService 网站信息配置服务实现类
type settingWebsiteService struct{}

//Detail 获取网站信息
func (wSrv settingWebsiteService) Detail() (res map[string]string, e error) {
	data, err := util.ConfigUtil.Get("website")
	if e = response.CheckErr(err, "Detail Get err"); e != nil {
		return
	}
	return map[string]string{
		"name":     data["name"],
		"logo":     util.UrlUtil.ToAbsoluteUrl(data["logo"]),
		"favicon":  util.UrlUtil.ToAbsoluteUrl(data["favicon"]),
		"backdrop": util.UrlUtil.ToAbsoluteUrl(data["backdrop"]),
		"shopName": data["shopName"],
		"shopLogo": util.UrlUtil.ToAbsoluteUrl(data["shopLogo"]),
	}, nil
}

//Save 保存网站信息
func (wSrv settingWebsiteService) Save(wsReq req.SettingWebsiteReq) (e error) {
	err := util.ConfigUtil.Set("website", "name", wsReq.Name)
	if e = response.CheckErr(err, "Save Set name err"); e != nil {
		return
	}
	err = util.ConfigUtil.Set("website", "logo", util.UrlUtil.ToRelativeUrl(wsReq.Logo))
	if e = response.CheckErr(err, "Save Set logo err"); e != nil {
		return
	}
	err = util.ConfigUtil.Set("website", "favicon", util.UrlUtil.ToRelativeUrl(wsReq.Favicon))
	if e = response.CheckErr(err, "Save Set favicon err"); e != nil {
		return
	}
	err = util.ConfigUtil.Set("website", "backdrop", util.UrlUtil.ToRelativeUrl(wsReq.Backdrop))
	if e = response.CheckErr(err, "Save Set backdrop err"); e != nil {
		return
	}
	err = util.ConfigUtil.Set("website", "shopName", wsReq.ShopName)
	if e = response.CheckErr(err, "Save Set shopName err"); e != nil {
		return
	}
	err = util.ConfigUtil.Set("website", "shopLogo", util.UrlUtil.ToRelativeUrl(wsReq.ShopLogo))
	e = response.CheckErr(err, "Save Set shopLogo err")
	return
}
