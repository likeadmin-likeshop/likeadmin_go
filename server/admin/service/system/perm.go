package system

import (
	"gorm.io/gorm"
	"likeadmin/config"
	"likeadmin/core/response"
	"likeadmin/model/system"
	"likeadmin/util"
	"strconv"
	"strings"
)

type ISystemAuthPermService interface {
	SelectMenuIdsByRoleId(roleId uint) (menuIds []uint, e error)
	CacheRoleMenusByRoleId(roleId uint) (e error)
	BatchSaveByMenuIds(roleId uint, menuIds string, db *gorm.DB) (e error)
	BatchDeleteByRoleId(roleId uint, db *gorm.DB) (e error)
	BatchDeleteByMenuId(menuId uint) (e error)
}

//NewSystemAuthPermService 初始化
func NewSystemAuthPermService(db *gorm.DB) ISystemAuthPermService {
	return &systemAuthPermService{db: db}
}

//systemAuthPermService 系统权限服务实现类
type systemAuthPermService struct {
	db *gorm.DB
}

//SelectMenuIdsByRoleId 根据角色ID获取菜单ID
func (permSrv systemAuthPermService) SelectMenuIdsByRoleId(roleId uint) (menuIds []uint, e error) {
	var role system.SystemAuthRole
	err := permSrv.db.Where("id = ? AND is_disable = ?", roleId, 0).Limit(1).First(&role).Error
	if e = response.CheckErr(err, "SelectMenuIdsByRoleId First err"); e != nil {
		return []uint{}, e
	}
	var perms []system.SystemAuthPerm
	err = permSrv.db.Where("role_id = ?", role.ID).Find(&perms).Error
	if e = response.CheckErr(err, "SelectMenuIdsByRoleId Find err"); e != nil {
		return []uint{}, e
	}
	for _, perm := range perms {
		menuIds = append(menuIds, perm.MenuId)
	}
	return
}

//CacheRoleMenusByRoleId 缓存角色菜单
func (permSrv systemAuthPermService) CacheRoleMenusByRoleId(roleId uint) (e error) {
	var perms []system.SystemAuthPerm
	err := permSrv.db.Where("role_id = ?", roleId).Find(&perms).Error
	if e = response.CheckErr(err, "CacheRoleMenusByRoleId Find perms err"); e != nil {
		return
	}
	var menuIds []uint
	for _, perm := range perms {
		menuIds = append(menuIds, perm.MenuId)
	}
	var menus []system.SystemAuthMenu
	err = permSrv.db.Where(
		"is_disable = ? and id in ? and menu_type in ?", 0, menuIds, []string{"C", "A"}).Order(
		"menu_sort, id").Find(&menus).Error
	if e = response.CheckErr(err, "CacheRoleMenusByRoleId Find menus err"); e != nil {
		return
	}
	var menuArray []string
	for _, menu := range menus {
		if menu.Perms != "" {
			menuArray = append(menuArray, strings.Trim(menu.Perms, ""))
		}
	}
	util.RedisUtil.HSet(config.AdminConfig.BackstageRolesKey, strconv.FormatUint(uint64(roleId), 10), strings.Join(menuArray, ","), 0)
	return
}

//BatchSaveByMenuIds 批量写入角色菜单
func (permSrv systemAuthPermService) BatchSaveByMenuIds(roleId uint, menuIds string, db *gorm.DB) (e error) {
	if menuIds == "" {
		return
	}
	if db == nil {
		db = permSrv.db
	}
	err := db.Transaction(func(tx *gorm.DB) error {
		var perms []system.SystemAuthPerm
		for _, menuIdStr := range strings.Split(menuIds, ",") {
			menuId, _ := strconv.ParseUint(menuIdStr, 10, 32)
			perms = append(perms, system.SystemAuthPerm{ID: util.ToolsUtil.MakeUuid(), RoleId: roleId, MenuId: uint(menuId)})
		}
		txErr := tx.Create(&perms).Error
		var te error
		te = response.CheckErr(txErr, "BatchSaveByMenuIds Create in tx err")
		return te
	})
	e = response.CheckErr(err, "BatchSaveByMenuIds Transaction err")
	return
}

//BatchDeleteByRoleId 批量删除角色菜单(根据角色ID)
func (permSrv systemAuthPermService) BatchDeleteByRoleId(roleId uint, db *gorm.DB) (e error) {
	if db == nil {
		db = permSrv.db
	}
	err := db.Delete(&system.SystemAuthPerm{}, "role_id = ?", roleId).Error
	e = response.CheckErr(err, "BatchDeleteByRoleId Delete err")
	return
}

//BatchDeleteByMenuId 批量删除角色菜单(根据菜单ID)
func (permSrv systemAuthPermService) BatchDeleteByMenuId(menuId uint) (e error) {
	err := permSrv.db.Delete(&system.SystemAuthPerm{}, "menu_id = ?", menuId).Error
	e = response.CheckErr(err, "BatchDeleteByMenuId Delete err")
	return
}
