package route

import (
	"taurus-backend/api/handler"
	"taurus-backend/constant"
)

func getSignaturesRoutes() *NodeRoute {
	routers := []*Route{
		NewRoute(constant.HTTPMethodPost, "", handler.CreateSignature),
		NewRoute(constant.HTTPMethodGet, "/count", handler.GetSignatureCount),
	}

	return NewNodeRoute("/signatures", routers...)
}
