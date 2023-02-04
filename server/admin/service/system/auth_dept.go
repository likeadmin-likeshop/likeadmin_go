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
func (deptSrv systemAuthDeptService) All() (res []resp.SystemAuthDeptResp, e error) {
	var depts []system.SystemAuthDept
	err := core.DB.Where("pid > ? AND is_delete = ?", 0, 0).Order("sort desc, id desc").Find(&depts).Error
	if e = response.CheckErr(err, "All Find err"); e != nil {
		return
	}
	res = []resp.SystemAuthDeptResp{}
	response.Copy(&res, depts)
	return
}

//List 部门列表
func (deptSrv systemAuthDeptService) List(listReq req.SystemAuthDeptListReq) (mapList []interface{}, e error) {
	deptModel := core.DB.Where("is_delete = ?", 0)
	if listReq.Name != "" {
		deptModel = deptModel.Where("name like ?", "%"+listReq.Name+"%")
	}
	if listReq.IsStop >= 0 {
		deptModel = deptModel.Where("is_stop = ?", listReq.IsStop)
	}
	var depts []system.SystemAuthDept
	err := deptModel.Order("sort desc, id desc").Find(&depts).Error
	if e = response.CheckErr(err, "List Find err"); e != nil {
		return
	}
	deptResps := []resp.SystemAuthDeptResp{}
	response.Copy(&deptResps, depts)
	mapList = util.ArrayUtil.ListToTree(
		util.ConvertUtil.StructsToMaps(deptResps), "id", "pid", "children")
	return
}

//Detail 部门详情
func (deptSrv systemAuthDeptService) Detail(id uint) (res resp.SystemAuthDeptResp, e error) {
	var dept system.SystemAuthDept
	err := core.DB.Where("id = ? AND is_delete = ?", id, 0).Limit(1).First(&dept).Error
	if e = response.CheckErrDBNotRecord(err, "部门已不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "Detail First err"); e != nil {
		return
	}
	response.Copy(&res, dept)
	return
}

//Add 部门新增
func (deptSrv systemAuthDeptService) Add(addReq req.SystemAuthDeptAddReq) (e error) {
	if addReq.Pid == 0 {
		r := core.DB.Where("pid = ? AND is_delete = ?", 0, 0).Limit(1).Find(&system.SystemAuthDept{})
		if e = response.CheckErr(r.Error, "Add Find err"); e != nil {
			return
		}
		if r.RowsAffected > 0 {
			return response.AssertArgumentError.Make("顶级部门只允许有一个!")
		}
	}
	var dept system.SystemAuthDept
	response.Copy(&dept, addReq)
	err := core.DB.Create(&dept).Error
	e = response.CheckErr(err, "Add Create err")
	return
}

//Edit 部门编辑
func (deptSrv systemAuthDeptService) Edit(editReq req.SystemAuthDeptEditReq) (e error) {
	var dept system.SystemAuthDept
	err := core.DB.Where("id = ? AND is_delete = ?", editReq.ID, 0).Limit(1).First(&dept).Error
	// 校验
	if e = response.CheckErrDBNotRecord(err, "部门不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "Edit First err"); e != nil {
		return
	}
	if dept.Pid == 0 && editReq.Pid > 0 {
		return response.AssertArgumentError.Make("顶级部门不能修改上级!")
	}
	if editReq.ID == editReq.Pid {
		return response.AssertArgumentError.Make("上级部门不能是自己!")
	}
	// 更新
	response.Copy(&dept, editReq)
	err = core.DB.Model(&dept).Updates(dept).Error
	e = response.CheckErr(err, "Edit Updates err")
	return
}

//Del 部门删除
func (deptSrv systemAuthDeptService) Del(id uint) (e error) {
	var dept system.SystemAuthDept
	err := core.DB.Where("id = ? AND is_delete = ?", id, 0).Limit(1).First(&dept).Error
	// 校验
	if e = response.CheckErrDBNotRecord(err, "部门不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "Del First err"); e != nil {
		return
	}
	if dept.Pid == 0 {
		return response.AssertArgumentError.Make("顶级部门不能删除!")
	}
	r := core.DB.Where("pid = ? AND is_delete = ?", id, 0).Limit(1).Find(&system.SystemAuthDept{})
	if e = response.CheckErr(r.Error, "Del Find dept err"); e != nil {
		return
	}
	if r.RowsAffected > 0 {
		return response.AssertArgumentError.Make("请先删除子级部门!")
	}
	r = core.DB.Where("dept_id = ? AND is_delete = ?", id, 0).Limit(1).Find(&system.SystemAuthAdmin{})
	if e = response.CheckErr(r.Error, "Del Find admin err"); e != nil {
		return
	}
	if r.RowsAffected > 0 {
		return response.AssertArgumentError.Make("该部门已被管理员使用,请先移除!")
	}
	dept.IsDelete = 1
	err = core.DB.Save(&dept).Error
	e = response.CheckErr(err, "Del Save err")
	return
}
