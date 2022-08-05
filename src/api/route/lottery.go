package route

import (
	"taurus-backend/api/handler"
	"taurus-backend/constant"
)

func getLotteriesRoutes() *NodeRoute {
	routers := []*Route{
		NewRoute(constant.HTTPMethodGet, "/chance", handler.GetLotteryChance),
		NewRoute(constant.HTTPMethodPost, "", handler.WinLottery),
	}

	return NewNodeRoute("/lotteries", routers...)
}
