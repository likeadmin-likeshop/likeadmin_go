package system

import (
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
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
func (menuSrv systemAuthMenuService) SelectMenuByRoleId(c *gin.Context, roleId uint) (mapList []interface{}, e error) {
	adminId := config.AdminConfig.GetAdminId(c)
	var menuIds []uint
	if menuIds, e = SystemAuthPermService.SelectMenuIdsByRoleId(roleId); e != nil {
		return
	}
	if len(menuIds) == 0 {
		menuIds = []uint{0}
	}
	chain := core.DB.Where("menu_type in ? AND is_disable = ?", []string{"M", "C"}, 0)
	if adminId != config.AdminConfig.SuperAdminId {
		chain = chain.Where("id in ?", menuIds)
	}
	var menus []system.SystemAuthMenu
	err := chain.Order("menu_sort desc, id").Find(&menus).Error
	if e = response.CheckErr(err, "SelectMenuByRoleId Find err"); e != nil {
		return
	}
	var menuResps []resp.SystemAuthMenuResp
	response.Copy(&menuResps, menus)
	mapList = util.ArrayUtil.ListToTree(
		util.ConvertUtil.StructsToMaps(menuResps), "id", "pid", "children")
	return
}

//List 菜单列表
func (menuSrv systemAuthMenuService) List() (res []interface{}, e error) {
	var menus []system.SystemAuthMenu
	err := core.DB.Order("menu_sort desc, id").Find(&menus).Error
	if e = response.CheckErr(err, "List Find err"); e != nil {
		return
	}
	var menuResps []resp.SystemAuthMenuResp
	response.Copy(&menuResps, menus)
	return util.ArrayUtil.ListToTree(
		util.ConvertUtil.StructsToMaps(menuResps), "id", "pid", "children"), nil
}

//Detail 菜单详情
func (menuSrv systemAuthMenuService) Detail(id uint) (res resp.SystemAuthMenuResp, e error) {
	var menu system.SystemAuthMenu
	err := core.DB.Where("id = ?", id).Limit(1).First(&menu).Error
	if e = response.CheckErrDBNotRecord(err, "菜单已不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "Detail First err"); e != nil {
		return
	}
	response.Copy(&res, menu)
	return
}

func (menuSrv systemAuthMenuService) Add(addReq req.SystemAuthMenuAddReq) (e error) {
	var menu system.SystemAuthMenu
	response.Copy(&menu, addReq)
	err := core.DB.Create(&menu).Error
	if e = response.CheckErr(err, "Add Create err"); e != nil {
		return
	}
	util.RedisUtil.Del(config.AdminConfig.BackstageRolesKey)
	return
}

func (menuSrv systemAuthMenuService) Edit(editReq req.SystemAuthMenuEditReq) (e error) {
	var menu system.SystemAuthMenu
	err := core.DB.Where("id = ?", editReq.ID).Limit(1).Find(&menu).Error
	if e = response.CheckErrDBNotRecord(err, "菜单已不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "Edit Find err"); e != nil {
		return
	}
	response.Copy(&menu, editReq)
	err = core.DB.Model(&menu).Updates(structs.Map(menu)).Error
	if e = response.CheckErr(err, "Edit Updates err"); e != nil {
		return
	}
	util.RedisUtil.Del(config.AdminConfig.BackstageRolesKey)
	return
}

//Del 删除菜单
func (menuSrv systemAuthMenuService) Del(id uint) (e error) {
	var menu system.SystemAuthMenu
	err := core.DB.Where("id = ?", id).Limit(1).First(&menu).Error
	if e = response.CheckErrDBNotRecord(err, "菜单已不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "Delete First err"); e != nil {
		return
	}
	r := core.DB.Where("pid = ?", id).Limit(1).Find(&system.SystemAuthMenu{})
	err = r.Error
	if e = response.CheckErr(err, "Delete Find by pid err"); e != nil {
		return
	}
	if r.RowsAffected > 0 {
		return response.AssertArgumentError.Make("请先删除子菜单再操作！")
	}
	err = core.DB.Delete(&menu).Error
	e = response.CheckErr(err, "Delete Delete err")
	return
}
