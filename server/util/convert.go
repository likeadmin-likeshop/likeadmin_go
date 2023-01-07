package util

import (
	"github.com/fatih/structs"
	"github.com/jinzhu/copier"
	"likeadmin/core"
)

var ConvertUtil = convertUtil{}

//convertUtil 转换工具
type convertUtil struct{}

//StructsToMaps 将结构体转换成Map列表
func (cu convertUtil) StructsToMaps(objs interface{}) (data []map[string]interface{}) {
	var objList []interface{}
	err := copier.Copy(&objList, objs)
	if err != nil {
		core.Logger.Errorf("convertUtil.StructsToMaps err: err=[%+v]", err)
		return nil
	}
	for _, v := range objList {
		data = append(data, structs.Map(v))
	}
	return data
}
