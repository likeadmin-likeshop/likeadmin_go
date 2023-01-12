package common

import (
	"likeadmin/config"
	"likeadmin/core"
	"likeadmin/core/response"
	"likeadmin/util"
	"time"
)

var IndexService = indexService{}

//indexService 主页服务实现类
type indexService struct{}

//Console 控制台数据
func (iSrv indexService) Console() map[string]interface{} {
	// 版本信息
	name, err := util.ConfigUtil.GetVal("website", "name", "LikeAdmin-Go")
	if err != nil {
		core.Logger.Errorf("Console Get err: err=[%+v]", err)
		panic(response.SystemError)
	}
	version := map[string]interface{}{
		"name":    name,
		"version": config.Config.Version,
		"website": "www.likeadmin.cn",
		"based":   "Vue3.x、ElementUI、MySQL",
		"channel": map[string]string{
			"gitee":   "https://gitee.com/likeadmin/likeadmin_python",
			"website": "https://www.likeadmin.cn",
		},
	}
	// 今日数据
	today := map[string]interface{}{
		"time":        "2022-08-11 15:08:29",
		"todayVisits": 10,  // 访问量(人)
		"totalVisits": 100, // 总访问量
		"todaySales":  30,  // 销售额(元)
		"totalSales":  65,  // 总销售额
		"todayOrder":  12,  // 订单量(笔)
		"totalOrder":  255, // 总订单量
		"todayUsers":  120, // 新增用户
		"totalUsers":  360, // 总访用户
	}
	// 访客图表
	now := time.Now()
	var date []string
	for i := 14; i >= 0; i-- {
		date = append(date, now.AddDate(0, 0, -i).Format(core.DateFormat))
	}
	visitor := map[string]interface{}{
		"date": date,
		"list": []int{12, 13, 11, 5, 8, 22, 14, 9, 456, 62, 78, 12, 18, 22, 46},
	}
	return map[string]interface{}{
		"version": version,
		"today":   today,
		"visitor": visitor,
	}
}

//Config 公共配置
func (iSrv indexService) Config() map[string]interface{} {
	website, err := util.ConfigUtil.Get("website")
	if err != nil {
		core.Logger.Errorf("Config Get err: err=[%+v]", err)
		panic(response.SystemError)
	}
	copyrightStr, err := util.ConfigUtil.GetVal("website", "copyright", "")
	if err != nil {
		core.Logger.Errorf("Config GetVal err: err=[%+v]", err)
		panic(response.SystemError)
	}
	var copyright []map[string]string
	if copyrightStr != "" {
		if err = util.ToolsUtil.JsonToObj(copyrightStr, &copyright); err != nil {
			core.Logger.Errorf("Config JsonToObj err: err=[%+v]", err)
			panic(response.SystemError)
		}
	} else {
		copyright = []map[string]string{}
	}
	return map[string]interface{}{
		"webName":     website["name"],
		"webLogo":     util.UrlUtil.ToAbsoluteUrl(website["logo"]),
		"webFavicon":  util.UrlUtil.ToAbsoluteUrl(website["favicon"]),
		"webBackdrop": util.UrlUtil.ToAbsoluteUrl(website["backdrop"]),
		"ossDomain":   config.Config.PublicUrl,
		"copyright":   copyright,
	}
}
