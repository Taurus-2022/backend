package response

type GetLotteryChanceResponse struct {
	CanParticipateLottery bool `json:"canParticipateLottery"`
}

type WinLotteryResponse struct {
	IsWinLottery bool `json:"isWinLottery"`
	LotteryType  int  `json:"lotteryType"`
}
