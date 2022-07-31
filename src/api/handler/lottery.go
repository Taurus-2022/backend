package handler

import (
	"github.com/gin-gonic/gin"
	"taurus-backend/api/request"
	"taurus-backend/api/response"
	"taurus-backend/constant"
)

func GetLotteryChance(c *gin.Context) {
	req := &request.GetLotteryChangeRequest{}
	if err := c.ShouldBindQuery(req); err != nil {
		c.AbortWithStatusJSON(400, response.GetErrorResponse(constant.ErrorHttpParamInvalid, err.Error()))
		return
	}
	resp := &response.GetLotteryChanceResponse{CanParticipateLottery: true}
	c.JSON(200, resp)
}

func WinLottery(c *gin.Context) {
	req := &request.WinLotteryRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.AbortWithStatusJSON(400, response.GetErrorResponse(constant.ErrorHttpParamInvalid, err.Error()))
		return
	}
	resp := &response.WinLotteryResponse{IsWinLottery: true, AwardType: constant.MEITUAN}
	c.JSON(200, resp)
}
