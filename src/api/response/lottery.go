package response

type GetLotteryChanceResponse struct {
	CanParticipateLottery bool `json:"canParticipateLottery"`
	HasWinLottery         bool `json:"hasWinLottery"`
}

type WinLotteryResponse struct {
	IsWinLottery bool `json:"isWinLottery"`
	AwardType    int  `json:"awardType"`
}
