package logic

import (
	"log"
	"taurus-backend/constant"
	"taurus-backend/db"
	"taurus-backend/sms"
	"time"
)

func SendLotteryMessage(phone string, awardType int, awardCode string) error {
	serialNo, err := sms.GetSMSClient().SendSMS(phone, awardType, awardCode)
	var smsSendStatus int
	if err != nil {
		log.Println("send sms fail, err: ", err)
		smsSendStatus = constant.SmsSendStatusFail
	} else {
		smsSendStatus = constant.SmsSendStatusSuccess
	}
	err = db.CreateSms(phone, awardType, awardCode, smsSendStatus, serialNo)
	if err != nil {
		log.Println("create sms record fail, err: ", err)
		return err
	}
	return nil
}

func LoopAndResend() {
	log.Println("loopAndResend failed message start")
	for {
		tasks, err := db.GetAllFailedSMSTask()
		if err != nil {
			log.Println("get all failed sms task error:", err)
			// retry after 30 seconds
			continue
		}
		log.Printf("get all failed sms task success,len: %v, tasks: %v", len(tasks), tasks)
		for _, task := range tasks {
			err := db.UpdateSmsStatusWithLock(task)
			if err != nil {
				log.Printf("update sms status fail, task: %v, err: %v", task, err)
				continue
			}
		}
		log.Println("loopAndResend failed message success")
		time.Sleep(time.Second * 5)
	}
}
