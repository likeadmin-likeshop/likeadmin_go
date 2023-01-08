package system

import (
	"fmt"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"likeadmin/admin/schemas/req"
	"likeadmin/admin/schemas/resp"
	"likeadmin/config"
	"likeadmin/core"
	"likeadmin/core/request"
	"likeadmin/core/response"
	"likeadmin/model/system"
	"likeadmin/util"
	"strconv"
	"strings"
	"time"
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
	util.CheckUtil.CheckErr(err, "Self First err")
	// 角色权限
	var auths []string
	if adminId > 1 {
		roleId, _ := strconv.ParseUint(sysAdmin.Role, 10, 32)
		menuIds := SystemAuthPermService.SelectMenuIdsByRoleId(uint(roleId))
		if len(menuIds) > 0 {
			var menus []system.SystemAuthMenu
			err := core.DB.Where(
				"id in ? AND is_disable = ? AND menu_type in ?", menuIds, 0, []string{"C", "A"}).Order(
				"menu_sort, id").Find(&menus).Error
			util.CheckUtil.CheckErr(err, "Self SystemAuthMenu Find err")
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
	admin.Dept = strconv.FormatUint(uint64(sysAdmin.DeptId), 10)
	admin.Avatar = util.UrlUtil.ToAbsoluteUrl(sysAdmin.Avatar)
	return resp.SystemAuthAdminSelfResp{User: admin, Permissions: auths}
}

//List 管理员列表
func (adminSrv systemAuthAdminService) List(page request.PageReq, listReq req.SystemAuthAdminListReq) response.PageResp {
	// 分页信息
	var res response.PageResp
	response.Copy(&res, page)
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)
	// 查询
	admin := system.SystemAuthAdmin{}
	adminTbName := core.DBTableName(&admin)
	roleTbName := core.DBTableName(&system.SystemAuthRole{})
	deptTbName := core.DBTableName(&system.SystemAuthDept{})
	adminModel := core.DB.Model(&admin).Joins(
		fmt.Sprintf("LEFT JOIN %s ON %s.role = %s.id", roleTbName, adminTbName, roleTbName)).Joins(
		fmt.Sprintf("LEFT JOIN %s ON %s.dept_id = %s.id", deptTbName, adminTbName, deptTbName)).Select(
		fmt.Sprintf("%s.*, %s.name as dept, %s.name as role", adminTbName, deptTbName, roleTbName))
	// 条件
	if listReq.Username != "" {
		adminModel = adminModel.Where("username like ?", "%"+listReq.Username+"%")
	}
	if listReq.Nickname != "" {
		adminModel = adminModel.Where("nickname like ?", "%"+listReq.Nickname+"%")
	}
	if listReq.Role >= 0 {
		adminModel = adminModel.Where("role = ?", listReq.Role)
	}
	// 总数
	var count int64
	err := adminModel.Count(&count).Error
	util.CheckUtil.CheckErr(err, "List Count err")
	// 数据
	var adminResp []resp.SystemAuthAdminResp
	err = adminModel.Limit(limit).Offset(offset).Order("id desc, sort desc").Find(&adminResp).Error
	util.CheckUtil.CheckErr(err, "List Find err")
	for i := 0; i < len(adminResp); i++ {
		adminResp[i].Avatar = util.UrlUtil.ToAbsoluteUrl(adminResp[i].Avatar)
		if adminResp[i].ID == 1 {
			adminResp[i].Role = "系统管理员"
		}
	}
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Lists:    adminResp,
	}
}

//Detail 管理员详细
func (adminSrv systemAuthAdminService) Detail(id uint) (res resp.SystemAuthAdminResp) {
	var sysAdmin system.SystemAuthAdmin
	err := core.DB.Where("id = ? AND is_delete = ?", id, 0).Limit(1).First(&sysAdmin).Error
	util.CheckUtil.CheckErrDBNotRecord(err, "账号已不存在！")
	util.CheckUtil.CheckErr(err, "Detail First err")
	response.Copy(&res, sysAdmin)
	res.Avatar = util.UrlUtil.ToAbsoluteUrl(res.Avatar)
	if res.Dept == "" {
		res.Dept = strconv.FormatUint(uint64(res.DeptId), 10)
	}
	return
}

//Add 管理员新增
func (adminSrv systemAuthAdminService) Add(addReq req.SystemAuthAdminAddReq) {
	var sysAdmin system.SystemAuthAdmin
	// 检查username
	r := core.DB.Where("username = ? AND is_delete = ?", addReq.Username, 0).Limit(1).Find(&sysAdmin)
	err := r.Error
	util.CheckUtil.CheckErr(err, "Add Find by username err")
	if r.RowsAffected > 0 {
		panic(response.AssertArgumentError.Make("账号已存在换一个吧！"))
	}
	// 检查nickname
	r = core.DB.Where("nickname = ? AND is_delete = ?", addReq.Nickname, 0).Limit(1).Find(&sysAdmin)
	err = r.Error
	util.CheckUtil.CheckErr(err, "Add Find by nickname err")
	if r.RowsAffected > 0 {
		panic(response.AssertArgumentError.Make("昵称已存在换一个吧！"))
	}
	roleResp := SystemAuthRoleService.Detail(addReq.Role)
	if roleResp.IsDisable > 0 {
		panic(response.AssertArgumentError.Make("当前角色已被禁用!"))
	}
	passwdLen := len(addReq.Password)
	if !(passwdLen >= 6 && passwdLen <= 20) {
		panic(response.Failed.Make("密码必须在6~20位"))
	}
	salt := util.ToolsUtil.RandomString(5)
	response.Copy(&sysAdmin, addReq)
	sysAdmin.Role = strconv.FormatUint(uint64(addReq.Role), 10)
	sysAdmin.Salt = salt
	sysAdmin.Password = util.ToolsUtil.MakeMd5(strings.Trim(addReq.Password, " ") + salt)
	if addReq.Avatar == "" {
		addReq.Avatar = "/api/static/backend_avatar.png"
	}
	sysAdmin.Avatar = util.UrlUtil.ToRelativeUrl(addReq.Avatar)
	err = core.DB.Create(&sysAdmin).Error
	util.CheckUtil.CheckErr(err, "Add Create err")
}

//Edit 管理员编辑
func (adminSrv systemAuthAdminService) Edit(c *gin.Context, editReq req.SystemAuthAdminEditReq) {
	// 检查id
	err := core.DB.Where("id = ? AND is_delete = ?", editReq.ID, 0).Limit(1).First(&system.SystemAuthAdmin{}).Error
	util.CheckUtil.CheckErrDBNotRecord(err, "账号不存在了!")
	util.CheckUtil.CheckErr(err, "Edit First err")
	// 检查username
	var admin system.SystemAuthAdmin
	r := core.DB.Where("username = ? AND is_delete = ? AND id != ?", editReq.Username, 0, editReq.ID).Find(&admin)
	err = r.Error
	util.CheckUtil.CheckErr(err, "Edit Find by username err")
	if r.RowsAffected > 0 {
		panic(response.AssertArgumentError.Make("账号已存在换一个吧！"))
	}
	// 检查nickname
	r = core.DB.Where("nickname = ? AND is_delete = ? AND id != ?", editReq.Nickname, 0, editReq.ID).Find(&admin)
	err = r.Error
	util.CheckUtil.CheckErr(err, "Edit Find by nickname err")
	if r.RowsAffected > 0 {
		panic(response.AssertArgumentError.Make("昵称已存在换一个吧！"))
	}
	// 检查role
	if editReq.Role > 0 && editReq.ID != 1 {
		SystemAuthRoleService.Detail(editReq.Role)
	}
	// 更新管理员信息
	adminMap := structs.Map(editReq)
	delete(adminMap, "ID")
	adminMap["Avatar"] = util.UrlUtil.ToRelativeUrl(editReq.Avatar)
	role := editReq.Role
	if editReq.ID == 1 {
		role = 0
	}
	adminMap["Role"] = strconv.FormatUint(uint64(role), 10)
	if editReq.ID == 1 {
		delete(adminMap, "Username")
	}
	if editReq.Password != "" {
		passwdLen := len(editReq.Password)
		if !(passwdLen >= 6 && passwdLen <= 20) {
			panic(response.Failed.Make("密码必须在6~20位"))
		}
		salt := util.ToolsUtil.RandomString(5)
		adminMap["Salt"] = salt
		adminMap["Password"] = util.ToolsUtil.MakeMd5(strings.Trim(editReq.Password, "") + salt)
	} else {
		delete(adminMap, "Password")
	}
	err = core.DB.Model(&admin).Where("id = ?", editReq.ID).Updates(adminMap).Error
	util.CheckUtil.CheckErr(err, "Edit Updates err")
	adminSrv.CacheAdminUserByUid(editReq.ID)
	// 如果更改自己的密码,则删除旧缓存
	adminId := config.AdminConfig.GetAdminId(c)
	if editReq.Password != "" && editReq.ID == adminId {
		token := c.Request.Header.Get("token")
		util.RedisUtil.Del(config.AdminConfig.BackstageTokenKey + token)
		adminSetKey := config.AdminConfig.BackstageTokenSet + strconv.FormatUint(uint64(adminId), 10)
		ts := util.RedisUtil.SGet(adminSetKey)
		if len(ts) > 0 {
			var tokenKeys []string
			for _, t := range ts {
				tokenKeys = append(tokenKeys, config.AdminConfig.BackstageTokenKey+t)
			}
			util.RedisUtil.Del(tokenKeys...)
		}
		util.RedisUtil.Del(adminSetKey)
		util.RedisUtil.SSet(adminSetKey, token)
	}
}

//Update 管理员更新
func (adminSrv systemAuthAdminService) Update(c *gin.Context, updateReq req.SystemAuthAdminUpdateReq, adminId uint) {
	// 检查id
	var admin system.SystemAuthAdmin
	err := core.DB.Where("id = ? AND is_delete = ?", adminId, 0).Limit(1).First(&admin).Error
	util.CheckUtil.CheckErrDBNotRecord(err, "账号不存在了!")
	util.CheckUtil.CheckErr(err, "Update First err")
	// 更新管理员信息
	adminMap := structs.Map(updateReq)
	delete(adminMap, "CurrPassword")
	avatar := "/api/static/backend_avatar.png"
	if updateReq.Avatar != "" {
		avatar = updateReq.Avatar
	}
	adminMap["Avatar"] = util.UrlUtil.ToRelativeUrl(avatar)
	delete(adminMap, "aaa")
	if updateReq.Password != "" {
		currPass := util.ToolsUtil.MakeMd5(updateReq.CurrPassword + admin.Salt)
		if currPass != admin.Password {
			panic(response.Failed.Make("当前密码不正确!"))
		}
		passwdLen := len(updateReq.Password)
		if !(passwdLen >= 6 && passwdLen <= 20) {
			panic(response.Failed.Make("密码必须在6~20位"))
		}
		salt := util.ToolsUtil.RandomString(5)
		adminMap["Salt"] = salt
		adminMap["Password"] = util.ToolsUtil.MakeMd5(strings.Trim(updateReq.Password, " ") + salt)
	} else {
		delete(adminMap, "Password")
	}
	err = core.DB.Model(&admin).Updates(adminMap).Error
	util.CheckUtil.CheckErr(err, "Update Updates err")
	adminSrv.CacheAdminUserByUid(adminId)
	// 如果更改自己的密码,则删除旧缓存
	if updateReq.Password != "" {
		token := c.Request.Header.Get("token")
		util.RedisUtil.Del(config.AdminConfig.BackstageTokenKey + token)
		adminSetKey := config.AdminConfig.BackstageTokenSet + strconv.FormatUint(uint64(adminId), 10)
		ts := util.RedisUtil.SGet(adminSetKey)
		if len(ts) > 0 {
			var tokenKeys []string
			for _, t := range ts {
				tokenKeys = append(tokenKeys, config.AdminConfig.BackstageTokenKey+t)
			}
			util.RedisUtil.Del(tokenKeys...)
		}
		util.RedisUtil.Del(adminSetKey)
		util.RedisUtil.SSet(adminSetKey, token)
	}
}

//Del 管理员删除
func (adminSrv systemAuthAdminService) Del(c *gin.Context, id uint) {
	var admin system.SystemAuthAdmin
	err := core.DB.Where("id = ? AND is_delete = ?", id, 0).Limit(1).First(&admin).Error
	util.CheckUtil.CheckErrDBNotRecord(err, "账号已不存在!")
	util.CheckUtil.CheckErr(err, "Del First err")
	if id == 1 {
		panic(response.AssertArgumentError.Make("系统管理员不允许删除!"))
	}
	if id == config.AdminConfig.GetAdminId(c) {
		panic(response.AssertArgumentError.Make("不能删除自己!"))
	}
	err = core.DB.Model(&admin).Updates(system.SystemAuthAdmin{IsDelete: 1, DeleteTime: time.Now().Unix()}).Error
	util.CheckUtil.CheckErr(err, "Del Updates err")
}

//Disable 管理员状态切换
func (adminSrv systemAuthAdminService) Disable(c *gin.Context, id uint) {
	var admin system.SystemAuthAdmin
	err := core.DB.Where("id = ? AND is_delete = ?", id, 0).Limit(1).Find(&admin).Error
	util.CheckUtil.CheckErr(err, "Disable Find err")
	if admin.ID == 0 {
		panic(response.AssertArgumentError.Make("账号已不存在!"))
	}
	if id == config.AdminConfig.GetAdminId(c) {
		panic(response.AssertArgumentError.Make("不能禁用自己!"))
	}
	var isDisable uint8
	if admin.IsDisable == 0 {
		isDisable = 1
	}
	err = core.DB.Model(&admin).Updates(system.SystemAuthAdmin{IsDisable: isDisable, UpdateTime: time.Now().Unix()}).Error
	util.CheckUtil.CheckErr(err, "Disable Updates err")
}

//CacheAdminUserByUid 缓存管理员
func (adminSrv systemAuthAdminService) CacheAdminUserByUid(id uint) (err error) {
	var admin system.SystemAuthAdmin
	err = core.DB.Where("id = ?", id).Limit(1).First(&admin).Error
	if err != nil {
		return
	}
	str, err := util.ToolsUtil.ObjToJson(&admin)
	if err != nil {
		return
	}
	util.RedisUtil.HSet(config.AdminConfig.BackstageManageKey, strconv.FormatUint(uint64(admin.ID), 10), str, 0)
	return nil
}
