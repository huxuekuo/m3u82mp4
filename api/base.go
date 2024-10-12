package api

import (
	"fmt"
	"m3u82mp4/consts"
	"m3u82mp4/consts/errcode"
	"m3u82mp4/library"
	"m3u82mp4/model/db"
	"m3u82mp4/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type BaseApi struct {
	*gin.RouterGroup
	Token   string
	Logger  *zap.Logger
	MysqlDB *gorm.DB
	User    *db.User
	URK     string
}

type HandleFunc func(*gin.Context) any

type Respone struct {
	Data any    `json:"data"`
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (r *Respone) OK(data any) {
	r.Code = 200
	r.Data = data
}

func (r *Respone) OK2() *Respone {
	r.Code = 200
	return r
}

func (r *Respone) Set(e errcode.ErrorCode) *Respone {
	r.Code = e.Code
	r.Msg = e.Msg
	return r
}

func (b *BaseApi) Api(httpMethod, relativePath string, h HandleFunc, befor ...gin.HandlerFunc) {
	befor = append(befor, func(c *gin.Context) {
		err := b.LoadSysConf(c)
		if err != nil {
			c.JSON(200, err)
			return
		}
		res := h(c)
		c.JSON(200, res)
	})
	b.Handle(httpMethod, relativePath, befor...)
}

func (b *BaseApi) ApiByte(httpMethod, relativePath string, h HandleFunc) {
	b.Handle(httpMethod, relativePath, func(c *gin.Context) {
		err := b.LoadSysConf(c)
		if err != nil {
			c.JSON(200, err)
			return
		}
		res := h(c)
		c.Writer.Write(res.([]byte))
	})
}

func (b *BaseApi) LoadSysConf(c *gin.Context) *errcode.ErrorCode {
	_, ignoreOk := c.Get("ignoreUser")
	if !ignoreOk {
		URK, err := c.Cookie(consts.SYS_AUTH_USER_REDIS_KEY_NAME)
		if err != nil {
			return &errcode.SYS_USER_READ_ERR
		}
		b.URK = URK
	}
	for k, _ := range consts.LoadSysConf {
		if cv, ok := c.Keys[k]; ok {
			if k == consts.SYS_LOGGER_VAR_NAME {
				b.Logger = cv.(*zap.Logger)
			} else if k == consts.SYS_MYSQL_VAR_NAME {
				b.MysqlDB = cv.(*gorm.DB)
			} else if k == consts.SYS_AUTH_USER_INFO_VAR_NAME {
				jwtTokenCustom := cv.(*utils.JwtTokenCustom)
				if jwtTokenCustom == nil {
					b.Logger.Sugar().Error("解析Token为nil")
				} else {
					userdb := db.NewUserDB(library.MysqlDB)
					userInfo := userdb.GetUserInfoById(jwtTokenCustom.Uid)
					if userInfo == nil {
						b.Logger.Sugar().Error("根据token内容获取用户信息失败")
					} else {
						b.User = userInfo
						b.URK = fmt.Sprint(b.User.Id)
					}
				}
			}
		}
	}
	return nil
}
