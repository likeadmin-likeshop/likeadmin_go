package system

import (
	"likeadmin/core"
	"likeadmin/models/system"
)

var SystemAuthMenuService = systemAuthMenuService{}

//systemAuthMenuService 系统菜单服务实现类
type systemAuthMenuService struct{}

func (menuSrv systemAuthMenuService) SelectMenuByRoleId(menus []system.SystemAuthMenu) {
	// TODO:
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
