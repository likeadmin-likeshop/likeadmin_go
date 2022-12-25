package middleware

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"likeadmin/admin/service/system"
	"likeadmin/config"
	"likeadmin/core"
	"likeadmin/core/response"
	sysModel "likeadmin/models/system"
	"likeadmin/utils"
	"strconv"
	"strings"
)

//TokenAuth Token认证中间件
func TokenAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 路由转权限
		auths := strings.ReplaceAll(strings.Replace(c.Request.URL.Path, "/api/", "", 1), "/", ":")

		// 免登录接口
		if utils.ToolsUtil.Contains(config.AdminConfig.NotLoginUri, auths) {
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
		existCnt := utils.RedisUtil.Exists(token)
		if existCnt == 0 {
			response.Fail(c, response.TokenInvalid)
			c.Abort()
			return
		}

		// 用户信息缓存
		uidStr := utils.RedisUtil.Get(token)
		var uid uint
		if uidStr != "" {
			i, err := strconv.Atoi(uidStr)
			if err != nil {
				core.Logger.Errorf("TokenAuth Atoi uidStr err: err=[%+v]", err)
				response.Fail(c, response.TokenInvalid)
				c.Abort()
				return
			}
			uid = uint(i)
		}
		if !utils.RedisUtil.HExists(config.AdminConfig.BackstageManageKey, uidStr) {
			err := system.SystemAuthAdminService.CacheAdminUserByUid(uid)
			if err != nil {
				core.Logger.Errorf("TokenAuth CacheAdminUserByUid err: err=[%+v]", err)
				response.Fail(c, response.SystemError)
				c.Abort()
				return
			}
		}

		// 校验用户被删除
		var mapping sysModel.SystemAuthAdmin
		err := json.Unmarshal([]byte(utils.RedisUtil.HGet(config.AdminConfig.BackstageManageKey, uidStr)), &mapping)
		if err != nil {
			core.Logger.Errorf("TokenAuth Unmarshal err: err=[%+v]", err)
			response.Fail(c, response.SystemError)
			c.Abort()
			return
		}
		if mapping.IsDelete == 1 {
			utils.RedisUtil.Del(token)
			utils.RedisUtil.HDel(config.AdminConfig.BackstageManageKey + uidStr)
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
		if utils.RedisUtil.TTL(token) < 1800 {
			utils.RedisUtil.Expire(token, 7200)
		}

		// 单次请求信息保存
		c.Set(config.AdminConfig.ReqAdminIdKey, uid)
		c.Set(config.AdminConfig.ReqRoleIdKey, mapping.Role)
		c.Set(config.AdminConfig.ReqUsernameKey, mapping.Username)
		c.Set(config.AdminConfig.ReqNicknameKey, mapping.Nickname)

		// 免权限验证接口
		if utils.ToolsUtil.Contains(config.AdminConfig.NotAuthUri, auths) || uid == 1 {
			c.Next()
			return
		}

		// 校验角色权限是否存在
		roleId := mapping.Role
		if utils.RedisUtil.HExists(config.AdminConfig.BackstageRolesKey, roleId) {
			i, err := strconv.Atoi(roleId)
			if err != nil {
				core.Logger.Errorf("TokenAuth Atoi roleId err: err=[%+v]", err)
				response.Fail(c, response.SystemError)
				c.Abort()
				return
			}
			err = system.SystemAuthPermService.CacheRoleMenusByRoleId(uint(i))
			if err != nil {
				core.Logger.Errorf("TokenAuth CacheRoleMenusByRoleId err: err=[%+v]", err)
				response.Fail(c, response.SystemError)
				c.Abort()
				return
			}
		}

		// 验证是否有权限操作
		menus := utils.RedisUtil.HGet(config.AdminConfig.BackstageRolesKey, roleId)
		if !(menus != "" && utils.ToolsUtil.Contains(strings.Split(menus, ","), auths)) {
			response.Fail(c, response.NoPermission)
			c.Abort()
			return
		}
	}
}
