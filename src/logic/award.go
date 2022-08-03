package logic

import (
	"taurus-backend/constant"
	"taurus-backend/db"
)

// GetRemainAwardCountByType 获取某种类型剩余奖券数量
func GetRemainAwardCountByType(awardType int) (remainCount int, err error) {
	total, err := db.GetTodayLotteryCountByAwardType(awardType)
	if err != nil {
		return 0, err
	}
	switch awardType {
	case constant.DIDI:
		remainCount = DIDI_AWARD_COUNT - total
	case constant.TENCENT:
		remainCount = TENCENT_AWARD_COUNT - total
	case constant.MEITUAN:
		remainCount = MEITUAN_AWARD_COUNT - total
	}

	if remainCount < 1 {
		return 0, nil
	}
	return remainCount, nil
}
