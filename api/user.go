package api

import (
	"m3u82mp4/api/middleware"
	"m3u82mp4/consts/errcode"
	"m3u82mp4/model/db"
	"m3u82mp4/model/user"
	"m3u82mp4/utils"

	"github.com/gin-gonic/gin"
)

type UserApi struct {
	BaseApi
}

func InitUserRouter(r *gin.RouterGroup) {
	api := &UserApi{}
	api.RouterGroup = r.Group("/user", middleware.SetDB)
	api.Api("POST", "/register", api.Register, middleware.IgnoreUser)
	api.Api("POST", "/login", api.login, middleware.IgnoreUser)
}

// Register 账号注册
func (u *UserApi) Register(c *gin.Context) any {
	res := &Respone{}
	var param user.RegisterParam
	c.ShouldBindJSON(&param)
	if b, msg := param.Check(); b {
		u.Logger.Sugar().Error(msg)
		return res.Set(errcode.PARAM_ERR)
	}
	userdb := db.NewUserDB(u.MysqlDB)
	count := userdb.GetAccountCount(param.Account)
	if count > 0 {
		userInfo := userdb.GetUserInfoByAccount(param.Account)
		if userInfo.Password != param.PassWord {
			return res.Set(errcode.USER_PASSWORD_ERR)
		}
		token, err := utils.Token(userInfo.Id)
		if err != nil {
			return res.Set(errcode.USER_LOGIN_ERR)
		}
		res.Data = map[string]string{
			"t": token,
		}
		return res.OK2()
	}
	err, id := utils.ID()
	if err != nil {
		return res.Set(errcode.USER_SAVE_ERR)
	}
	errb := userdb.Save(&db.User{
		Id:       id,
		Account:  param.Account,
		Password: param.PassWord,
	})
	if err != nil {
		return res.Set(*errb)
	}
	token, err := utils.Token(id)
	if err != nil {
		return res.Set(errcode.USER_LOGIN_ERR)
	}
	res.Data = map[string]string{
		"t": token,
	}
	return res.OK2()
}

// login 登录
func (u *UserApi) login(c *gin.Context) any {
	res := &Respone{}
	var param user.LoginParam
	c.ShouldBindJSON(&param)
	if b, msg := param.Check(); b {
		u.Logger.Sugar().Error(msg)
		return res.Set(errcode.PARAM_ERR)
	}
	userDB := db.NewUserDB(u.MysqlDB)
	userInfo := userDB.GetUserInfoByAccount(param.Account)
	if userInfo.Password != param.PassWord {
		return res.Set(errcode.USER_PASSWORD_ERR)
	}
	token, err := utils.Token(userInfo.Id)
	if err != nil {
		return res.Set(errcode.USER_LOGIN_ERR)
	}
	res.Data = map[string]string{
		"t": token,
	}
	return res.OK2()
}
