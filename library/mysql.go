package library

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MysqlDB *gorm.DB

func InitMysql() *gorm.DB {
	MysqlDB, _ = gorm.Open(mysql.Open("root:yiqizou89.@tcp(152.136.33.217:3510)/video_server?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	return MysqlDB
}
