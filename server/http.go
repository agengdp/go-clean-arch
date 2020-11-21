package server

import (
	"anwarmedika/auth/config"
	"fmt"
	"github.com/gin-gonic/gin"
)

type httpServer struct{
	router *gin.Engine
}

func NewServer(e *gin.Engine) *httpServer{
	return &httpServer{
		router: e,
	}
}

func (hs *httpServer) Start() error{
	var cfg *config.Config

	return hs.router.Run(fmt.Sprintf(":%s", cfg.Port))
}