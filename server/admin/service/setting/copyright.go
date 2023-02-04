package setting

import (
	"likeadmin/admin/schemas/req"
	"likeadmin/core/response"
	"likeadmin/util"
)

var SettingCopyrightService = settingCopyrightService{}

//settingCopyrightService 网站备案服务实现类
type settingCopyrightService struct{}

//Detail 获取网站备案信息
func (cSrv settingCopyrightService) Detail() (res []map[string]interface{}, e error) {
	data, err := util.ConfigUtil.GetVal("website", "copyright", "[]")
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
	err = util.ConfigUtil.Set("website", "copyright", json)
	e = response.CheckErr(err, "Save Set err")
	return
}
