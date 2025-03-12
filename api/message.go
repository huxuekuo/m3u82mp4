package api

import (
	"m3u82mp4/api/middleware"
	messageser "m3u82mp4/model/service/message"

	"github.com/gin-gonic/gin"
)

type MessageAPI struct {
	BaseApi
}

func InitMessageRouter(r *gin.RouterGroup) {
	api := &MessageAPI{}
	api.RouterGroup = r.Group("/message", middleware.SetDB, middleware.SetUserInfo)
	api.Api("POST", "/add", api.add)
	api.Api("POST", "/delete", api.delete)
	api.Api("POST", "/reply", api.reply)
	api.Api("GET", "/list", api.list)
}

func (a *MessageAPI) add(c *gin.Context) any {
	type Param struct {
		Content  string `json:"content"`
		ParentId int64  `json:"parentId"`
	}
	var param Param
	c.ShouldBindJSON(&param)

	user := a.User
	messageService := messageser.NewMessageService(a.MysqlDB)
	return messageService.Add(param.Content, param.ParentId, user.Id, user.Username, user.Avatar)
}

func (a *MessageAPI) delete(c *gin.Context) any {
	type Param struct {
		Id int64 `json:"id"`
	}
	var param Param
	c.ShouldBindJSON(&param)

	user := a.User
	messageService := messageser.NewMessageService(a.MysqlDB)
	return messageService.Delete(param.Id, user.Id)
}

func (a *MessageAPI) reply(c *gin.Context) any {
	type Param struct {
		Content  string `json:"content"`
		ParentId int64  `json:"parentId"`
	}
	var param Param
	c.ShouldBindJSON(&param)

	user := a.User
	messageService := messageser.NewMessageService(a.MysqlDB)
	return messageService.Add(param.Content, param.ParentId, user.Id, user.Username, user.Avatar)
}

func (a *MessageAPI) list(c *gin.Context) any {
	messageService := messageser.NewMessageService(a.MysqlDB)
	return messageService.QueryList()
}
