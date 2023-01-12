package routers

import (
	"github.com/gin-gonic/gin"
	"likeadmin/admin/schemas/req"
	"likeadmin/admin/service/common"
	"likeadmin/config"
	"likeadmin/core"
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
