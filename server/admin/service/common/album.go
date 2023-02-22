package common

import (
	"gorm.io/gorm"
	"likeadmin/admin/schemas/req"
	"likeadmin/admin/schemas/resp"
	"likeadmin/config"
	"likeadmin/core/request"
	"likeadmin/core/response"
	"likeadmin/model/common"
	"likeadmin/util"
	"path"
	"time"
)

type IAlbumService interface {
	AlbumList(page request.PageReq, listReq req.CommonAlbumListReq) (res response.PageResp, e error)
	AlbumRename(id uint, name string) (e error)
	AlbumMove(ids []uint, cid int) (e error)
	AlbumAdd(addReq req.CommonAlbumAddReq) (res uint, e error)
	AlbumDel(ids []uint) (e error)
	CateList(listReq req.CommonCateListReq) (mapList []interface{}, e error)
	CateAdd(addReq req.CommonCateAddReq) (e error)
	CateRename(id uint, name string) (e error)
	CateDel(id uint) (e error)
}

//NewAlbumService 初始化
func NewAlbumService(db *gorm.DB) IAlbumService {
	return &AlbumService{db: db}
}

//AlbumService 相册服务实现类
type AlbumService struct {
	db *gorm.DB
}

//AlbumList 相册文件列表
func (albSrv AlbumService) AlbumList(page request.PageReq, listReq req.CommonAlbumListReq) (res response.PageResp, e error) {
	// 分页信息
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)
	// 查询
	albumModel := albSrv.db.Model(&common.Album{}).Where("is_delete = ?", 0)
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
	if e = response.CheckErr(err, "AlbumList Count err"); e != nil {
		return
	}
	// 数据
	var albums []common.Album
	err = albumModel.Limit(limit).Offset(offset).Order("id desc").Find(&albums).Error
	if e = response.CheckErr(err, "AlbumList Find err"); e != nil {
		return
	}
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
	}, nil
}

//AlbumRename 相册文件重命名
func (albSrv AlbumService) AlbumRename(id uint, name string) (e error) {
	var album common.Album
	err := albSrv.db.Where("id = ? AND is_delete = ?", id, 0).Limit(1).First(&album).Error
	if e = response.CheckErrDBNotRecord(err, "文件丢失！"); e != nil {
		return
	}
	if e = response.CheckErr(err, "AlbumRename First err"); e != nil {
		return
	}
	album.Name = name
	err = albSrv.db.Save(&album).Error
	e = response.CheckErr(err, "AlbumRename Save err")
	return
}

//AlbumMove 相册文件移动
func (albSrv AlbumService) AlbumMove(ids []uint, cid int) (e error) {
	var albums []common.Album
	err := albSrv.db.Where("id in ? AND is_delete = ?", ids, 0).Find(&albums).Error
	if e = response.CheckErr(err, "AlbumMove Find err"); e != nil {
		return
	}
	if len(albums) == 0 {
		return response.AssertArgumentError.Make("文件丢失！")
	}
	if cid > 0 {
		err = albSrv.db.Where("id = ? AND is_delete = ?", cid, 0).Limit(1).First(&common.AlbumCate{}).Error
		if e = response.CheckErrDBNotRecord(err, "类目已不存在！"); e != nil {
			return
		}
		if e = response.CheckErr(err, "AlbumMove First err"); e != nil {
			return
		}
	}
	err = albSrv.db.Model(&common.Album{}).Where("id in ?", ids).UpdateColumn("cid", cid).Error
	e = response.CheckErr(err, "AlbumMove UpdateColumn err")
	return
}

//AlbumAdd 相册文件新增
func (albSrv AlbumService) AlbumAdd(addReq req.CommonAlbumAddReq) (res uint, e error) {
	var alb common.Album
	//var params map[string]interface{}
	//if err := mapstructure.Decode(params, &alb); err != nil {
	//	core.Logger.Errorf("AlbumAdd Decode err: err=[%+v]", err)
	//	return response.SystemError
	//}
	response.Copy(&alb, addReq)
	err := albSrv.db.Create(&alb).Error
	if e = response.CheckErr(err, "AlbumAdd Create err"); e != nil {
		return
	}
	return alb.ID, nil
}

//AlbumDel 相册文件删除
func (albSrv AlbumService) AlbumDel(ids []uint) (e error) {
	var albums []common.Album
	err := albSrv.db.Where("id in ? AND is_delete = ?", ids, 0).Find(&albums).Error
	if e = response.CheckErr(err, "AlbumDel Find err"); e != nil {
		return
	}
	if len(albums) == 0 {
		return response.AssertArgumentError.Make("文件丢失！")
	}
	err = albSrv.db.Model(&common.Album{}).Where("id in ?", ids).Updates(
		common.Album{IsDelete: 1, DeleteTime: time.Now().Unix()}).Error
	e = response.CheckErr(err, "AlbumDel UpdateColumn err")
	return
}

//CateList 相册分类列表
func (albSrv AlbumService) CateList(listReq req.CommonCateListReq) (mapList []interface{}, e error) {
	var cates []common.AlbumCate
	cateModel := albSrv.db.Where("is_delete = ?", 0).Order("id desc")
	if listReq.Type > 0 {
		cateModel = cateModel.Where("type = ?", listReq.Type)
	}
	if listReq.Name != "" {
		cateModel = cateModel.Where("name like ?", "%"+listReq.Name+"%")
	}
	err := cateModel.Find(&cates).Error
	if e = response.CheckErr(err, "CateList Find err"); e != nil {
		return
	}
	cateResps := []resp.CommonCateListResp{}
	response.Copy(&cateResps, cates)
	return util.ArrayUtil.ListToTree(
		util.ConvertUtil.StructsToMaps(cateResps), "id", "pid", "children"), nil
}

//CateAdd 分类新增
func (albSrv AlbumService) CateAdd(addReq req.CommonCateAddReq) (e error) {
	var cate common.AlbumCate
	response.Copy(&cate, addReq)
	err := albSrv.db.Create(&cate).Error
	e = response.CheckErr(err, "CateAdd Create err")
	return
}

//CateRename 分类重命名
func (albSrv AlbumService) CateRename(id uint, name string) (e error) {
	var cate common.AlbumCate
	err := albSrv.db.Where("id = ? AND is_delete = ?", id, 0).Limit(1).First(&cate).Error
	if e = response.CheckErrDBNotRecord(err, "分类已不存在！"); e != nil {
		return
	}
	if e = response.CheckErr(err, "CateRename First err"); e != nil {
		return
	}
	cate.Name = name
	err = albSrv.db.Save(&cate).Error
	e = response.CheckErr(err, "CateRename Save err")
	return
}

//CateDel 分类删除
func (albSrv AlbumService) CateDel(id uint) (e error) {
	var cate common.AlbumCate
	err := albSrv.db.Where("id = ? AND is_delete = ?", id, 0).Limit(1).First(&cate).Error
	if e = response.CheckErrDBNotRecord(err, "分类已不存在！"); e != nil {
		return
	}
	if e = response.CheckErr(err, "CateDel First err"); e != nil {
		return
	}
	r := albSrv.db.Where("cid = ? AND is_delete = ?", id, 0).Limit(1).Find(&common.Album{})
	if e = response.CheckErr(r.Error, "CateDel Find err"); e != nil {
		return
	}
	if r.RowsAffected > 0 {
		return response.AssertArgumentError.Make("当前分类正被使用中,不能删除！")
	}
	cate.IsDelete = 1
	cate.DeleteTime = time.Now().Unix()
	err = albSrv.db.Save(&cate).Error
	e = response.CheckErr(err, "CateDel Save err")
	return
}
