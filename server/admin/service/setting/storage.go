package setting

import "likeadmin/util"

var SettingStorageService = settingStorageService{}

//settingWebsiteService 存储配置服务实现类
type settingStorageService struct{}

var storageList = []map[string]interface{}{
	{"name": "本地存储", "alias": "local", "describe": "存储在本地服务器", "status": 0},
}

//List 存储列表
func (sSrv settingStorageService) List() []map[string]interface{} {
	// TODO: engine默认local
	engine := "local"
	mapList := storageList
	for i := 0; i < len(mapList); i++ {
		if engine == mapList[i]["alias"] {
			mapList[i]["status"] = 1
		}
	}
	return mapList
}

//Detail 存储详情
func (sSrv settingStorageService) Detail(alias string) map[string]interface{} {
	// TODO: engine默认local
	engine := "local"
	cnf, err := util.ConfigUtil.GetMap("storage", alias)
	util.CheckUtil.CheckErr(err, "Detail GetMap err")
	status := 0
	if engine == alias {
		status = 1
	}
	return map[string]interface{}{
		"name":   cnf["name"],
		"alias":  alias,
		"status": status,
	}
}
