package utils

import (
	"likeadmin/config"
	"likeadmin/core"
	"net/url"
	"path"
	"strings"
)

var (
	UrlUtil      = urlUtil{}
	publicUrl    = config.Config.PublicUrl
	publicPrefix = config.Config.PublicPrefix
)

//urlUtil 文件路径处理工具
type urlUtil struct{}

//ToAbsoluteUrl 转绝对路径
func (uu urlUtil) ToAbsoluteUrl(u string) string {
	// TODO: engine默认local
	if u == "" {
		return ""
	}
	up, err := url.Parse(publicUrl)
	if err != nil {
		core.Logger.Errorf("ToAbsoluteUrl Parse err: err=[%+v]", err)
		return u
	}
	if strings.HasPrefix(u, "/api/static/") {
		up.Path = path.Join(up.Path, u)
		return up.String()
	}
	engine := "local"
	if engine == "local" {
		up.Path = path.Join(up.Path, publicPrefix, u)
		return up.String()
	}
	// TODO: 其他engine
	return u
}

func (uu urlUtil) ToRelativeUrl(u string) string {
	// TODO: engine默认local
	if u == "" {
		return ""
	}
	up, err := url.Parse(u)
	if err != nil {
		core.Logger.Errorf("ToRelativeUrl Parse err: err=[%+v]", err)
		return u
	}
	engine := "local"
	if engine == "local" {
		lu := up.String()
		return strings.Replace(
			strings.Replace(lu, publicUrl, "", 1),
			publicPrefix, "", 1)
	}
	// TODO: 其他engine
	return u
}
