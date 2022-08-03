package logic

import (
	"taurus-backend/db"
)

func GetTodayUserIsSigned(phone string) (isSigned bool, err error) {
	total, err := db.GetTodaySignatureCountByPhone(phone)
	if err != nil {
		return false, err
	}
	return total > 0, nil
}
