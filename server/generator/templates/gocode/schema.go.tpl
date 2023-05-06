package {{{ .PackageName }}}

import (
	"likeadmin/core"
)

//{{{ title (toCamelCase .EntityName) }}}ListReq {{{ .FunctionName }}}列表参数
type {{{ title (toCamelCase .EntityName) }}}ListReq struct {
    {{{- range .Columns }}}
    {{{- if .IsQuery }}}
    {{{ title (toCamelCase .JavaField) }}} {{{ .JavaType }}} `form:"{{{ toCamelCase .JavaField }}}"` // {{{ .ColumnComment }}}
    {{{- end }}}
    {{{- end }}}
}

//{{{ title (toCamelCase .EntityName) }}}DetailReq {{{ .FunctionName }}}详情参数
type {{{ title (toCamelCase .EntityName) }}}DetailReq struct {
    {{{- range .Columns }}}
    {{{- if .IsPk }}}
    {{{ title (toCamelCase .JavaField) }}} {{{ .JavaType }}} `form:"{{{ toCamelCase .JavaField }}}"` // {{{ .ColumnComment }}}
    {{{- end }}}
    {{{- end }}}
}

//{{{ title (toCamelCase .EntityName) }}}AddReq {{{ .FunctionName }}}新增参数
type {{{ title (toCamelCase .EntityName) }}}AddReq struct {
    {{{- range .Columns }}}
    {{{- if .IsInsert }}}
    {{{ title (toCamelCase .JavaField) }}} {{{ .JavaType }}} `form:"{{{ toCamelCase .JavaField }}}"` // {{{ .ColumnComment }}}
    {{{- end }}}
    {{{- end }}}
}

//{{{ title (toCamelCase .EntityName) }}}EditReq {{{ .FunctionName }}}新增参数
type {{{ title (toCamelCase .EntityName) }}}EditReq struct {
    {{{- range .Columns }}}
    {{{- if .IsEdit }}}
    {{{ title (toCamelCase .JavaField) }}} {{{ .JavaType }}} `form:"{{{ toCamelCase .JavaField }}}"` // {{{ .ColumnComment }}}
    {{{- end }}}
    {{{- end }}}
}

//{{{ title (toCamelCase .EntityName) }}}DelReq {{{ .FunctionName }}}新增参数
type {{{ title (toCamelCase .EntityName) }}}DelReq struct {
    {{{- range .Columns }}}
    {{{- if .IsPk }}}
    {{{ title (toCamelCase .JavaField) }}} {{{ .JavaType }}} `form:"{{{ toCamelCase .JavaField }}}"` // {{{ .ColumnComment }}}
    {{{- end }}}
    {{{- end }}}
}

//{{{ title (toCamelCase .EntityName) }}}Resp {{{ .FunctionName }}}返回信息
type {{{ title (toCamelCase .EntityName) }}}Resp struct {
	{{{- range .Columns }}}
    {{{- if or .IsList .IsPk }}}
    {{{ title (toCamelCase .JavaField) }}} {{{ .JavaType }}} `json:"{{{ toCamelCase .JavaField }}}" structs:"{{{ toCamelCase .JavaField }}}"` // {{{ .ColumnComment }}}
    {{{- end }}}
    {{{- end }}}
}
