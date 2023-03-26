package routers

import (
	"likeadmin/core"
	"likeadmin/generator/routers/gen"
)

var InitRouters = []*core.GroupBase{
	// gen
	gen.GenGroup,
}
