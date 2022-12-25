package utils

var ArrayUtil = arrayUtil{}

//arrayUtil 数组工具类
type arrayUtil struct{}

//ListToTree 字典列表转树形结构
func (au arrayUtil) ListToTree(arr []map[string]interface{}, id string, pid string, child string) (mapList []interface{}) {
	// 遍历以id_为key生成map
	idValMap := make(map[uint]interface{})
	for _, m := range arr {
		if idVal, ok := m[id]; ok {
			idValMap[idVal.(uint)] = m
		}
	}
	// 遍历
	for _, m := range arr {
		// 获取父节点
		if pidVal, ok := m[pid]; ok {
			if pNode, pok := idValMap[pidVal.(uint)]; pok {
				// 有父节点则添加到父节点子集
				if cVal, cok := pNode.(map[string]interface{})[child]; cok {
					if cVal == nil {
						cVal = []interface{}{m}
					} else {
						cVal = append(cVal.([]interface{}), m)
					}
					pNode.(map[string]interface{})[child] = cVal
					continue
				}
			}
		}
		mapList = append(mapList, m)
	}
	return
}
