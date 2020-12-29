package server

import (
	"github.com/gin-gonic/gin"
)

type Server struct{}

func New() *Server {
	return &Server{}
}

func (s *Server) Start() *gin.Engine {
	app := gin.New()
	Routes(app)
	return app
}
