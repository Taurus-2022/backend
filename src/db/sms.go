package db

import (
	"database/sql"
	"errors"
	"log"
	"taurus-backend/constant"
	"taurus-backend/sms"
)

func GetAllFailedSMSTask() (tasks []*sms.Task, err error) {
	rows, err := GetDB().Query("SELECT phone, award_type, award_code FROM sms WHERE is_sms_sent = ?", constant.SmsSendStatusFail)
	if err != nil {
		log.Println("get all failed sms task fail, err: ", err)
		return nil, nil
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Println("close rows fail, err: ", err)
		}
	}(rows)

	tasks = make([]*sms.Task, 0)
	for rows.Next() {
		task := &sms.Task{}
		err := rows.Scan(&task.Phone, &task.AwardType, &task.AwardCode)
		if err != nil {
			log.Println("sms scan task fail, err: ", err)
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func UpdateSmsStatusWithLock(task *sms.Task) error {
	tx, err := GetDB().Begin()
	if err != nil {
		log.Println("begin transaction fail, err: ", err)
		return err
	}
	defer tx.Rollback()

	var serialNo string
	err = tx.QueryRow("SELECT serial_no FROM sms WHERE phone = ? AND is_sms_sent = ? FOR UPDATE", task.Phone, constant.SmsSendStatusFail).Scan(&serialNo)

	if err != nil {
		log.Println("lock sms record fail or can not find this failed task, err: ", err)
		return err
	}
	serialNo, err = sms.GetSMSClient().SendSMS(task.Phone, task.AwardType, task.AwardCode)
	if err != nil {
		log.Println("send sms fail, err: ", err)
		return err
	}
	result, err := tx.Exec("UPDATE sms SET is_sms_sent = ?, serial_no = ? WHERE phone = ? AND is_sms_sent = ?", constant.SmsSendStatusSuccess, serialNo, task.Phone, constant.SmsSendStatusFail)
	if err != nil {
		log.Println("update sms record fail, err: ", err)
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		log.Println("get rows affected fail, err: ", err)
		return err
	}
	if rows != 1 {
		log.Println("update sms record fail, rows affected: ", rows)
		return errors.New("update sms record fail")
	}
	return tx.Commit()
}
