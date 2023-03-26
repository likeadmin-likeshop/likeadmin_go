package gen

import (
	"gorm.io/gorm"
	"likeadmin/core/request"
	"likeadmin/core/response"
	"likeadmin/generator"
	"likeadmin/generator/schemas/req"
	"likeadmin/generator/schemas/resp"
)

type IGenerateService interface {
	DbTables(page request.PageReq, req req.DbTablesReq) (res response.PageResp, e error)
	//List
	//Detail
	//ImportTable
	//SyncTable
	//EditTable
	//DelTable
	//PreviewCode
	//DownloadCode
	//GenCode
	//GenZipCode
}

//NewGenerateService 初始化
func NewGenerateService(db *gorm.DB) IGenerateService {
	return &generateService{db: db}
}

//GenerateService 代码生成服务实现类
type generateService struct {
	db *gorm.DB
}

//DbTables 库表列表
func (genSrv generateService) DbTables(page request.PageReq, req req.DbTablesReq) (res response.PageResp, e error) {
	// 分页信息
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)
	tbModel := generator.GenUtil.GetDbTablesQuery(genSrv.db, req.TableName, req.TableComment)
	// 总数
	var count int64
	err := tbModel.Count(&count).Error
	if e = response.CheckErr(err, "DbTables Count err"); e != nil {
		return
	}
	// 数据
	var tbResp []resp.DbTablesResp
	err = tbModel.Limit(limit).Offset(offset).Find(&tbResp).Error
	if e = response.CheckErr(err, "DbTables Find err"); e != nil {
		return
	}
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Lists:    tbResp,
	}, nil
}
