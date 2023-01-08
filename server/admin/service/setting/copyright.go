package setting

import (
	"likeadmin/admin/schemas/req"
	"likeadmin/util"
)

var SettingCopyrightService = settingCopyrightService{}

//settingWebsiteService 网站备案服务实现类
type settingCopyrightService struct{}

//Detail 获取网站备案信息
func (cSrv settingCopyrightService) Detail() (res []map[string]interface{}) {
	data, err := util.ConfigUtil.GetVal("website", "copyright", "[]")
	util.CheckUtil.CheckErr(err, "Detail GetVal err")
	util.CheckUtil.CheckErr(util.ToolsUtil.JsonToObj(data, &res), "Detail JsonToObj err")
	return res
}

//Save 保存网站备案信息
func (cSrv settingCopyrightService) Save(cReqs []req.SettingCopyrightItemReq) {
	json, err := util.ToolsUtil.ObjToJson(cReqs)
	util.CheckUtil.CheckErr(err, "Save ObjToJson err")
	err = util.ConfigUtil.Set("website", "copyright", json)
	util.CheckUtil.CheckErr(err, "Save Set err")
}
