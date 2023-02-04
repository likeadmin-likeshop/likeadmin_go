package common

import (
	"likeadmin/admin/schemas/req"
	"likeadmin/admin/schemas/resp"
	"likeadmin/core/response"
	"likeadmin/plugin"
	"mime/multipart"
)

var UploadService = uploadService{}

//uploadService 上传服务实现类
type uploadService struct{}

//UploadImage 上传图片
func (upSrv uploadService) UploadImage(file *multipart.FileHeader, cid uint, aid uint) (res resp.CommonUploadFileResp, e error) {
	return upSrv.uploadFile(file, "image", 10, cid, aid)
}

//UploadVideo 上传视频
func (upSrv uploadService) UploadVideo(file *multipart.FileHeader, cid uint, aid uint) (res resp.CommonUploadFileResp, e error) {
	return upSrv.uploadFile(file, "video", 20, cid, aid)
}

//uploadFile 上传文件
func (upSrv uploadService) uploadFile(file *multipart.FileHeader, folder string, fileType int, cid uint, aid uint) (res resp.CommonUploadFileResp, e error) {
	var upRes *plugin.UploadFile
	if upRes, e = plugin.StorageDriver.Upload(file, folder, fileType); e != nil {
		return
	}
	var addReq req.CommonAlbumAddReq
	response.Copy(&addReq, upRes)
	addReq.Aid = aid
	addReq.Cid = cid
	var albumId uint
	if albumId, e = AlbumService.AlbumAdd(addReq); e != nil {
		return
	}
	response.Copy(&res, addReq)
	res.ID = albumId
	res.Path = upRes.Path
	return res, nil
}
