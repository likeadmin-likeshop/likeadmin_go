package system

import (
	"likeadmin/config"
	"likeadmin/core"
	"likeadmin/models/system"
	"likeadmin/utils"
	"strconv"
	"strings"
)

var SystemAuthPermService = systemAuthPermService{}

//systemAuthPermService 系统权限服务实现类
type systemAuthPermService struct{}

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