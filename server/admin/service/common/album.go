package common

import (
	"github.com/mitchellh/mapstructure"
	"likeadmin/core"
	"likeadmin/core/response"
	"likeadmin/model/common"
)

var AlbumService = albumService{}

//albumService 相册服务实现类
type albumService struct{}

//AlbumAdd 相册文件新增
func (albSrv albumService) AlbumAdd(params map[string]interface{}) uint {
	// TODO: AlbumAdd
	var alb common.Album
	if err := mapstructure.Decode(params, &alb); err != nil {
		core.Logger.Errorf("AlbumAdd Decode err: err=[%+v]", err)
		panic(response.SystemError)
	}
	if err := core.DB.Create(&alb).Error; err != nil {
		core.Logger.Errorf("AlbumAdd Create err: err=[%+v]", err)
		panic(response.SystemError)
	}
	return alb.ID
}
