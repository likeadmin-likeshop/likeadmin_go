package gen

import (
	"archive/zip"
	"bytes"
	"gorm.io/gorm"
	"likeadmin/config"
	"likeadmin/core/request"
	"likeadmin/core/response"
	"likeadmin/generator"
	"likeadmin/generator/schemas/req"
	"likeadmin/generator/schemas/resp"
	"likeadmin/model/gen"
	"likeadmin/util"
	"strings"
)

type IGenerateService interface {
	DbTables(page request.PageReq, req req.DbTablesReq) (res response.PageResp, e error)
	List(page request.PageReq, listReq req.ListTableReq) (res response.PageResp, e error)
	Detail(id uint) (res resp.GenTableDetailResp, e error)
	ImportTable(tableNames []string) (e error)
	SyncTable(id uint) (e error)
	EditTable(editReq req.EditTableReq) (e error)
	DelTable(ids []uint) (e error)
	PreviewCode(id uint) (res map[string]string, e error)
	GenCode(tableName string) (e error)
	DownloadCode(tableNames []string) ([]byte, error)
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
		Base:   base,
		Gen:    gen,
		Column: colResp,
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

//SyncTable 同步表结构
func (genSrv generateService) SyncTable(id uint) (e error) {
	//旧数据
	var genTable gen.GenTable
	err := genSrv.db.Where("id = ?", id).Limit(1).First(&genTable).Error
	if e = response.CheckErrDBNotRecord(err, "生成数据不存在！"); e != nil {
		return
	}
	if e = response.CheckErr(err, "SyncTable First err"); e != nil {
		return
	}
	var genTableCols []gen.GenTableColumn
	err = genSrv.db.Where("table_id = ?", id).Order("sort").Find(&genTableCols).Error
	if e = response.CheckErr(err, "SyncTable Find err"); e != nil {
		return
	}
	if len(genTableCols) <= 0 {
		e = response.AssertArgumentError.Make("旧数据异常！")
		return
	}
	prevColMap := make(map[string]gen.GenTableColumn)
	for i := 0; i < len(genTableCols); i++ {
		prevColMap[genTableCols[i].ColumnName] = genTableCols[i]
	}
	//新数据
	var columns []gen.GenTableColumn
	err = generator.GenUtil.GetDbTableColumnsQueryByName(genSrv.db, genTable.TableName).Find(&columns).Error
	if e = response.CheckErr(err, "SyncTable Find new err"); e != nil {
		return
	}
	if len(columns) <= 0 {
		e = response.AssertArgumentError.Make("同步结构失败,原表结构不存在！")
		return
	}
	//事务处理
	err = genSrv.db.Transaction(func(tx *gorm.DB) error {
		//处理新增和更新
		for i := 0; i < len(columns); i++ {
			col := generator.GenUtil.InitColumn(id, columns[i])
			if prevCol, ok := prevColMap[columns[i].ColumnName]; ok {
				//更新
				col.ID = prevCol.ID
				if col.IsList == 0 {
					col.DictType = prevCol.DictType
					col.QueryType = prevCol.QueryType
				}
				if prevCol.IsRequired == 1 && prevCol.IsPk == 0 && prevCol.IsInsert == 1 || prevCol.IsEdit == 1 {
					col.HtmlType = prevCol.HtmlType
					col.IsRequired = prevCol.IsRequired
				}
				txErr := tx.Save(&col).Error
				if te := response.CheckErr(txErr, "SyncTable Save err"); te != nil {
					return te
				}
			} else {
				//新增
				txErr := tx.Create(&col).Error
				if te := response.CheckErr(txErr, "SyncTable Create err"); te != nil {
					return te
				}
			}
		}
		//处理删除
		colNames := make([]string, len(columns))
		for i := 0; i < len(columns); i++ {
			colNames[i] = columns[i].ColumnName
		}
		delColIds := make([]uint, 0)
		for _, prevCol := range prevColMap {
			if !util.ToolsUtil.Contains(colNames, prevCol.ColumnName) {
				delColIds = append(delColIds, prevCol.ID)
			}
		}
		txErr := tx.Delete(&gen.GenTableColumn{}, "id in ?", delColIds).Error
		if te := response.CheckErr(txErr, "SyncTable Delete err"); te != nil {
			return te
		}
		return nil
	})
	e = response.CheckErr(err, "SyncTable Transaction err")
	return nil
}

//EditTable 编辑表结构
func (genSrv generateService) EditTable(editReq req.EditTableReq) (e error) {
	if editReq.GenTpl == generator.GenConstants.TplTree {
		if editReq.TreePrimary == "" {
			e = response.AssertArgumentError.Make("树主ID不能为空！")
			return
		}
		if editReq.TreeParent == "" {
			e = response.AssertArgumentError.Make("树父ID不能为空！")
			return
		}
	}
	var genTable gen.GenTable
	err := genSrv.db.Where("id = ?", editReq.ID).Limit(1).First(&genTable).Error
	if e = response.CheckErrDBNotRecord(err, "数据已丢失！"); e != nil {
		return
	}
	if e = response.CheckErr(err, "EditTable First err"); e != nil {
		return
	}
	response.Copy(&genTable, editReq)
	err = genSrv.db.Transaction(func(tx *gorm.DB) error {
		genTable.SubTableName = strings.Replace(editReq.SubTableName, config.Config.DbTablePrefix, "", 1)
		txErr := tx.Save(&genTable).Error
		if te := response.CheckErr(txErr, "EditTable Save GenTable err"); te != nil {
			return te
		}
		for i := 0; i < len(editReq.Columns); i++ {
			var col gen.GenTableColumn
			response.Copy(&col, editReq.Columns[i])
			txErr = tx.Save(&col).Error
			if te := response.CheckErr(txErr, "EditTable Save GenTableColumn err"); te != nil {
				return te
			}
		}
		return nil
	})
	e = response.CheckErr(err, "EditTable Transaction err")
	return
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

//getSubTableInfo 根据主表获取子表主键和列信息
func (genSrv generateService) getSubTableInfo(genTable gen.GenTable) (pkCol gen.GenTableColumn, cols []gen.GenTableColumn, e error) {
	if genTable.SubTableName == "" || genTable.SubTableFk == "" {
		return
	}
	var table gen.GenTable
	err := genSrv.db.Where("table_name = ?", genTable.SubTableName).Limit(1).First(&table).Error
	if e = response.CheckErrDBNotRecord(err, "子表记录丢失！"); e != nil {
		return
	}
	if e = response.CheckErr(err, "getSubTableInfo First err"); e != nil {
		return
	}
	err = generator.GenUtil.GetDbTableColumnsQueryByName(genSrv.db, genTable.SubTableName).Find(&cols).Error
	if e = response.CheckErr(err, "getSubTableInfo Find err"); e != nil {
		return
	}
	pkCol = generator.GenUtil.InitColumn(table.ID, generator.GenUtil.GetTablePriCol(cols))
	return
}

//renderCodeByTable 根据主表和模板文件渲染模板代码
func (genSrv generateService) renderCodeByTable(genTable gen.GenTable) (res map[string]string, e error) {
	var columns []gen.GenTableColumn
	err := genSrv.db.Where("table_id = ?", genTable.ID).Order("sort").Find(&columns).Error
	if e = response.CheckErr(err, "renderCodeByTable Find err"); e != nil {
		return
	}
	//获取子表信息
	pkCol, cols, err := genSrv.getSubTableInfo(genTable)
	if e = response.CheckErr(err, "renderCodeByTable getSubTableInfo err"); e != nil {
		return
	}
	//获取模板变量信息
	vars := generator.TemplateUtil.PrepareVars(genTable, columns, pkCol, cols)
	//生成模板内容
	res = make(map[string]string)
	for _, tplPath := range generator.TemplateUtil.GetTemplatePaths(genTable.GenTpl) {
		res[tplPath], err = generator.TemplateUtil.Render(tplPath, vars)
		if e = response.CheckErr(err, "renderCodeByTable Render err"); e != nil {
			return
		}
	}
	return
}

//PreviewCode 预览代码
func (genSrv generateService) PreviewCode(id uint) (res map[string]string, e error) {
	var genTable gen.GenTable
	err := genSrv.db.Where("id = ?", id).Limit(1).First(&genTable).Error
	if e = response.CheckErrDBNotRecord(err, "记录丢失！"); e != nil {
		return
	}
	if e = response.CheckErr(err, "PreviewCode First err"); e != nil {
		return
	}
	//获取模板内容
	tplCodeMap, err := genSrv.renderCodeByTable(genTable)
	if e = response.CheckErr(err, "PreviewCode renderCodeByTable err"); e != nil {
		return
	}
	res = make(map[string]string)
	for tplPath, tplCode := range tplCodeMap {
		res[strings.ReplaceAll(tplPath, ".tpl", "")] = tplCode
	}
	return
}

//GenCode 生成代码 (自定义路径)
func (genSrv generateService) GenCode(tableName string) (e error) {
	var genTable gen.GenTable
	err := genSrv.db.Where("table_name = ?", tableName).Order("id desc").Limit(1).First(&genTable).Error
	if e = response.CheckErrDBNotRecord(err, "记录丢失！"); e != nil {
		return
	}
	if e = response.CheckErr(err, "GenCode First err"); e != nil {
		return
	}
	//获取模板内容
	tplCodeMap, err := genSrv.renderCodeByTable(genTable)
	if e = response.CheckErr(err, "GenCode renderCodeByTable err"); e != nil {
		return
	}
	//获取生成根路径
	basePath := generator.TemplateUtil.GetGenPath(genTable)
	//生成代码文件
	err = generator.TemplateUtil.GenCodeFiles(tplCodeMap, genTable.ModuleName, basePath)
	if e = response.CheckErr(err, "GenCode GenCodeFiles err"); e != nil {
		return
	}
	return
}

//genZipCode 生成代码 (压缩包下载)
func (genSrv generateService) genZipCode(zipWriter *zip.Writer, tableName string) (e error) {
	var genTable gen.GenTable
	err := genSrv.db.Where("table_name = ?", tableName).Order("id desc").Limit(1).First(&genTable).Error
	if e = response.CheckErrDBNotRecord(err, "记录丢失！"); e != nil {
		return
	}
	if e = response.CheckErr(err, "genZipCode First err"); e != nil {
		return
	}
	//获取模板内容
	tplCodeMap, err := genSrv.renderCodeByTable(genTable)
	if e = response.CheckErr(err, "genZipCode renderCodeByTable err"); e != nil {
		return
	}
	//压缩文件
	err = generator.TemplateUtil.GenZip(zipWriter, tplCodeMap, genTable.ModuleName)
	if e = response.CheckErr(err, "genZipCode GenZip err"); e != nil {
		return
	}
	return
}

//DownloadCode 下载代码
func (genSrv generateService) DownloadCode(tableNames []string) ([]byte, error) {
	buf := new(bytes.Buffer)
	zipWriter := zip.NewWriter(buf)
	for _, tableName := range tableNames {
		err := genSrv.genZipCode(zipWriter, tableName)
		if err != nil {
			return nil, response.CheckErr(err, "DownloadCode genZipCode for %s err", tableName)
		}
	}
	err := zipWriter.Close()
	if err != nil {
		return nil, response.CheckErr(err, "DownloadCode zipWriter.Close err")
	}
	return buf.Bytes(), nil
}
