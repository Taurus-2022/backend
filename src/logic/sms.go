package logic

import (
	"log"
	"taurus-backend/db"
	"time"
)

func LoopAndResend() {
	log.Println("loopAndResend failed message start")
	for {
		log.Println("Pull task after 30s ...")
		time.Sleep(time.Second * 30)
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
	}
}
