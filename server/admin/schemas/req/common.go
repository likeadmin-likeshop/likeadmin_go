package req

//CommonUploadImageReq 上传图片参数
type CommonUploadImageReq struct {
	Cid uint `form:"cid" binding:"gte=0"` // 主键
}

//CommonAlbumListReq 相册文件列表参数
type CommonAlbumListReq struct {
	Cid  int    `form:"cid,default=-1"`                       // 类目ID
	Type int    `form:"type" binding:"omitempty,oneof=10 20"` // 文件类型: [10=图片, 20=视频]
	Name string `form:"keyword"`                              // 文件名称
}

//CommonAlbumRenameReq 相册文件重命名参数
type CommonAlbumRenameReq struct {
	ID   uint   `form:"id" binding:"required,gt=0"`              // 主键
	Name string `form:"keyword" binding:"required,min=1,max=30"` // 文件名称
}

//CommonAlbumMoveReq 相册文件移动参数
type CommonAlbumMoveReq struct {
	Ids []uint `form:"ids" binding:"required"` // 主键
	Cid int    `form:"cid,default=-1"`         // 类目ID
}

//CommonAlbumDelReq 相册文件删除参数
type CommonAlbumDelReq struct {
	Ids []uint `form:"ids" binding:"required"` // 主键
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
