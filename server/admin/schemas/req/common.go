package req

//CommonUploadImageReq 上传图片参数
type CommonUploadImageReq struct {
	Cid uint `form:"cid" binding:"gte=0"` // 主键
}

//CommonAlbumAddReq 相册文件新增参数
type CommonAlbumAddReq struct {
	Cid  uint   `form:"cid" binding:"gte=0"`        // 类目ID
	Aid  uint   `form:"aid" binding:"gte=0"`        // 管理ID
	Uid  uint   `form:"uid" binding:"gte=0"`        // 用户ID
	Type int    `form:"type" binding:"oneof=10 20"` // 文件类型: [10=图片, 20=视频]
	Name string `form:"name"`                       // 文件名称
	Uri  string `form:"uri"`                        // 文件路径
	Ext  string `form:"ext"`                        // 文件扩展
	Size int64  `form:"size"`                       // 文件大小
}
