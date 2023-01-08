package system

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"likeadmin/admin/schemas/req"
	"likeadmin/admin/schemas/resp"
	"likeadmin/config"
	"likeadmin/core"
	"likeadmin/core/response"
	"likeadmin/model/system"
	"likeadmin/util"
	"runtime/debug"
	"strconv"
	"time"
)

var SystemLoginService = systemLoginService{}

//systemLoginService 系统登录服务实现类
type systemLoginService struct{}

//Login 登录
func (loginSrv systemLoginService) Login(c *gin.Context, req *req.SystemLoginReq) resp.SystemLoginResp {
	sysAdmin, err := SystemAuthAdminService.FindByUsername(req.Username)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		loginSrv.RecordLoginLog(c, 0, req.Username, response.LoginAccountError.Msg())
		panic(response.LoginAccountError)
	} else if err != nil {
		core.Logger.Errorf("Login FindByUsername err: err=[%+v]", err)
		loginSrv.RecordLoginLog(c, 0, req.Username, response.Failed.Msg())
		panic(response.Failed)
	}
	if sysAdmin.IsDelete == 1 {
		loginSrv.RecordLoginLog(c, 0, req.Username, response.LoginAccountError.Msg())
		panic(response.LoginAccountError)
	}
	if sysAdmin.IsDisable == 1 {
		loginSrv.RecordLoginLog(c, sysAdmin.ID, req.Username, response.LoginDisableError.Msg())
		panic(response.LoginDisableError)
	}
	md5Pwd := util.ToolsUtil.MakeMd5(req.Password + sysAdmin.Salt)
	if sysAdmin.Password != md5Pwd {
		loginSrv.RecordLoginLog(c, sysAdmin.ID, req.Username, response.LoginAccountError.Msg())
		panic(response.LoginAccountError)
	}
	defer func() {
		if r := recover(); r != nil {
			switch r.(type) {
			// 自定义类型
			case response.RespType:
				panic(r)
			// 其他类型
			default:
				core.Logger.Errorf("stacktrace from panic: %+v\n%s", r, string(debug.Stack()))
				loginSrv.RecordLoginLog(c, sysAdmin.ID, req.Username, response.Failed.Msg())
				panic(response.Failed)
			}
		}
	}()
	token := util.ToolsUtil.MakeToken()
	adminIdStr := strconv.FormatUint(uint64(sysAdmin.ID), 10)

	//非多处登录
	if sysAdmin.IsMultipoint == 0 {
		sysAdminSetKey := config.AdminConfig.BackstageTokenSet + adminIdStr
		ts := util.RedisUtil.SGet(sysAdminSetKey)
		if len(ts) > 0 {
			var keys []string
			for _, t := range ts {
				keys = append(keys, t)
			}
			util.RedisUtil.Del(keys...)
		}
		util.RedisUtil.Del(sysAdminSetKey)
		util.RedisUtil.SSet(sysAdminSetKey, token)
	}

	// 缓存登录信息
	util.RedisUtil.Set(config.AdminConfig.BackstageTokenKey+token, adminIdStr, 7200)
	SystemAuthAdminService.CacheAdminUserByUid(sysAdmin.ID)

	// 更新登录信息
	err = core.DB.Model(&sysAdmin).Updates(
		system.SystemAuthAdmin{LastLoginIp: c.ClientIP(), LastLoginTime: time.Now().Unix()}).Error
	if err != nil {
		loginSrv.RecordLoginLog(c, sysAdmin.ID, req.Username, response.SystemError.Msg())
		util.CheckUtil.CheckErr(err, "Login Updates err")
	}
	// 记录登录日志
	loginSrv.RecordLoginLog(c, sysAdmin.ID, req.Username, "")
	// 返回登录信息
	return resp.SystemLoginResp{Token: token}
}

//Logout 退出
func (loginSrv systemLoginService) Logout(req *req.SystemLogoutReq) {
	util.RedisUtil.Del(config.AdminConfig.BackstageTokenKey + req.Token)
}

//RecordLoginLog 记录登录日志
func (loginSrv systemLoginService) RecordLoginLog(c *gin.Context, adminId uint, username string, errStr string) {
	ua := core.UAParser.Parse(c.GetHeader("user-agent"))
	var status uint8
	if errStr == "" {
		status = 1
	}
	err := core.DB.Create(&system.SystemLogLogin{
		AdminId: adminId, Username: username, Ip: c.ClientIP(), Os: ua.Os.Family,
		Browser: ua.UserAgent.Family, Status: status}).Error
	util.CheckUtil.CheckErr(err, "RecordLoginLog Create err")
}
