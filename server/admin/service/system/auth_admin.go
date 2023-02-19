package system

import (
	"fmt"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

//NewSystemAuthAdminService 初始化
func NewSystemAuthAdminService(db *gorm.DB, permSrv *SystemAuthPermService, roleSrv *SystemAuthRoleService) *SystemAuthAdminService {
	return &SystemAuthAdminService{db: db, permSrv: permSrv, roleSrv: roleSrv}
}

//SystemAuthAdminService 系统管理员服务实现类
type SystemAuthAdminService struct {
	db      *gorm.DB
	permSrv *SystemAuthPermService
	roleSrv *SystemAuthRoleService
}

//FindByUsername 根据账号查找管理员
func (adminSrv SystemAuthAdminService) FindByUsername(username string) (admin system.SystemAuthAdmin, err error) {
	err = adminSrv.db.Where("username = ?", username).Limit(1).First(&admin).Error
	return
}

//Self 当前管理员
func (adminSrv SystemAuthAdminService) Self(adminId uint) (res resp.SystemAuthAdminSelfResp, e error) {
	// 管理员信息
	var sysAdmin system.SystemAuthAdmin
	err := adminSrv.db.Where("id = ? AND is_delete = ?", adminId, 0).Limit(1).First(&sysAdmin).Error
	if e = response.CheckErr(err, "Self First err"); e != nil {
		return
	}
	// 角色权限
	var auths []string
	if adminId > 1 {
		roleId, _ := strconv.ParseUint(sysAdmin.Role, 10, 32)
		var menuIds []uint
		if menuIds, e = adminSrv.permSrv.SelectMenuIdsByRoleId(uint(roleId)); e != nil {
			return
		}
		if len(menuIds) > 0 {
			var menus []system.SystemAuthMenu
			err := adminSrv.db.Where(
				"id in ? AND is_disable = ? AND menu_type in ?", menuIds, 0, []string{"C", "A"}).Order(
				"menu_sort, id").Find(&menus).Error
			if e = response.CheckErr(err, "Self SystemAuthMenu Find err"); e != nil {
				return
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
	admin.Dept = strconv.FormatUint(uint64(sysAdmin.DeptId), 10)
	admin.Avatar = util.UrlUtil.ToAbsoluteUrl(sysAdmin.Avatar)
	return resp.SystemAuthAdminSelfResp{User: admin, Permissions: auths}, nil
}

//List 管理员列表
func (adminSrv SystemAuthAdminService) List(page request.PageReq, listReq req.SystemAuthAdminListReq) (res response.PageResp, e error) {
	// 分页信息
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)
	// 查询
	adminTbName := core.DBTableName(&system.SystemAuthAdmin{})
	roleTbName := core.DBTableName(&system.SystemAuthRole{})
	deptTbName := core.DBTableName(&system.SystemAuthDept{})
	adminModel := adminSrv.db.Table(adminTbName+" AS admin").Where("admin.is_delete = ?", 0).Joins(
		fmt.Sprintf("LEFT JOIN %s ON admin.role = %s.id", roleTbName, roleTbName)).Joins(
		fmt.Sprintf("LEFT JOIN %s ON admin.dept_id = %s.id", deptTbName, deptTbName)).Select(
		fmt.Sprintf("admin.*, %s.name as dept, %s.name as role", deptTbName, roleTbName))
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
	if e = response.CheckErr(err, "List Count err"); e != nil {
		return
	}
	// 数据
	var adminResp []resp.SystemAuthAdminResp
	err = adminModel.Limit(limit).Offset(offset).Order("id desc, sort desc").Find(&adminResp).Error
	if e = response.CheckErr(err, "List Find err"); e != nil {
		return
	}
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
	}, nil
}

//Detail 管理员详细
func (adminSrv SystemAuthAdminService) Detail(id uint) (res resp.SystemAuthAdminResp, e error) {
	var sysAdmin system.SystemAuthAdmin
	err := adminSrv.db.Where("id = ? AND is_delete = ?", id, 0).Limit(1).First(&sysAdmin).Error
	if e = response.CheckErrDBNotRecord(err, "账号已不存在！"); e != nil {
		return
	}
	if e = response.CheckErr(err, "Detail First err"); e != nil {
		return
	}
	response.Copy(&res, sysAdmin)
	res.Avatar = util.UrlUtil.ToAbsoluteUrl(res.Avatar)
	if res.Dept == "" {
		res.Dept = strconv.FormatUint(uint64(res.DeptId), 10)
	}
	return
}

//Add 管理员新增
func (adminSrv SystemAuthAdminService) Add(addReq req.SystemAuthAdminAddReq) (e error) {
	var sysAdmin system.SystemAuthAdmin
	// 检查username
	r := adminSrv.db.Where("username = ? AND is_delete = ?", addReq.Username, 0).Limit(1).Find(&sysAdmin)
	err := r.Error
	if e = response.CheckErr(err, "Add Find by username err"); e != nil {
		return
	}
	if r.RowsAffected > 0 {
		return response.AssertArgumentError.Make("账号已存在换一个吧！")
	}
	// 检查nickname
	r = adminSrv.db.Where("nickname = ? AND is_delete = ?", addReq.Nickname, 0).Limit(1).Find(&sysAdmin)
	err = r.Error
	if e = response.CheckErr(err, "Add Find by nickname err"); e != nil {
		return
	}
	if r.RowsAffected > 0 {
		return response.AssertArgumentError.Make("昵称已存在换一个吧！")
	}
	var roleResp resp.SystemAuthRoleResp
	if roleResp, e = adminSrv.roleSrv.Detail(addReq.Role); e != nil {
		return
	}
	if roleResp.IsDisable > 0 {
		return response.AssertArgumentError.Make("当前角色已被禁用!")
	}
	passwdLen := len(addReq.Password)
	if !(passwdLen >= 6 && passwdLen <= 20) {
		return response.Failed.Make("密码必须在6~20位")
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
	err = adminSrv.db.Create(&sysAdmin).Error
	e = response.CheckErr(err, "Add Create err")
	return
}

//Edit 管理员编辑
func (adminSrv SystemAuthAdminService) Edit(c *gin.Context, editReq req.SystemAuthAdminEditReq) (e error) {
	// 检查id
	err := adminSrv.db.Where("id = ? AND is_delete = ?", editReq.ID, 0).Limit(1).First(&system.SystemAuthAdmin{}).Error
	if e = response.CheckErrDBNotRecord(err, "账号不存在了!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "Edit First err"); e != nil {
		return
	}
	// 检查username
	var admin system.SystemAuthAdmin
	r := adminSrv.db.Where("username = ? AND is_delete = ? AND id != ?", editReq.Username, 0, editReq.ID).Find(&admin)
	err = r.Error
	if e = response.CheckErr(err, "Edit Find by username err"); e != nil {
		return
	}
	if r.RowsAffected > 0 {
		return response.AssertArgumentError.Make("账号已存在换一个吧！")
	}
	// 检查nickname
	r = adminSrv.db.Where("nickname = ? AND is_delete = ? AND id != ?", editReq.Nickname, 0, editReq.ID).Find(&admin)
	err = r.Error
	if e = response.CheckErr(err, "Edit Find by nickname err"); e != nil {
		return
	}
	if r.RowsAffected > 0 {
		return response.AssertArgumentError.Make("昵称已存在换一个吧！")
	}
	// 检查role
	if editReq.Role > 0 && editReq.ID != 1 {
		if _, e = adminSrv.roleSrv.Detail(editReq.Role); e != nil {
			return
		}
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
			return response.Failed.Make("密码必须在6~20位")
		}
		salt := util.ToolsUtil.RandomString(5)
		adminMap["Salt"] = salt
		adminMap["Password"] = util.ToolsUtil.MakeMd5(strings.Trim(editReq.Password, "") + salt)
	} else {
		delete(adminMap, "Password")
	}
	err = adminSrv.db.Model(&admin).Where("id = ?", editReq.ID).Updates(adminMap).Error
	if e = response.CheckErr(err, "Edit Updates err"); e != nil {
		return
	}
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
	return
}

//Update 管理员更新
func (adminSrv SystemAuthAdminService) Update(c *gin.Context, updateReq req.SystemAuthAdminUpdateReq, adminId uint) (e error) {
	// 检查id
	var admin system.SystemAuthAdmin
	err := adminSrv.db.Where("id = ? AND is_delete = ?", adminId, 0).Limit(1).First(&admin).Error
	if e = response.CheckErrDBNotRecord(err, "账号不存在了!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "Update First err"); e != nil {
		return
	}
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
			return response.Failed.Make("当前密码不正确!")
		}
		passwdLen := len(updateReq.Password)
		if !(passwdLen >= 6 && passwdLen <= 20) {
			return response.Failed.Make("密码必须在6~20位")
		}
		salt := util.ToolsUtil.RandomString(5)
		adminMap["Salt"] = salt
		adminMap["Password"] = util.ToolsUtil.MakeMd5(strings.Trim(updateReq.Password, " ") + salt)
	} else {
		delete(adminMap, "Password")
	}
	err = adminSrv.db.Model(&admin).Updates(adminMap).Error
	if e = response.CheckErr(err, "Update Updates err"); e != nil {
		return
	}
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
	return
}

//Del 管理员删除
func (adminSrv SystemAuthAdminService) Del(c *gin.Context, id uint) (e error) {
	var admin system.SystemAuthAdmin
	err := adminSrv.db.Where("id = ? AND is_delete = ?", id, 0).Limit(1).First(&admin).Error
	if e = response.CheckErrDBNotRecord(err, "账号已不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "Del First err"); e != nil {
		return
	}
	if id == 1 {
		return response.AssertArgumentError.Make("系统管理员不允许删除!")
	}
	if id == config.AdminConfig.GetAdminId(c) {
		return response.AssertArgumentError.Make("不能删除自己!")
	}
	err = adminSrv.db.Model(&admin).Updates(system.SystemAuthAdmin{IsDelete: 1, DeleteTime: time.Now().Unix()}).Error
	e = response.CheckErr(err, "Del Updates err")
	return
}

//Disable 管理员状态切换
func (adminSrv SystemAuthAdminService) Disable(c *gin.Context, id uint) (e error) {
	var admin system.SystemAuthAdmin
	err := adminSrv.db.Where("id = ? AND is_delete = ?", id, 0).Limit(1).Find(&admin).Error
	if e = response.CheckErr(err, "Disable Find err"); e != nil {
		return
	}
	if admin.ID == 0 {
		return response.AssertArgumentError.Make("账号已不存在!")
	}
	if id == config.AdminConfig.GetAdminId(c) {
		return response.AssertArgumentError.Make("不能禁用自己!")
	}
	var isDisable uint8
	if admin.IsDisable == 0 {
		isDisable = 1
	}
	err = adminSrv.db.Model(&admin).Updates(system.SystemAuthAdmin{IsDisable: isDisable, UpdateTime: time.Now().Unix()}).Error
	e = response.CheckErr(err, "Disable Updates err")
	return
}

//CacheAdminUserByUid 缓存管理员
func (adminSrv SystemAuthAdminService) CacheAdminUserByUid(id uint) (err error) {
	var admin system.SystemAuthAdmin
	err = adminSrv.db.Where("id = ?", id).Limit(1).First(&admin).Error
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
