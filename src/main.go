package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"taurus-backend/api/route"
	"taurus-backend/constant"
	"taurus-backend/db"
)

var (
	host = "0.0.0.0"
	port = 9000
)

func init() {
	if err := db.GetDB().Ping(); err != nil {
		log.Fatal("open database fail")
		return
	}
	log.Println("database connect success")

	if constant.LocalStage != os.Getenv(constant.Stage) {
		gin.SetMode(gin.ReleaseMode)
	}
}

func main() {
	gin.DefaultWriter = log.Writer()
	r := gin.Default()
	route.InitAllRouters(r)
	err := r.Run(fmt.Sprintf("%s:%d", host, port))
	log.Println("server starting...")
	if err != nil {
		log.Fatal("setup server fatal:", err)
		return
	}
}
