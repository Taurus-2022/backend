package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"taurus-backend/api/request"
	"taurus-backend/api/response"
	"taurus-backend/constant"
	"taurus-backend/db"
	"taurus-backend/logic"
)

func GetLotteryChance(c *gin.Context) {
	req := &request.GetLotteryChangeRequest{}
	if err := c.ShouldBindQuery(req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.GetErrorResponse(constant.ErrorHttpParamInvalid, err.Error()))
		return
	}
	total, err := db.GetWinLotteryCountByPhone(req.Phone)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.GetErrorResponse(constant.ErrorDbInnerError, err.Error()))
		return
	}
	if total > 0 {
		// 中过奖
		c.JSON(http.StatusOK, &response.GetLotteryChanceResponse{CanParticipateLottery: false, HasWinLottery: true})
		return
	}
	total, err = db.GetTodayLotteryCountByPhone(req.Phone)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.GetErrorResponse(constant.ErrorDbInnerError, err.Error()))
		return
	}
	resp := &response.GetLotteryChanceResponse{CanParticipateLottery: total < 1, HasWinLottery: false}
	c.JSON(http.StatusOK, resp)
}

func WinLottery(c *gin.Context) {
	req := &request.WinLotteryRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.GetErrorResponse(constant.ErrorHttpParamInvalid, err.Error()))
		return
	}
	isWinLottery, awardType, err := logic.WinLottery(req.Phone)
	if err.IsNotOK() {
		if err.IsDBError() {
			c.AbortWithStatusJSON(http.StatusInternalServerError, response.ToErrorResponse(err))
			return
		}
	}
	c.JSON(http.StatusOK, response.WinLotteryResponse{IsWinLottery: isWinLottery, AwardType: awardType})
}
