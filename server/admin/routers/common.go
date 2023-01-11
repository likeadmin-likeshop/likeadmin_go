package routers

import (
	"github.com/gin-gonic/gin"
	"likeadmin/admin/schemas/req"
	"likeadmin/admin/service/common"
	"likeadmin/config"
	"likeadmin/core"
	"likeadmin/core/request"
	"likeadmin/core/response"
	"likeadmin/middleware"
	"likeadmin/util"
)

var CommonGroup = core.Group("/common")

func init() {
	group := CommonGroup
	group.AddGET("/index/console", indexConsole)
	group.AddGET("/index/config", indexConfig)
	group.AddPOST("/upload/image", uploadImage, middleware.RecordLog("上传图片", middleware.RequestFile))
	group.AddPOST("/upload/video", uploadVideo, middleware.RecordLog("上传视频", middleware.RequestFile))
	group.AddGET("/album/albumList", albumList)
	group.AddPOST("/album/albumRename", albumRename, middleware.RecordLog("相册文件重命名"))
	group.AddPOST("/album/albumMove", albumMove, middleware.RecordLog("相册文件移动"))
	group.AddPOST("/album/albumDel", albumDel, middleware.RecordLog("相册文件删除"))
}

//indexConsole 控制台
func indexConsole(c *gin.Context) {
	response.OkWithData(c, common.IndexService.Console())
}

//indexConfig 公共配置
func indexConfig(c *gin.Context) {
	response.OkWithData(c, common.IndexService.Config())
}

//uploadImage 上传图片
func uploadImage(c *gin.Context) {
	var uReq req.CommonUploadImageReq
	util.VerifyUtil.VerifyBody(c, &uReq)
	file := util.VerifyUtil.VerifyFile(c, "file")
	response.OkWithData(c, common.UploadService.UploadImage(file, uReq.Cid, config.AdminConfig.GetAdminId(c)))
}

//uploadVideo 上传视频
func uploadVideo(c *gin.Context) {
	var uReq req.CommonUploadImageReq
	util.VerifyUtil.VerifyBody(c, &uReq)
	file := util.VerifyUtil.VerifyFile(c, "file")
	response.OkWithData(c, common.UploadService.UploadVideo(file, uReq.Cid, config.AdminConfig.GetAdminId(c)))
}

//albumList 相册文件列表
func albumList(c *gin.Context) {
	var page request.PageReq
	var listReq req.CommonAlbumListReq
	util.VerifyUtil.VerifyQuery(c, &page)
	util.VerifyUtil.VerifyQuery(c, &listReq)
	response.OkWithData(c, common.AlbumService.AlbumList(page, listReq))
}

//albumRename 相册文件重命名
func albumRename(c *gin.Context) {
	var rnReq req.CommonAlbumRenameReq
	util.VerifyUtil.VerifyJSON(c, &rnReq)
	common.AlbumService.AlbumRename(rnReq.ID, rnReq.Name)
	response.Ok(c)
}

//albumMove 相册文件移动
func albumMove(c *gin.Context) {
	var mvReq req.CommonAlbumMoveReq
	util.VerifyUtil.VerifyJSON(c, &mvReq)
	common.AlbumService.AlbumMove(mvReq.Ids, mvReq.Cid)
	response.Ok(c)
}

//albumDel 相册文件删除
func albumDel(c *gin.Context) {
	var delReq req.CommonAlbumDelReq
	util.VerifyUtil.VerifyJSON(c, &delReq)
	common.AlbumService.AlbumDel(delReq.Ids)
	response.Ok(c)
}
