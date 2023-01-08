package system

import (
	"fmt"
	"likeadmin/admin/schemas/req"
	"likeadmin/admin/schemas/resp"
	"likeadmin/core"
	"likeadmin/core/request"
	"likeadmin/core/response"
	"likeadmin/model/system"
)

var SystemLogsServer = systemLogsServer{}

//systemAuthMenuService 系统日志服务实现类
type systemLogsServer struct{}

//Operate 系统操作日志
func (logSrv systemLogsServer) Operate(page request.PageReq, logReq req.SystemLogOperateReq) response.PageResp {
	// 分页信息
	var res response.PageResp
	response.Copy(&res, page)
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)
	// 查询
	logTbName := core.DBTableName(&system.SystemLogOperate{})
	adminTbName := core.DBTableName(&system.SystemAuthAdmin{})
	logModel := core.DB.Table(logTbName + " AS log").Joins(
		fmt.Sprintf("LEFT JOIN %s AS admin ON log.admin_id = admin.id", adminTbName)).Select(
		"log.*, admin.username, admin.nickname")
	// 条件
	if logReq.Title != "" {
		logModel = logModel.Where("title like ?", "%"+logReq.Title+"%")
	}
	if logReq.Username != "" {
		logModel = logModel.Where("username like ?", "%"+logReq.Username+"%")
	}
	if logReq.Ip != "" {
		logModel = logModel.Where("ip like ?", "%"+logReq.Ip+"%")
	}
	if logReq.Type != "" {
		logModel = logModel.Where("type = ?", logReq.Type)
	}
	if logReq.Status > 0 {
		logModel = logModel.Where("status = ?", logReq.Status)
	}
	if logReq.Url != "" {
		logModel = logModel.Where("url = ?", logReq.Url)
	}
	if !logReq.StartTime.IsZero() {
		logModel = logModel.Where("log.create_time >= ?", logReq.StartTime.Unix())
	}
	if !logReq.EndTime.IsZero() {
		logModel = logModel.Where("log.create_time <= ?", logReq.EndTime.Unix())
	}
	// 总数
	var count int64
	if err := logModel.Count(&count).Error; err != nil {
		core.Logger.Errorf("Operate Count err: err=[%+v]", err)
		panic(response.SystemError)
	}
	// 数据
	var logResp []resp.SystemLogOperateResp
	if err := logModel.Limit(limit).Offset(offset).Order("id desc").Find(&logResp).Error; err != nil {
		core.Logger.Errorf("Operate Find err: err=[%+v]", err)
		panic(response.SystemError)
	}
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Lists:    logResp,
	}
}

//Login 系统登录日志
func (logSrv systemLogsServer) Login(page request.PageReq, logReq req.SystemLogLoginReq) response.PageResp {
	// 分页信息
	var res response.PageResp
	response.Copy(&res, page)
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)
	// 查询
	logModel := core.DB.Model(&system.SystemLogLogin{})
	// 条件
	if logReq.Username != "" {
		logModel = logModel.Where("username like ?", "%"+logReq.Username+"%")
	}
	if logReq.Status > 0 {
		logModel = logModel.Where("status = ?", logReq.Status)
	}
	if !logReq.StartTime.IsZero() {
		logModel = logModel.Where("create_time >= ?", logReq.StartTime.Unix())
	}
	if !logReq.EndTime.IsZero() {
		logModel = logModel.Where("create_time <= ?", logReq.EndTime.Unix())
	}
	// 总数
	var count int64
	if err := logModel.Count(&count).Error; err != nil {
		core.Logger.Errorf("Login Count err: err=[%+v]", err)
		panic(response.SystemError)
	}
	// 数据
	var logResp []resp.SystemLogLoginResp
	if err := logModel.Limit(limit).Offset(offset).Order("id desc").Find(&logResp).Error; err != nil {
		core.Logger.Errorf("Login Find err: err=[%+v]", err)
		panic(response.SystemError)
	}
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Lists:    logResp,
	}
}
