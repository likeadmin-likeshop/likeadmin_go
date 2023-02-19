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
	group.AddGET("/album/cateList", cateList)
	group.AddPOST("/album/cateAdd", cateAdd, middleware.RecordLog("相册分类新增"))
	group.AddPOST("/album/cateRename", cateRename, middleware.RecordLog("相册分类重命名"))
	group.AddPOST("/album/cateDel", cateDel, middleware.RecordLog("相册分类删除"))
}

//indexConsole 控制台
func indexConsole(c *gin.Context) {
	res, err := common.NewIndexService(core.DB).Console()
	response.CheckAndRespWithData(c, res, err)
}

//indexConfig 公共配置
func indexConfig(c *gin.Context) {
	res, err := common.NewIndexService(core.DB).Config()
	response.CheckAndRespWithData(c, res, err)
}

//uploadImage 上传图片
func uploadImage(c *gin.Context) {
	var uReq req.CommonUploadImageReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &uReq)) {
		return
	}
	file, ve := util.VerifyUtil.VerifyFile(c, "file")
	if response.IsFailWithResp(c, ve) {
		return
	}
	srv := common.NewUploadService(common.NewAlbumService(core.DB))
	res, err := srv.UploadImage(file, uReq.Cid, config.AdminConfig.GetAdminId(c))
	response.CheckAndRespWithData(c, res, err)
}

//uploadVideo 上传视频
func uploadVideo(c *gin.Context) {
	var uReq req.CommonUploadImageReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &uReq)) {
		return
	}
	file, ve := util.VerifyUtil.VerifyFile(c, "file")
	if response.IsFailWithResp(c, ve) {
		return
	}
	srv := common.NewUploadService(common.NewAlbumService(core.DB))
	res, err := srv.UploadVideo(file, uReq.Cid, config.AdminConfig.GetAdminId(c))
	response.CheckAndRespWithData(c, res, err)
}

//albumList 相册文件列表
func albumList(c *gin.Context) {
	var page request.PageReq
	var listReq req.CommonAlbumListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := common.NewAlbumService(core.DB).AlbumList(page, listReq)
	response.CheckAndRespWithData(c, res, err)
}

//albumRename 相册文件重命名
func albumRename(c *gin.Context) {
	var rnReq req.CommonAlbumRenameReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &rnReq)) {
		return
	}
	response.CheckAndResp(c, common.NewAlbumService(core.DB).AlbumRename(rnReq.ID, rnReq.Name))
}

//albumMove 相册文件移动
func albumMove(c *gin.Context) {
	var mvReq req.CommonAlbumMoveReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &mvReq)) {
		return
	}
	response.CheckAndResp(c, common.NewAlbumService(core.DB).AlbumMove(mvReq.Ids, mvReq.Cid))
}

//albumDel 相册文件删除
func albumDel(c *gin.Context) {
	var delReq req.CommonAlbumDelReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &delReq)) {
		return
	}
	response.CheckAndResp(c, common.NewAlbumService(core.DB).AlbumDel(delReq.Ids))
}

//cateList 类目列表
func cateList(c *gin.Context) {
	var listReq req.CommonCateListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := common.NewAlbumService(core.DB).CateList(listReq)
	response.CheckAndRespWithData(c, res, err)
}

//cateAdd 类目新增
func cateAdd(c *gin.Context) {
	var addReq req.CommonCateAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &addReq)) {
		return
	}
	response.CheckAndResp(c, common.NewAlbumService(core.DB).CateAdd(addReq))
}

//cateRename 类目命名
func cateRename(c *gin.Context) {
	var rnReq req.CommonCateRenameReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &rnReq)) {
		return
	}
	response.CheckAndResp(c, common.NewAlbumService(core.DB).CateRename(rnReq.ID, rnReq.Name))
}

//cateDel 类目删除
func cateDel(c *gin.Context) {
	var delReq req.CommonCateDelReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &delReq)) {
		return
	}
	response.CheckAndResp(c, common.NewAlbumService(core.DB).CateDel(delReq.ID))
}
