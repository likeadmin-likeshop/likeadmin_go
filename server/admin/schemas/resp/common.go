package resp

import "likeadmin/core"

//CommonUploadFileResp 上传图片返回信息
type CommonUploadFileResp struct {
	ID   uint   `json:"id" structs:"id"`     // 主键
	Cid  uint   `json:"cid" structs:"cid"`   // 类目ID
	Aid  uint   `json:"aid" structs:"aid"`   // 管理ID
	Uid  uint   `json:"uid" structs:"uid"`   // 用户ID
	Type int    `json:"type" structs:"type"` // 文件类型: [10=图片, 20=视频]
	Name string `json:"name" structs:"name"` // 文件名称
	Uri  string `json:"url" structs:"url"`   // 文件路径
	Path string `json:"path" structs:"path"` // 访问地址
	Ext  string `json:"ext" structs:"ext"`   // 文件扩展
	Size int64  `json:"size" structs:"size"` // 文件大小
}

//CommonAlbumListResp 相册文件列表返回信息
type CommonAlbumListResp struct {
	ID         uint        `json:"id" structs:"id"`                 // 主键
	Cid        uint        `json:"cid" structs:"cid"`               // 所属类目
	Name       string      `json:"name" structs:"name"`             // 文件名称
	Path       string      `json:"path" structs:"path"`             // 相对路径
	Uri        string      `json:"uri" structs:"uri"`               // 文件路径
	Ext        string      `json:"ext" structs:"ext"`               // 文件扩展
	Size       string      `json:"size" structs:"size"`             // 文件大小
	CreateTime core.TsTime `json:"createTime" structs:"createTime"` // 创建时间
	UpdateTime core.TsTime `json:"updateTime" structs:"updateTime"` // 更新时间
}
