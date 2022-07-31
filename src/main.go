package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"taurus-backend/api/route"
)

var (
	host_ = "0.0.0.0"
	port_ = 9000
)

func main() {
	var r = gin.Default()
	route.InitAllRouters(r)
	err := r.Run(fmt.Sprintf("%s:%d", host_, port_))
	log.Println("server starting...")
	if err != nil {
		log.Fatal("setup server fatal:", err)
		return
	}
}
