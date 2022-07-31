package handler

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"taurus-backend/api/request"
	"taurus-backend/api/response"
	"taurus-backend/constant"
	"taurus-backend/db"
	"time"
)

func GetLotteryChance(c *gin.Context) {
	req := &request.GetLotteryChangeRequest{}
	if err := c.ShouldBindQuery(req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.GetErrorResponse(constant.ErrorHttpParamInvalid, err.Error()))
		return
	}
	total, err := db.GetTodayLotteryCountByPhone(req.Phone)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.GetErrorResponse(constant.ErrorDbInnerError, err.Error()))
		return
	}
	resp := &response.GetLotteryChanceResponse{CanParticipateLottery: total < 1}
	c.JSON(http.StatusOK, resp)
}

func WinLottery(c *gin.Context) {
	req := &request.WinLotteryRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.GetErrorResponse(constant.ErrorHttpParamInvalid, err.Error()))
		return
	}
	canDoLottery, err := getCanDoLottery(req.Phone)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.GetErrorResponse(constant.ErrorDbInnerError, err.Error()))
		return
	}
	if !canDoLottery {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.GetErrorResponse(constant.ErrorLotteryNoChance, ""))
		return
	}
	isWinLottery, awardType := WinLotteryFunc()
	var saveLotteryErr error
	if isWinLottery {
		saveLotteryErr = CreateAwardLottery(req.Phone, isWinLottery, awardType)
	} else {
		saveLotteryErr = db.CreateLottery(req.Phone, isWinLottery, constant.NONE, "")
	}
	var resp *response.WinLotteryResponse
	if saveLotteryErr != nil {
		// 提示用户可以重复抽奖
		resp = &response.WinLotteryResponse{IsWinLottery: false, AwardType: constant.NONE}
	} else {
		resp = &response.WinLotteryResponse{IsWinLottery: isWinLottery, AwardType: awardType}
	}
	c.JSON(http.StatusOK, resp)
}

func getCanDoLottery(phone string) (bool, error) {
	remainAward, err := db.GetRemainAwardCount()
	if err != nil || remainAward < 1 {
		// 全部没有奖券
		return false, err
	}
	total, err := db.GetTodayLotteryCountByPhone(phone)
	if err != nil {
		// 今天没有抽奖机会了
		return false, err
	}
	return total < 1, err
}

func CreateAwardLottery(phone string, lottery bool, awardType int) error {
	awardCode, err := db.CreateAwardLottery(phone, lottery, awardType)
	if err != nil {
		return err
	}
	go func(code string) {
		// TODO 异步发送短信
	}(awardCode)
	return nil
}

// WinLotteryFunc 计算用户是否中奖.
func WinLotteryFunc() (isWinLottery bool, awardType int) {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(100)
	if r < 10 {
		// 美团奖券
		return true, constant.MEITUAN
	} else if r >= 10 && r < 20 {
		// 腾讯奖券
		return true, constant.TENCENT
	} else if r >= 20 && r < 30 {
		// 青桔奖券
		return true, constant.DIDI
	}
	return false, constant.NONE
}
