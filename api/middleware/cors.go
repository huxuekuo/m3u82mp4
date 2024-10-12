package middleware

import (
	"m3u82mp4/consts"
	"m3u82mp4/library"
	"m3u82mp4/utils"

	"github.com/gin-gonic/gin"
)

func SetLogger(c *gin.Context) {
	c.Set(consts.SYS_LOGGER_VAR_NAME, library.Logger)
	c.Next()
}

func SetDB(c *gin.Context) {
	c.Set(consts.SYS_MYSQL_VAR_NAME, library.MysqlDB)
	c.Next()
}

func SetUserInfo(c *gin.Context) {
	token, err := c.Cookie("urk")
	if err != nil {
		library.Logger.Sugar().Error(err)
	} else {
		authToken := utils.TokenParse(token)
		c.Set(consts.SYS_AUTH_USER_INFO_VAR_NAME, authToken)
	}
	c.Next()
}

func IgnoreUser(c *gin.Context) {
	c.Set("ignoreUser", 1)
}

func Cors(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "POST,GET,PUT,DELETE,OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type,Content-Length,Token")
	c.Header("Access-Control-Expose-Headers", "Access-Control-Allow-Headers,Token")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Next()
}
