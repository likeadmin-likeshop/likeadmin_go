package system

import (
	"likeadmin/admin/schemas/req"
	"likeadmin/admin/schemas/resp"
	"likeadmin/core"
	"likeadmin/core/request"
	"likeadmin/core/response"
	"likeadmin/model/system"
	"likeadmin/util"
)

var SystemAuthPostService = systemAuthPostService{}

//systemAuthPostService 系统岗位服务实现类
type systemAuthPostService struct{}

//All 岗位所有
func (postSrv systemAuthPostService) All() []resp.SystemAuthPostResp {
	var posts []system.SystemAuthPost
	err := core.DB.Where("is_delete = ?", 0).Order("sort desc, id desc").Find(&posts).Error
	util.CheckUtil.CheckErr(err, "All Find err")
	res := []resp.SystemAuthPostResp{}
	response.Copy(&res, posts)
	return res
}

//List 岗位列表
func (postSrv systemAuthPostService) List(page request.PageReq, listReq req.SystemAuthPostListReq) response.PageResp {
	// 分页信息
	var res response.PageResp
	response.Copy(&res, page)
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)
	// 查询
	postModel := core.DB.Model(&system.SystemAuthPost{}).Where("is_delete = ?", 0)
	if listReq.Code != "" {
		postModel = postModel.Where("code like ?", "%"+listReq.Code+"%")
	}
	if listReq.Name != "" {
		postModel = postModel.Where("name like ?", "%"+listReq.Name+"%")
	}
	if listReq.IsStop >= 0 {
		postModel = postModel.Where("is_stop = ?", listReq.IsStop)
	}
	// 总数
	var count int64
	err := postModel.Count(&count).Error
	util.CheckUtil.CheckErr(err, "List Count err")
	// 数据
	var posts []system.SystemAuthPost
	err = postModel.Limit(limit).Offset(offset).Order("id desc").Find(&posts).Error
	util.CheckUtil.CheckErr(err, "List Find err")
	postResps := []resp.SystemAuthPostResp{}
	response.Copy(&postResps, posts)
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Lists:    postResps,
	}
}

//Detail 部门详情
func (postSrv systemAuthPostService) Detail(id uint) (res resp.SystemAuthPostResp) {
	var post system.SystemAuthPost
	err := core.DB.Where("id = ? AND is_delete = ?", id, 0).Limit(1).First(&post).Error
	util.CheckUtil.CheckErrDBNotRecord(err, "岗位不存在!")
	util.CheckUtil.CheckErr(err, "Detail First err")
	response.Copy(&res, post)
	return
}

//Add 部门新增
func (postSrv systemAuthPostService) Add(addReq req.SystemAuthPostAddReq) {
	r := core.DB.Where("(code = ? OR name = ?) AND is_delete = ?", addReq.Code, addReq.Name, 0).Limit(1).Find(&system.SystemAuthPost{})
	util.CheckUtil.CheckErr(r.Error, "Add Find err")
	if r.RowsAffected > 0 {
		panic(response.AssertArgumentError.Make("该岗位已存在!"))
	}
	var post system.SystemAuthPost
	response.Copy(&post, addReq)
	err := core.DB.Create(&post).Error
	util.CheckUtil.CheckErr(err, "Add Create err")
}

//Edit 部门编辑
func (postSrv systemAuthPostService) Edit(editReq req.SystemAuthPostEditReq) {
	var post system.SystemAuthPost
	err := core.DB.Where("id = ? AND is_delete = ?", editReq.ID, 0).Limit(1).First(&post).Error
	// 校验
	util.CheckUtil.CheckErrDBNotRecord(err, "部门不存在!")
	util.CheckUtil.CheckErr(err, "Edit First err")
	r := core.DB.Where("(code = ? OR name = ?) AND id != ? AND is_delete = ?", editReq.Code, editReq.Name, editReq.ID, 0).Limit(1).Find(&system.SystemAuthPost{})
	util.CheckUtil.CheckErr(r.Error, "Add Find err")
	if r.RowsAffected > 0 {
		panic(response.AssertArgumentError.Make("该岗位已存在!"))
	}
	// 更新
	response.Copy(&post, editReq)
	err = core.DB.Model(&post).Updates(post).Error
	util.CheckUtil.CheckErr(err, "Edit Updates err")
}

//Del 部门删除
func (postSrv systemAuthPostService) Del(id uint) {
	var post system.SystemAuthPost
	err := core.DB.Where("id = ? AND is_delete = ?", id, 0).Limit(1).First(&post).Error
	// 校验
	util.CheckUtil.CheckErrDBNotRecord(err, "部门不存在!")
	util.CheckUtil.CheckErr(err, "Del First err")
	r := core.DB.Where("post_id = ? AND is_delete = ?", id, 0).Limit(1).Find(&system.SystemAuthAdmin{})
	util.CheckUtil.CheckErr(r.Error, "Del Find err")
	if r.RowsAffected > 0 {
		panic(response.AssertArgumentError.Make("该岗位已被管理员使用,请先移除!"))
	}
	post.IsDelete = 1
	err = core.DB.Save(&post).Error
	util.CheckUtil.CheckErr(err, "Del Save err")
}
