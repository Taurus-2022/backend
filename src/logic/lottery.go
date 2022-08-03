package logic

import (
	"log"
	"math/rand"
	"os"
	"strconv"
	"taurus-backend/constant"
	"taurus-backend/db"
	"time"
)

const (
	MEITUAN_RATE = 0.0045402985074627
	TENCENT_RATE = 0.0022701492537314
	DIDI_RATE    = 0.00075671641791045

	MEITUAN_AWARD_COUNT = 300
	TENCENT_AWARD_COUNT = 150
	DIDI_AWARD_COUNT    = 57

	// MULTIPLE 最大值为130
	MULTIPLE = 1
)

func WinLottery(phone string) (isWinLottery bool, awardType int, err constant.StatusError) {
	// 校验是否能抽奖
	errCode := GetCanDoLottery(phone)
	if errCode != constant.ErrorCodeOK {
		// 不能抽奖， 返回为什么不能抽奖
		return false, constant.NONE, constant.NewStatusError(errCode, nil)
	}
	isWinLottery, awardType, err = CalculateWinLottery()
	if err.IsNotOK() {
		// 异常情况
		return false, constant.NONE, err
	}

	var dbError error
	if isWinLottery {
		// 抽中
		dbError = CreateAwardLottery(phone, isWinLottery, awardType)
		if dbError == nil {
			// 中奖减库存成功
			log.Println("win lottery success, phone: ", phone, ", awardType: ", awardType)
			return isWinLottery, awardType, constant.NewOKStatusError()
		}
		// db 处理失败 转为未中奖
		log.Printf("win but consume award lottery fail, err: %v", err)
		isWinLottery = false
		awardType = constant.NONE
	}
	// 插入未中奖记录
	dbError = db.CreateLottery(phone, isWinLottery, constant.NONE, "")
	if dbError != nil {
		log.Printf("not win but insert award lottery fail, err: %v", err)
		return isWinLottery, awardType, constant.NewStatusError(constant.ErrorDbInnerError, dbError)
	}
	return isWinLottery, awardType, constant.NewOKStatusError()
}

func CalculateWinLottery() (isWinLottery bool, awardType int, err constant.StatusError) {
	// 计算是否中奖
	isWinLottery, awardType = WinLotteryFunc()
	if !isWinLottery {
		// 没有中奖
		return false, constant.NONE, constant.NewOKStatusError()
	}
	remainAward, error := GetRemainAwardCountByType(awardType)
	if error != nil {
		return false, constant.NONE, constant.NewStatusError(constant.ErrorWinLotteryFailed, error)
	}
	if remainAward < 1 {
		// 该类奖券已经抽完，算作没抽中
		return false, constant.NONE, constant.NewOKStatusError()
	}
	return true, awardType, constant.NewOKStatusError()
}

// WinLotteryFunc 计算用户是否中奖.
func WinLotteryFunc() (isWinLottery bool, awardType int) {
	rand.Seed(time.Now().UnixNano())
	r := rand.Float64() / GetMultiple()
	if r < DIDI_RATE {
		// 滴滴
		return true, constant.DIDI
	} else if r >= DIDI_RATE && r < DIDI_RATE+TENCENT_RATE {
		// 腾讯奖券
		return true, constant.TENCENT
	} else if r >= DIDI_RATE+TENCENT_RATE && r < DIDI_RATE+TENCENT_RATE+MEITUAN_RATE {
		// 美团奖券
		return true, constant.MEITUAN
	}
	return false, constant.NONE
}

func GetMultiple() float64 {
	var multiple float64
	if os.Getenv("MULTIPLE") != "" {
		multiple, _ = strconv.ParseFloat(os.Getenv("MULTIPLE"), 64)
	}
	// 最大限制为130
	if multiple == 0 || multiple > 130 {
		multiple = MULTIPLE
	}
	return multiple
}

// GetCanDoLottery 校验是否能抽奖
func GetCanDoLottery(phone string) (errCode int) {
	hasWinLottery, err := HasWinLottery(phone)
	if err != nil {
		return constant.ErrorWinLotteryFailed
	}
	if hasWinLottery {
		return constant.ErrorHasWinLottery
	}
	remainAward, err := db.GetRemainAwardCount()
	if err != nil {
		return constant.ErrorWinLotteryFailed
	}
	if remainAward < 1 {
		// 全部没有奖券
		return constant.ErrorNoMoreAward
	}
	total, err := db.GetTodayLotteryCountByPhone(phone)
	if err != nil {
		return constant.ErrorWinLotteryFailed
	}
	if total > 1 {
		// 今天没有抽奖机会了
		return constant.ErrorTodayNoMoreLotteryChance
	}
	return constant.ErrorCodeOK
}

func CreateAwardLottery(phone string, isWinLottery bool, awardType int) error {
	awardCode, err := db.CreateAwardLottery(phone, isWinLottery, awardType)
	if err != nil {
		return err
	}
	go func(code string) {
		// TODO 异步发送短信
	}(awardCode)
	return nil
}

func HasWinLottery(phone string) (bool, error) {
	total, err := db.GetWinLotteryCountByPhone(phone)
	if err != nil {
		return false, err
	}
	return total > 0, nil
}
