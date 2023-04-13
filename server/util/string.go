package util

import (
	"bytes"
	"strings"
	"unicode"
)

var StringUtil = stringUtil{}

//arrayUtil 数组工具类
type stringUtil struct{}

func (su stringUtil) ToSnakeCase(s string) string {
	buf := bytes.Buffer{}
	for i, r := range s {
		if unicode.IsUpper(r) {
			if i > 0 {
				buf.WriteRune('_')
			}
			buf.WriteRune(unicode.ToLower(r))
		} else {
			buf.WriteRune(r)
		}
	}
	return buf.String()
}
func (su stringUtil) ToCamelCase(s string) string {
	words := strings.Split(s, "_")
	for i := 1; i < len(words); i++ {
		words[i] = strings.Title(words[i])
	}
	return strings.Join(words, "")
}
