package system

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"likeadmin/admin/schemas/resp"
	"likeadmin/config"
	"likeadmin/core"
	"likeadmin/core/response"
	"likeadmin/models/system"
	"likeadmin/utils"
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
	if err != nil {
		core.Logger.Errorf("SelectMenuByRoleId Find err: err=[%+v]", err)
		panic(response.SystemError)
	}
	var menuResps []resp.SystemAuthMenuResp
	response.Copy(&menuResps, menus)
	mapList = utils.ArrayUtil.ListToTree(
		utils.ConvertUtil.StructsToMaps(menuResps), "id", "pid", "children")
	return
}

//List 菜单列表
func (menuSrv systemAuthMenuService) List() []interface{} {
	var menus []system.SystemAuthMenu
	err := core.DB.Order("menu_sort desc, id").Find(&menus).Error
	if err != nil {
		core.Logger.Errorf("List Find err: err=[%+v]", err)
		panic(response.SystemError)
	}
	var menuResps []resp.SystemAuthMenuResp
	response.Copy(&menuResps, menus)
	return utils.ArrayUtil.ListToTree(
		utils.ConvertUtil.StructsToMaps(menuResps), "id", "pid", "children")
}

func (menuSrv systemAuthMenuService) Detail(id uint) (res resp.SystemAuthMenuResp) {
	var menu system.SystemAuthMenu
	err := core.DB.Where("id = ?", id).Limit(1).First(&menu).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		panic(response.AssertArgumentError.Make("菜单已不存在!"))
	} else if err != nil {
		core.Logger.Errorf("Detail Find err: err=[%+v]", err)
		panic(response.SystemError)
	}
	response.Copy(&res, menu)
	return
}

func (menuSrv systemAuthMenuService) Add(menus []system.SystemAuthMenu) {
	// TODO: Add
}

func (menuSrv systemAuthMenuService) Edit(menus []system.SystemAuthMenu) {
	// TODO: Edit
}

func (menuSrv systemAuthMenuService) Delete(menus []system.SystemAuthMenu) {
	// TODO: Delete
}
