package common

import (
	"likeadmin/admin/schemas/req"
	"likeadmin/core"
	"likeadmin/core/response"
	"likeadmin/model/common"
)

var AlbumService = albumService{}

//albumService 相册服务实现类
type albumService struct{}

//AlbumAdd 相册文件新增
func (albSrv albumService) AlbumAdd(addReq req.CommonAlbumAddReq) uint {
	var alb common.Album
	//var params map[string]interface{}
	//if err := mapstructure.Decode(params, &alb); err != nil {
	//	core.Logger.Errorf("AlbumAdd Decode err: err=[%+v]", err)
	//	panic(response.SystemError)
	//}
	response.Copy(&alb, addReq)
	if err := core.DB.Create(&alb).Error; err != nil {
		core.Logger.Errorf("AlbumAdd Create err: err=[%+v]", err)
		panic(response.SystemError)
	}
	return alb.ID
}
