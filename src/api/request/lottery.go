package request

type GetLotteryChangeRequest struct {
	Phone string `form:"phone" binding:"required"`
}

type WinLotteryRequest struct {
	Phone string `json:"phone" binding:"required"`
}
