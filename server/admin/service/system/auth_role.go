package system

import (
	"errors"
	"github.com/fatih/structs"
	"gorm.io/gorm"
	"likeadmin/admin/schemas/req"
	"likeadmin/admin/schemas/resp"
	"likeadmin/config"
	"likeadmin/core"
	"likeadmin/core/request"
	"likeadmin/core/response"
	"likeadmin/models/system"
	"likeadmin/utils"
	"strconv"
	"strings"
)

var SystemAuthRoleService = systemAuthRoleService{}

//systemAuthRoleService 系统角色服务实现类
type systemAuthRoleService struct{}

//All 角色所有
func (roleSrv systemAuthRoleService) All() (res []resp.SystemAuthRoleSimpleResp) {
	var roles []system.SystemAuthRole
	err := core.DB.Order("sort desc, id desc").Find(&roles).Error
	if err != nil {
		core.Logger.Errorf("All Find err: err=[%+v]", err)
		panic(response.SystemError)
	}
	response.Copy(&res, roles)
	return
}

//List 根据角色ID获取菜单ID
func (roleSrv systemAuthRoleService) List(page request.PageReq) response.PageResp {
	var res response.PageResp
	response.Copy(&res, page)
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)
	roleModel := core.DB.Model(&system.SystemAuthRole{})
	var count int64
	err := roleModel.Count(&count).Error
	if err != nil {
		core.Logger.Errorf("List Count err: err=[%+v]", err)
		panic(response.SystemError)
	}
	var roles []system.SystemAuthRole
	err = roleModel.Limit(limit).Offset(offset).Order("sort desc, id desc").Find(&roles).Error
	if err != nil {
		core.Logger.Errorf("List Find err: err=[%+v]", err)
		panic(response.SystemError)
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
	}
}

//Detail 角色详情
func (roleSrv systemAuthRoleService) Detail(id uint) (res resp.SystemAuthRoleResp) {
	var role system.SystemAuthRole
	err := core.DB.Where("id = ?", id).Limit(1).First(&role).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		panic(response.AssertArgumentError.Make("角色已不存在!"))
	} else if err != nil {
		core.Logger.Errorf("Detail First err: err=[%+v]", err)
		panic(response.SystemError)
	}
	response.Copy(&res, role)
	res.Member = roleSrv.getMemberCnt(role.ID)
	res.Menus = SystemAuthPermService.SelectMenuIdsByRoleId(role.ID)
	return
}

//getMemberCnt 根据角色ID获取成员数量
func (roleSrv systemAuthRoleService) getMemberCnt(roleId uint) (count int64) {
	core.DB.Model(&system.SystemAuthAdmin{}).Where(
		"role = ? AND is_delete = ?", roleId, 0).Count(&count)
	return
}

//Add 新增角色
func (roleSrv systemAuthRoleService) Add(addReq req.SystemAuthRoleAddReq) {
	var role system.SystemAuthRole
	r := core.DB.Where("name = ?", strings.Trim(addReq.Name, " ")).Limit(1).First(&role)
	if r.RowsAffected > 0 {
		panic(response.AssertArgumentError.Make("角色名称已存在!"))
	}
	response.Copy(&role, addReq)
	role.Name = strings.Trim(addReq.Name, " ")
	// 事务
	err := core.DB.Transaction(func(tx *gorm.DB) error {
		txErr := tx.Create(&role).Error
		if txErr != nil {
			core.Logger.Errorf("Add Create err: txErr=[%+v]", txErr)
			return txErr
		}
		SystemAuthPermService.BatchSaveByMenuIds(role.ID, addReq.MenuIds, tx)
		return nil
	})
	if err != nil {
		core.Logger.Errorf("Add Transaction err: err=[%+v]", err)
		panic(response.SystemError)
	}
}

//Edit 编辑角色
func (roleSrv systemAuthRoleService) Edit(editReq req.SystemAuthRoleEditReq) {
	err := core.DB.Where("id = ?", editReq.ID).Limit(1).First(&system.SystemAuthRole{}).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		panic(response.AssertArgumentError.Make("角色已不存在!"))
	} else if err != nil {
		core.Logger.Errorf("Edit First err: err=[%+v]", err)
		panic(response.SystemError)
	}
	var role system.SystemAuthRole
	r := core.DB.Where("id != ? AND name = ?", editReq.ID, strings.Trim(editReq.Name, " ")).Limit(1).First(&role)
	if r.RowsAffected > 0 {
		panic(response.AssertArgumentError.Make("角色名称已存在!"))
	}
	role.ID = editReq.ID
	roleMap := structs.Map(editReq)
	delete(roleMap, "ID")
	delete(roleMap, "MenuIds")
	roleMap["Name"] = strings.Trim(editReq.Name, " ")
	// 事务
	err = core.DB.Transaction(func(tx *gorm.DB) error {
		txErr := tx.Model(&role).Updates(roleMap).Error
		if txErr != nil {
			core.Logger.Errorf("Edit Updates err: txErr=[%+v]", txErr)
			return txErr
		}
		SystemAuthPermService.BatchDeleteByRoleId(editReq.ID, tx)
		SystemAuthPermService.BatchSaveByMenuIds(editReq.ID, editReq.MenuIds, tx)
		SystemAuthPermService.CacheRoleMenusByRoleId(editReq.ID)
		return nil
	})
	if err != nil {
		core.Logger.Errorf("Edit Transaction err: err=[%+v]", err)
		panic(response.SystemError)
	}
}

//Del 删除角色
func (roleSrv systemAuthRoleService) Del(id uint) {
	err := core.DB.Where("id = ?", id).Limit(1).First(&system.SystemAuthRole{}).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		panic(response.AssertArgumentError.Make("角色已不存在!"))
	} else if err != nil {
		core.Logger.Errorf("Del First err: err=[%+v]", err)
		panic(response.SystemError)
	}
	r := core.DB.Where("role = ? AND is_delete = ?", id, 0).Limit(1).Find(&system.SystemAuthAdmin{})
	if r.RowsAffected > 0 {
		panic(response.AssertArgumentError.Make("角色已被管理员使用,请先移除!"))
	}
	// 事务
	err = core.DB.Transaction(func(tx *gorm.DB) error {
		txErr := tx.Delete(&system.SystemAuthRole{}, "id = ?", id).Error
		if txErr != nil {
			core.Logger.Errorf("Del Delete err: txErr=[%+v]", txErr)
			return txErr
		}
		SystemAuthPermService.BatchDeleteByRoleId(id, tx)
		utils.RedisUtil.HDel(config.AdminConfig.BackstageRolesKey, strconv.Itoa(int(id)))
		return nil
	})
	if err != nil {
		core.Logger.Errorf("Edit Transaction err: err=[%+v]", err)
		panic(response.SystemError)
	}
}
