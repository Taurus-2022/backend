package api

import (
	"log"
	"taurus-backend/api/route"
	"taurus-backend/constant"

	"github.com/gin-gonic/gin"
)

type Server struct {
	srv *gin.Engine
}

func NewServer() *Server {
	a := &Server{}
	return a
}

func (s *Server) Init(e *constant.Env) {
	gin.DefaultWriter = log.Writer()
	r := gin.Default()
	route.InitAllRouters(r)
	s.srv = r
}

func (s *Server) Run() {
	// Web Func硬性要求
	err := s.srv.Run("0.0.0.0:9000")
	log.Println("server starting...")
	if err != nil {
		log.Fatal("setup server fatal:", err)
		return
	}
}
