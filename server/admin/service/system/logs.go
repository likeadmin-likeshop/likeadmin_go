package system

import (
	"fmt"
	"gorm.io/gorm"
	"likeadmin/admin/schemas/req"
	"likeadmin/admin/schemas/resp"
	"likeadmin/core"
	"likeadmin/core/request"
	"likeadmin/core/response"
	"likeadmin/model/system"
)

//NewSystemLogsServer 初始化
func NewSystemLogsServer(db *gorm.DB) *SystemLogsServer {
	return &SystemLogsServer{db: db}
}

//SystemLogsServer 系统日志服务实现类
type SystemLogsServer struct {
	db *gorm.DB
}

//Operate 系统操作日志
func (logSrv SystemLogsServer) Operate(page request.PageReq, logReq req.SystemLogOperateReq) (res response.PageResp, e error) {
	// 分页信息
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)
	// 查询
	logTbName := core.DBTableName(&system.SystemLogOperate{})
	adminTbName := core.DBTableName(&system.SystemAuthAdmin{})
	logModel := logSrv.db.Table(logTbName + " AS log").Joins(
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
	err := logModel.Count(&count).Error
	if e = response.CheckErr(err, "Operate Count err"); e != nil {
		return
	}
	// 数据
	var logResp []resp.SystemLogOperateResp
	err = logModel.Limit(limit).Offset(offset).Order("id desc").Find(&logResp).Error
	if e = response.CheckErr(err, "Operate Find err"); e != nil {
		return
	}
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Lists:    logResp,
	}, nil
}

//Login 系统登录日志
func (logSrv SystemLogsServer) Login(page request.PageReq, logReq req.SystemLogLoginReq) (res response.PageResp, e error) {
	// 分页信息
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)
	// 查询
	logModel := logSrv.db.Model(&system.SystemLogLogin{})
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
	err := logModel.Count(&count).Error
	if e = response.CheckErr(err, "Login Count err"); e != nil {
		return
	}
	// 数据
	var logResp []resp.SystemLogLoginResp
	err = logModel.Limit(limit).Offset(offset).Order("id desc").Find(&logResp).Error
	if e = response.CheckErr(err, "Login Find err"); e != nil {
		return
	}
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Lists:    logResp,
	}, nil
}
