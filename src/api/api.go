package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"taurus-backend/api/route"
)

type Server struct {
	srv *gin.Engine
}

func NewServer() *Server {
	a := &Server{}
	return a
}

func (s *Server) Init() {
	gin.DefaultWriter = log.Writer()
	r := gin.Default()
	route.InitAllRouters(r)
	s.srv = r
}

func (s *Server) Run() {
	// Web Func硬性要求
	go func() {
		err := s.srv.Run("0.0.0.0:9000")
		if err != nil {
			log.Fatal("server run fail", err)
		}
	}()
	log.Println("server starting...")
}
