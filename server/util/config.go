package util

import (
	"errors"
	"gorm.io/gorm"
	"likeadmin/core"
	"likeadmin/model/system"
)

var ConfigUtil = configUtil{}

//convertUtil 数据库配置操作工具
type configUtil struct{}

//Get 根据类型和名称获取配置字典
func (cu configUtil) Get(cnfType string, names ...string) (data map[string]string, err error) {
	chain := core.DB.Where("type = ?", cnfType)
	if len(names) > 0 {
		chain.Where("name = ?", names[0])
	}
	var configs []system.SystemConfig
	err = chain.Find(&configs).Error
	if err != nil {
		return nil, err
	}
	data = make(map[string]string)
	for i := 0; i < len(configs); i++ {
		data[configs[i].Name] = configs[i].Value
	}
	return data, nil
}

//GetVal 根据类型和名称获取配置值
func (cu configUtil) GetVal(cnfType string, name string, defaultVal string) (data string, err error) {
	config, err := cu.Get(cnfType, name)
	if err != nil {
		return data, err
	}
	data, ok := config[name]
	if !ok {
		data = defaultVal
	}
	return data, nil
}

//GetMap 根据类型和名称获取配置值(Json字符串转dict)
func (cu configUtil) GetMap(cnfType string, name string) (data map[string]string, err error) {
	val, err := cu.GetVal(cnfType, name, "")
	if err != nil {
		return data, err
	}
	if val == "" {
		return map[string]string{}, nil
	}
	err = ToolsUtil.JsonToObj(val, &data)
	return data, err
}

//Set 设置配置的值
func (cu configUtil) Set(cnfType string, name string, val string) (err error) {
	var config system.SystemConfig
	err = core.DB.Where("type = ? AND name = ?", cnfType, name).First(&config).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		if err = core.DB.Create(&config).Error; err != nil {
			return err
		}
		return nil
	} else if err != nil {
		return err
	}
	if err = core.DB.Model(&config).Update("value", val).Error; err != nil {
		return err
	}
	return nil
}
