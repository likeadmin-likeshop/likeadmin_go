package core

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
	"log"
)

type GroupBase struct {
	basePath    string
	initHandle  interface{}
	regHandle   func(rg *gin.RouterGroup, group *GroupBase) error
	middlewares []gin.HandlerFunc
}

// Group creates a new router group
func Group(relativePath string, initHandle interface{}, regHandle func(rg *gin.RouterGroup, group *GroupBase) error, middlewares ...gin.HandlerFunc) *GroupBase {
	return &GroupBase{
		basePath:    relativePath,
		initHandle:  initHandle,
		regHandle:   regHandle,
		middlewares: middlewares,
	}
}

//RegisterGroup registers all handle of group to gin
func RegisterGroup(rg *gin.RouterGroup, group *GroupBase) {
	r := rg.Group(group.basePath)
	if len(group.middlewares) > 0 {
		r.Use(group.middlewares...)
	}
	if err := ProvideForDI(group.initHandle); err != nil {
		log.Fatalln(err)
	}
	if err := group.regHandle(r, group); err != nil {
		log.Fatalln(err)
	}
}

//Reg registers handle by DI
func (group GroupBase) Reg(function interface{}, opts ...dig.InvokeOption) error {
	return DI(function, opts...)
}
