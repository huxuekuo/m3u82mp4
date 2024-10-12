package router

import (
	"m3u82mp4/api"
	"m3u82mp4/api/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.RouterGroup) {
	r.Use(middleware.Cors, middleware.SetLogger, middleware.SetDB)
	api.InitVideoRouter(r)
	api.InitUserRouter(r)
}
