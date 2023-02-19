package middleware

import (
	"github.com/gin-gonic/gin"
	"likeadmin/admin/service/system"
	"likeadmin/config"
	"likeadmin/core"
	"likeadmin/core/response"
	sysModel "likeadmin/model/system"
	"likeadmin/util"
	"strconv"
	"strings"
)

//TokenAuth Token认证中间件
func TokenAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 路由转权限
		auths := strings.ReplaceAll(strings.Replace(c.Request.URL.Path, "/api/", "", 1), "/", ":")

		// 免登录接口
		if util.ToolsUtil.Contains(config.AdminConfig.NotLoginUri, auths) {
			c.Next()
			return
		}

		// Token是否为空
		token := c.Request.Header.Get("token")
		if token == "" {
			response.Fail(c, response.TokenEmpty)
			c.Abort()
			return
		}

		// Token是否过期
		token = config.AdminConfig.BackstageTokenKey + token
		existCnt := util.RedisUtil.Exists(token)
		if existCnt < 0 {
			response.Fail(c, response.SystemError)
			c.Abort()
			return
		} else if existCnt == 0 {
			response.Fail(c, response.TokenInvalid)
			c.Abort()
			return
		}

		// 用户信息缓存
		uidStr := util.RedisUtil.Get(token)
		var uid uint
		if uidStr != "" {
			i, err := strconv.ParseUint(uidStr, 10, 32)
			if err != nil {
				core.Logger.Errorf("TokenAuth Atoi uidStr err: err=[%+v]", err)
				response.Fail(c, response.TokenInvalid)
				c.Abort()
				return
			}
			uid = uint(i)
		}
		permSrv := system.NewSystemAuthPermService(core.DB)
		roleSrv := system.NewSystemAuthRoleService(core.DB, permSrv)
		adminSrv := system.NewSystemAuthAdminService(core.DB, permSrv, roleSrv)
		if !util.RedisUtil.HExists(config.AdminConfig.BackstageManageKey, uidStr) {
			err := adminSrv.CacheAdminUserByUid(uid)
			if err != nil {
				core.Logger.Errorf("TokenAuth CacheAdminUserByUid err: err=[%+v]", err)
				response.Fail(c, response.SystemError)
				c.Abort()
				return
			}
		}

		// 校验用户被删除
		var mapping sysModel.SystemAuthAdmin
		err := util.ToolsUtil.JsonToObj(util.RedisUtil.HGet(config.AdminConfig.BackstageManageKey, uidStr), &mapping)
		if err != nil {
			core.Logger.Errorf("TokenAuth Unmarshal err: err=[%+v]", err)
			response.Fail(c, response.SystemError)
			c.Abort()
			return
		}
		if mapping.IsDelete == 1 {
			util.RedisUtil.Del(token)
			util.RedisUtil.HDel(config.AdminConfig.BackstageManageKey + uidStr)
			response.Fail(c, response.TokenInvalid)
			c.Abort()
			return
		}

		// 校验用户被禁用
		if mapping.IsDisable == 1 {
			response.Fail(c, response.LoginDisableError)
			c.Abort()
			return
		}

		// 令牌剩余30分钟自动续签
		if util.RedisUtil.TTL(token) < 1800 {
			util.RedisUtil.Expire(token, 7200)
		}

		// 单次请求信息保存
		c.Set(config.AdminConfig.ReqAdminIdKey, uid)
		c.Set(config.AdminConfig.ReqRoleIdKey, mapping.Role)
		c.Set(config.AdminConfig.ReqUsernameKey, mapping.Username)
		c.Set(config.AdminConfig.ReqNicknameKey, mapping.Nickname)

		// 免权限验证接口
		if util.ToolsUtil.Contains(config.AdminConfig.NotAuthUri, auths) || uid == 1 {
			c.Next()
			return
		}

		// 校验角色权限是否存在
		roleId := mapping.Role
		if util.RedisUtil.HExists(config.AdminConfig.BackstageRolesKey, roleId) {
			i, err := strconv.ParseUint(roleId, 10, 32)
			if err != nil {
				core.Logger.Errorf("TokenAuth Atoi roleId err: err=[%+v]", err)
				response.Fail(c, response.SystemError)
				c.Abort()
				return
			}
			err = permSrv.CacheRoleMenusByRoleId(uint(i))
			if err != nil {
				core.Logger.Errorf("TokenAuth CacheRoleMenusByRoleId err: err=[%+v]", err)
				response.Fail(c, response.SystemError)
				c.Abort()
				return
			}
		}

		// 验证是否有权限操作
		menus := util.RedisUtil.HGet(config.AdminConfig.BackstageRolesKey, roleId)
		if !(menus != "" && util.ToolsUtil.Contains(strings.Split(menus, ","), auths)) {
			response.Fail(c, response.NoPermission)
			c.Abort()
			return
		}
	}
}
