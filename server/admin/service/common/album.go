package common

import (
	"likeadmin/admin/schemas/req"
	"likeadmin/admin/schemas/resp"
	"likeadmin/config"
	"likeadmin/core"
	"likeadmin/core/request"
	"likeadmin/core/response"
	"likeadmin/model/common"
	"likeadmin/util"
	"path"
	"time"
)

var AlbumService = albumService{}

//albumService 相册服务实现类
type albumService struct{}

//AlbumList 相册文件列表
func (albSrv albumService) AlbumList(page request.PageReq, listReq req.CommonAlbumListReq) response.PageResp {
	// 分页信息
	var res response.PageResp
	response.Copy(&res, page)
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)
	// 查询
	albumModel := core.DB.Model(&common.Album{}).Where("is_delete = ?", 0)
	if listReq.Cid > 0 {
		albumModel = albumModel.Where("cid = ?", listReq.Cid)
	}
	if listReq.Name != "" {
		albumModel = albumModel.Where("name like ?", "%"+listReq.Name+"%")
	}
	if listReq.Type > 0 {
		albumModel = albumModel.Where("type = ?", listReq.Type)
	}
	// 总数
	var count int64
	err := albumModel.Count(&count).Error
	util.CheckUtil.CheckErr(err, "AlbumList Count err")
	// 数据
	var albums []common.Album
	err = albumModel.Limit(limit).Offset(offset).Order("id desc").Find(&albums).Error
	util.CheckUtil.CheckErr(err, "AlbumList Find err")
	albumResps := []resp.CommonAlbumListResp{}
	response.Copy(&albumResps, albums)
	// TODO: engine默认local
	engine := "local"
	for i := 0; i < len(albumResps); i++ {
		if engine == "local" {
			albumResps[i].Path = path.Join(config.Config.PublicPrefix, albums[i].Uri)
		} else {
			// TODO: 其他engine
		}
		albumResps[i].Uri = util.UrlUtil.ToAbsoluteUrl(albums[i].Uri)
		albumResps[i].Size = util.ServerUtil.GetFmtSize(uint64(albums[i].Size))
	}
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Lists:    albumResps,
	}
}

//AlbumRename 相册文件重命名
func (albSrv albumService) AlbumRename(id uint, name string) {
	var album common.Album
	err := core.DB.Where("id = ? AND is_delete = ?", id, 0).Limit(1).First(&album).Error
	util.CheckUtil.CheckErrDBNotRecord(err, "文件丢失！")
	util.CheckUtil.CheckErr(err, "AlbumRename First err")
	album.Name = name
	err = core.DB.Save(&album).Error
	util.CheckUtil.CheckErr(err, "AlbumRename Save err")
}

//AlbumMove 相册文件移动
func (albSrv albumService) AlbumMove(ids []uint, cid int) {
	var albums []common.Album
	err := core.DB.Where("id in ? AND is_delete = ?", ids, 0).Find(&albums).Error
	util.CheckUtil.CheckErr(err, "AlbumMove Find err")
	if len(albums) == 0 {
		panic(response.AssertArgumentError.Make("文件丢失！"))
	}
	if cid > 0 {
		err = core.DB.Where("id = ? AND is_delete = ?", cid, 0).Limit(1).First(&common.AlbumCate{}).Error
		util.CheckUtil.CheckErrDBNotRecord(err, "类目已不存在！")
		util.CheckUtil.CheckErr(err, "AlbumMove First err")
	}
	err = core.DB.Model(&common.Album{}).Where("id in ?", ids).UpdateColumn("cid", cid).Error
	util.CheckUtil.CheckErr(err, "AlbumMove UpdateColumn err")
}

//AlbumAdd 相册文件新增
func (albSrv albumService) AlbumAdd(addReq req.CommonAlbumAddReq) uint {
	var alb common.Album
	//var params map[string]interface{}
	//if err := mapstructure.Decode(params, &alb); err != nil {
	//	core.Logger.Errorf("AlbumAdd Decode err: err=[%+v]", err)
	//	panic(response.SystemError)
	//}
	response.Copy(&alb, addReq)
	err := core.DB.Create(&alb).Error
	util.CheckUtil.CheckErr(err, "AlbumAdd Create err")
	return alb.ID
}

//AlbumDel 相册文件删除
func (albSrv albumService) AlbumDel(ids []uint) {
	var albums []common.Album
	err := core.DB.Where("id in ? AND is_delete = ?", ids, 0).Find(&albums).Error
	util.CheckUtil.CheckErr(err, "AlbumDel Find err")
	if len(albums) == 0 {
		panic(response.AssertArgumentError.Make("文件丢失！"))
	}
	err = core.DB.Model(&common.Album{}).Where("id in ?", ids).Updates(
		common.Album{IsDelete: 1, DeleteTime: time.Now().Unix()}).Error
	util.CheckUtil.CheckErr(err, "AlbumDel UpdateColumn err")
}

//CateList 相册分类列表
func (albSrv albumService) CateList(listReq req.CommonCateListReq) (mapList []interface{}) {
	var cates []common.AlbumCate
	cateModel := core.DB.Where("is_delete = ?", 0).Order("id desc")
	if listReq.Type > 0 {
		cateModel = cateModel.Where("type = ?", listReq.Type)
	}
	if listReq.Name != "" {
		cateModel = cateModel.Where("name like ?", "%"+listReq.Name+"%")
	}
	err := cateModel.Find(&cates).Error
	util.CheckUtil.CheckErr(err, "CateList Find err")
	cateResps := []resp.CommonCateListResp{}
	response.Copy(&cateResps, cates)
	return util.ArrayUtil.ListToTree(
		util.ConvertUtil.StructsToMaps(cateResps), "id", "pid", "children")
}

//CateAdd 分类新增
func (albSrv albumService) CateAdd(addReq req.CommonCateAddReq) {
	var cate common.AlbumCate
	response.Copy(&cate, addReq)
	err := core.DB.Create(&cate).Error
	util.CheckUtil.CheckErr(err, "CateAdd Create err")
}

//CateRename 分类重命名
func (albSrv albumService) CateRename(id uint, name string) {
	var cate common.AlbumCate
	err := core.DB.Where("id = ? AND is_delete = ?", id, 0).Limit(1).First(&cate).Error
	util.CheckUtil.CheckErrDBNotRecord(err, "分类已不存在！")
	util.CheckUtil.CheckErr(err, "CateRename First err")
	cate.Name = name
	err = core.DB.Save(&cate).Error
	util.CheckUtil.CheckErr(err, "CateRename Save err")
}

//CateDel 分类删除
func (albSrv albumService) CateDel(id uint) {
	var cate common.AlbumCate
	err := core.DB.Where("id = ? AND is_delete = ?", id, 0).Limit(1).First(&cate).Error
	util.CheckUtil.CheckErrDBNotRecord(err, "分类已不存在！")
	util.CheckUtil.CheckErr(err, "CateDel First err")
	r := core.DB.Where("cid = ? AND is_delete = ?", id, 0).Limit(1).Find(&common.Album{})
	util.CheckUtil.CheckErr(r.Error, "CateDel Find err")
	if r.RowsAffected > 0 {
		panic(response.AssertArgumentError.Make("当前分类正被使用中,不能删除！"))
	}
	cate.IsDelete = 1
	cate.DeleteTime = time.Now().Unix()
	err = core.DB.Save(&cate).Error
	util.CheckUtil.CheckErr(err, "CateDel Save err")
}
