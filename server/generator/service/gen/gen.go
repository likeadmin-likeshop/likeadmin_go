package gen

import (
	"gorm.io/gorm"
	"likeadmin/core/request"
	"likeadmin/core/response"
	"likeadmin/generator"
	"likeadmin/generator/schemas/req"
	"likeadmin/generator/schemas/resp"
	"likeadmin/model/gen"
)

type IGenerateService interface {
	DbTables(page request.PageReq, req req.DbTablesReq) (res response.PageResp, e error)
	List(page request.PageReq, listReq req.ListTableReq) (res response.PageResp, e error)
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
func (genSrv generateService) DbTables(page request.PageReq, dbReq req.DbTablesReq) (res response.PageResp, e error) {
	// 分页信息
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)
	tbModel := generator.GenUtil.GetDbTablesQuery(genSrv.db, dbReq.TableName, dbReq.TableComment)
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

//List 生成列表
func (genSrv generateService) List(page request.PageReq, listReq req.ListTableReq) (res response.PageResp, e error) {
	// 分页信息
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)
	genModel := genSrv.db.Model(&gen.GenTable{})
	if listReq.TableName != "" {
		genModel = genModel.Where("table_name like ?", "%"+listReq.TableName+"%")
	}
	if listReq.TableComment != "" {
		genModel = genModel.Where("table_comment like ?", "%"+listReq.TableComment+"%")
	}
	if !listReq.StartTime.IsZero() {
		genModel = genModel.Where("create_time >= ?", listReq.StartTime.Unix())
	}
	if !listReq.EndTime.IsZero() {
		genModel = genModel.Where("create_time <= ?", listReq.EndTime.Unix())
	}
	// 总数
	var count int64
	err := genModel.Count(&count).Error
	if e = response.CheckErr(err, "List Count err"); e != nil {
		return
	}
	// 数据
	var genResp []resp.GenTableResp
	err = genModel.Limit(limit).Offset(offset).Find(&genResp).Error
	if e = response.CheckErr(err, "List Find err"); e != nil {
		return
	}
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Lists:    genResp,
	}, nil
}
