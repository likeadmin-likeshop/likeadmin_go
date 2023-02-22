package common

import (
	"likeadmin/admin/schemas/req"
	"likeadmin/admin/schemas/resp"
	"likeadmin/core/response"
	"likeadmin/plugin"
	"mime/multipart"
)

type IUploadService interface {
	UploadImage(file *multipart.FileHeader, cid uint, aid uint) (res resp.CommonUploadFileResp, e error)
	UploadVideo(file *multipart.FileHeader, cid uint, aid uint) (res resp.CommonUploadFileResp, e error)
}

//NewUploadService 初始化
func NewUploadService(albSrv *AlbumService) IUploadService {
	return &UploadService{albSrv}
}

//UploadService 上传服务实现类
type UploadService struct {
	albSrv *AlbumService
}

//UploadImage 上传图片
func (upSrv UploadService) UploadImage(file *multipart.FileHeader, cid uint, aid uint) (res resp.CommonUploadFileResp, e error) {
	return upSrv.uploadFile(file, "image", 10, cid, aid)
}

//UploadVideo 上传视频
func (upSrv UploadService) UploadVideo(file *multipart.FileHeader, cid uint, aid uint) (res resp.CommonUploadFileResp, e error) {
	return upSrv.uploadFile(file, "video", 20, cid, aid)
}

//uploadFile 上传文件
func (upSrv UploadService) uploadFile(file *multipart.FileHeader, folder string, fileType int, cid uint, aid uint) (res resp.CommonUploadFileResp, e error) {
	var upRes *plugin.UploadFile
	if upRes, e = plugin.StorageDriver.Upload(file, folder, fileType); e != nil {
		return
	}
	var addReq req.CommonAlbumAddReq
	response.Copy(&addReq, upRes)
	addReq.Aid = aid
	addReq.Cid = cid
	var albumId uint
	if albumId, e = upSrv.albSrv.AlbumAdd(addReq); e != nil {
		return
	}
	response.Copy(&res, addReq)
	res.ID = albumId
	res.Path = upRes.Path
	return res, nil
}
