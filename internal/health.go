package internal

import (
	"github.com/gin-gonic/gin"
	"net/http"
	)

type healthCtrl struct{}

func (h healthCtrl) Ping( ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{"status": "healthy"})
}
