package consts

import "m3u82mp4/library"

var LoadSysConf = map[string]any{
	SYS_LOGGER_VAR_NAME:         library.Logger,
	SYS_MYSQL_VAR_NAME:          library.MysqlDB,
	SYS_AUTH_USER_INFO_VAR_NAME: nil,
}

const (
	SYS_LOGGER_VAR_NAME          = "logger"
	SYS_MYSQL_VAR_NAME           = "mysql"
	SYS_AUTH_USER_INFO_VAR_NAME  = "authUserInfo"
	SYS_AUTH_USER_REDIS_KEY_NAME = "urk"
)

const (
	SYS_VIDEO_QUERY         = "query"
	SYS_VIDEO_QUERY_ADD     = "add"
	SYS_VIDEO_QUERY_GET     = "get"
	SYS_VIDEO_QUERY_DELETED = "deleted"
)

const SYS_SEQUENCE_ID = 1
