package generator

import (
	"bytes"
	"likeadmin/config"
	"likeadmin/core/response"
	"likeadmin/model/gen"
	"likeadmin/util"
	"path"
	"text/template"
)

var TemplateUtil = templateUtil{
	basePath: "generator/templates",
	tpl: template.New("").Delims("{{{", "}}}").Funcs(
		template.FuncMap{
			"sub": sub,
		}),
}

func sub(a, b int) int {
	return a - b
}

//TplVars 模板变量
type TplVars struct {
	GenTpl          string
	TableName       string
	AuthorName      string
	PackageName     string
	EntityName      string
	EntitySnakeName string
	ModuleName      string
	FunctionName    string
	JavaCamelField  string
	DateFields      []string
	PrimaryKey      string
	PrimaryField    string
	AllFields       []string
	SubPriCol       gen.GenTableColumn
	SubPriField     string
	SubTableFields  []string
	ListFields      []string
	DetailFields    []string
	DictFields      []string
	IsSearch        bool
	ModelOprMap     map[string]string
	Table           gen.GenTable
	Columns         []gen.GenTableColumn
	SubColumns      []gen.GenTableColumn
	//ModelTypeMap    map[string]string
}

//genUtil 模板工具
type templateUtil struct {
	basePath string
	tpl      *template.Template
}

//PrepareVars 获取模板变量信息
func (tu templateUtil) PrepareVars(table gen.GenTable, columns []gen.GenTableColumn,
	oriSubPriCol gen.GenTableColumn, oriSubCols []gen.GenTableColumn) TplVars {
	subPriField := "id"
	isSearch := false
	primaryKey := "id"
	primaryField := "id"
	functionName := "【请填写功能名称】"
	var allFields []string
	var subTableFields []string
	var listFields []string
	var detailFields []string
	var dictFields []string
	var subColumns []gen.GenTableColumn
	var oriSubColNames []string
	for _, column := range oriSubCols {
		oriSubColNames = append(oriSubColNames, column.ColumnName)
	}
	if oriSubPriCol.ID > 0 {
		subPriField = oriSubPriCol.ColumnName
		subColumns = append(subColumns, oriSubPriCol)
	}
	for _, column := range columns {
		allFields = append(allFields, column.ColumnName)
		if !util.ToolsUtil.Contains(oriSubColNames, column.ColumnName) {
			subTableFields = append(subTableFields, column.ColumnName)
			subColumns = append(subColumns, column)
		}
		if column.IsList == 1 {
			listFields = append(listFields, column.ColumnName)
		}
		if column.IsEdit == 1 {
			detailFields = append(detailFields, column.ColumnName)
		}
		if column.IsQuery == 1 {
			isSearch = true
		}
		if column.IsPk == 1 {
			primaryKey = column.JavaField
			primaryField = column.ColumnName
		}
		if column.DictType != "" && !util.ToolsUtil.Contains(dictFields, column.DictType) {
			dictFields = append(dictFields, column.DictType)
		}
	}
	//QueryType转换查询比较运算符
	modelOprMap := map[string]string{
		"=":    "==",
		"LIKE": "like",
	}
	if table.FunctionName != "" {
		functionName = table.FunctionName
	}
	return TplVars{
		GenTpl:          table.GenTpl,
		TableName:       table.TableName,
		AuthorName:      table.AuthorName,
		PackageName:     config.GenConfig.PackageName,
		EntityName:      table.EntityName,
		EntitySnakeName: util.StringUtil.ToSnakeCase(table.EntityName),
		ModuleName:      table.ModuleName,
		FunctionName:    functionName,
		DateFields:      SqlConstants.ColumnTimeName,
		PrimaryKey:      primaryKey,
		PrimaryField:    primaryField,
		AllFields:       allFields,
		SubPriCol:       oriSubPriCol,
		SubPriField:     subPriField,
		SubTableFields:  subTableFields,
		ListFields:      listFields,
		DetailFields:    detailFields,
		DictFields:      dictFields,
		IsSearch:        isSearch,
		ModelOprMap:     modelOprMap,
		Columns:         columns,
		SubColumns:      subColumns,
	}
}

//GetTemplatePaths 获取模板路径
func (tu templateUtil) GetTemplatePaths(genTpl string) []string {
	tplPaths := []string{
		"vue/api.ts.tpl",
		"vue/edit.vue.tpl",
	}
	if genTpl == GenConstants.TplCrud {
		tplPaths = append(tplPaths, "vue/index.vue.tpl")
	} else if genTpl == GenConstants.TplTree {
		tplPaths = append(tplPaths, "vue/index-tree.vue.tpl")
	}
	return tplPaths
}

//Render 渲染模板
func (tu templateUtil) Render(tplPath string, tplVars TplVars) (res string, e error) {
	tpl, err := tu.tpl.ParseFiles(path.Join(config.Config.RootPath, tu.basePath, tplPath))
	if e = response.CheckErr(err, "TemplateUtil.Render ParseFiles err"); e != nil {
		return "", e
	}
	buf := &bytes.Buffer{}
	err = tpl.ExecuteTemplate(buf, path.Base(tplPath), tplVars)
	if e = response.CheckErr(err, "TemplateUtil.Render Execute err"); e != nil {
		return "", e
	}
	return buf.String(), nil
}
