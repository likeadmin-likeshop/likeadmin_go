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
func (sSrv settingStorageService) List() ([]map[string]interface{}, error) {
	// TODO: engine默认local
	engine := "local"
	mapList := storageList
	for i := 0; i < len(mapList); i++ {
		if engine == mapList[i]["alias"] {
			mapList[i]["status"] = 1
		}
	}
	return mapList, nil
}

//Detail 存储详情
func (sSrv settingStorageService) Detail(alias string) (res map[string]interface{}, e error) {
	// TODO: engine默认local
	engine := "local"
	cnf, err := util.ConfigUtil.GetMap("storage", alias)
	if e = response.CheckErr(err, "Detail GetMap err"); e != nil {
		return
	}
	status := 0
	if engine == alias {
		status = 1
	}
	return map[string]interface{}{
		"name":   cnf["name"],
		"alias":  alias,
		"status": status,
	}, nil
}

//Edit 存储编辑
func (sSrv settingStorageService) Edit(editReq req.SettingStorageEditReq) (e error) {
	// TODO: engine默认local
	engine := "local"
	if engine != editReq.Alias {
		return response.Failed.Make(fmt.Sprintf("engine:%s 暂时不支持", editReq.Alias))
	}
	json, err := util.ToolsUtil.ObjToJson(map[string]interface{}{"name": "本地存储"})
	if e = response.CheckErr(err, "Edit ObjToJson err"); e != nil {
		return
	}
	err = util.ConfigUtil.Set("storage", editReq.Alias, json)
	if e = response.CheckErr(err, "Edit Set alias err"); e != nil {
		return
	}
	if editReq.Status == 1 {
		err = util.ConfigUtil.Set("storage", "default", editReq.Alias)
	} else {
		util.ConfigUtil.Set("storage", "default", "")
	}
	e = response.CheckErr(err, "Edit Set default err")
	return
}

//Change 存储切换
func (sSrv settingStorageService) Change(alias string, status int) (e error) {
	// TODO: engine默认local
	engine := "local"
	if engine != alias {
		return response.Failed.Make(fmt.Sprintf("engine:%s 暂时不支持", alias))
	}
	var err error
	if engine == alias && status == 0 {
		err = util.ConfigUtil.Set("storage", "default", "")
	} else {
		err = util.ConfigUtil.Set("storage", "default", alias)
	}
	e = response.CheckErr(err, "Change Set err")
	return
}
