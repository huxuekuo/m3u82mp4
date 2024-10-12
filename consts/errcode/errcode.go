package errcode

type ErrorCode struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

var (
	SYS_USER_READ_ERR = ErrorCode{Code: 10001, Msg: "读取用户信息错误,请刷新页面！"}
)

var (
	PARAM_ERR = ErrorCode{Code: 5001, Msg: "参数错误"}
)

var (
	USER_ACCOUNT_REPEAT = ErrorCode{Code: 6001, Msg: "账号重复"}
	USER_SAVE_ERR       = ErrorCode{Code: 6002, Msg: "注册失败"}
	USER_PASSWORD_ERR   = ErrorCode{Code: 6003, Msg: "密码错误"}
	USER_LOGIN_ERR      = ErrorCode{Code: 6003, Msg: "登录失败"}
)
