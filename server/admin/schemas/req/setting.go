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
