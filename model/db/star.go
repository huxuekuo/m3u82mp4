package db

import (
	"m3u82mp4/consts/errcode"

	"gorm.io/gorm"
)

// Star 收藏表
type Star struct {
	Id         int64  `gorm:"column:id;primaryKey" json:"id"`
	TvID       string `gorm:"column:tv_id" json:"tvId"`
	Name       string `gorm:"column:name" json:"name"`
	Url        string `gorm:"column:url" json:"url"`
	IsDel      bool   `gorm:"column:is_del"  json:"isDel"`
	UserID     int64  `gorm:"column:user_id"  json:"userId"`
	Createtime int64  `gorm:"column:createtime;autoCreateTime:milli" json:"createTime"`
	Updatetime int64  `gorm:"column:updatetime;autoUpdateTime:milli" json:"updateTime"`
}

type StarDB struct {
	db *gorm.DB
}

func (Star) TableName() string {
	return "v_star"
}

func NewStarDB(db *gorm.DB) *StarDB {
	stardb := new(StarDB)
	stardb.db = db
	return stardb
}

// Create _
func (s *StarDB) Create(star *Star) *errcode.ErrorCode {
	rt := s.db.Table("v_star").Save(star)
	if rt.Error != nil {
		er := errcode.DBCustom(rt.Error.Error())
		return &er
	}
	return nil
}

// ListAll 查询全部内容
func (s *StarDB) ListAll(userID int64) (*errcode.ErrorCode, []Star) {
	var stars []Star
	rt := s.db.Table("v_star").Where("user_id = ?", userID).Find(&stars)
	if rt.Error != nil {
		er := errcode.DBCustom(rt.Error.Error())
		return &er, nil
	}
	return nil, stars
}

// Exists 验证是否以重复添加
func (s *StarDB) Exists(UserID int64, tvID string) (*errcode.ErrorCode, bool) {
	var count int64
	rt := s.db.Table("v_star").Where("user_id=? and tv_id=?", UserID, tvID).Count(&count)
	if rt.Error != nil {
		er := errcode.DBCustom(rt.Error.Error())
		return &er, false
	}
	if count > 0 {
		return nil, true
	}
	return nil, false
}
