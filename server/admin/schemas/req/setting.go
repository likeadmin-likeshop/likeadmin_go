package req

//SettingWebsiteReq 保存网站信息参数
type SettingWebsiteReq struct {
	Name     string `form:"name"`     // 网站名称
	Logo     string `form:"logo"`     // 网站图标
	Favicon  string `form:"favicon"`  // 网站LOGO
	Backdrop string `form:"backdrop"` // 登录页广告图
	ShopName string `form:"shopName"` // 商城名称
	ShopLogo string `form:"shopLogo"` // 商城Logo
}

//SettingCopyrightItemReq 保存备案信息参数
type SettingCopyrightItemReq struct {
	Name string `form:"name" json:"name"`  // 名称
	Link string `form:"link"  json:"link"` // 链接
}

//SettingProtocolItem 政策通用参数
type SettingProtocolItem struct {
	Name    string `form:"name" json:"name"`        // 名称
	Content string `form:"content"  json:"content"` // 内容
}

//SettingProtocolReq 保存政策信息参数
type SettingProtocolReq struct {
	Service SettingProtocolItem `form:"service" json:"service"`  // 服务协议
	Privacy SettingProtocolItem `form:"privacy"  json:"privacy"` // 隐私协议
}

//SettingStorageDetailReq 存储详情参数
type SettingStorageDetailReq struct {
	Alias string `form:"alias" binding:"required,oneof=local qiniu qcloud aliyun"` // 别名: [local,qiniu,qcloud,aliyun]
}

//SettingStorageEditReq 存储编辑参数
type SettingStorageEditReq struct {
	Alias     string `form:"alias" binding:"required,oneof=local qiniu qcloud aliyun"` // 别名: [local,qiniu,qcloud,aliyun]
	Status    int    `form:"status" binding:"oneof=0 1"`                               // 状态: 0/1
	Bucket    string `form:"bucket"`                                                   // 存储空间名
	SecretKey string `form:"secretKey"`                                                // SK
	AccessKey string `form:"accessKey"`                                                // AK
	Domain    string `form:"domain"`                                                   // 访问域名
	Region    string `form:"region"`                                                   // 地区,腾讯存储特有
}

//SettingStorageChangeReq 存储切换参数
type SettingStorageChangeReq struct {
	Alias  string `form:"alias" binding:"required,oneof=local qiniu qcloud aliyun"` // 别名: [local,qiniu,qcloud,aliyun]
	Status int    `form:"status" binding:"oneof=0 1"`                               // 状态: 0/1
}

//SettingDictTypeListReq 字典类型新增参数
type SettingDictTypeListReq struct {
	DictName   string `form:"dictName" binding:"max=200"`                   // 字典名称
	DictType   string `form:"dictType" binding:"max=200"`                   // 字典类型
	DictStatus int8   `form:"dictStatus,default=-1" binding:"oneof=-1 0 1"` // 字典状态: 0/1
}

//SettingDictTypeDetailReq 字典类型详情参数
type SettingDictTypeDetailReq struct {
	ID uint `form:"id" binding:"required,gt=0"` // 主键
}

//SettingDictTypeAddReq 字典类型新增参数
type SettingDictTypeAddReq struct {
	DictName   string `form:"dictName" binding:"required,max=200"`     // 字典名称
	DictType   string `form:"dictType" binding:"required,max=200"`     // 字典类型
	DictRemark string `form:"dictRemark" binding:"max=200"`            // 字典备注
	DictStatus int8   `form:"dictStatus" binding:"required,oneof=0 1"` // 字典状态: 0/1
}

//SettingDictTypeEditReq 字典类型编辑参数
type SettingDictTypeEditReq struct {
	ID         uint   `form:"id" binding:"required,gt=0"`              // 主键
	DictName   string `form:"dictName" binding:"required,max=200"`     // 字典名称
	DictType   string `form:"dictType" binding:"required,max=200"`     // 字典类型
	DictRemark string `form:"dictRemark" binding:"max=200"`            // 字典备注
	DictStatus int8   `form:"dictStatus" binding:"required,oneof=0 1"` // 字典状态: 0/1
}

//SettingDictTypeDelReq 字典类型删除参数
type SettingDictTypeDelReq struct {
	Ids []uint `form:"ids" binding:"required"` // 主键列表
}

//SettingDictDataListReq 字典数据列表参数
type SettingDictDataListReq struct {
	DictType string `form:"dictType" binding:"max=200"`               // 字典类型
	Name     string `form:"name" binding:"max=100"`                   // 键
	Value    string `form:"value" binding:"max=200"`                  // 值
	Status   int8   `form:"status,default=-1" binding:"oneof=-1 0 1"` // 状态: 0=停用,1=启用
}

//SettingDictDataDetailReq 字典数据详情参数
type SettingDictDataDetailReq struct {
	ID uint `form:"id" binding:"required,gt=0"` // 主键
}

//SettingDictDataAddReq 字典数据新增参数
type SettingDictDataAddReq struct {
	TypeId uint   `form:"typeId" binding:"required,gt=0"`           // 类型
	Name   string `form:"name" binding:"required,max=100"`          // 键
	Value  string `form:"value" binding:"required,max=200"`         // 值
	remark string `form:"remark" binding:"max=200"`                 // 备注
	Sort   int    `form:"sort" binding:"gte=0"`                     // 排序
	Status int8   `form:"status,default=-1" binding:"oneof=-1 0 1"` // 状态: 0=停用,1=启用
}

//SettingDictDataEditReq 字典数据编辑参数
type SettingDictDataEditReq struct {
	ID     uint   `form:"id" binding:"required,gt=0"`               // 主键
	TypeId uint   `form:"typeId" binding:"required,gte=0"`          // 类型
	Name   string `form:"name" binding:"required,max=100"`          // 键
	Value  string `form:"value" binding:"required,max=200"`         // 值
	remark string `form:"remark" binding:"max=200"`                 // 备注
	Sort   int    `form:"sort" binding:"gte=0"`                     // 排序
	Status int8   `form:"status,default=-1" binding:"oneof=-1 0 1"` // 状态: 0=停用,1=启用
}

//SettingDictDataDelReq 字典数据删除参数
type SettingDictDataDelReq struct {
	Ids []uint `form:"ids" binding:"required"` // 主键列表
}
