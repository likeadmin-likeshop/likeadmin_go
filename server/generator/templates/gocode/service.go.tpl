package {{{ .PackageName }}}

import (
	"gorm.io/gorm"
	"likeadmin/core/request"
	"likeadmin/core/response"
	"likeadmin/util"
)

type I{{{ title (toCamelCase .EntityName) }}}Service interface {
	List(page request.PageReq, listReq {{{ title (toCamelCase .EntityName) }}}ListReq) (res response.PageResp, e error)
	Detail(id uint) (res {{{ title (toCamelCase .EntityName) }}}Resp, e error)
	Add(addReq {{{ title (toCamelCase .EntityName) }}}AddReq) (e error)
	Edit(editReq {{{ title (toCamelCase .EntityName) }}}EditReq) (e error)
	Del(id uint) (e error)
}

//New{{{ title (toCamelCase .EntityName) }}}Service 初始化
func New{{{ title (toCamelCase .EntityName) }}}Service(db *gorm.DB) I{{{ title (toCamelCase .EntityName) }}}Service {
	return &{{{ toCamelCase .EntityName }}}Service{db: db}
}

//{{{ toCamelCase .EntityName }}}Service {{{ .FunctionName }}}服务实现类
type {{{ toCamelCase .EntityName }}}Service struct {
	db *gorm.DB
}

//List {{{ .FunctionName }}}列表
func (srv {{{ toCamelCase .EntityName }}}Service) List(page request.PageReq, listReq {{{ title (toCamelCase .EntityName) }}}ListReq) (res response.PageResp, e error) {
	// 分页信息
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)
	// 查询
	model := srv.db.Model(&{{{ title (toCamelCase .EntityName) }}}{})
	{{{- range .Columns }}}
	{{{- if .IsQuery }}}
	{{{- $queryOpr := index $.ModelOprMap .QueryType }}}
	{{{- if and (eq .JavaType "string") (eq $queryOpr "like") }}}
	if listReq.{{{ title (toCamelCase .ColumnName) }}} != "" {
        model = model.Where("{{{ .ColumnName }}} like ?", "%"+listReq.{{{ title (toCamelCase .ColumnName) }}}+"%")
    }
    {{{- else }}}
    if listReq.{{{ title (toCamelCase .ColumnName) }}} {{{ if eq .JavaType "string" }}}!= ""{{{ else }}}>=0{{{ end }}} {
        model = model.Where("{{{ .ColumnName }}} = ?", listReq.{{{ title (toCamelCase .ColumnName) }}})
    }
    {{{- end }}}
    {{{- end }}}
    {{{- end }}}
	{{{- if contains .AllFields "is_delete" }}}
	model = model.Where("is_delete = ?", 0)
	{{{- end }}}
	// 总数
	var count int64
	err := model.Count(&count).Error
	if e = response.CheckErr(err, "List Count err"); e != nil {
		return
	}
	// 数据
	var objs []{{{ title (toCamelCase .EntityName) }}}
	err = model.Limit(limit).Offset(offset).Order("id desc").Find(&objs).Error
	if e = response.CheckErr(err, "List Find err"); e != nil {
		return
	}
	resps := []{{{ title (toCamelCase .EntityName) }}}Resp{}
	response.Copy(&resps, objs)
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Lists:    resps,
	}, nil
}

//Detail {{{ .FunctionName }}}详情
func (srv {{{ toCamelCase .EntityName }}}Service) Detail(id uint) (res {{{ title (toCamelCase .EntityName) }}}Resp, e error) {
	var obj {{{ title (toCamelCase .EntityName) }}}
	err := srv.db.Where("{{{ $.PrimaryKey }}} = ?{{{ if contains .AllFields "is_delete" }}} AND is_delete = ?{{{ end }}}", id{{{ if contains .AllFields "is_delete" }}}, 0{{{ end }}}).Limit(1).First(&obj).Error
	if e = response.CheckErrDBNotRecord(err, "数据不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "Detail First err"); e != nil {
		return
	}
	response.Copy(&res, obj)
	{{{- range .Columns }}}
    {{{- if and .IsEdit (contains (slice "image" "avatar" "logo" "img") .JavaField) }}}
    res.Avatar = util.UrlUtil.ToAbsoluteUrl(res.Avatar)
    {{{- end }}}
    {{{- end }}}
	return
}

//Add {{{ .FunctionName }}}新增
func (srv {{{ toCamelCase .EntityName }}}Service) Add(addReq {{{ title (toCamelCase .EntityName) }}}AddReq) (e error) {
	var obj {{{ title (toCamelCase .EntityName) }}}
	response.Copy(&obj, addReq)
	err := srv.db.Create(&obj).Error
	e = response.CheckErr(err, "Add Create err")
	return
}

//Edit {{{ .FunctionName }}}编辑
func (srv {{{ toCamelCase .EntityName }}}Service) Edit(editReq {{{ title (toCamelCase .EntityName) }}}EditReq) (e error) {
	var obj {{{ title (toCamelCase .EntityName) }}}
	err := srv.db.Where("{{{ $.PrimaryKey }}} = ?{{{ if contains .AllFields "is_delete" }}} AND is_delete = ?{{{ end }}}", editReq.ID{{{ if contains .AllFields "is_delete" }}}, 0{{{ end }}}).Limit(1).First(&obj).Error
	// 校验
	if e = response.CheckErrDBNotRecord(err, "数据不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "Edit First err"); e != nil {
		return
	}
	// 更新
	response.Copy(&obj, editReq)
	err = srv.db.Model(&obj).Updates(obj).Error
	e = response.CheckErr(err, "Edit Updates err")
	return
}

//Del {{{ .FunctionName }}}删除
func (srv {{{ toCamelCase .EntityName }}}Service) Del(id uint) (e error) {
	var obj {{{ title (toCamelCase .EntityName) }}}
	err := srv.db.Where("{{{ $.PrimaryKey }}} = ?{{{ if contains .AllFields "is_delete" }}} AND is_delete = ?{{{ end }}}", id{{{ if contains .AllFields "is_delete" }}}, 0{{{ end }}}).Limit(1).First(&obj).Error
	// 校验
	if e = response.CheckErrDBNotRecord(err, "数据不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "Del First err"); e != nil {
		return
	}
    // 删除
    {{{- if contains .AllFields "is_delete" }}}
    obj.IsDelete = 1
    err = srv.db.Save(&obj).Error
    e = response.CheckErr(err, "Del Save err")
    {{{- else }}}
    err = srv.db.Delete(&obj).Error
    e = response.CheckErr(err, "Del Delete err")
    {{{- end }}}
	return
}
