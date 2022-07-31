package route

import (
	"github.com/gin-gonic/gin"
	"taurus-backend/constant"
)

type Route struct {
	method  string
	path    string
	handler gin.HandlerFunc
}

type NodeRoute struct {
	path   string
	routes []*Route
}

func (n *NodeRoute) registerRoutes(r *gin.RouterGroup) {
	group := r.Group(n.path)
	for _, route := range n.routes {
		methodMapper(group, route.method)(route.path, route.handler)
	}
}
func NewRoute(method, path string, handler gin.HandlerFunc) *Route {
	return &Route{
		method:  method,
		path:    path,
		handler: handler,
	}
}

func NewNodeRoute(path string, routers ...*Route) *NodeRoute {
	return &NodeRoute{
		path:   path,
		routes: routers,
	}
}
func methodMapper(group *gin.RouterGroup, method string) func(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	switch method {
	case constant.HTTPMethodGet:
		return group.GET
	case constant.HTTPMethodPost:
		return group.POST
	case constant.HTTPMethodPut:
		return group.PUT
	case constant.HTTPMethodDelete:
		return group.DELETE
	case constant.HTTPMethodPatch:
		return group.PATCH
	case constant.HTTPMethodHead:
		return group.HEAD
	default:
		return group.Any
	}
}

func mountRoutes(r *gin.Engine, routers ...*NodeRoute) {
	router := r.Group("/")
	for _, node := range routers {
		node.registerRoutes(router)
	}
}

func InitAllRouters(r *gin.Engine) {
	nodes := []*NodeRoute{
		getPingRoutes(),
		getSignaturesRoutes(),
		getLotteriesRoutes(),
	}
	mountRoutes(r, nodes...)
}
