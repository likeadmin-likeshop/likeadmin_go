package system

import (
	"gorm.io/gorm"
	"likeadmin/config"
	"likeadmin/core"
	"likeadmin/core/response"
	"likeadmin/model/system"
	"likeadmin/util"
	"strconv"
	"strings"
)

var SystemAuthPermService = systemAuthPermService{}

//systemAuthPermService 系统权限服务实现类
type systemAuthPermService struct{}

//SelectMenuIdsByRoleId 根据角色ID获取菜单ID
func (permSrv systemAuthPermService) SelectMenuIdsByRoleId(roleId uint) (menuIds []uint) {
	var role system.SystemAuthRole
	err := core.DB.Where("id = ? AND is_disable = ?", roleId, 0).Limit(1).First(&role).Error
	if err != nil {
		return []uint{}
	}
	var perms []system.SystemAuthPerm
	err = core.DB.Where("role_id = ?", role.ID).Find(&perms).Error
	if err != nil {
		return []uint{}
	}
	for _, perm := range perms {
		menuIds = append(menuIds, perm.MenuId)
	}
	return
}

//CacheRoleMenusByRoleId 缓存角色菜单
func (permSrv systemAuthPermService) CacheRoleMenusByRoleId(roleId uint) (err error) {
	var perms []system.SystemAuthPerm
	err = core.DB.Where("role_id = ?", roleId).Find(&perms).Error
	if err != nil {
		core.Logger.Errorf("CacheRoleMenusByRoleId find SystemAuthPerm err: err=[%+v]", err)
		return err
	}
	var menuIds []uint
	for _, perm := range perms {
		menuIds = append(menuIds, perm.MenuId)
	}
	var menus []system.SystemAuthMenu
	err = core.DB.Where(
		"is_disable = ? and id in ? and menu_type in ?", 0, menuIds, []string{"C", "A"}).Order(
		"menu_sort, id").Find(&menus).Error
	if err != nil {
		core.Logger.Errorf("CacheRoleMenusByRoleId find SystemAuthMenu err: err=[%+v]", err)
		return err
	}
	var menuArray []string
	for _, menu := range menus {
		if menu.Perms != "" {
			menuArray = append(menuArray, strings.Trim(menu.Perms, ""))
		}
	}
	util.RedisUtil.HSet(config.AdminConfig.BackstageRolesKey, strconv.Itoa(int(roleId)), strings.Join(menuArray, ","), 0)
	return
}

//BatchSaveByMenuIds 批量写入角色菜单
func (permSrv systemAuthPermService) BatchSaveByMenuIds(roleId uint, menuIds string, db *gorm.DB) {
	if menuIds == "" {
		return
	}
	if db == nil {
		db = core.DB
	}
	err := db.Transaction(func(tx *gorm.DB) error {
		var perms []system.SystemAuthPerm
		for _, menuIdStr := range strings.Split(menuIds, ",") {
			menuId, _ := strconv.Atoi(menuIdStr)
			perms = append(perms, system.SystemAuthPerm{ID: util.ToolsUtil.MakeUuid(), RoleId: roleId, MenuId: uint(menuId)})
		}
		txErr := tx.Create(&perms).Error
		if txErr != nil {
			core.Logger.Errorf("BatchSaveByMenuIds Create err: txErr=[%+v]", txErr)
			return txErr
		}
		return nil
	})
	if err != nil {
		core.Logger.Errorf("BatchSaveByMenuIds Transaction err: err=[%+v]", err)
		panic(response.SystemError)
	}
}

//BatchDeleteByRoleId 批量删除角色菜单(根据角色ID)
func (permSrv systemAuthPermService) BatchDeleteByRoleId(roleId uint, db *gorm.DB) {
	if db == nil {
		db = core.DB
	}
	if err := db.Delete(&system.SystemAuthPerm{}, "role_id = ?", roleId).Error; err != nil {
		core.Logger.Errorf("BatchDeleteByRoleId Delete err: err=[%+v]", err)
		panic(response.SystemError)
	}
	return
}

//BatchDeleteByMenuId 批量删除角色菜单(根据菜单ID)
func (permSrv systemAuthPermService) BatchDeleteByMenuId(menuId uint) {
	if err := core.DB.Delete(&system.SystemAuthPerm{}, "menu_id = ?", menuId).Error; err != nil {
		core.Logger.Errorf("BatchDeleteByMenuId Delete err: err=[%+v]", err)
		panic(response.SystemError)
	}
	return
}
