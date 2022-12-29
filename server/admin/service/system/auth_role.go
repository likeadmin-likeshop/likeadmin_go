package system

import (
	"likeadmin/admin/schemas/resp"
	"likeadmin/core"
	"likeadmin/core/request"
	"likeadmin/core/response"
	"likeadmin/models/system"
)

var SystemAuthRoleService = systemAuthRoleService{}

//systemAuthRoleService 系统角色服务实现类
type systemAuthRoleService struct{}

// List 根据角色ID获取菜单ID
func (roleSrv systemAuthRoleService) List(page request.PageReq) response.PageResp {
	var res response.PageResp
	response.Copy(&res, page)
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)
	roleModel := core.DB.Model(&system.SystemAuthRole{})
	var count int64
	err := roleModel.Count(&count).Error
	if err != nil {
		core.Logger.Errorf("List Count err: err=[%+v]", err)
		panic(response.SystemError)
	}
	var roles []system.SystemAuthRole
	err = roleModel.Limit(limit).Offset(offset).Order("sort desc, id desc").Find(&roles).Error
	if err != nil {
		core.Logger.Errorf("List Find err: err=[%+v]", err)
		panic(response.SystemError)
	}
	var roleResp []resp.SystemAuthRoleResp
	response.Copy(&roleResp, roles)
	for i := 0; i < len(roleResp); i++ {
		roleResp[i].Menus = []uint{}
		roleResp[i].Member = roleSrv.getMemberCnt(roleResp[i].ID)
	}
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Lists:    roleResp,
	}
}

// getMemberCnt 根据角色ID获取成员数量
func (roleSrv systemAuthRoleService) getMemberCnt(roleId uint) (count int64) {
	core.DB.Model(&system.SystemAuthAdmin{}).Where(
		"role = ? AND is_delete = ?", roleId, 0).Count(&count)
	return
}
