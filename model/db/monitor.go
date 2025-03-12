package db

import (
	"m3u82mp4/consts/errcode"

	"gorm.io/gorm"
)

type Monitor struct {
	Id         int64 `gorm:"column:id;primaryKey" json:"id"`
	Max        int   `gorm:"column:max" json:"max"`
	VideoId    int64 `gorm:"column:video_id"  json:"video_id"`
	IsDel      bool  `gorm:"column:is_del"  json:"isDel"`
	UserID     int64 `gorm:"column:user_id"  json:"userId"`
	Createtime int64 `gorm:"column:createtime;autoCreateTime:milli" json:"createTime"`
	Updatetime int64 `gorm:"column:updatetime;autoUpdateTime:milli" json:"updateTime"`
}

type MonitorDB struct {
	db *gorm.DB
}

func NewMonitorDB(db *gorm.DB) *MonitorDB {
	instance := new(MonitorDB)
	instance.db = db
	return instance
}

func (m *MonitorDB) List() ([]Monitor, *errcode.ErrorCode) {
	var list []Monitor
	err := m.db.Where("is_del=?", 0).Find(&list).Error
	if err != nil {
		errc := errcode.DBCustom(err.Error())
		return nil, &errc
	}
	return list, nil
}
