package system

import (
	"likeadmin/admin/schemas/req"
	"likeadmin/admin/schemas/resp"
	"likeadmin/core"
	"likeadmin/core/response"
	"likeadmin/model/system"
	"likeadmin/util"
)

var SystemAuthDeptService = systemAuthDeptService{}

//systemAuthDeptService 系统部门服务实现类
type systemAuthDeptService struct{}

//All 部门所有
func (deptSrv systemAuthDeptService) All() []resp.SystemAuthDeptResp {
	var depts []system.SystemAuthDept
	err := core.DB.Where("pid > ? AND is_delete = ?", 0, 0).Order("sort desc, id desc").Find(&depts).Error
	util.CheckUtil.CheckErr(err, "All Find err")
	res := []resp.SystemAuthDeptResp{}
	response.Copy(&res, depts)
	return res
}

//List 部门列表
func (deptSrv systemAuthDeptService) List(listReq req.SystemAuthDeptListReq) (mapList []interface{}) {
	deptModel := core.DB.Where("is_delete = ?", 0)
	if listReq.Name != "" {
		deptModel = deptModel.Where("name like ?", "%"+listReq.Name+"%")
	}
	if listReq.IsStop >= 0 {
		deptModel = deptModel.Where("is_stop = ?", listReq.IsStop)
	}
	var depts []system.SystemAuthDept
	err := deptModel.Order("sort desc, id desc").Find(&depts).Error
	util.CheckUtil.CheckErr(err, "List Find err")
	deptResps := []resp.SystemAuthDeptResp{}
	response.Copy(&deptResps, depts)
	mapList = util.ArrayUtil.ListToTree(
		util.ConvertUtil.StructsToMaps(deptResps), "id", "pid", "children")
	return
}

//Detail 部门详情
func (deptSrv systemAuthDeptService) Detail(id uint) (res resp.SystemAuthDeptResp) {
	var dept system.SystemAuthDept
	err := core.DB.Where("id = ? AND is_delete = ?", id, 0).Limit(1).First(&dept).Error
	util.CheckUtil.CheckErrDBNotRecord(err, "部门已不存在!")
	util.CheckUtil.CheckErr(err, "Detail First err")
	response.Copy(&res, dept)
	return
}

//Add 部门新增
func (deptSrv systemAuthDeptService) Add(addReq req.SystemAuthDeptAddReq) {
	if addReq.Pid == 0 {
		r := core.DB.Where("pid = ? AND is_delete = ?", 0, 0).Limit(1).Find(&system.SystemAuthDept{})
		util.CheckUtil.CheckErr(r.Error, "Add Find err")
		if r.RowsAffected > 0 {
			panic(response.AssertArgumentError.Make("顶级部门只允许有一个!"))
		}
	}
	var dept system.SystemAuthDept
	response.Copy(&dept, addReq)
	err := core.DB.Create(&dept).Error
	util.CheckUtil.CheckErr(err, "Add Create err")
}

//Edit 部门编辑
func (deptSrv systemAuthDeptService) Edit(editReq req.SystemAuthDeptEditReq) {
	var dept system.SystemAuthDept
	err := core.DB.Where("id = ? AND is_delete = ?", editReq.ID, 0).Limit(1).First(&dept).Error
	// 校验
	util.CheckUtil.CheckErrDBNotRecord(err, "部门不存在!")
	util.CheckUtil.CheckErr(err, "Edit First err")
	if dept.Pid == 0 && editReq.Pid > 0 {
		panic(response.AssertArgumentError.Make("顶级部门不能修改上级!"))
	}
	if editReq.ID == editReq.Pid {
		panic(response.AssertArgumentError.Make("上级部门不能是自己!"))
	}
	// 更新
	response.Copy(&dept, editReq)
	err = core.DB.Model(&dept).Updates(dept).Error
	util.CheckUtil.CheckErr(err, "Edit Updates err")
}

//Del 部门删除
func (deptSrv systemAuthDeptService) Del(id uint) {
	var dept system.SystemAuthDept
	err := core.DB.Where("id = ? AND is_delete = ?", id, 0).Limit(1).First(&dept).Error
	// 校验
	util.CheckUtil.CheckErrDBNotRecord(err, "部门不存在!")
	util.CheckUtil.CheckErr(err, "Del First err")
	if dept.Pid == 0 {
		panic(response.AssertArgumentError.Make("顶级部门不能删除!"))
	}
	r := core.DB.Where("pid = ? AND is_delete = ?", id, 0).Limit(1).Find(&system.SystemAuthDept{})
	util.CheckUtil.CheckErr(r.Error, "Del Find dept err")
	if r.RowsAffected > 0 {
		panic(response.AssertArgumentError.Make("请先删除子级部门!"))
	}
	r = core.DB.Where("dept_id = ? AND is_delete = ?", id, 0).Limit(1).Find(&system.SystemAuthAdmin{})
	util.CheckUtil.CheckErr(r.Error, "Del Find admin err")
	if r.RowsAffected > 0 {
		panic(response.AssertArgumentError.Make("该部门已被管理员使用,请先移除!"))
	}
	dept.IsDelete = 1
	err = core.DB.Save(&dept).Error
	util.CheckUtil.CheckErr(err, "Del Save err")
}
