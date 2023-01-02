package core

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type routerBase struct {
	method      string
	path        string
	handler     gin.HandlerFunc
	middlewares []gin.HandlerFunc
}

type groupBase struct {
	basePath  string
	routerMap map[string]routerBase
}

// Group creates a new router group
func Group(relativePath string) *groupBase {
	return &groupBase{
		basePath:  relativePath,
		routerMap: make(map[string]routerBase),
	}
}

//RegisterGroup registers all handle to gin
func RegisterGroup(rg *gin.RouterGroup, group *groupBase, useFunc func(g *gin.RouterGroup)) {
	r := rg.Group(group.basePath)
	if useFunc != nil {
		useFunc(r)
	}
	for _, item := range group.routerMap {
		var handlers []gin.HandlerFunc
		if item.middlewares != nil {
			handlers = item.middlewares
		}
		handlers = append(handlers, item.handler)
		r.Handle(item.method, item.path, handlers...)
	}
}

//AddHandle registers a new request handle
func (group *groupBase) AddHandle(httpMethod, relativePath string, handler gin.HandlerFunc, middlewares []gin.HandlerFunc) *groupBase {
	group.routerMap[relativePath] = routerBase{
		method: httpMethod, path: relativePath, handler: handler, middlewares: middlewares}
	return group
}

// AddPOST is a shortcut for router.AddHandle("POST", path, handle).
func (group *groupBase) AddPOST(relativePath string, handler gin.HandlerFunc, middlewares ...gin.HandlerFunc) *groupBase {
	group.AddHandle(http.MethodPost, relativePath, handler, middlewares)
	return group
}

// AddGET is a shortcut for router.AddHandle("GET", path, handle).
func (group *groupBase) AddGET(relativePath string, handler gin.HandlerFunc, middlewares ...gin.HandlerFunc) *groupBase {
	group.AddHandle(http.MethodGet, relativePath, handler, middlewares)
	return group
}

// AddDELETE is a shortcut for router.AddHandle("DELETE", path, handle).
func (group *groupBase) AddDELETE(relativePath string, handler gin.HandlerFunc, middlewares ...gin.HandlerFunc) *groupBase {
	group.AddHandle(http.MethodDelete, relativePath, handler, middlewares)
	return group
}

// AddPATCH is a shortcut for router.AddHandle("PATCH", path, handle).
func (group *groupBase) AddPATCH(relativePath string, handler gin.HandlerFunc, middlewares ...gin.HandlerFunc) *groupBase {
	group.AddHandle(http.MethodPatch, relativePath, handler, middlewares)
	return group
}

// AddPUT is a shortcut for router.AddHandle("PUT", path, handle).
func (group *groupBase) AddPUT(relativePath string, handler gin.HandlerFunc, middlewares ...gin.HandlerFunc) *groupBase {
	group.AddHandle(http.MethodPut, relativePath, handler, middlewares)
	return group
}
