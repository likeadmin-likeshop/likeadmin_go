package setting

import (
	"gorm.io/gorm"
	"likeadmin/admin/schemas/req"
	"likeadmin/core/response"
	"likeadmin/util"
)

type ISettingWebsiteService interface {
	Detail() (res map[string]string, e error)
	Save(wsReq req.SettingWebsiteReq) (e error)
}

//NewSettingWebsiteService 初始化
func NewSettingWebsiteService(db *gorm.DB) ISettingWebsiteService {
	return &SettingWebsiteService{db: db}
}

//SettingWebsiteService 网站信息配置服务实现类
type SettingWebsiteService struct {
	db *gorm.DB
}

//Detail 获取网站信息
func (wSrv SettingWebsiteService) Detail() (res map[string]string, e error) {
	data, err := util.ConfigUtil.Get(wSrv.db, "website")
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
func (wSrv SettingWebsiteService) Save(wsReq req.SettingWebsiteReq) (e error) {
	err := util.ConfigUtil.Set(wSrv.db, "website", "name", wsReq.Name)
	if e = response.CheckErr(err, "Save Set name err"); e != nil {
		return
	}
	err = util.ConfigUtil.Set(wSrv.db, "website", "logo", util.UrlUtil.ToRelativeUrl(wsReq.Logo))
	if e = response.CheckErr(err, "Save Set logo err"); e != nil {
		return
	}
	err = util.ConfigUtil.Set(wSrv.db, "website", "favicon", util.UrlUtil.ToRelativeUrl(wsReq.Favicon))
	if e = response.CheckErr(err, "Save Set favicon err"); e != nil {
		return
	}
	err = util.ConfigUtil.Set(wSrv.db, "website", "backdrop", util.UrlUtil.ToRelativeUrl(wsReq.Backdrop))
	if e = response.CheckErr(err, "Save Set backdrop err"); e != nil {
		return
	}
	err = util.ConfigUtil.Set(wSrv.db, "website", "shopName", wsReq.ShopName)
	if e = response.CheckErr(err, "Save Set shopName err"); e != nil {
		return
	}
	err = util.ConfigUtil.Set(wSrv.db, "website", "shopLogo", util.UrlUtil.ToRelativeUrl(wsReq.ShopLogo))
	e = response.CheckErr(err, "Save Set shopLogo err")
	return
}
