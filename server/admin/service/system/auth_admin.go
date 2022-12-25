package system

import (
	"encoding/json"
	"likeadmin/config"
	"likeadmin/core"
	"likeadmin/models/system"
	"likeadmin/utils"
	"strconv"
)

var SystemAuthAdminService = systemAuthAdminService{}

//systemAuthAdminService 系统管理员服务实现类
type systemAuthAdminService struct{}

//FindByUsername 根据账号查找管理员
func (adminSrv systemAuthAdminService) FindByUsername(username string) (admin system.SystemAuthAdmin, err error) {
	err = core.DB.Where("username = ?", username).Limit(1).First(&admin).Error
	return
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
