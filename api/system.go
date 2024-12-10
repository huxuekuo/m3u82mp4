package api

import (
	"m3u82mp4/api/middleware"
	"m3u82mp4/consts"
	systemser "m3u82mp4/model/service/system"
	"m3u82mp4/model/system"

	"github.com/gin-gonic/gin"
)

type SystemAPI struct {
	BaseApi
}

// method 地址探针options
var method = map[string]func(*gin.Context, *system.VideoURLProbeParam) any{
	consts.SYS_VIDEO_QUERY:         systemser.QueryProbe,
	consts.SYS_VIDEO_QUERY_ADD:     systemser.AddProbe,
	consts.SYS_VIDEO_QUERY_GET:     systemser.GetURLApi,
	consts.SYS_VIDEO_QUERY_DELETED: systemser.Deleted,
}

func InitSysRouter(r *gin.RouterGroup) {
	api := &SystemAPI{}
	api.RouterGroup = r.Group("sys", middleware.IgnoreUser)
	api.Api("POST", "/probeOptions", api.probeOptions)

}

func (s SystemAPI) probeOptions(c *gin.Context) any {
	var param system.VideoURLProbeParam
	c.ShouldBindJSON(&param)
	return method[param.Type](c, &param)
}
