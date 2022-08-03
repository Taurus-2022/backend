package db

import "log"

func GetRemainAwardCount() (total int, err error) {
	err = GetDB().QueryRow("SELECT COUNT(*) AS total FROM award WHERE is_used = 0").Scan(&total)
	if err != nil {
		log.Fatalf("get remain award count fail, err: %v", err)
		return 0, err
	}
	return total, nil
}
