package common

import (
	"gorm.io/gorm"
	"likeadmin/config"
	"likeadmin/core"
	"likeadmin/core/response"
	"likeadmin/util"
	"time"
)

type IIndexService interface {
	Console() (res map[string]interface{}, e error)
	Config() (res map[string]interface{}, e error)
}

//NewIndexService 初始化
func NewIndexService(db *gorm.DB) IIndexService {
	return &indexService{db: db}
}

//indexService 主页服务实现类
type indexService struct {
	db *gorm.DB
}

//Console 控制台数据
func (iSrv indexService) Console() (res map[string]interface{}, e error) {
	// 版本信息
	name, err := util.ConfigUtil.GetVal(iSrv.db, "website", "name", "LikeAdmin-Go")
	if e = response.CheckErr(err, "Console Get err"); e != nil {
		return
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
	}, nil
}

//Config 公共配置
func (iSrv indexService) Config() (res map[string]interface{}, e error) {
	website, err := util.ConfigUtil.Get(iSrv.db, "website")
	if e = response.CheckErr(err, "Config Get err"); e != nil {
		return
	}
	copyrightStr, err := util.ConfigUtil.GetVal(iSrv.db, "website", "copyright", "")
	if e = response.CheckErr(err, "Config GetVal err"); e != nil {
		return
	}
	var copyright []map[string]string
	if copyrightStr != "" {
		err = util.ToolsUtil.JsonToObj(copyrightStr, &copyright)
		if e = response.CheckErr(err, "Config JsonToObj err"); e != nil {
			return
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
	}, nil
}
