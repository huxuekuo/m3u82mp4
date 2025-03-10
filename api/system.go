package api

import (
	"m3u82mp4/api/middleware"
	systemser "m3u82mp4/model/service/system"
	"m3u82mp4/model/system"

	"github.com/gin-gonic/gin"
)

type SystemAPI struct {
	BaseApi
}

func InitSysRouter(r *gin.RouterGroup) {
	api := &SystemAPI{}
	api.RouterGroup = r.Group("sys", middleware.IgnoreUser, middleware.SetDB)
	api.Api("POST", "/config", api.config)

}
func (s *SystemAPI) config(c *gin.Context) any {
	var param system.ConfigParam
	c.ShouldBindJSON(&param)
	configService := systemser.NewConfigService(s.MysqlDB)
	method := map[string]func(*system.ConfigParam) any{
		"add":       configService.APIAdd,
		"queryList": configService.APIQueryList,
	}
	return method[param.MT](&param)
}
