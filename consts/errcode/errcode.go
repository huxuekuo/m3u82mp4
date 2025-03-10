package errcode

type ErrorCode struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

var (
	SYS_USER_READ_ERR = ErrorCode{Code: 1001, Msg: "读取用户信息错误,请刷新页面！"}
	SYS_RETRY         = ErrorCode{Code: 1002, Msg: "系统错误请重试"}
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

var (
	VIDEO_QUERY_KEY_LEN_ERR = ErrorCode{Code: 7001, Msg: "关键词长度要大于1哦~"}
	VIDEO_DOWNLOAD_ERR      = ErrorCode{Code: 7002, Msg: "请重试，当前下载失败~"}
)

func DBCustom(err string) ErrorCode {
	return ErrorCode{Code: 1001, Msg: err}
}

func ApiCustom(err string) ErrorCode {
	return ErrorCode{Code: 1002, Msg: err}
}
