package library

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MysqlDB *gorm.DB

func InitMysql() *gorm.DB {
	MysqlDB, _ = gorm.Open(mysql.Open("root:yiqizou89.@tcp(62.234.47.210:3150)/video_server?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	return MysqlDB
}
