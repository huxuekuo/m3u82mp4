package db

import (
	"m3u82mp4/consts/errcode"

	"gorm.io/gorm"
)

type SystemConfig struct {
	Id         int64  `gorm:"column:id;primaryKey" json:"id"`
	Key        string `gorm:"column:key" json:"key"`
	Value      string `gorm:"column:value" json:"value"`
	Type       string `gorm:"column:type" json:"type"`
	Comment    string `gorm:"column:comment" json:"comment"`
	State      uint8  `gorm:"column:state" json:"state"`
	IsDel      bool   `gorm:"column:is_del"  json:"isDel"`
	Createtime int64  `gorm:"column:createtime;autoCreateTime:milli" json:"createTime"`
	Updatetime int64  `gorm:"column:updatetime;autoUpdateTime:milli" json:"updateTime"`
}

type SystemConfigDB struct {
	db *gorm.DB
}

func (SystemConfig) table() string {
	return "v_system_config"
}

func NewSystemConfigDB(db *gorm.DB) *SystemConfigDB {
	return &SystemConfigDB{
		db: db,
	}
}

func (s SystemConfigDB) Create(config *SystemConfig) (*errcode.ErrorCode, bool) {
	err := s.db.Table("v_system_config").Save(config).Error
	if err != nil {
		custome := errcode.DBCustom(err.Error())
		return &custome, false
	}
	return nil, true
}

func (s SystemConfigDB) QueryList(types, key string) (*errcode.ErrorCode, []*SystemConfig) {
	var list []*SystemConfig
	db := s.db.Table("v_system_config").Where("is_del = ?", false)

	if types != "" {
		db = db.Where("type = ?", types)
	}
	if key != "" {
		db = db.Where("key LIKE ?", "%"+key+"%")
	}

	err := db.Find(&list).Error
	if err != nil {
		custome := errcode.DBCustom(err.Error())
		return &custome, nil
	}
	return nil, list
}

func (s SystemConfigDB) QueryOne(types, key string) (*errcode.ErrorCode, *SystemConfig) {
	var config *SystemConfig
	err := s.db.Table("v_system_config").Limit(1).Where("type = ? and key = ?", types, key, &config).Error
	if err != nil {
		custome := errcode.DBCustom(err.Error())
		return &custome, nil
	}
	return nil, config
}

func (s SystemConfigDB) Delete(id int64) *errcode.ErrorCode {
	err := s.db.Table("v_system_config").Where("id = ?", id).Update("is_del", true).Error
	if err != nil {
		custome := errcode.DBCustom(err.Error())
		return &custome
	}
	return nil
}
