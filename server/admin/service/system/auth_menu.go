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
	err := chain.Order("menu_sort desc, id").Find(&menus).Error
	util.CheckUtil.CheckErr(err, "SelectMenuByRoleId Find err")
	var menuResps []resp.SystemAuthMenuResp
	response.Copy(&menuResps, menus)
	mapList = util.ArrayUtil.ListToTree(
		util.ConvertUtil.StructsToMaps(menuResps), "id", "pid", "children")
	return
}

//List 菜单列表
func (menuSrv systemAuthMenuService) List() []interface{} {
	var menus []system.SystemAuthMenu
	err := core.DB.Order("menu_sort desc, id").Find(&menus).Error
	util.CheckUtil.CheckErr(err, "List Find err")
	var menuResps []resp.SystemAuthMenuResp
	response.Copy(&menuResps, menus)
	return util.ArrayUtil.ListToTree(
		util.ConvertUtil.StructsToMaps(menuResps), "id", "pid", "children")
}

//Detail 菜单详情
func (menuSrv systemAuthMenuService) Detail(id uint) (res resp.SystemAuthMenuResp) {
	var menu system.SystemAuthMenu
	err := core.DB.Where("id = ?", id).Limit(1).First(&menu).Error
	util.CheckUtil.CheckErrDBNotRecord(err, "菜单已不存在!")
	util.CheckUtil.CheckErr(err, "Detail First err")
	response.Copy(&res, menu)
	return
}

func (menuSrv systemAuthMenuService) Add(addReq req.SystemAuthMenuAddReq) {
	var menu system.SystemAuthMenu
	response.Copy(&menu, addReq)
	err := core.DB.Create(&menu).Error
	util.CheckUtil.CheckErr(err, "Add Create err")
	util.RedisUtil.Del(config.AdminConfig.BackstageRolesKey)
}

func (menuSrv systemAuthMenuService) Edit(editReq req.SystemAuthMenuEditReq) {
	var menu system.SystemAuthMenu
	err := core.DB.Where("id = ?", editReq.ID).Limit(1).Find(&menu).Error
	util.CheckUtil.CheckErrDBNotRecord(err, "菜单已不存在!")
	util.CheckUtil.CheckErr(err, "Edit Find err")
	response.Copy(&menu, editReq)
	err = core.DB.Model(&menu).Updates(structs.Map(menu)).Error
	util.CheckUtil.CheckErr(err, "Edit Updates err")
	util.RedisUtil.Del(config.AdminConfig.BackstageRolesKey)
}

//Del 删除菜单
func (menuSrv systemAuthMenuService) Del(id uint) {
	var menu system.SystemAuthMenu
	err := core.DB.Where("id = ?", id).Limit(1).First(&menu).Error
	util.CheckUtil.CheckErrDBNotRecord(err, "菜单已不存在!")
	util.CheckUtil.CheckErr(err, "Delete First err")
	r := core.DB.Where("pid = ?", id).Limit(1).Find(&system.SystemAuthMenu{})
	err = r.Error
	util.CheckUtil.CheckErr(err, "Delete Find by pid err")
	if r.RowsAffected > 0 {
		panic(response.AssertArgumentError.Make("请先删除子菜单再操作！"))
	}
	err = core.DB.Delete(&menu).Error
	util.CheckUtil.CheckErr(err, "Delete Delete err")
}
