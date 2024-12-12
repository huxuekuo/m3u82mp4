package main

import (
	"m3u82mp4/api/middleware"
	"m3u82mp4/api/router"
	"m3u82mp4/library"
	"m3u82mp4/task"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	library.Logger, _ = zap.NewProduction()
	library.InitMysql()
	r := gin.New()
	r.Use(middleware.GinLogger(library.Logger), middleware.GinRecovery(library.Logger, true))
	router.InitRouter(&r.RouterGroup)
	// 加载任务
	task.InitTask()
	r.Run(":8080")
}
