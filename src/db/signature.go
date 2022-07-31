package db

import "log"

func GetSignatureCountByPhone(phone string) (total int, err error) {
	err = GetDB().QueryRow("SELECT COUNT(*) AS total FROM signature WHERE phone = ?", phone).Scan(&total)
	if err != nil {
		log.Printf("get count signature by phone fail: %v ,err: %v", phone, err)
		return total, err
	}
	return total, nil
}
