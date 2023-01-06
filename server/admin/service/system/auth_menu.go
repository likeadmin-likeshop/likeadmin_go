package system

import (
	"errors"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"likeadmin/admin/schemas/req"
	"likeadmin/admin/schemas/resp"
	"likeadmin/config"
	"likeadmin/core"
	"likeadmin/core/response"
	"likeadmin/model/system"
	"likeadmin/util"
)

var SystemAuthMenuService = systemAuthMenuService{}

//systemAuthMenuService 系统菜单服务实现类
type systemAuthMenuService struct{}

//SelectMenuByRoleId 根据角色ID获取菜单
func (menuSrv systemAuthMenuService) SelectMenuByRoleId(c *gin.Context, roleId uint) (mapList []interface{}) {
	adminId := config.AdminConfig.GetAdminId(c)
	menuIds := SystemAuthPermService.SelectMenuIdsByRoleId(roleId)
	if len(menuIds) == 0 {
		menuIds = []uint{0}
	}
	chain := core.DB.Where("menu_type in ? AND is_disable = ?", []string{"M", "C"}, 0)
	if adminId != config.AdminConfig.SuperAdminId {
		chain = chain.Where("id in ?", menuIds)
	}
	var menus []system.SystemAuthMenu
	if err := chain.Order("menu_sort desc, id").Find(&menus).Error; err != nil {
		core.Logger.Errorf("SelectMenuByRoleId Find err: err=[%+v]", err)
		panic(response.SystemError)
	}
	var menuResps []resp.SystemAuthMenuResp
	response.Copy(&menuResps, menus)
	mapList = util.ArrayUtil.ListToTree(
		util.ConvertUtil.StructsToMaps(menuResps), "id", "pid", "children")
	return
}

//List 菜单列表
func (menuSrv systemAuthMenuService) List() []interface{} {
	var menus []system.SystemAuthMenu
	if err := core.DB.Order("menu_sort desc, id").Find(&menus).Error; err != nil {
		core.Logger.Errorf("List Find err: err=[%+v]", err)
		panic(response.SystemError)
	}
	var menuResps []resp.SystemAuthMenuResp
	response.Copy(&menuResps, menus)
	return util.ArrayUtil.ListToTree(
		util.ConvertUtil.StructsToMaps(menuResps), "id", "pid", "children")
}

//Detail 菜单详情
func (menuSrv systemAuthMenuService) Detail(id uint) (res resp.SystemAuthMenuResp) {
	var menu system.SystemAuthMenu
	err := core.DB.Where("id = ?", id).Limit(1).First(&menu).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		panic(response.AssertArgumentError.Make("菜单已不存在!"))
	} else if err != nil {
		core.Logger.Errorf("Detail First err: err=[%+v]", err)
		panic(response.SystemError)
	}
	response.Copy(&res, menu)
	return
}

func (menuSrv systemAuthMenuService) Add(addReq req.SystemAuthMenuAddReq) {
	var menu system.SystemAuthMenu
	response.Copy(&menu, addReq)
	if err := core.DB.Create(&menu).Error; err != nil {
		core.Logger.Errorf("Add Create err: err=[%+v]", err)
		panic(response.SystemError)
	}
	util.RedisUtil.Del(config.AdminConfig.BackstageRolesKey)
}

func (menuSrv systemAuthMenuService) Edit(editReq req.SystemAuthMenuEditReq) {
	var menu system.SystemAuthMenu
	err := core.DB.Where("id = ?", editReq.ID).Limit(1).Find(&menu).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		panic(response.AssertArgumentError.Make("菜单已不存在!"))
	} else if err != nil {
		core.Logger.Errorf("Edit Find err: err=[%+v]", err)
		panic(response.SystemError)
	}
	response.Copy(&menu, editReq)
	if err = core.DB.Model(&menu).Updates(structs.Map(menu)).Error; err != nil {
		core.Logger.Errorf("Edit Updates err: err=[%+v]", err)
		panic(response.SystemError)
	}
	util.RedisUtil.Del(config.AdminConfig.BackstageRolesKey)
}

//Del 删除菜单
func (menuSrv systemAuthMenuService) Del(id uint) {
	var menu system.SystemAuthMenu
	err := core.DB.Where("id = ?", id).Limit(1).First(&menu).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		panic(response.AssertArgumentError.Make("菜单已不存在!"))
	} else if err != nil {
		core.Logger.Errorf("Delete First err: err=[%+v]", err)
		panic(response.SystemError)
	}
	r := core.DB.Where("pid = ?", id).Limit(1).Find(&system.SystemAuthMenu{})
	err = r.Error
	if err != nil {
		core.Logger.Errorf("Delete Find by pid err: err=[%+v]", err)
		panic(response.SystemError)
	}
	if r.RowsAffected > 0 {
		panic(response.AssertArgumentError.Make("请先删除子菜单再操作！"))
	}
	if err = core.DB.Delete(&menu).Error; err != nil {
		core.Logger.Errorf("Delete Delete err: err=[%+v]", err)
		panic(response.SystemError)
	}
}
