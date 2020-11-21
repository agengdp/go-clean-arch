package main

import (
	"anwarmedika/auth/middleware"
	"anwarmedika/auth/server"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(-1)
	}
}

func run() error {
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())

	svr := server.NewServer(r)
	svr.MapRoutes()

	return svr.Start()
}