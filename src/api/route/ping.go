package route

import (
	"taurus-backend/api/handler"
	"taurus-backend/constant"
)

func getPingRoutes() *NodeRoute {
	routers := []*Route{
		NewRoute(constant.HTTPMethodGet, "/ping", handler.Ping),
	}

	return NewNodeRoute("/", routers...)
}
