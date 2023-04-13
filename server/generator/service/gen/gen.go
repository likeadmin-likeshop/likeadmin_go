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
	Detail(id uint) (res resp.GenTableDetailResp, e error)
	ImportTable(tableNames []string) (e error)
	DelTable(ids []uint) (e error)
	//SyncTable
	//EditTable
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
	var tbResp []resp.DbTableResp
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

//Detail 生成详情
func (genSrv generateService) Detail(id uint) (res resp.GenTableDetailResp, e error) {
	var genTb gen.GenTable
	err := genSrv.db.Where("id = ?", id).Limit(1).First(&genTb).Error
	if e = response.CheckErrDBNotRecord(err, "查询的数据不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "Detail Find err"); e != nil {
		return
	}
	var columns []gen.GenTableColumn
	err = genSrv.db.Where("table_id = ?", id).Order("sort").Find(&columns).Error
	if e = response.CheckErr(err, "Detail Find err"); e != nil {
		return
	}
	var base resp.GenTableBaseResp
	response.Copy(&base, genTb)
	var gen resp.GenTableGenResp
	response.Copy(&gen, genTb)
	var colResp []resp.GenColumnResp
	response.Copy(&colResp, columns)
	return resp.GenTableDetailResp{
		Base:    base,
		Gen:     gen,
		Columns: colResp,
	}, e
}

//ImportTable 导入表结构
func (genSrv generateService) ImportTable(tableNames []string) (e error) {
	var dbTbs []resp.DbTableResp
	err := generator.GenUtil.GetDbTablesQueryByNames(genSrv.db, tableNames).Find(&dbTbs).Error
	if e = response.CheckErr(err, "ImportTable Find tables err"); e != nil {
		return
	}
	var tables []gen.GenTable
	response.Copy(&tables, dbTbs)
	if len(tables) == 0 {
		e = response.AssertArgumentError.Make("表不存在!")
		return
	}
	err = genSrv.db.Transaction(func(tx *gorm.DB) error {
		for i := 0; i < len(tables); i++ {
			//生成表信息
			genTable := generator.GenUtil.InitTable(tables[i])
			txErr := tx.Create(&genTable).Error
			if te := response.CheckErr(txErr, "ImportTable Create table err"); te != nil {
				return te
			}
			// 生成列信息
			var columns []gen.GenTableColumn
			txErr = generator.GenUtil.GetDbTableColumnsQueryByName(genSrv.db, tables[i].TableName).Find(&columns).Error
			if te := response.CheckErr(txErr, "ImportTable Find columns err"); te != nil {
				return te
			}
			for j := 0; j < len(columns); j++ {
				column := generator.GenUtil.InitColumn(genTable.ID, columns[j])
				txErr = tx.Create(&column).Error
				if te := response.CheckErr(txErr, "ImportTable Create column err"); te != nil {
					return te
				}
			}
		}
		return nil
	})
	e = response.CheckErr(err, "ImportTable Transaction err")
	return nil
}

//DelTable 删除表结构
func (genSrv generateService) DelTable(ids []uint) (e error) {
	err := genSrv.db.Transaction(func(tx *gorm.DB) error {
		txErr := tx.Delete(&gen.GenTable{}, "id in ?", ids).Error
		if te := response.CheckErr(txErr, "DelTable Delete GenTable err"); te != nil {
			return te
		}
		txErr = tx.Delete(&gen.GenTableColumn{}, "table_id in ?", ids).Error
		if te := response.CheckErr(txErr, "DelTable Delete GenTableColumn err"); te != nil {
			return te
		}
		return nil
	})
	e = response.CheckErr(err, "DelTable Transaction err")
	return
}
