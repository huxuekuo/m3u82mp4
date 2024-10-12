package main

import (
	"m3u82mp4/api/router"
	"m3u82mp4/library"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	library.Logger, _ = zap.NewProduction()
	library.InitMysql()
	r := gin.Default()
	router.InitRouter(&r.RouterGroup)
	r.Run(":8080")
}
