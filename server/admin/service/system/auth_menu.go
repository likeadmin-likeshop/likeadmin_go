package system

import (
	"github.com/gin-gonic/gin"
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
		return
	}
	var menuResps []resp.SystemAuthMenuResp
	response.Copy(c, &menuResps, menus)
	mapList = utils.ArrayUtil.ListToTree(
		utils.ConvertUtil.StructsToMaps(menuResps), "id", "pid", "children")
	return
}

//List 菜单列表
func (menuSrv systemAuthMenuService) List() (menus []system.SystemAuthMenu) {
	core.DB.Order("menu_sort desc, id").Find(&menus)
	return
}

func (menuSrv systemAuthMenuService) Detail(menus []system.SystemAuthMenu) {
	// TODO:
}

func (menuSrv systemAuthMenuService) Add(menus []system.SystemAuthMenu) {
	// TODO:
}

func (menuSrv systemAuthMenuService) Edit(menus []system.SystemAuthMenu) {
	// TODO:
}

func (menuSrv systemAuthMenuService) Delete(menus []system.SystemAuthMenu) {
	// TODO:
}
