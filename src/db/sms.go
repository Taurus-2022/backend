package db

import "log"

func CreateSms(phone string, awardType int, awardCode string, status int, serialNo string) (err error) {
	_, err = GetDB().Exec("INSERT INTO sms (phone, award_type, award_code, is_sms_sent, serial_no) VALUES (?, ?, ?, ?, ?)", phone, awardType, awardCode, status, serialNo)
	if err != nil {
		log.Println("create sms fail, err: ", err)
		return err
	}
	return nil
}
