package setting

import (
	"fmt"
	"likeadmin/admin/schemas/req"
	"likeadmin/core/response"
	"likeadmin/util"
)

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

//Edit 存储编辑
func (sSrv settingStorageService) Edit(editReq req.SettingStorageEditReq) {
	// TODO: engine默认local
	engine := "local"
	if engine != editReq.Alias {
		panic(response.Failed.Make(fmt.Sprintf("engine:%s 暂时不支持", editReq.Alias)))
	}
	json, err := util.ToolsUtil.ObjToJson(map[string]interface{}{"name": "本地存储"})
	util.CheckUtil.CheckErr(err, "Edit ObjToJson err")
	err = util.ConfigUtil.Set("storage", editReq.Alias, json)
	util.CheckUtil.CheckErr(err, "Edit Set alias err")
	if editReq.Status == 1 {
		err = util.ConfigUtil.Set("storage", "default", editReq.Alias)
	} else {
		util.ConfigUtil.Set("storage", "default", "")
	}
	util.CheckUtil.CheckErr(err, "Edit Set default err")
}

//Change 存储切换
func (sSrv settingStorageService) Change(alias string, status int) {
	// TODO: engine默认local
	engine := "local"
	if engine != alias {
		panic(response.Failed.Make(fmt.Sprintf("engine:%s 暂时不支持", alias)))
	}
	var err error
	if engine == alias && status == 0 {
		err = util.ConfigUtil.Set("storage", "default", "")
	} else {
		err = util.ConfigUtil.Set("storage", "default", alias)
	}
	util.CheckUtil.CheckErr(err, "Change Set err")
}
