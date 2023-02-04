package system

import (
	"github.com/fatih/structs"
	"gorm.io/gorm"
	"likeadmin/admin/schemas/req"
	"likeadmin/admin/schemas/resp"
	"likeadmin/config"
	"likeadmin/core"
	"likeadmin/core/request"
	"likeadmin/core/response"
	"likeadmin/model/system"
	"likeadmin/util"
	"strconv"
	"strings"
)

var SystemAuthRoleService = systemAuthRoleService{}

//systemAuthRoleService 系统角色服务实现类
type systemAuthRoleService struct{}

//All 角色所有
func (roleSrv systemAuthRoleService) All() (res []resp.SystemAuthRoleSimpleResp, e error) {
	var roles []system.SystemAuthRole
	err := core.DB.Order("sort desc, id desc").Find(&roles).Error
	if e = response.CheckErr(err, "All Find err"); e != nil {
		return
	}
	response.Copy(&res, roles)
	return
}

//List 根据角色ID获取菜单ID
func (roleSrv systemAuthRoleService) List(page request.PageReq) (res response.PageResp, e error) {
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)
	roleModel := core.DB.Model(&system.SystemAuthRole{})
	var count int64
	err := roleModel.Count(&count).Error
	if e = response.CheckErr(err, "List Count err"); e != nil {
		return
	}
	var roles []system.SystemAuthRole
	err = roleModel.Limit(limit).Offset(offset).Order("sort desc, id desc").Find(&roles).Error
	if e = response.CheckErr(err, "List Find err"); e != nil {
		return
	}
	var roleResp []resp.SystemAuthRoleResp
	response.Copy(&roleResp, roles)
	for i := 0; i < len(roleResp); i++ {
		roleResp[i].Menus = []uint{}
		roleResp[i].Member = roleSrv.getMemberCnt(roleResp[i].ID)
	}
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Lists:    roleResp,
	}, nil
}

//Detail 角色详情
func (roleSrv systemAuthRoleService) Detail(id uint) (res resp.SystemAuthRoleResp, e error) {
	var role system.SystemAuthRole
	err := core.DB.Where("id = ?", id).Limit(1).First(&role).Error
	if e = response.CheckErrDBNotRecord(err, "角色已不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "Detail First err"); e != nil {
		return
	}
	response.Copy(&res, role)
	res.Member = roleSrv.getMemberCnt(role.ID)
	res.Menus, e = SystemAuthPermService.SelectMenuIdsByRoleId(role.ID)
	return
}

//getMemberCnt 根据角色ID获取成员数量
func (roleSrv systemAuthRoleService) getMemberCnt(roleId uint) (count int64) {
	core.DB.Model(&system.SystemAuthAdmin{}).Where(
		"role = ? AND is_delete = ?", roleId, 0).Count(&count)
	return
}

//Add 新增角色
func (roleSrv systemAuthRoleService) Add(addReq req.SystemAuthRoleAddReq) (e error) {
	var role system.SystemAuthRole
	if r := core.DB.Where("name = ?", strings.Trim(addReq.Name, " ")).Limit(1).First(&role); r.RowsAffected > 0 {
		return response.AssertArgumentError.Make("角色名称已存在!")
	}
	response.Copy(&role, addReq)
	role.Name = strings.Trim(addReq.Name, " ")
	// 事务
	err := core.DB.Transaction(func(tx *gorm.DB) error {
		txErr := tx.Create(&role).Error
		var te error
		if te = response.CheckErr(txErr, "Add Create in tx err"); te != nil {
			return te
		}
		te = SystemAuthPermService.BatchSaveByMenuIds(role.ID, addReq.MenuIds, tx)
		return te
	})
	e = response.CheckErr(err, "Add Transaction err")
	return
}

//Edit 编辑角色
func (roleSrv systemAuthRoleService) Edit(editReq req.SystemAuthRoleEditReq) (e error) {
	err := core.DB.Where("id = ?", editReq.ID).Limit(1).First(&system.SystemAuthRole{}).Error
	if e = response.CheckErrDBNotRecord(err, "角色已不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "Edit First err"); e != nil {
		return
	}
	var role system.SystemAuthRole
	if r := core.DB.Where("id != ? AND name = ?", editReq.ID, strings.Trim(editReq.Name, " ")).Limit(1).First(&role); r.RowsAffected > 0 {
		return response.AssertArgumentError.Make("角色名称已存在!")
	}
	role.ID = editReq.ID
	roleMap := structs.Map(editReq)
	delete(roleMap, "ID")
	delete(roleMap, "MenuIds")
	roleMap["Name"] = strings.Trim(editReq.Name, " ")
	// 事务
	err = core.DB.Transaction(func(tx *gorm.DB) error {
		txErr := tx.Model(&role).Updates(roleMap).Error
		var te error
		if te = response.CheckErr(txErr, "Edit Updates in tx err"); te != nil {
			return te
		}
		if te = SystemAuthPermService.BatchDeleteByRoleId(editReq.ID, tx); te != nil {
			return te
		}
		if te = SystemAuthPermService.BatchSaveByMenuIds(editReq.ID, editReq.MenuIds, tx); te != nil {
			return te
		}
		te = SystemAuthPermService.CacheRoleMenusByRoleId(editReq.ID)
		return te
	})
	e = response.CheckErr(err, "Edit Transaction err")
	return
}

//Del 删除角色
func (roleSrv systemAuthRoleService) Del(id uint) (e error) {
	err := core.DB.Where("id = ?", id).Limit(1).First(&system.SystemAuthRole{}).Error
	if e = response.CheckErrDBNotRecord(err, "角色已不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "Del First err"); e != nil {
		return
	}
	if r := core.DB.Where("role = ? AND is_delete = ?", id, 0).Limit(1).Find(&system.SystemAuthAdmin{}); r.RowsAffected > 0 {
		return response.AssertArgumentError.Make("角色已被管理员使用,请先移除!")
	}
	// 事务
	err = core.DB.Transaction(func(tx *gorm.DB) error {
		txErr := tx.Delete(&system.SystemAuthRole{}, "id = ?", id).Error
		var te error
		if te = response.CheckErr(txErr, "Del Delete in tx err"); te != nil {
			return te
		}
		if te = SystemAuthPermService.BatchDeleteByRoleId(id, tx); te != nil {
			return te
		}
		util.RedisUtil.HDel(config.AdminConfig.BackstageRolesKey, strconv.FormatUint(uint64(id), 10))
		return nil
	})
	e = response.CheckErr(err, "Del Transaction err")
	return
}
