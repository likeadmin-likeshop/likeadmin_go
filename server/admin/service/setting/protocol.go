package setting

import (
	"likeadmin/admin/schemas/req"
	"likeadmin/util"
)

var SettingProtocolService = settingProtocolService{}

//settingProtocolService 政策协议服务实现类
type settingProtocolService struct{}

//Detail 获取政策协议信息
func (cSrv settingProtocolService) Detail() map[string]interface{} {
	defaultVal := `{"name":"","content":""}`
	json, err := util.ConfigUtil.GetVal("protocol", "service", defaultVal)
	util.CheckUtil.CheckErr(err, "Detail GetVal service err")
	var service map[string]interface{}
	util.CheckUtil.CheckErr(util.ToolsUtil.JsonToObj(json, &service), "Detail JsonToObj service err")
	json, err = util.ConfigUtil.GetVal("protocol", "privacy", defaultVal)
	util.CheckUtil.CheckErr(err, "Detail GetVal privacy err")
	var privacy map[string]interface{}
	util.CheckUtil.CheckErr(util.ToolsUtil.JsonToObj(json, &privacy), "Detail JsonToObj privacy err")
	return map[string]interface{}{
		"service": service,
		"privacy": privacy,
	}
}

//Save 保存政策协议信息
func (cSrv settingProtocolService) Save(pReq req.SettingProtocolReq) {
	serviceJson, err := util.ToolsUtil.ObjToJson(pReq.Service)
	util.CheckUtil.CheckErr(err, "Save ObjToJson service err")
	privacyJson, err := util.ToolsUtil.ObjToJson(pReq.Privacy)
	util.CheckUtil.CheckErr(err, "Save ObjToJson privacy err")
	err = util.ConfigUtil.Set("protocol", "service", serviceJson)
	util.CheckUtil.CheckErr(err, "Save Set service err")
	err = util.ConfigUtil.Set("protocol", "privacy", privacyJson)
	util.CheckUtil.CheckErr(err, "Save Set privacy err")
}
