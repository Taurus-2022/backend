package route

import (
	"taurus-backend/constant"
	"taurus-backend/handler"
)

func getPingRoutes() *NodeRoute {
	routers := []*Route{
		NewRoute(constant.HTTPMethodGet, "ping", handler.Ping),
	}

	return NewNodeRoute("/", routers...)
}
