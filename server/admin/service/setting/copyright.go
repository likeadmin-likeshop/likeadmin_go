package setting

import (
	"gorm.io/gorm"
	"likeadmin/admin/schemas/req"
	"likeadmin/core/response"
	"likeadmin/util"
)

type ISettingCopyrightService interface {
	Detail() (res []map[string]interface{}, e error)
	Save(cReqs []req.SettingCopyrightItemReq) (e error)
}

//NewSettingCopyrightService 初始化
func NewSettingCopyrightService(db *gorm.DB) ISettingCopyrightService {
	return &settingCopyrightService{db: db}
}

//settingCopyrightService 网站备案服务实现类
type settingCopyrightService struct {
	db *gorm.DB
}

//Detail 获取网站备案信息
func (cSrv settingCopyrightService) Detail() (res []map[string]interface{}, e error) {
	data, err := util.ConfigUtil.GetVal(cSrv.db, "website", "copyright", "[]")
	if e = response.CheckErr(err, "Detail GetVal err"); e != nil {
		return
	}
	e = response.CheckErr(util.ToolsUtil.JsonToObj(data, &res), "Detail JsonToObj err")
	return
}

//Save 保存网站备案信息
func (cSrv settingCopyrightService) Save(cReqs []req.SettingCopyrightItemReq) (e error) {
	json, err := util.ToolsUtil.ObjToJson(cReqs)
	if e = response.CheckErr(err, "Save ObjToJson err"); e != nil {
		return
	}
	err = util.ConfigUtil.Set(cSrv.db, "website", "copyright", json)
	e = response.CheckErr(err, "Save Set err")
	return
}
