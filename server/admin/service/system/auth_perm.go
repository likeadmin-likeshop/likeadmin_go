package system

import (
	"likeadmin/config"
	"likeadmin/core"
	"likeadmin/core/response"
	"likeadmin/models/system"
	"likeadmin/utils"
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
	utils.RedisUtil.HSet(config.AdminConfig.BackstageRolesKey, strconv.Itoa(int(roleId)), strings.Join(menuArray, ","), 0)
	return
}

//BatchSaveByMenuIds 批量写入角色菜单
func (permSrv systemAuthPermService) BatchSaveByMenuIds(roleId uint, menuIds string) {
	if menuIds == "" {
		return
	}
	var perms []system.SystemAuthPerm
	for _, menuIdStr := range strings.Split(menuIds, ",") {
		menuId, _ := strconv.Atoi(menuIdStr)
		perms = append(perms, system.SystemAuthPerm{ID: utils.ToolsUtil.MakeUuid(), RoleId: roleId, MenuId: uint(menuId)})
	}
	err := core.DB.Create(&perms).Error
	if err != nil {
		core.Logger.Errorf("BatchSaveByMenuIds Create err: err=[%+v]", err)
		panic(response.SystemError)
	}
}

//BatchDeleteByRoleId 批量删除角色菜单(根据角色ID)
func (permSrv systemAuthPermService) BatchDeleteByRoleId(roleId uint) (err error) {
	err = core.DB.Delete(&system.SystemAuthPerm{}, "role_id = ?", roleId).Error
	return
}

//BatchDeleteByMenuId 批量删除角色菜单(根据菜单ID)
func (permSrv systemAuthPermService) BatchDeleteByMenuId(menuId uint) (err error) {
	err = core.DB.Delete(&system.SystemAuthPerm{}, "menu_id = ?", menuId).Error
	return
}
