package system

import (
	"encoding/json"
	"likeadmin/admin/schemas/resp"
	"likeadmin/config"
	"likeadmin/core"
	"likeadmin/core/response"
	"likeadmin/models/system"
	"likeadmin/utils"
	"strconv"
	"strings"
)

var SystemAuthAdminService = systemAuthAdminService{}

//systemAuthAdminService 系统管理员服务实现类
type systemAuthAdminService struct{}

//FindByUsername 根据账号查找管理员
func (adminSrv systemAuthAdminService) FindByUsername(username string) (admin system.SystemAuthAdmin, err error) {
	err = core.DB.Where("username = ?", username).Limit(1).First(&admin).Error
	return
}

//Self 当前管理员
func (adminSrv systemAuthAdminService) Self(adminId uint) (res resp.SystemAuthAdminSelfResp) {
	// 管理员信息
	var sysAdmin system.SystemAuthAdmin
	err := core.DB.Where("id = ? AND is_delete = ?", adminId, 0).Limit(1).First(&sysAdmin).Error
	if err != nil {
		core.Logger.Errorf("Self First err: err=[%+v]", err)
		panic(response.SystemError)
	}
	// 角色权限
	var auths []string
	if adminId > 1 {
		roleId, _ := strconv.Atoi(sysAdmin.Role)
		menuIds := SystemAuthPermService.SelectMenuIdsByRoleId(uint(roleId))
		if len(menuIds) > 0 {
			var menus []system.SystemAuthMenu
			err = core.DB.Where(
				"id in ? AND is_disable = ? AND menu_type in ?", menuIds, 0, []string{"C", "A"}).Order(
				"menu_sort, id").Find(&menus).Error
			if err != nil {
				core.Logger.Errorf("Self SystemAuthMenu Find err: err=[%+v]", err)
				panic(response.SystemError)
			}
			if len(menus) > 0 {
				for _, v := range menus {
					auths = append(auths, strings.Trim(v.Perms, " "))
				}
			}
		}
		if len(auths) > 0 {
			auths = append(auths, "")
		}
	} else {
		auths = append(auths, "*")
	}
	var admin resp.SystemAuthAdminSelfOneResp
	response.Copy(&admin, sysAdmin)
	admin.Dept = strconv.Itoa(int(sysAdmin.DeptId))
	admin.Avatar = utils.UrlUtil.ToAbsoluteUrl(sysAdmin.Avatar)
	return resp.SystemAuthAdminSelfResp{User: admin, Permissions: auths}
}

//CacheAdminUserByUid 缓存管理员
func (adminSrv systemAuthAdminService) CacheAdminUserByUid(id uint) (err error) {
	var admin system.SystemAuthAdmin
	err = core.DB.Where("id = ?", id).Limit(1).First(&admin).Error
	if err != nil {
		return
	}
	b, err := json.Marshal(admin)
	if err != nil {
		return
	}
	utils.RedisUtil.HSet(config.AdminConfig.BackstageManageKey, strconv.Itoa(int(admin.ID)), string(b), 0)
	return nil
}
