package {{{ .PackageName }}}

//{{{ title (toCamelCase .EntityName) }}} {{{ .FunctionName }}}实体
type {{{ title (toCamelCase .EntityName) }}} struct {
	{{{- range .Columns }}}
    {{{- if not (contains $.SubTableFields .ColumnName) }}}
    {{{ title (toCamelCase .JavaField) }}} {{{ if eq .JavaType "core.TsTime" }}} int64 {{{ else }}} {{{ .JavaType }}} {{{ end }}} `gorm:"{{{ if .IsPk }}}primarykey;{{{ end }}}comment:'{{{ .ColumnComment }}}'"` // {{{ .ColumnComment }}}
    {{{- end }}}
    {{{- end }}}
}
